package content

import (
	"io/ioutil"
	"log"
	"strings"
)

//File Content - структура, содержащая путь файла
// для считывания данных  расширение самого файла
type FileContent struct {
	filePath  string
	extension string
}

// структура FILECONTENT реализует инерфейс CONTENTSOURCE
func (f *FileContent) Name() string {
	var strs = strings.Split(f.filePath, "\\")
	strslen := strs[len(strs)-1]
	name := strings.Split(strslen, ".")
	return name[0]
}
func (f *FileContent) Extension() string {
	var strs = strings.Split(f.filePath, ".")
	strslen := strs[len(strs)-1]
	f.extension = strslen
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
