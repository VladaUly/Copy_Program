package main

import (
	"bufio"
	"copy_program/content"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Дефолтные значения для флагов.
const (
	defaultSourceFilePathFlag     = "src.txt"
	defaultDestinationDirPathFlag = "C:\\dir_to_copy"
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
	fmt.Println(sourceFilePath, destinationDirPath)
	// Проверка флага на пустую строку.
	if sourceFilePath == "" && destinationDirPath == "" {
		// Если проверка не пройдена, то завершение программы. os.Exit(1).
		fmt.Println("Было введено недопустимое значение.")
		os.Exit(1)
	}
	// ReadSourceFile - функция открытия заданного файла и построчного считывания данных из него
	readSourceFile(sourceFilePath)
	// runCopying - функция копирования данных в заданную директорию
	//runCopying(source, destinationDirPath)
}

func readSourceFile(sourceFilePath string) *content.SourceCollection {
	collection := &content.SourceCollection{}
	// Открыть файл, путь к которому задан в флаге sourceFilePath
	f, err := os.Open(sourceFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		// Прочитать каждую строчку в файле,
		bline, isPrefix, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		line := string(bline)
		if strings.HasPrefix(line, "https") {
			// и добавить каждую ссылку через AppendURL в collection,
			collection.AppendURL(line)
		} else {
			// а каждый путь к файлу через AppendFile.
			collection.AppendFile(line)
		}
		if !isPrefix {
			break
		}
	}
	fmt.Printf("collection = %v", collection)
	return collection
}

// func runCopying(source *content.SourceCollection, dirPath string) {
// 	crawler.Crawler().SetPath(dirPath)
// 	crawler.Crawler().SetCollection(source)
// 	crawler.Crawler().CopyAll()
// }
