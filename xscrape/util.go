package xscrape

import (
	"net/url"

	"golang.org/x/net/html"
)

func attr(t *html.Node, name string) (bool, string) {
	for _, a := range t.Attr {
		if a.Key == name {
			return true, a.Val
		}
	}
	return false, ""
}

func httpStatusIsError(status int) bool {
	return status == 0 || status >= 400
}

func appendPageIfNotPresent(pages []string, pageURL string) []string {
	for _, existingPageURL := range pages {
		if existingPageURL == pageURL {
			return pages
		}
	}
	return append(pages, pageURL)
}

func resolveURL(val string, relativeTo *url.URL) (*url.URL, error) {
	parsedVal, err := url.Parse(val)
	if err != nil {
		return nil, err
	}

	parsedVal = relativeTo.ResolveReference(parsedVal)
	return parsedVal, nil
}
