package loggergen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	objectgen "github.com/chas3air/StructureGenerator/services/objectGen"
)

var loggerDirsNamesDefault = []string{
	"internal/lib/",
	"internal/lib/logger/",
	"internal/lib/logger/handler/",
	"internal/lib/logger/handler/slogpretty/",
	"internal/lib/logger/sl/",
}

var loggerFilesNamesDefault = []string{
	"internal/lib/logger/handler/slogpretty/slogpretty.go",
	"internal/lib/logger/sl/sl.go",
}

type LoggerDataGenerator struct {
	ConfigDirsNames  []string
	ConfigFilesNames []string
}

func (ldg *LoggerDataGenerator) GenerateDirectory() {
	for _, v := range ldg.ConfigDirsNames {
		objectgen.MakingDirProcedure(v)
	}
}

func (ldg *LoggerDataGenerator) GenerateFiles() {
	for _, v := range ldg.ConfigFilesNames {
		objectgen.MakingFileProcedure(v)
	}
}

func (ldg *LoggerDataGenerator) GetDirsNames() []string {
	return ldg.ConfigDirsNames
}
func (ldg *LoggerDataGenerator) GetFilesNames() []string {
	return ldg.ConfigFilesNames
}

func NewLoggerDataGenerator(dirsNames, filesNames []string) packagesgen.CommonGenFactory {
	if dirsNames == nil {
		dirsNames = make([]string, len(loggerDirsNamesDefault))
		copy(dirsNames, loggerDirsNamesDefault)
	}
	if filesNames == nil {
		filesNames = make([]string, len(loggerFilesNamesDefault))
		copy(filesNames, loggerFilesNamesDefault)
	}

	return &LoggerDataGenerator{
		ConfigDirsNames:  dirsNames,
		ConfigFilesNames: filesNames,
	}
}
