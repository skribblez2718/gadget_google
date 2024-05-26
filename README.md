# Gadget Google
gadget_google provides a mostly automated way to get URLs from Google search results. These URLs can then be used to enumerate subdomains, endpoints, etc. and be fed into other tools for additional scope discovery. This mostly automated because Google may or may not pick this tool up as a bot. In this case the user will need to solvce the CAPTCHA before the remaining automation can take over.

##  Install
```sh
git clone https://github.com/skribblez2718/gadget_google.git
cd gadget_google
go build
```

## Usage
```sh
gadget_google --searchTerm "<search_term>" --output "<output_file>" [--maxPages <int>]
```

**Note**: By default this tool will attempt to fetch all results. The upper limit of search results possible has not been tested. Currently Google dynamically displays more results as you scroll or hit the "More Results" button. There is very likely an upper limit in content the browser will handle before crashing. This can easily be an issue if a large number of results are returned from the search. Based on my understanding of the Custom Search API this method only allows for 1K results max. This tool should be able to easily exceed that, but heed this note and browse responsibly my friends. 

## TO DO
- Research CAPTCHA solving to hopefully make the tool reliably full automated
- Create similar tools for other engines such as Bing, Duck Duck Go and others
    - Likely combine then into a single tool
