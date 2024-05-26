package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

func search(searchTerm string, chromeDpContext context.Context, maxPages int) []string {
	/*
		- get the search URL with URL encoded q param
		- navigate to https://www.google.com/search?q={searchTerm}
		- expand the search results by clicking "More Results"
		- get the URLs from the <cite> tags in the search results
	*/

	url := getURL(searchTerm)

	chromeDpCtx := navigateToGoogle(chromeDpContext, url)
	chromeDpCtx = expandSearchResults(chromeDpCtx, maxPages)

	searchResultUrls := getSearchResultUrls(chromeDpCtx)

	return searchResultUrls
}

func getURL(searchTerm string) string {
	/*
		- url encode the search term
		- return the google search url
	*/

	encodedSearchTerm := url.QueryEscape(searchTerm)

	return fmt.Sprintf("https://www.google.com/search?q=%s", encodedSearchTerm)
}

func navigateToGoogle(chromeDpCtx context.Context, url string) context.Context {
	/*
		- navigate to https://www.google.com/search?q={searchTerm}
			- wait for search results to appear
				- they may not initially appear since a Captcha will need to be solved
				- this should hold the window open for the user to solve the Captcha
			- wait two seconds to ensure the "More Results" button is loaded
		- possible alternative approach would be to download the CAPTCHA audio and solve?
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`a[jsname="UWckNb"]`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	return chromeDpCtx
}

func expandSearchResults(chromeDpCtx context.Context, maxPages int) context.Context {
	/*
		- get the CSS selector for the "More Results" anchor tag
		- check the page to see if "More Results" exists
			- if does not exist, we have reach the final result page
			- else, keeping clicking until maxPages is reach or all results are found
				- can too many clicks kill the browser?
	*/

	pageCount := maxPages

	var moreResults bool
	for {
		chromeDpCtx, moreResults = checkForMoreResults(chromeDpCtx, moreResults)

		if (!moreResults) || (pageCount == 0) {
			break
		}

		chromeDpCtx = getMoreResults(chromeDpCtx)

		pageCount -= 1
	}

	return chromeDpCtx
}

func getMoreResults(chromeDpCtx context.Context) context.Context {
	/*
		- if the "More Results" anchor tag exists, click it to get more results
		- wait .5 seconds for new results to load
			- can this time be reduced?
	*/

	moreResultsSelector := getMoreResultsSelector()
	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Click(moreResultsSelector),
	); err != nil {
		log.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	return chromeDpCtx
}

func getSearchResultUrls(chromeDpCtx context.Context) []string {
	/*
		- get the <cite> elements' InnerText (these contain URLs)
		- iterate through the InnerTexts
			- if InnerText contains a domain, transform to URL format
				- not all InnerTexts will contain a domain
	*/

	resultLinks := getResultLinks(chromeDpCtx)

	var searchResultsURLs []string
	searchResultsURLs = append(searchResultsURLs, resultLinks...)

	return searchResultsURLs
}
