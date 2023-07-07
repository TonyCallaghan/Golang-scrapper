package core

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Plant struct {
	Name string
	Url  string
}

func GetPlants(sectionUrl string) ([]Plant, error) {
	var plants []Plant

	doc, err := goquery.NewDocument(sectionUrl)
	if err != nil {
		return plants, err
	}

	doc.Find(".plantbox").Each(func(i int, s *goquery.Selection) {
		plant := Plant{}

		plant.Name = s.Find(".pltitle").First().Text()
		plant.Url, _ = s.Find(".dbseemore").First().Attr("href")

		plants = append(plants, plant)
	})

	return plants, nil
}

func Execute() {
	doc, err := goquery.NewDocument("https://www.thetortoisetable.org.uk")
	if err != nil {
		panic(err)
	}

	doc.Find(".homepagebox").Each(func(i int, s *goquery.Selection) {
		sectionUrl, _ := s.Find(".boxpic a").First().Attr("href")
		sectionName := strings.ToUpper(s.Find(".boxtitle").First().Text())

		fmt.Println("Section:", sectionName)

		plants, err := GetPlants(sectionUrl)
		if err != nil {
			panic(err)
		}

		for _, plant := range plants {
			fmt.Println("Plant:", plant.Name)
			fmt.Println("URL:", plant.Url)
			fmt.Println()
		}
	})
}
