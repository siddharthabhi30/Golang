package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (fakeResultParam fakeResult, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}
	fakeResultRes, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wgInternal sync.WaitGroup
	fmt.Printf("found: %s %q\n", url, fakeResultRes)
	for _, u := range fakeResultRes.urls {
		wgInternal.Add(1)
		go Crawl(u, depth-1, fetcher, &wgInternal)
	}
	wgInternal.Wait()
	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	mu           sync.Mutex
	result       map[string]*fakeResult
	cachedResult map[string]*fakeResult
}

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (fakeResult, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if cachedRes, okCached := f.cachedResult[url]; okCached {
		return fakeResult{cachedRes.body, cachedRes.urls}, nil
	}
	if res, ok := f.result[url]; ok {
		f.cachedResult[url] = &fakeResult{res.body, res.urls}
		return fakeResult{body: res.body, urls: res.urls}, nil
	}
	return fakeResult{}, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	result: map[string]*fakeResult{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	},
	cachedResult: make(map[string]*fakeResult),
}
