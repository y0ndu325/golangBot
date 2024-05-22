package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

// func Parse() {
// 	geziyor.NewGeziyor(&geziyor.Options{
// 		StartURLs: []string{"https://www.championat.com/hockey/_nhl/tournament/5918/calendar/"},
// 		ParseFunc: FetchDataParse,
// 		Exporters: []export.Exporter{&export.JSON{}},
// 	}).Start()
// }

func parseWeb(g *geziyor.Geziyor, r *client.Response, results *[]string) {
	r.HTMLDoc.Find("tr.stat-results__row").Each(func(i int, s *goquery.Selection) {

		title := s.Find("span.table-item__name").Text()
		result := s.Find("span.stat-results__count-main").Text()

		output := title + " " + result
		g.Exports <- map[string]interface{}{
			"output": output,
		}
		*results = append(*results, output)
	})

	if href, ok := r.HTMLDoc.Find("a.next-page").Attr("href"); ok {
		g.Get(r.JoinURL(href), func(g *geziyor.Geziyor, r *client.Response) {
			parseWeb(g,r, results)
		})
	}
}

func FetchDataParse(url string) []string {
	var results []string

	g := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			parseWeb(g,r, &results)
		},
		Exporters: []export.Exporter{&export.JSON{}},
	})
	g.Start()
	return results
}

