package generator

import (
	"fmt"
	"log"

	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	basementgen "github.com/chas3air/StructureGenerator/packagesGen/basementGen"
	configgen "github.com/chas3air/StructureGenerator/packagesGen/configGen"
	loggergen "github.com/chas3air/StructureGenerator/packagesGen/loggerGen"
)

const nameOfConfigPackage = "config"
const nameOfLoggerPackage = "logger"

func WorkPart() {
	log.Println("Structure directory generator...")
	var genFactory packagesgen.CommonGenFactory = basementgen.NewBasementDataGenerator(nil, nil)

	log.Println("Generate default dirs and files")
	genFactory.GenerateDirectory()
	genFactory.GenerateFiles()

	fmt.Println("At this point you need to choose whether you need modules or not")

	if isNeed, err := packagesgen.GenPackageByName(nameOfConfigPackage); isNeed && err != nil {
		genFactory = configgen.NewConfigDataGenerator(nil, nil)
		log.Println("config factory with parameters", genFactory)
		genFactory.GenerateDirectory()
		genFactory.GenerateFiles()
	}

	if isNeed, err := packagesgen.GenPackageByName(nameOfLoggerPackage); isNeed && err != nil {
		genFactory = loggergen.NewLoggerDataGenerator(nil, nil)
		log.Println("logger factory with parameters", genFactory)
		genFactory.GenerateDirectory()
		genFactory.GenerateFiles()
	}

	log.Println("Full directory structure was generated")
}
