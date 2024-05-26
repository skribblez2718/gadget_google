package main

import (
	"flag"
	"fmt"
	"os"
)

func getArgs() (string, string, int) {
	flag.Usage = displayUsage

	var (
		searchTerm string
		output     string
		maxPages   int
	)
	flag.StringVar(&searchTerm, "searchTerm", "", "Google search term")
	flag.StringVar(&output, "output", "", "Path to file to write results to")
	flag.IntVar(&maxPages, "maxPages", -1, "Maximum number of result pages to search")

	flag.Parse()

	if searchTerm == "" || output == "" {
		flag.Usage()
		os.Exit(1)
	}

	return searchTerm, output, maxPages
}

func displayUsage() {
	usageMessage := "Description:\n"
	usageMessage += "\tUses ChromeDP to mostly automate extracting URLs from Google search results.\n"
	usageMessage += "Usage:\n"
	usageMessage += "\tgadget_google --searchTerm \"<search_term>\" --output \"<output_file>\" [--maxPages <int>]\n"

	fmt.Print(usageMessage)
}
