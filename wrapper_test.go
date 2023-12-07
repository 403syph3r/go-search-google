package searchwrapper_test

import (
	"testing"

	searchwrapper "github.com/403syph3r/go-search-google"
)

func TestMainSearch(t *testing.T) {
	// Default Aggression level, US country code, 2 search queries
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
			if l < c.minWant[i] {
				t.Errorf("Not enough Results collected. Got %v results. Wanted %v", l, c.minWant[i])
			}
		}
	}
}

func TestAggressiveSearch(t *testing.T) {
	// High (H) Aggression level, 10 queries
	cases := []struct {
		in      []string
		minWant []int
	}{
		{[]string{"query1", "query2", "query3", "query4", "query5", "query6", "query7", "query8", "query9", "query10"}, []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
	}
	options := searchwrapper.SearchParameters{
		CountryCode: "us",
		Aggression:  "H",
	}
	for _, c := range cases {
		got, err := searchwrapper.SearchMultiple(nil, c.in, options)
		if err != nil {
			t.Errorf("Error Encountered: %v", err)
		}
		lengths := searchwrapper.GetResultSetCounts(got)
		for i, l := range lengths {
			if l < c.minWant[i] {
				t.Errorf("Not enough Results collected. Got %v results. Wanted %v", l, c.minWant[i])
			}
		}
	}
}
