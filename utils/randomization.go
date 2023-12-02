package utils

import (
	"math/rand"
	"time"
)

func GetRandomizedUserAgent(excludeFromList string) string {
	// Sample User-agents for chome/Firefox/Edge
	sampleUserAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/119.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 14.1; rv:109.0) Gecko/20100101 Firefox/119.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.2151.72",
	}
	// Pull a random user-agent. If if happened to choose the provided excludeFromList, try again.
	var rand_num int
	for {
		rand_num = rand.Intn(len(sampleUserAgents))
		if sampleUserAgents[rand_num] != excludeFromList {
			return sampleUserAgents[rand_num]
		}
	}
}

func WaitForTime(waitSeconds int) {
	time.Sleep(time.Second * time.Duration(waitSeconds))
}

func GetRandomPaginationTime(aggression_level string) int {
	//Time before going to the "next" page in Google Search results
	var min_s, max_s int
	switch aggression_level {
	case "H":
		min_s = 2
		max_s = 5
	case "M":
		min_s = 6
		max_s = 15
	case "L":
		min_s = 10
		max_s = 30
	default:
		min_s = 5
		max_s = 20
	}
	return rand.Intn(max_s-min_s) + min_s
}

func GetRandomNewSearchTime(aggression_level string) int {
	//Time between search queries
	var min_s, max_s int
	switch aggression_level {
	case "H":
		min_s = 5
		max_s = 15
	case "M":
		min_s = 30
		max_s = 60
	case "L":
		min_s = 60
		max_s = 120
	default:
		min_s = 25
		max_s = 70
	}
	return rand.Intn(max_s-min_s) + min_s
}

func WaitRandomNewSearchTime(aggression_level string) {
	rand_num := GetRandomNewSearchTime(aggression_level)
	WaitForTime(rand_num)
}
