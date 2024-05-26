package main

func main() {
	/*
		- get arguments from command line
		- create a chrome devtools protocol context
		- search google and return the results
		- write the results to output file
	*/

	searchTerm, output, maxPages := getArgs()

	emptyCtx, cancelEmptyCtx := getEmptyContext()
	defer cancelEmptyCtx()

	allocCtx, cancelAllocCtx := getAllocContext(emptyCtx)
	defer cancelAllocCtx()

	chromeDpCtx, cancelChromeDpCtx := getChromeDpContext(allocCtx)
	defer cancelChromeDpCtx()

	searchResultsURLs := search(searchTerm, chromeDpCtx, maxPages)

	writeResults(searchResultsURLs, output)
}
