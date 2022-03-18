package content

import (
	"io/ioutil"
	"log"
	"strings"
)

// File Content - структура, содержащая путь файла
// для считывания данных и расширение самого файла
type FileContent struct {
	filePath  string
	extension string
}

// структура FILECONTENT реализует интерфейс CONTENTSOURCE
func (f *FileContent) Name() string {
	return f.filePath
}
func (f *FileContent) Extension() string {
	ext := strings.HasPrefix(f.filePath, f.extension)
	if ext {
		log.Fatalln(ext)
	}
	return f.extension
}
func (f *FileContent) ReadContent() []byte {
	// чтение заданного файла
	body, err := ioutil.ReadFile(f.filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return body
}

////////////////////////////////////////////////////////
