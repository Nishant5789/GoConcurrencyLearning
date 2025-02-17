package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get environment variable
	homeDir := os.Getenv("part1")
	fmt.Println("Home Directory:", homeDir)

	// Create a new file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write to the file
	_, err = file.WriteString("Hello, Golang OS Package!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully!")

	// Read the file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:", string(content))

	// Create a new directory
	err = os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created successfully!")

	// Remove file and directory
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed successfully!")

	err = os.Remove("testdir")
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
	fmt.Println("Directory removed successfully!")
}
