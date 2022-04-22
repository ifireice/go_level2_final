package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alexgo92/go_level2_final/cli/internal/filesystem"
)

// что бы работало с обычной fs нужно запустить c -fs real
var fsKind = flag.String("fs", "real", "file system: mock or real")

func main() {
	flag.Parse()

	fmt.Printf("Enter directory path:\n")

	var directory string
	fmt.Scanf("%s", &directory)

	fs, err := filesystem.NewFileSystem(*fsKind, directory)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// ListFiles - возвращает структуру FileStats - список файлов
	// но у FileStats есть метод FindDuplicates, который тоже вернет список файлов-дубликатов
	fileStats, err := fs.ListFiles(directory)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("\nFound files:\n\n%s", fileStats)

	duplicatesStats := fileStats.FindDuplicates()
	if len(duplicatesStats.List) > 0 {
		fmt.Printf("\nFound duplicates:\n\n%s\n", duplicatesStats)
		fmt.Printf("Which files you would like to delete (enter comma-separated list):\n")
		var filesToDelete string
		fmt.Scanf("%s", &filesToDelete)
		if filesToDelete != "" {
			fmt.Printf("\n")
			var deletedFilesCounter int
			files := strings.Split(filesToDelete, ",")
			for _, fileIndStr := range files {
				fileInd, _ := strconv.Atoi(fileIndStr)
				dirPath := duplicatesStats.List[fileInd-1].ParentDir
				fileName := duplicatesStats.List[fileInd-1].Name
				fileSize := duplicatesStats.List[fileInd-1].SizeBytes
				fmt.Printf("Deleting file [%s] from directory [%s]...\n", fileName, dirPath)
				err := fs.DeleteFile(dirPath, fileName, fileSize)
				if err != nil {
					fmt.Printf("Failed to delete file [%s] from directory [%s], error [%s], skipping...\n", fileName, dirPath, err)
				} else {
					fmt.Printf("Deleted file [%s] from directory [%s]\n", fileName, dirPath)
					deletedFilesCounter++
				}
			}
			// отображаем что удалили и что осталось
			if deletedFilesCounter > 0 {
				fmt.Printf("Successfuly deleted %d duplicates\n\n", deletedFilesCounter)
				fileStats, _ = fs.ListFiles(directory)
				fmt.Printf("Current directory file list:\n\n%s", fileStats)
			}
		}
	}
}
