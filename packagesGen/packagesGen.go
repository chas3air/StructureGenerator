package packagesgen

import (
	"fmt"
	"log"
)

func GenPackageByName(packageName string) (bool, error) {
	var isPackageNeeds string = "yes"

	log.Println("Do you need " + packageName + "?(default: yes) yes/no")

	i, err := fmt.Scan(&isPackageNeeds)
	if err != nil {
		return false, err
	} else if i != 1 {
		return false, fmt.Errorf("error scanning")
	}

	if isPackageNeeds == "yes" {
		return true, nil
	} else {
		return false, nil
	}
}
