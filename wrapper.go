package searchwrapper

import (
	"context"
	"fmt"

	googlesearch "github.com/403syph3r/go-search-google/google-search" // including locally because at the time of releasing this project (Dec 2023) there was an unfixed bug in google-search which has been resolved in this repo

	// googlesearch "github.com/rocketlaunchr/google-search"
	"time"

	"github.com/403syph3r/go-search-google/utils"

	"golang.org/x/time/rate"
)

type SearchParameters struct {
	//Aggression:
	//   - How quickly/aggressive the search wrapper should function.
	//   - Options are "H","M","L", but can be left blank. See randomization.go for more info
	Aggression string `json:"aggression"`

	//MaxResults:
	//   - How many search results should be returned per query
	//   - Passed through to google-search option
	//   - Default is 30
	MaxResults int `json:"max_results"`

	//CountryCode:
	//   - Passed through to google-search
	//   - Helps identify which Google domain to utilize
	//   - google-search sets "us" as default
	CountryCode string `json:"country_code"`

	//MaxRetries:
	//   - If blocked by Google, how many times tp keep trying before giving up
	//   - Default is 10 retries
	MaxRetries int `json:"max_retries"`

	//BlkStartWait:
	//   - When blocked by Google, the base number of seconds to wait before trying again
	//   - Default is 60 seconds
	BlkStartWait int `json:"block_start_wait_seconds"`

	//BlkInc:
	//   - When blocked by Google, the number of seconds added to the wait time
	//   - Default is 30 seconds
	BlkInc int `json:"block_increment_seconds"`
}

type ResultSet struct {
	//Query:
	//   - The querystring utilized
	Query string `json:"query"`

	//Results:
	//   - A slice of "Result" type from google-search package
	Results []googlesearch.Result `json:"results"`
}

func SearchMultiple(ctx context.Context, queries []string, searchOptions SearchParameters) ([]ResultSet, error) {
	if ctx == nil {
		//Set a timeout of 30 minutes if no context is provided
		new_ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute*30))
		defer cancel()
		ctx = new_ctx
	}
	lastUserAgent := ""
	searchResults := []ResultSet{}
	gotBlocked := false
	for i, query := range queries {
		if i > 0 {
			//Wait for a random amount of time before performing a new search
			utils.WaitRandomNewSearchTime(searchOptions.Aggression)
		}
		// fmt.Printf("Running Query: %v\n", query)
		options := GetRandomizedOption(lastUserAgent)
		if searchOptions.CountryCode != "" {
			options.CountryCode = searchOptions.CountryCode
		}
		if searchOptions.MaxResults != 0 {
			options.Limit = searchOptions.MaxResults
		} else {
			options.Limit = 30
		}
		lastUserAgent = options.UserAgent
		pagination_rate := utils.GetRandomPaginationTime(searchOptions.Aggression)
		googlesearch.RateLimit.SetLimit(rate.Every(time.Second * time.Duration(pagination_rate)))
		googlesearch.RateLimit.SetBurst(1)
		res, err := googlesearch.Search(ctx, query, options)
		j := 0
		waitTime := searchOptions.BlkStartWait
		if searchOptions.MaxRetries == 0 {
			//Default parameters would error out if Google is still blocking after 360 seconds (6 minutes)
			searchOptions.MaxRetries = 10
		}
		if searchOptions.BlkStartWait == 0 {
			searchOptions.BlkStartWait = 60
		}
		if searchOptions.BlkInc == 0 {
			searchOptions.BlkInc = 30
		}
		if err == googlesearch.ErrBlocked {
			gotBlocked = true
		}
		//If search was too aggressive, wait a given amount of increasing time until okay to proceed
		for gotBlocked && j < searchOptions.MaxRetries {
			waitTime = waitTime + j*searchOptions.BlkInc
			fmt.Printf("Got blocked by Google. Taking a %v second break", waitTime)
			utils.WaitForTime(waitTime)
			// If blocked, get a new randomized option package
			options := GetRandomizedOption(lastUserAgent)
			res, err = googlesearch.Search(ctx, query, options)
			if err == googlesearch.ErrBlocked {
				gotBlocked = true
			} else {
				gotBlocked = false
			}
			j++
		}
		if err != nil {
			fmt.Printf("Error encountered during search: %v\n", err.Error())
			return nil, err
		}
		searchResults = append(searchResults, ResultSet{Query: query, Results: res})
	}

	return searchResults, nil
}

func GetResultSetCounts(results []ResultSet) []int {
	//Used for testing
	result_lengths := []int{}
	for _, r := range results {
		result_lengths = append(result_lengths, len(r.Results))
	}
	return result_lengths
}

func PrintResultSetPreview(results []ResultSet) {
	//Can be utilized to quicklky preview the results
	for _, r := range results {
		fmt.Printf("Query: %v", r.Query)
		for _, s := range r.Results {
			fmt.Printf("- %v (%v)", s.Title, s.URL)
		}
	}
}

func GetRandomizedOption(excludeFromList string) googlesearch.SearchOptions {
	return googlesearch.SearchOptions{
		UserAgent: utils.GetRandomizedUserAgent(excludeFromList),
	}
}
