package basementgen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	"github.com/chas3air/StructureGenerator/services/filesGen"
)

var basementDirsNamesDefautl = []string{
	"cmd/",
	"cmd/app/",
	"internal/",
	"internal/app/",
	"internal/config/",
	"internal/models/",
	"internal/services/",
}

var basementFilesNamesDefautl = []string{
	"cmd/app/main.go",
	"internal/app/app.go",
	"internal/config/config.go",
	"internal/models/models.go",
	"internal/services/services.go",
}

type BasementDataGenerator struct {
	ConfigDirNamesDefault   []string
	ConfigFilesNamesDefault []string
}

func (bdg *BasementDataGenerator) GenerateDirectory() {
	for _, v := range bdg.ConfigDirNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
}

func (bdg *BasementDataGenerator) GenerateFiles() {
	for _, v := range bdg.ConfigFilesNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
}

func NewBasementDataGenerator(dirsNames, filesNames []string) packagesgen.CommonGenFactory {
	if dirsNames == nil {
		dirsNames = make([]string, len(basementDirsNamesDefautl))
		copy(dirsNames, basementDirsNamesDefautl)
	}
	if filesNames == nil {
		filesNames = make([]string, len(basementFilesNamesDefautl))
		copy(filesNames, basementFilesNamesDefautl)
	}

	return &BasementDataGenerator{
		ConfigDirNamesDefault:   dirsNames,
		ConfigFilesNamesDefault: filesNames,
	}
}
