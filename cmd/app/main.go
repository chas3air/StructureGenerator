package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	internalDirName := "internal/"
	MakingDirProcedure(internalDirName)

	appDirName := internalDirName + "app/"
	MakingDirProcedure(appDirName)
	CreateFileInTheFolderOfTheSameName(appDirName + "app.go")

	config_iDirName := internalDirName + "config/"
	MakingDirProcedure(config_iDirName)
	CreateFileInTheFolderOfTheSameName(config_iDirName + "config.go")

	modelsDirName := internalDirName + "models/"
	MakingDirProcedure(modelsDirName)
	CreateFileInTheFolderOfTheSameName(modelsDirName + "models.go")

	servicesDirName := internalDirName + "services/"
	MakingDirProcedure(servicesDirName)
	CreateFileInTheFolderOfTheSameName(servicesDirName + "services.go")

	config_gDirName := "config/"
	MakingDirProcedure(config_gDirName)
	CreateFileInTheFolderOfTheSameName(appDirName + "config.yaml")

	fmt.Println("Success")
}

func MakingDirProcedure(dirName string) {
	err := os.Mkdir(dirName, 0750)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Dir '" + dirName + "' already exists")
		} else {
			fmt.Println("Error:", err)
		}
	}

	log.Println("Success")
}

func CreateFileInTheFolderOfTheSameName(nameOfFile string) {
	_, err := os.Create(nameOfFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
