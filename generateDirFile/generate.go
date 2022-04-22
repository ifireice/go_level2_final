package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	var basicDir string
	// example
	// basicDir = "/home/alexgo/testDir"
	fmt.Scanln(&basicDir)

	// создаем папки только на верхнем уровне
	for i := 0; i < 5; i++ {
		err := os.MkdirAll(basicDir+"/"+"dir"+strconv.Itoa(rand.Int()%10), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup

	// заполняем папки файлами, включая и basicDir
	err := filepath.Walk(basicDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup) {
					fullPath := path + "/file" + strconv.Itoa(rand.Int()%10)
					f, err := os.Create(fullPath)
					defer f.Close()
					if err != nil {
						log.Fatal(err)
					}
					wg.Done()
				}(&wg)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
