package main

import (
	"log"
	"os"
	"testing"

	"github.com/alexgo92/go_level2_final/cli/internal/filesystem"
)

func Test_main(t *testing.T) {
	// создаем тестовую дерикторию с файлами дубликатами
	generateForTestDirFile()

	// тут код из main
	fs, err := filesystem.NewFileSystem("real", "./testDir")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fileStats, err := fs.ListFiles("./testDir")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// ожидаемый результат по файлам-дубликатам
	testCase := map[string]string{"testDir": "file1", "testDir/dir1": "file1"}
	// сюда запишем реальный результат по файлам-дубликатам
	realCase := make(map[string]string, 2)

	duplicatesStats := fileStats.FindDuplicates()

	for _, i := range duplicatesStats.List {
		realCase[i.ParentDir] = i.Name
	}

	for path, nme := range testCase {
		name, ok := realCase[path]
		if name == nme && ok {
			continue
		}
		t.Error(realCase)
	}
	os.RemoveAll("./testDir")
}

func generateForTestDirFile() {
	os.MkdirAll("./testDir/dir1", 0777)
	f, _ := os.Create("./testDir/file1")
	f.Close()
	f, _ = os.Create("./testDir/file2")
	f.Close()
	f, _ = os.Create("./testDir/dir1/file1")
	f.Close()
}
