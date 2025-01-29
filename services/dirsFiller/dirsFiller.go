package dirsfiller

import (
	"log"
	"os"
	"path/filepath"
)

type DirFiller struct {
}

// fileSources это мапа.
// Ключ это название места куда нужно скопировать.
// Значение это место где лежит то что нужно копировать
func (df *DirFiller) Fill(fileSources map[string]string) error {
	for dest, source := range fileSources {
		if source == "" {
			log.Println("Error source")
			continue
		}

		sourceText, err := os.ReadFile(source)
		if err != nil {
			log.Println("Error of reading:", source, ", error:", err.Error())
			continue
		}

		dir := filepath.Dir(dest)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Println("Error creating directory:", dir, ", error:", err.Error())
			continue
		}

		err = os.WriteFile(dest, sourceText, 0644)
		if err != nil {
			log.Println("Error of filling:", dest, ", error:", err.Error())
		}
	}

	return nil
}

func CreateMapFromDestAndSource(dest, source []string) map[string]string {
	if len(dest) != len(source) {
		return nil
	}

	res := make(map[string]string, 0)
	for i := range len(dest) {
		res[dest[i]] = source[i]
	}

	return res
}
