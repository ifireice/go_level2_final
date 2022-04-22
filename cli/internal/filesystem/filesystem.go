package filesystem

import (
	"github.com/alexgo92/go_level2_final/cli/internal/filesystem/mock"
	"github.com/alexgo92/go_level2_final/cli/internal/filesystem/real"
	model "github.com/alexgo92/go_level2_final/cli/internal/model/filesystem"
)

type Filesystem interface {
	ListFiles(dirPath string) (*model.FileStats, error)
	DeleteFile(dirPath, name string, fileSize int) error
}

func NewFileSystem(kind string, basicPath string) (Filesystem, error) {
	switch kind {
	case mock.FileSystemKind:
		return mock.NewFileSystem(), nil
	case real.FileSystemKind:
		return real.NewFileSystem(basicPath), nil
	}
	return nil, model.ErrInvalidFilesystem{FilesystemKind: kind}
}
