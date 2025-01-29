package basementgen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	objectgen "github.com/chas3air/StructureGenerator/services/objectGen"
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
}

type BasementDataGenerator struct {
	ConfigDirNames   []string
	ConfigFilesNames []string
}

func (bdg *BasementDataGenerator) GenerateDirectory() {
	for _, v := range bdg.ConfigDirNames {
		objectgen.MakingDirProcedure(v)
	}
}

func (bdg *BasementDataGenerator) GenerateFiles() {
	for _, v := range bdg.ConfigFilesNames {
		objectgen.MakingFileProcedure(v)
	}
}

func (bdg *BasementDataGenerator) GetDirsNames() []string {
	return bdg.ConfigDirNames
}
func (bdg *BasementDataGenerator) GetFilesNames() []string {
	return bdg.ConfigFilesNames
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
		ConfigDirNames:   dirsNames,
		ConfigFilesNames: filesNames,
	}
}
