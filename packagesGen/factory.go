package packagesgen

type GenDirectoryFactory interface {
	GenerateDirectory([]string)
}

type GenFilesFactory interface {
	GenerateFiles([]string)
}

type CommonGenFactory interface {
	GenerateDirectory()
	GenerateFiles()

	GetDirsNames() []string
	GetFilesNames() []string
}
