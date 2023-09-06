package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fullSite bool
)

func init() {
	flag.BoolVar(&fullSite, "f", false, "Скачать сайт полностью")

	flag.Parse()
}

func downloadFile(filepath, url string) error {
	var writingFile *os.File
	pattern := `\.\w{3}$`
	reg := regexp.MustCompile(pattern)
	founds := reg.FindStringSubmatch(filepath)

	// Если совпадения не найдены, то необходимо создать html документ для загрузки страницы
	if len(founds) == 0 {
		err := os.MkdirAll(filepath, os.ModePerm)
		if err != nil {
			return err
		}
		writingFile, err = os.Create(filepath + "/index.html")
		if err != nil {
			return err
		}
	} else {
		// Иначе создаем директорию для конечного файла
		dir := path.Dir(filepath)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
		writingFile, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}

	defer writingFile.Close()

	// Делаем запрос к переданному url
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response with status: %s", res.Status)
	}

	// Записываем в файл тело ответа
	_, err = io.Copy(writingFile, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// Загрузка полного сайта
func siteDownload(history map[string]struct{}, folederPath, URL string) error {
	fmt.Println("Link Processing:", URL)
	urlObject, err := url.Parse(URL)
	if err != nil {
		return err
	}
	// Скачивание конкретного файла
	err = downloadFile(folederPath+urlObject.Path, URL)
	if err != nil {
		return err
	}

	// Добавляем путь, чтобы не скачивать заново
	history[urlObject.Path] = struct{}{}

	// Читаем данные из html документа для просмотра наличия ссылок на файлы
	content, err := os.ReadFile(folederPath + urlObject.Path + "/index.html")
	if err != nil {
		return err
	}
	// Создаем регулярку для путей типа: /css...., https://hostname/....
	reg := regexp.MustCompile(`"(/css[a-z, /, .]+|` + urlObject.Scheme + "://" + urlObject.Host + `.+?)"`)
	slcOfLinks := reg.FindAllStringSubmatch(string(content), -1)

	// Проходимся по каждой найденной ссылке
	for _, link := range slcOfLinks {
		u, err := url.Parse(link[1])
		if err != nil {
			return err
		}
		// Если это ссылка типа /css, то необходимо добавить схему и хост
		if u.Scheme == "" {
			link[1] = urlObject.Scheme + "://" + urlObject.Host + link[1]
		}
		// Если ссылка не была обработана, то обрабатываем
		if _, ok := history[u.Path]; !ok {
			siteDownload(history, folederPath, link[1])
		}
	}
	return nil
}
func main() {
	urlPath := flag.Args()[0]                                         // Получаем ссылку на сайт
	downloadDir := "E:/goprojects/L2/develop/dev09/downloaded_files/" // Куда сохранять
	// Если сайт необходимо скачать полностью, то запускаем рекурсивную функцию, иначе скачиваем необходимый файл
	if fullSite {
		history := make(map[string]struct{})
		u, err := url.Parse(urlPath)
		if err != nil {
			log.Fatalf("Cannot parse URL: %s, error: %s", urlPath, err)
		}
		err = siteDownload(history, downloadDir+u.Host, urlPath)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	} else {
		u, err := url.Parse(urlPath)
		if err != nil {
			log.Fatalf("Cannot parse URL: %s, error: %s", urlPath, err)
		}

		err = downloadFile(downloadDir+u.Host, urlPath)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
	fmt.Println("Успешно!")
}
