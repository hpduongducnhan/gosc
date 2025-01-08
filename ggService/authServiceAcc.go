package ggservice

import (
	"context"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewGgSheetService(ctx context.Context, credsPath string) (*sheets.Service, error) {
	sheetsService, err := sheets.NewService(
		ctx,
		option.WithCredentialsFile(credsPath),
	)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
		return nil, err
	}
	return sheetsService, nil
}

func NewGgDriveService(ctx context.Context, credsPath string) (*drive.Service, error) {
	driveService, err := drive.NewService(
		ctx,
		option.WithCredentialsFile(credsPath),
	)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
		return nil, err
	}
	return driveService, nil
}
