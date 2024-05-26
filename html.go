package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func checkForMoreResults(chromeDpCtx context.Context, moreResults bool) (context.Context, bool) {
	/*
		- check to see if the "More Results" button is on the page
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Evaluate(
			fmt.Sprintf(`document.querySelector("%s") !== null`, getMoreResultsSelector()),
			&moreResults,
		),
	); err != nil {
		log.Fatal(err)
	}

	return chromeDpCtx, moreResults
}

func getMoreResultsSelector() string {
	/*
		- this is currently the easiest way to find the "More Results" button
			- separated out for easy changes in the future if need be
	*/

	return "a[aria-label='More results']"
}

func getResultLinks(chromeDpCtx context.Context) []string {
	/*
		- find all anchor tags with jsname="UWckNb"
			- this appears to be the element that holds all links to urls
	*/
	var resultLinks []string
	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.EvaluateAsDevTools(
			fmt.Sprintf(`Array.from(document.querySelectorAll('%s')).map(el => el.href)`, getResultAnchorSelector()),
			&resultLinks,
		),
	); err != nil {
		log.Fatal(err)
	}

	return resultLinks
}

func getResultAnchorSelector() string {
	/*
		- this is currently the easiest way to find result links/urls
			- separated out for easy changes in the future if need be
	*/
	return "a[jsname=\"UWckNb\"]"
}
