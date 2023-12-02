# go-search-google

## About
`go-search-google` is an open-source repository for performing multiple Google searches programmatically. This library includes wait times between pagination/searches and checks to ensure that Google is not blocking the search. If Google does block the communication, an increasingly-long timeout period is utilized to prevent further blocks. This means this search method can take a long time. For quicker and more reliable searching, use Google's search API.

**Use library at your own risk. Using this library inappropriately could cause unexpected blocks from Google**

## Usage
This repository is to be used as a package, not as a standalone program.

## Example Usage
```
package main

import (
    searchwrapper "github.com/403syph3r/go-search-google"
    "fmt"
)

func main() {
    queries := []string{"Query 1","Query 2"}
    options := searchwrapper.SearchParameters{}
    rec, err := searchwrapper.SearchMultiple(nil, queries, options)
    if err != nil {
        fmt.Printf("Error Encountered: %v\n", err)
    } else {
        searchwrapper.PrintResultSetPreview(rec)
    }
}
```