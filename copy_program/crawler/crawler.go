package crawler

import (
	"copy_program/content"
)

var instance *crawler

type crawler struct {
	sources *content.SourceCollection
	dirPath string
}

// Crawler - единственная функция, через которую
// можно получить доступ к структуре crawler
func Crawler() *crawler {
	if instance == nil {
		instance = &crawler{}
	}
	return instance
}

// SetCollection - обновляет данные структуры crawler
func (c *crawler) SetCollection(sources *content.SourceCollection) {
	c.sources = sources
}

// SetPath - функция, задающая путь к директории,
// в которую будут копироваться данные
func (c *crawler) SetPath(sourceFilePath string) {
	c.dirPath = sourceFilePath
}

// // CopyAll - копирует данные в заданную директорию
// // и сохраняеет под определенным именем
// func (c *crawler) CopyAll() {
// 	var wg sync.WaitGroup
// 	for _, source := range c.sources {

// 		wg.Add(1)
// 		go func(src content.ContentSource) {
// 			contentSource := c.sources.Next()
// 			// инициализация методов интерфейса ContentSource
// 			name := src.Name()
// 			extention := src.Extension()
// 			content := src.ReadContent()
// 			// копирование данных в заданную директорию и
// 			// сохранение под заранее сформированным именем
// 			fileInDir := c.dirPath + name + extention
// 			err := os.WriteFile(fileInDir, content, 0666)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("Success")
// 			wg.Done()
// 		}(source)

// 	}
// 	wg.Wait()
// }
