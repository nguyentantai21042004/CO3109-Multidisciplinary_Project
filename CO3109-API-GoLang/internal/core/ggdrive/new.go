package ggdrive

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type googleDriveClient struct {
	service *drive.Service
	config  *oauth2.Config
}

var _ Usecase = (*googleDriveClient)(nil)

// New creates a new Google Drive client instance
func New(ctx context.Context, oauthConfig *oauth2.Config) (Usecase, error) {
	if oauthConfig.ClientID == "" || oauthConfig.ClientSecret == "" {
		return nil, fmt.Errorf("invalid oauth configuration")
	}

	// 1. Initialize OAuth2 configuration
	oauthConfig.Scopes = []string{drive.DriveScope}

	// 2. Create Drive service client without token source initially
	// The token will be obtained through the OAuth2 flow when needed
	service, err := drive.NewService(ctx, option.WithHTTPClient(oauthConfig.Client(ctx, nil)))
	if err != nil {
		return nil, fmt.Errorf("failed to create drive service: %w", err)
	}

	return &googleDriveClient{
		service: service,
		config:  oauthConfig,
	}, nil
}
