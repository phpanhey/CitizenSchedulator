package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: myprogram <URL>")
		return
	}
	url := os.Args[1]
	urlContent := getUrlContent(url)

	parts := strings.Split(urlContent, "Frühestmöglicher Termin in Bremen:")
	if len(parts) > 1 {
		doc := soup.HTMLParse(strings.Split(urlContent, "Frühestmöglicher Termin in Bremen:")[1])
		links := doc.FindAll("a")

		nextAvailableAppointmentDate := CleanDate(links[0].Text())
		fmt.Println("Next available appointment date: ", nextAvailableAppointmentDate)
		fmt.Println("Link to appointment: ", url)
	}
}

func CleanDate(date string) string {
	// also get rid of "um" in string
	date = strings.Replace(date, "um", "", -1)
	return strings.TrimSpace(date[4:])
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
