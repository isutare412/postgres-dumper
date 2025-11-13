package wire

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/isutare412/postgres-dumper/internal/aws"
	"github.com/isutare412/postgres-dumper/internal/backup"
	"github.com/isutare412/postgres-dumper/internal/config"
	corebackup "github.com/isutare412/postgres-dumper/internal/core/service/backup"
)

type App struct {
	cfg *config.Config

	backupService *corebackup.Service
}

func NewApp(cfg *config.Config) (*App, error) {
	backupExecutor := backup.NewExecutor()

	s3Client, err := aws.NewS3Client(cfg.ToAWSS3Config())
	if err != nil {
		return nil, fmt.Errorf("creating aws s3 client: %w", err)
	}

	backupService := corebackup.NewService(cfg.ToBackupConfig(), backupExecutor, s3Client)

	return &App{
		cfg:           cfg,
		backupService: backupService,
	}, nil
}

func (a *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ProcessTimeout)
	defer cancel()

	if err := a.backupService.BackupDatabase(ctx); err != nil {
		return fmt.Errorf("backup database: %w", err)
	}
	slog.Info("database backup completed successfully")

	if err := a.backupService.PruneDatabaseBackups(ctx); err != nil {
		return fmt.Errorf("pruning database backups: %w", err)
	}
	slog.Info("database backup pruning completed successfully")

	return nil
}
