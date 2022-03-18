package content

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// WebContent - структура, содержащая ссылку на веб-ресурс
// для считывания данных и расширение web-файла
type WebContent struct {
	url       string
	extension string
}

// структра WEBCONTENT реализует интерфейс CONTENTSOURCE
func (f *WebContent) Name() string {
	return f.url
}
func (f *WebContent) Extension() string {
	ext := strings.HasPrefix(f.url, f.extension)
	if ext {
		log.Fatalln(ext)
	}
	return f.extension
}
func (f *WebContent) ReadContent() []byte {
	// http запрос на чтение с определенного URL
	resp, err := http.Get(f.url)
	if err != nil {
		log.Fatalln(err)
	}
	// чтение данных
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

////////////////////////////////////////////////////////
