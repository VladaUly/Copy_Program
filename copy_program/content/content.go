package content

// SourceCollection - структура, содержащая копию интерфейса
// и индекс для считывания информации со всех методов интерфейса
type SourceCollection struct {
	register         []ContentSource
	currentitemindex int
}

// ContentSource - интерфейс,
// реализующий считывания данных из заданного источника
type ContentSource interface {
	Name() string
	Extension() string
	ReadContent() []byte
}

// AppendFile & AppendURL - функции, возвращающие ссылку на
// новые экземпляры структур FileContent & WebContent
func (s *SourceCollection) AppendFile(sourceFilPath string) {
	s.register = append(s.register, &FileContent{filePath: sourceFilPath, extension: ".txt"})
}
func (s *SourceCollection) AppendURL(sourceFilePath string) {
	s.register = append(s.register, &FileContent{filePath: sourceFilePath, extension: ".html"})
}

// Next возвращает интерфейс ContentSource из структуры SourceCollection
// для дальнейшей инициализации методов интерфейса в пакете crawler.go
func (s *SourceCollection) Next() ContentSource {
	var elem ContentSource
	for s.currentitemindex = 0; s.currentitemindex < len(s.register); s.currentitemindex++ {
		elem = s.register[s.currentitemindex]

	}
	return elem
}
