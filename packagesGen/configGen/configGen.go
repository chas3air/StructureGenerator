package configgen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	"github.com/chas3air/StructureGenerator/services/filesGen"
)

var configDirNamesDefault = []string{
	"config/",
}

var configFilesNamesDefault = []string{
	"config/local.yaml",
}

type ConfigDataGenerator struct {
	ConfigDirNamesDefault   []string
	ConfigFilesNamesDefault []string
}

func (cdg *ConfigDataGenerator) GenerateDirectory() {
	for _, v := range cdg.ConfigDirNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
}

func (cdg *ConfigDataGenerator) GenerateFiles() {
	for _, v := range cdg.ConfigFilesNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
}

func NewConfigDataGenerator(dirsNames, filesNames []string) packagesgen.CommonGenFactory {
	if dirsNames == nil {
		dirsNames = make([]string, 0, len(configDirNamesDefault))
		copy(dirsNames, configDirNamesDefault)
	}
	if filesNames == nil {
		filesNames = make([]string, 0, len(configFilesNamesDefault))
		copy(filesNames, configFilesNamesDefault)
	}

	return &ConfigDataGenerator{
		ConfigDirNamesDefault:   dirsNames,
		ConfigFilesNamesDefault: filesNames,
	}
}
