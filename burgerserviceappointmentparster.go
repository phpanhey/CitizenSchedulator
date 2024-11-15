package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func main() {
	url := "https://www.service.bremen.de/dienstleistungen/reisepass-beantragen-fuer-personen-unter-18-jahren-126498"
	urlContent := getUrlContent(url)
	// a very high nextAvailableAppointmentDate to make sure the parser gets the next nextAvailableAppointmentDate
	nextAvailableAppointmentDate := ParseTimeStamp("05.12.30 10:00")

	parts := strings.Split(urlContent, "Frühestmöglicher Termin in Bremen:")
	if len(parts) > 1 {
		doc := soup.HTMLParse(strings.Split(urlContent, "Frühestmöglicher Termin in Bremen:")[1])
		links := doc.FindAll("a")

		nextAvailableAppointmentDate = ParseTimeStamp(CleanDate(links[0].Text()))
	}

	if nextAvailableAppointmentDate.Before(TimeFromFile("date.txt")) {
		fmt.Println("Next available appointment: ", nextAvailableAppointmentDate)
		os.WriteFile("date.txt", []byte(nextAvailableAppointmentDate.Format("02.01.06 15:04")), 0644)
	}

}

func ParseTimeStamp(date string) time.Time {
	layout := "02.01.06 15:04"
	timeStamp, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return timeStamp
}

func TimeFromFile(s string) time.Time {
	content, err := os.ReadFile(s)
	if err != nil {
		os.WriteFile("date.txt", []byte("05.12.30 10:00"), 0644)
	}
	return ParseTimeStamp(string(content))
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
