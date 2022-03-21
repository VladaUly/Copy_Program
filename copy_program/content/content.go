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
func (s *SourceCollection) AppendFile(sourceFilePath string) {
	s.register = append(s.register, &FileContent{filePath: sourceFilePath})
}
func (s *SourceCollection) AppendURL(sourceFilePath string) {
	s.register = append(s.register, &WebContent{url: sourceFilePath, extension: "html"})
}

// Next функция, которая обходит все методы
// интерфейса ContentSource и возвращает его
func (s *SourceCollection) Next() (ContentSource, bool) {
	var source ContentSource
	if len(s.register) == s.currentitemindex {
		return source, false
	} else {
		sources := s.register[s.currentitemindex]
		s.currentitemindex++
		return sources, true
	}
}
