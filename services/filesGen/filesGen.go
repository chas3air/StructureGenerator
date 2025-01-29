package filesGen

import (
	"fmt"
	"log"
	"os"
)

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

	log.Println("Successfully created '" + dirName + "' directory")
}

// func CreateFileInTheFolderOfTheSameName(nameOfFile string) {
// 	_, err := os.Create(nameOfFile)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	log.Println("Successfully created '" + nameOfFile + "' file")
// }
