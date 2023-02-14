package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func isGoFile(name string) bool {
	return filepath.Ext(name) == ".go" || filepath.Ext(name) == ".mod"
}

const SKIP_PATH = ".scripts"

func main() {
	var name string
	var oldName string
	flag.StringVar(&name, "name", "", "project name")
	flag.StringVar(&oldName, "old", "go-template", "old project name")
	flag.Parse()
	if name == "" {
		log.Fatal("project name is empty")
	}
	wg := &sync.WaitGroup{}

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasPrefix(path, SKIP_PATH) || !isGoFile(info.Name()) {
				return nil
			}
			fmt.Printf("rename %s to %s\n", path, name)
			wg.Add(1)
			go func() {
				defer wg.Done()
				replaceProjectNameWith(path, oldName, name)
			}()
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("waiting...")
	wg.Wait()
	fmt.Println("done")
}

func replaceProjectNameWith(path, from, to string) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	s := string(b)
	s = strings.ReplaceAll(s, from, to)
	os.WriteFile(path, []byte(s), 0644)
}
