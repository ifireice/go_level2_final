package filesystem

import "fmt"

type ErrFileNotFound struct {
	FileName string
	DirPath  string
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("file [%s] not found in [%s]", e.FileName, e.DirPath)
}

type ErrDirNotFound struct {
	DirPath string
}

func (e ErrDirNotFound) Error() string {
	return fmt.Sprintf("directory [%s] not found", e.DirPath)
}

type ErrFaildToDeleteFile struct {
	FileName string
	DirPath  string
}

func (e ErrFaildToDeleteFile) Error() string {
	return fmt.Sprintf("failed to delete file [%s/%s]", e.DirPath, e.FileName)
}
