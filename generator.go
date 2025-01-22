package generator

import (
	"fmt"

	"github.com/chas3air/StructureGenerator/services/filesGen"
	packagesgen "github.com/chas3air/StructureGenerator/services/packagesGen"
)

const nameOfConfigPackage = "config"
const nameOfLoggerPackage = "logger"

var nondefaultDirNames = [][]string{configDirNames, loggerDirNames}
var nondefaultFileNames = [][]string{configFileNames, loggerFileNames}

func WorkPart() {
	fmt.Println("Structure directory generator...")

	for _, v := range defaultDirNames {
		filesGen.MakingDirProcedure(v)
	}

	for _, v := range defaultFileNames {
		filesGen.CreateFileInTheFolderOfTheSameName(v)
	}

	for i, v := range []string{nameOfConfigPackage, nameOfLoggerPackage} {
		isNeeds, err := packagesgen.GenPackageByName(v)
		if err != nil {
			fmt.Println("Package " + v + " will not created")
		} else if !isNeeds {
			fmt.Println("Package " + v + " will not created")
		} else {
			fmt.Println("Package " + v + " will created")

			for _, value := range nondefaultDirNames[i] {
				filesGen.MakingDirProcedure(value)
			}

			for _, value := range nondefaultFileNames[i] {
				filesGen.CreateFileInTheFolderOfTheSameName(value)
			}
		}
	}

	fmt.Println("Full directory structure was generated")
}

var defaultDirNames = []string{
	"cmd/",
	"cmd/app/",
	"internal/",
	"internal/app/",
	"internal/config/",
	"internal/models/",
	"internal/services/",
}

var defaultFileNames = []string{
	"cmd/app/main.go",
	"internal/app/app.go",
	"internal/config/config.go",
	"internal/models/models.go",
	"internal/services/services.go",
}

var configDirNames = []string{
	"config/",
}

var configFileNames = []string{
	"config/local.yaml",
}

var loggerDirNames = []string{
	"internal/lib",
	"internal/lib/logger",
	"internal/lib/logger/handler",
	"internal/lib/logger/handler/slogdiscard",
	"internal/lib/logger/handler/slogpretty",
	"internal/lib/logger/sl",
}

var loggerFileNames = []string{
	"internal/lib/logger/handler/slogdiscard/slogdiscard.go",
	"internal/lib/logger/handler/slogpretty/slogpretty.go",
	"internal/lib/logger/sl/sl.go",
}
