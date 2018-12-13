package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Нужная часть api radio-t.com (см. https://radio-t.com/api-docs/)
type Entry struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// rt представляет отображение номеров подкастов (и заголовков) на url
type rt struct {
	cache    map[string]string // забитые руками отклонения (см. newRt)
	podcasts map[int]string    // маппинг номеров подкастов на url выпусков
	prep     map[int]string    // маппинг номеров подкастов на url тем для выпусков
}

// podcast - путь до json файла с url выпусков подкаста
// prep - путь до json файла с url тем для выпусков подкаста
// (см. https://radio-t.com/api-docs/)
func newRt(podcast, prep string) rt {
	// Методом научного тыка были найдены несколько записей на новом сайте, шаблон названия которых отличается
	// от стандартного ("Радио-Т...%d" или "Темы для...%d"), их, опять же, оказалось проще вбить один раз,
	// чем сочинять хитрый велосипед для нахождения.
	r := rt{
		cache: map[string]string{
			"Радио-Т 150 (нам 3 года)": "https://radio-t.com/p/2009/08/18/prep-150/",
			"Запись и трансляция #43":  "https://radio-t.com/p/2007/07/08/podcast-43/",
			"Запись и трансляция #42":  "https://radio-t.com/p/2007/07/01/podcast-42/",
			"Запись и трансляция #40":  "https://radio-t.com/p/2007/06/17/podcast-40/",
			"Темы для РТ#114":          "https://radio-t.com/p/2008/11/25/prep-114/",
			"Темы для РТ#119":          "https://radio-t.com/p/2009/01/01/prep-119/",
		},
		podcasts: jsonToMap(podcast),
		prep:     jsonToMap(prep),
	}
	return r
}

func parseJsonFile(fpath string) []Entry {
	var data []Entry
	b := read(fpath)
	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// На вход передается имя json файла, в котором хранится одна категория ссылок (Темы для подкастов или сами подкасты).
// Из заголовка (поля title) извлекается номер подаста, который является ключем в хешмапе, и url - значением.
func jsonToMap(fpath string) map[int]string {
	data := parseJsonFile(fpath)
	res := map[int]string{}
	for i := range data {
		title := strings.TrimSpace(data[i].Title)
		num, err := toNumber(title)
		if err != nil {
			// Если не удалось извлечь из заголовка (title) записи в блоге номера подкаста,
			// то пропускаем такую запись, выведя ее в stdout
			log.Println(err)
			continue
		}
		url := data[i].URL
		res[num] = url
	}
	return res
}

func read(p string) []byte {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// toNumber пытается из заголовка блога (title) извлечь номер подкаста. Некоторые значения забиты вручную,
// так проще, чем парсить.
func toNumber(title string) (int, error) {
	if title == "" {
		return 0, errors.New(title)
	}
	if title == "Темы для РТ114" {
		return 114, nil
	}
	if title == "Темы для РТ119" {
		return 119, nil
	}
	f := strings.Fields(title)
	last := f[len(f)-1]
	if strings.HasPrefix(last, "#") {
		last = last[1:]
	}
	i, err := strconv.Atoi(last)
	if err != nil {
		msg := fmt.Sprintf("Post is not podcast or themes: Title - [%s]", title)
		return 0, errors.New(msg)
	}
	return i, nil
}

// url возвращется новый url записи блога, основываясь на содержании его (записи блога) заголовка.
func (r *rt) url(title string) (string, error) {
	if url, ok := r.cache[title]; ok {
		return url, nil
	}
	num, err := toNumber(title)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(title, "Радио") {
		url := r.podcasts[num]
		return url, nil
	}
	if strings.HasPrefix(title, "Темы") || // Первая буква Т - кириллица
		strings.HasPrefix(title, "Tемы") { // Первая буква T - латиница !!!
		url := r.prep[num]
		return url, nil
	}
	msg := fmt.Sprintf("Post was not found: Title - [%s]", title)
	return "", errors.New(msg)
}
