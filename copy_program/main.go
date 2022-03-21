package main

import (
	"bufio"
	"copy_program/content"
	"copy_program/crawler"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Дефолтные значения для флагов.
const (
	defaultSourceFilePathFlag     = "src.txt"
	defaultDestinationDirPathFlag = "C:\\Vlada\\src\\copy_program\\files_to_copy\\dir_to_copy"
)

func main() {
	// sourceFilePathFlag - путь к файлу с именами файлов и адресами странниц, котрые необходимо скопировать.
	sourceFilePathFlag := flag.String("source", defaultSourceFilePathFlag, "path of the file")
	//destinationDirPathFlag - путь к директории, в которую осуществляется копирование.
	destinationDirPathFlag := flag.String("to", defaultDestinationDirPathFlag, "path of the directory")
	// Парсинг флагов.
	flag.Parse()
	sourceFilePath := *sourceFilePathFlag
	destinationDirPath := *destinationDirPathFlag
	// Проверка флагов на ошибки ввода.
	fileInfo, err := os.Stat(destinationDirPath)
	if err != nil {
		err = os.MkdirAll(destinationDirPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if !fileInfo.IsDir() {
			fmt.Println("Не является директорией. Введите другое значение...")
			os.Exit(1)
		}
	}
	var strs = strings.Split(sourceFilePath, ".")
	strslen := strs[len(strs)-1]
	if strslen != "txt" {
		fmt.Println("Файл не имеет расширения txt. Попробуйте другой файл...")
		os.Exit(1)
	}
	if sourceFilePath == "" && destinationDirPath == "" {
		// Если проверка не пройдена, то завершение программы. os.Exit(1).
		fmt.Println("Было введено пустое значение. Попробуйте еще раз...")
		os.Exit(1)
	}

	// ReadSourceFile - функция открытия заданного файла и построчного считывания данных из него
	source := readSourceFile(sourceFilePath)
	// runCopying - функция копирования данных в заданную директорию
	runCopying(source, destinationDirPath)
}

// ReadSourceFile - функция открытия заданного файла и построчного считывания данных из него
func readSourceFile(sourceFilePath string) *content.SourceCollection {
	collection := &content.SourceCollection{}
	// Открыть файл, путь к которому задан в флаге sourceFilePath
	f, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Данного файла не существует.Попробуйте другой...")
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "https") {
			// и добавить каждую ссылку через AppendURL в collection,
			collection.AppendURL(scanner.Text())
		} else {
			// а каждый путь к файлу через AppendFile.
			collection.AppendFile(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return collection
}

// runCopying - функция копирования данных в заданную директорию
func runCopying(source *content.SourceCollection, dirPath string) {
	crawler.Crawler().SetPath(dirPath)
	crawler.Crawler().SetCollection(source)
	crawler.Crawler().CopyAll()
}
