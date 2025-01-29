package dirsfiller

import (
	"log"
	"os"
)

type DirFiller struct {
}

// fileSources это мапа.
// Ключ это название места куда нужно скопировать.
// Значение это место где лежит то что нужно копировать
func (df *DirFiller) Fill(fileSources map[string]string) error {
	for dest, source := range fileSources {
		file, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("Error opening file %s: %v", dest, err)
			continue
		}

		defer file.Close()

		if _, err := file.WriteString(source); err != nil {
			log.Printf("Error writing to file %s: %v", dest, err)
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
