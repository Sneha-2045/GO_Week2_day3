package main

import (
	"fmt"
	"os"
	"sync"
)

func processFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return
	}

	fmt.Printf("File: %s | Size: %d bytes\n", filename, len(data))
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go processFile(file, &wg)
	}

	wg.Wait()
	fmt.Println("All files processed successfully!")
}