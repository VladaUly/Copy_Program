package content

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// WebContent - структура, содержащая ссылку на веб-ресурс
// для считывания данных и расширение web-файла
type WebContent struct {
	url       string
	extension string
}

// структра WEBCONTENT реализует интерфейс CONTENTSOURCE
func (f *WebContent) Name() string {
	u, err := url.Parse(f.url)
	if err != nil {
		log.Fatal(err)
	}
	return u.Hostname()
}
func (f *WebContent) Extension() string {
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
