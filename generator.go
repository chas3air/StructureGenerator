package generator

import (
	"fmt"
	"log"

	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	basementgen "github.com/chas3air/StructureGenerator/packagesGen/basementGen"
	configgen "github.com/chas3air/StructureGenerator/packagesGen/configGen"
	loggergen "github.com/chas3air/StructureGenerator/packagesGen/loggerGen"
	dirfiller "github.com/chas3air/StructureGenerator/services/dirsFiller"
	"github.com/chas3air/StructureGenerator/source"
)

const nameOfConfigPackage = "config"
const nameOfLoggerPackage = "logger"

func WorkPart() {
	log.Println("Structure directory generator...")
	var genFactory packagesgen.CommonGenFactory = basementgen.NewBasementDataGenerator(nil, nil)
	var dirFiller *dirfiller.DirFiller = &dirfiller.DirFiller{}

	log.Println("Generate default dirs and files")
	genFactory.GenerateDirectory()
	genFactory.GenerateFiles()

	if err := dirFiller.Fill(
		dirfiller.CreateMapFromDestAndSource(
			genFactory.GetFilesNames(),
			[]string{
				source.SourceAppMain,
				source.SourceInternalApp,
				source.SourceInternalConfig,
			},
		),
	); err != nil {
		log.Println("Error filling basement stack of dirs")
	}

	fmt.Println("At this point you need to choose whether you need modules or not")

	if isNeed, err := packagesgen.GenPackageByName(nameOfConfigPackage); isNeed && err == nil {
		genFactory = configgen.NewConfigDataGenerator(nil, nil)
		genFactory.GenerateDirectory()
		genFactory.GenerateFiles()

		if err := dirFiller.Fill(
			dirfiller.CreateMapFromDestAndSource(
				genFactory.GetFilesNames(),
				[]string{
					source.SourceConfig,
				},
			),
		); err != nil {
			log.Println("Error filling config stack of dirs")
		}
	}

	if isNeed, err := packagesgen.GenPackageByName(nameOfLoggerPackage); isNeed && err == nil {
		genFactory = loggergen.NewLoggerDataGenerator(nil, nil)
		genFactory.GenerateDirectory()
		genFactory.GenerateFiles()

		if err := dirFiller.Fill(
			dirfiller.CreateMapFromDestAndSource(
				genFactory.GetFilesNames(),
				[]string{
					source.SourceLibLoggerHandlerSlogpretty,
					source.SourceLibLoggerSl,
				},
			),
		); err != nil {
			log.Println("Error filling logger` stack of dirs")
		}
	}

	log.Println("Full directory structure was generated")
}
