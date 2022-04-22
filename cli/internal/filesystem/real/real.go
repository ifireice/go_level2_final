package real

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	model "github.com/alexgo92/go_level2_final/cli/internal/model/filesystem"
)

const FileSystemKind = "real"

var prefix string

func NewFileSystem(basicPath string) *FileSystem {
	prefix = basicPath

	list := []*model.File{}

	err := filepath.Walk(basicPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {

			pathForParentDir := path[:len(path)-len(info.Name())-1]
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			contentByte, err := ioutil.ReadAll(f)
			f.Close()
			content := string(contentByte)
			if err != nil {
				log.Fatal(err)
			}

			list = append(list,
				&model.File{Name: info.Name(), ParentDir: pathForParentDir,
					SizeBytes: int(info.Size()), Content: content})
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	fileTree := map[string][]*model.File{
		basicPath: list,
	}
	return &FileSystem{fileTree: fileTree}
}

type FileSystem struct {
	fileTree map[string][]*model.File
}

func (fs *FileSystem) ListFiles(dirPath string) (*model.FileStats, error) {
	list := make([]*model.File, 0)
	dir, ok := fs.fileTree[dirPath]
	if !ok {
		return nil, model.ErrDirNotFound{DirPath: dirPath}
	}
	for _, file := range dir {
		list = append(list, file)
	}
	return &model.FileStats{List: list}, nil
}

func (fs *FileSystem) DeleteFile(dirPath, name string, SizeBytes int) error {
	var forDelete string
	if strings.HasPrefix(dirPath, prefix) {
		forDelete = dirPath
		dirPath = prefix
	}
	dir, ok := fs.fileTree[dirPath]
	if !ok {
		return model.ErrDirNotFound{DirPath: dirPath}
	}
	for fileInd := range dir {
		if dir[fileInd].Name == name && dir[fileInd].SizeBytes == SizeBytes {
			err := os.Remove(forDelete + "/" + name)
			if err != nil {
				return model.ErrFaildToDeleteFile{FileName: name, DirPath: forDelete}
			}
			copy(dir[fileInd:], dir[fileInd+1:])
			dir[len(dir)-1] = nil
			dir = dir[:len(dir)-1]
		} else {
			continue
		}
		fs.fileTree[dirPath] = dir
		return nil
	}
	return model.ErrFileNotFound{DirPath: dirPath, FileName: name}
}
