package filesGen

import (
	"fmt"
	"log"
	"os"
)

func MakingDirProcedure(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Dir '" + dirName + "' already exists")
		} else {
			fmt.Println("Error:", err)
		}

		return
	}

	log.Println("Successfully created '" + dirName + "' directory")
}

func MakingFileProcedure(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	log.Println("Successfully created '" + fileName + "' file")
}
