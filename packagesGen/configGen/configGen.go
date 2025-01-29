package configgen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	"github.com/chas3air/StructureGenerator/services/objectgen"
)

var configDirNamesDefault = []string{
	"config/",
}

var configFilesNamesDefault = []string{
	"config/local.yaml",
}

type ConfigDataGenerator struct {
	ConfigDirsNames  []string
	ConfigFilesNames []string
}

func (cdg *ConfigDataGenerator) GenerateDirectory() {
	for _, v := range cdg.ConfigDirsNames {
		objectgen.MakingDirProcedure(v)
	}
}

func (cdg *ConfigDataGenerator) GenerateFiles() {
	for _, v := range cdg.ConfigFilesNames {
		objectgen.MakingFileProcedure(v)
	}
}

func (cdg *ConfigDataGenerator) GetDirsNames() []string {
	return cdg.ConfigDirsNames
}
func (cdg *ConfigDataGenerator) GetFilesNames() []string {
	return cdg.ConfigFilesNames
}

func NewConfigDataGenerator(dirsNames, filesNames []string) packagesgen.CommonGenFactory {
	if dirsNames == nil {
		dirsNames = make([]string, len(configDirNamesDefault))
		copy(dirsNames, configDirNamesDefault)
	}
	if filesNames == nil {
		filesNames = make([]string, len(configFilesNamesDefault))
		copy(filesNames, configFilesNamesDefault)
	}

	return &ConfigDataGenerator{
		ConfigDirsNames:  dirsNames,
		ConfigFilesNames: filesNames,
	}
}
