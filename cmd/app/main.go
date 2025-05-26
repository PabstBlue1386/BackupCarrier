package main

import (
	"fmt"
	"os"

	"github.com/PabstBlue1386/BackupCarrier/internal/fileutils"
)

func main() {
	path := "C:\\soft\\hello.txt"
	exists, err := fileutils.FileExists(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File not found")
		} else if os.IsPermission(err) {
			fmt.Printf("No access")
		}
	}

	if exists {
		fmt.Printf("The file exists")
	}

}
