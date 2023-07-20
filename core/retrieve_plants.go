package core

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	LatinNamePrefix  = "Latin Name:"
	FamilyNamePrefix = "Family Name:"
)

type Plant struct {
	Name        string
	Safety      string
	LatinName   string
	FamilyName  string
	Description string
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func GetPlantData(plantUrl string) (Plant, error) {
	var plant Plant

	res, err := httpClient.Get(plantUrl)
	if err != nil {
		return plant, err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return plant, errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return plant, err
	}

	plant.Name = doc.Find("h1").First().Text()

	safetyClasses := []string{".greensign", ".redsign", ".orangesign1", ".orangesign2"}
	for _, className := range safetyClasses {
		safetyElement := doc.Find(className)
		if safetyElement.Length() > 0 {
			plant.Safety = strings.TrimSpace(safetyElement.Text())
			break
		}
	}

	doc.Find("#plantinfoouter ul li").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		switch i {
		case 1:
			plant.LatinName = strings.TrimPrefix(text, LatinNamePrefix)
		case 2:
			plant.FamilyName = strings.TrimPrefix(text, FamilyNamePrefix)
		}
	})

	ulSelection := doc.Find("#plantinfoouter ul")
	plant.Description = CleanupDescription(strings.TrimSpace(ulSelection.Parent().Next().Text()))

	return plant, nil
}

func GetPlants(sectionUrl string) ([]Plant, error) {
	var plants []Plant

	res, err := httpClient.Get(sectionUrl)
	if err != nil {
		return plants, err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return plants, errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return plants, err
	}

	doc.Find(".plantbox").Each(func(i int, s *goquery.Selection) {
		plantUrl, _ := s.Find(".dbseemore").First().Attr("href")

		plantData, err := GetPlantData(plantUrl)
		if err != nil {
			return
		}

		plants = append(plants, plantData)
	})

	return plants, nil
}

func Execute() error {
	res, err := httpClient.Get("https://www.thetortoisetable.org.uk")
	if err != nil {
		return err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".homepagebox").Each(func(i int, s *goquery.Selection) {
		sectionUrl, _ := s.Find(".boxpic a").First().Attr("href")
		sectionName := strings.ToUpper(s.Find(".boxtitle").First().Text())

		fmt.Println("Section:", sectionName)

		plants, err := GetPlants(sectionUrl)
		if err != nil {
			fmt.Println("Error getting plants for section", sectionName, ":", err)
			return
		}

		for _, plant := range plants {
			fmt.Println("Plant:", plant.Name)
			fmt.Println("Safety:", plant.Safety)
			fmt.Println("Latin Name:", plant.LatinName)
			fmt.Println("Family Name:", plant.FamilyName)
			fmt.Println("Description:", plant.Description)
			fmt.Println()
		}
	})

	return nil
}

func CloseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		fmt.Println("Error closing the body:", err)
	}
}
