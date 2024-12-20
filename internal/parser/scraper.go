package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/rostis232/parcelstrackingservice/models"
	"strings"
)

func ScrapeData(html string) (map[string]*models.Data, error) {
	result := make(map[string]*models.Data)

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return result, fmt.Errorf("error parsing response html: %w", err)
	}

	doc.Find(".scroll-floor").Each(func(i int, s *goquery.Selection) {
		data := models.Data{}
		number := strings.TrimSpace(s.Find("div.order-info h1").Text())
		number = strings.ReplaceAll(number, "\t", "")
		result[number] = &data

		s.Find(".country-col").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				country := strings.TrimSpace(s.Find(".country-row1").Text())
				country = strings.ReplaceAll(country, "\t", "")
				data.OriginCountry = country
			case 1:
				country := strings.TrimSpace(s.Find(".country-row1").Text())
				country = strings.ReplaceAll(country, "\t", "")
				data.DestinationCountry = country
			}
		})

		s.Find(".track-detail ul li").Each(func(i int, s *goquery.Selection) {
			date := strings.TrimSpace(s.Find(".date").Text())
			date = strings.ReplaceAll(date, "\t", "")
			time := strings.TrimSpace(s.Find(".time").Text())
			time = strings.ReplaceAll(time, "\t", "")
			text := strings.TrimSpace(s.Find(".text").Text())
			text = strings.ReplaceAll(text, "\t", "")
			text = strings.ReplaceAll(text, "\n", "")
			data.Checkpoints = append(data.Checkpoints, models.Checkpoint{
				Date:   date + " " + time + ":00",
				Status: text,
			})
		})
	})

	return result, err
}
