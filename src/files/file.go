package files

import (
	"log"
	"os"
)

type File struct {
	filePath string
}

func (*File) CreateFile(filePath string) *os.File {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("create file: ", err)
	}

	return f
}

func NewFile() *File {
	return new(File)
}
