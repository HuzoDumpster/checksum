package main

import (
	"bufio"
	"container/list"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Function to read a file and return the data
func ReadFileData(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string
	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		result += string(buf[:n])
	}

	return result, nil
}

// Hashing functions
func sha256file(filename string, checksum bool, path bool, date bool) (string, error) {
	fData, _ := ReadFileData(filename)
	h := sha256.New()

	h.Write([]byte(fData))
	bs := h.Sum(nil)

	var retstring string

	if checksum {
		retstring += fmt.Sprintf("%x", bs)
	}
	if path {
		if retstring != "" {
			retstring += " ; "
		}
		retstring += filename
	}
	if date {
		if retstring != "" {
			retstring += " ; "
		}
		retstring += time.Now().Format(time.RFC3339)
	}

	fmt.Println(retstring)
	return retstring, nil
}

func main() {
	fmt.Println("Checksum application v1.0")
	// Define Flags & Help Messages
	var fileFlag = flag.String("file", "", "Path to a single file.")
	var folderFlag = flag.String("folder", "", "Path to a folder with files to get the checksum of.")
	var formatFlag = flag.String("format", "cpd", "Choose output format flags: `c` = checksum, `p` = path, `d` = date (default = `cpd`)")

	// Flag Parsing
	flag.Parse()
	if (*fileFlag == "") && (*folderFlag == "") {
		fmt.Println("You're required to specify either a file or a folder.")
	}

	// Get a list of files
	filePaths := list.New()
	if *fileFlag != "" {
		// Check if the file exists
		if _, err := os.Stat(*fileFlag); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("The specified file to the flag --file, does not exist. Make sure the path is correct.")
				return
			}
		}

		// Add it to the list.
		filePaths.PushBack(*fileFlag)
	}
	if *folderFlag != "" {
		entries, err := os.ReadDir(*folderFlag)
		if err != nil {
			fmt.Println("Error reading folder:", err)
			return
		}
		// Make sure there are files in the folder.
		if len(entries) < 1 {
			fmt.Println("The specified folder path to the flag --folder, does not contain any files. Make sure the folder is not empty.")
		}

		// Add the files to the list.
		for _, e := range entries {
			if !e.IsDir() {
				fullPath := filepath.Join(*folderFlag, e.Name())
				filePaths.PushBack(fullPath)
			}
		}
	}

	for e := filePaths.Front(); e != nil; e = e.Next() {
		filePath := e.Value.(string)

		CHECKSUM := true
		PATH := true
		DATE := false

		if *formatFlag != "" {
			CHECKSUM, PATH, DATE = false, false, false

			if strings.Contains(*formatFlag, "c") {
				CHECKSUM = true
			}
			if strings.Contains(*formatFlag, "p") {
				PATH = true
			}
			if strings.Contains(*formatFlag, "d") {
				DATE = true
			}
		}

		sha256file(filePath, CHECKSUM, PATH, DATE)
	}

}
