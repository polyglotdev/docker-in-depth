package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	log.Println("Starting writedata...")
	if len(os.Args) <= 1 {
		log.Fatal("Usage: writedata /path/to/file")
		log.Fatal("No file specified")
	}

	filename := os.Args[1]
	if err := writeToFile(filename); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func writeToFile(filename string) error {
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("error getting hostname: %v", err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}()

	for i := 0; i < 20; i++ {
		data := fmt.Sprintf("Host: %s\nLoop Iteration: %d\n", hostname, i)
		if _, err := file.WriteString(data); err != nil {
			log.Printf("error writing to file during iteration %d: %v", i, err)
		}
	}

	return nil
}