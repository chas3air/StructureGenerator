package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	internalDirName := "internal/"
	MakingDirProcedure(internalDirName)

	for _, v := range dirNames {
		MakingDirProcedure(v)
	}

	for _, v := range fileNames {
		CreateFileInTheFolderOfTheSameName(v)
	}

	fmt.Println("Full directory structure was generated")
}

func MakingDirProcedure(dirName string) {
	err := os.Mkdir(dirName, 0750)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Dir '" + dirName + "' already exists")
		} else {
			fmt.Println("Error:", err)
		}

		return
	}

	log.Println("Successfully created'" + dirName + "' directory")
}

func CreateFileInTheFolderOfTheSameName(nameOfFile string) {
	_, err := os.Create(nameOfFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	log.Println("Successfully created'" + nameOfFile + "' file")
}

var dirNames = []string{
	"cmd/",
	"cmd/app/",
	"internal/",
	"internal/app/",
	"internal/config/",
	"internal/models/",
	"internal/services/",
	"config/",
}

var fileNames = []string{
	"cmd/app/main.go",
	"internal/app/app.go",
	"internal/config/config.go",
	"internal/models/models.go",
	"internal/services/services.go",
	"config/config.yaml",
}
