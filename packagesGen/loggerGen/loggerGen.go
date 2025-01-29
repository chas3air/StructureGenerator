package loggergen

import (
	packagesgen "github.com/chas3air/StructureGenerator/packagesGen"
	"github.com/chas3air/StructureGenerator/services/filesGen"
)

var loggerDirsNamesDefault = []string{
	"internal/lib",
	"internal/lib/logger",
	"internal/lib/logger/handler",
	"internal/lib/logger/handler/slogdiscard",
	"internal/lib/logger/sl",
}

var loggerFilesNamesDefault = []string{
	"internal/lib/logger/handler/slogpretty/slogpretty.go",
	"internal/lib/logger/sl/sl.go",
}

type LoggerDataGenerator struct {
	ConfigDirsNamesDefault  []string
	ConfigFilesNamesDefault []string
}

func (ldg *LoggerDataGenerator) GenerateDirectory() {
	for _, v := range ldg.ConfigDirsNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
}

func (ldg *LoggerDataGenerator) GenerateFiles() {
	for _, v := range ldg.ConfigFilesNamesDefault {
		filesGen.MakingDirProcedure(v)
	}
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
		ConfigDirsNamesDefault:  dirsNames,
		ConfigFilesNamesDefault: filesNames,
	}
}
