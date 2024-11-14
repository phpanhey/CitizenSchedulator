package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	url := "https://www.service.bremen.de/dienstleistungen/reisepass-beantragen-fuer-personen-unter-18-jahren-126498"
	urlContent := getUrlContent(url)

	doc := soup.HTMLParse(strings.Split(urlContent, "Frühestmöglicher Termin in Bremen:")[1])
	links := doc.FindAll("a")

	m := make(map[string]string)
	m["date"] = links[0].Text()
	m["link"] = links[0].Attrs()["href"]

	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(string(res))

}

func getUrlContent(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
