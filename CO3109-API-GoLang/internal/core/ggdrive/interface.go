package ggdrive

import (
	"context"

	"google.golang.org/api/drive/v3"
)

// Usecase defines the interface for Google Drive operations
type Usecase interface {
	UploadFile(ctx context.Context, filePath, fileName string, parentFolderID string) (string, error)
	ListFiles(ctx context.Context, pageSize int64) ([]*drive.File, error)
	DownloadFile(ctx context.Context, fileID string) ([]byte, error)
	ShareFile(ctx context.Context, fileID, email string) error
	UploadLargeFile(ctx context.Context, filePath, fileName string, chunkSize int) (string, error)
	CreateFolder(ctx context.Context, name string, parentFolderID string) (string, error)
	DeleteFile(ctx context.Context, fileID string) error
	GetAuthURL() string
}
