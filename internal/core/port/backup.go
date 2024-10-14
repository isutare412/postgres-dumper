package port

import (
	"context"
	"io"

	"github.com/isutare412/postgres-s3-dumper/internal/core/model"
)

type BackupExecutor interface {
	BackupDatabase(ctx context.Context, req model.DatabaseBackupRequest) (backup io.Reader, err error)
}
