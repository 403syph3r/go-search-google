package searchwrapper_test

import (
	searchwrapper "go-search-google"
	"testing"
)

func TestMainSearch(t *testing.T) {
	cases := []struct {
		in      []string
		minWant []int
	}{
		{[]string{"testing", "site:espn.com"}, []int{5, 1}},
	}
	options := searchwrapper.SearchParameters{
		CountryCode: "us",
		Aggression:  "",
	}
	for _, c := range cases {
		got, err := searchwrapper.SearchMultiple(nil, c.in, options)
		if err != nil {
			t.Errorf("Error Encountered: %v", err)
		}
		lengths := searchwrapper.GetResultSetCounts(got)
		for i, l := range lengths {
			//fmt.Printf("Got %v results. Wanted %v", l, c.minWant)
			if l < c.minWant[i] {
				t.Errorf("Not enough Results collected. Got %v results. Wanted %v", l, c.minWant[i])
			}
		}
		// searchwrapper.PrintResultSetPreview(got)
	}
}

// func TestMinimalSearch(t *testing.T) {
// 	opts := googlesearch.SearchOptions{
// 		Limit: 20,
// 	}
// 	returnLinks, err := googlesearch.Search(nil, "Hello World", opts)
// 	if err != nil {
// 		t.Errorf("something went wrong: %v", err)
// 		return
// 	}
// 	if len(returnLinks) == 0 {
// 		t.Errorf("no results returned: %v", returnLinks)
// 	}
// }

// func TestLargeSearch(t *testing.T) {
// 	cases := []struct {
// 		in      []string
// 		minWant []int
// 	}{
// 		{[]string{"filetype:pdf", "Testing", "Just Testing"}, []int{5, 5, 5}},
// 	}
// 	options := searchwrapper.SearchParameters{
// 		CountryCode: "us",
// 	}
// 	for _, c := range cases {
// 		got, err := searchwrapper.SearchMultiple(nil, c.in, options)
// 		if err != nil {
// 			t.Errorf("Error Encountered: %v", err)
// 		}
// 		lengths := searchwrapper.GetResultSetCounts(got)
// 		for i, l := range lengths {
// 			//fmt.Printf("Got %v results. Wanted %v", l, c.minWant)
// 			if l < c.minWant[i] {
// 				t.Errorf("Not enough Results collected. Got %v results. Wanted %v", l, c.minWant[i])
// 			}
// 		}
// 		// searchwrapper.PrintResultSetPreview(got)
// 	}
// }

// func TestAggressiveSearch(t *testing.T) {
// 	cases := []struct {
// 		in      []string
// 		minWant []int
// 	}{
// 		{[]string{"query1", "query2", "query3", "query4", "query5", "query6", "query7", "query8", "query9", "query10"}, []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
// 	}
// 	options := searchwrapper.SearchParameters{
// 		CountryCode: "us",
// 		Aggression:  "H",
// 	}
// 	for _, c := range cases {
// 		got, err := searchwrapper.SearchMultiple(nil, c.in, options)
// 		if err != nil {
// 			t.Errorf("Error Encountered: %v", err)
// 		}
// 		lengths := searchwrapper.GetResultSetCounts(got)
// 		for i, l := range lengths {
// 			//fmt.Printf("Got %v results. Wanted %v", l, c.minWant)
// 			if l < c.minWant[i] {
// 				t.Errorf("Not enough Results collected. Got %v results. Wanted %v", l, c.minWant[i])
// 			}
// 		}
// 		// searchwrapper.PrintResultSetPreview(got)
// 	}
// }
