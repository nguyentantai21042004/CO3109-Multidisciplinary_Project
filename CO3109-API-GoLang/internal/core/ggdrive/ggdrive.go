package ggdrive

import (
	"context"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// GetAuthURL returns the authorization URL for Google Drive
func (g *googleDriveClient) GetAuthURL() string {
	return g.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

// UploadFile uploads a file to Google Drive
func (g *googleDriveClient) UploadFile(ctx context.Context, filePath, fileName string, parentFolderID string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("warning: failed to close file: %v\n", closeErr)
		}
	}()

	driveFile := &drive.File{
		Name:    fileName,
		Parents: []string{parentFolderID}, // Empty parent ID will upload to root
	}

	res, err := g.service.Files.Create(driveFile).
		Context(ctx).
		Media(file).
		Do()
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return res.Id, nil
}

// ListFiles lists files in Google Drive with pagination
func (g *googleDriveClient) ListFiles(ctx context.Context, pageSize int64) ([]*drive.File, error) {
	call := g.service.Files.List().
		Context(ctx).
		PageSize(pageSize).
		Fields("nextPageToken, files(id, name, mimeType, size, webViewLink)")

	res, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	return res.Files, nil
}

// DownloadFile downloads a file from Google Drive
func (g *googleDriveClient) DownloadFile(ctx context.Context, fileID string) ([]byte, error) {
	res, err := g.service.Files.Get(fileID).
		Context(ctx).
		Download()
	if err != nil {
		return nil, fmt.Errorf("failed to initiate download: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

// ShareFile shares a file with a specific email
func (g *googleDriveClient) ShareFile(ctx context.Context, fileID, email string) error {
	permission := &drive.Permission{
		Type:         "user",
		Role:         "reader",
		EmailAddress: email,
	}

	_, err := g.service.Permissions.Create(fileID, permission).
		Context(ctx).
		SendNotificationEmail(false).
		Do()
	if err != nil {
		return fmt.Errorf("failed to share file: %w", err)
	}

	return nil
}

// UploadLargeFile handles resumable uploads for large files
func (g *googleDriveClient) UploadLargeFile(ctx context.Context, filePath, fileName string, chunkSize int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	driveFile := &drive.File{Name: fileName}
	res, err := g.service.Files.Create(driveFile).
		Context(ctx).
		Media(file, googleapi.ChunkSize(chunkSize)).
		Do()
	if err != nil {
		return "", fmt.Errorf("failed to upload large file: %w", err)
	}

	return res.Id, nil
}

// CreateFolder creates a new folder in Google Drive
func (g *googleDriveClient) CreateFolder(ctx context.Context, name string, parentFolderID string) (string, error) {
	folder := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentFolderID},
	}

	res, err := g.service.Files.Create(folder).
		Context(ctx).
		Do()
	if err != nil {
		return "", fmt.Errorf("failed to create folder: %w", err)
	}

	return res.Id, nil
}

// DeleteFile deletes a file from Google Drive
func (g *googleDriveClient) DeleteFile(ctx context.Context, fileID string) error {
	err := g.service.Files.Delete(fileID).
		Context(ctx).
		Do()
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
