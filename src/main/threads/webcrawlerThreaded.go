package main

import (
	"fmt"
	"sync"
)

func threadedFetcher(input string, fakeFetcher FakeFetcher) {
	fakeFetcher.mu.Lock()
	_, hasKey := fakeFetcher.taken[input]
	fakeFetcher.taken[input] = true
	fakeFetcher.mu.Unlock()

	if hasKey {
		return
	}

	fakeResultRes, err := fetcher.Fetch(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wgInternal sync.WaitGroup
	fmt.Printf("found: %s %q\n", input, fakeResultRes)

	for _, u := range fakeResultRes {
		wgInternal.Add(1)
		go func(u2 string) {
			defer wgInternal.Done()
			threadedFetcher(u2, fakeFetcher)
		}(u)
	}
	wgInternal.Wait()
	return
}

func worker(input string, fakeFetcher FakeFetcher, channel chan []string) {
	fakeFetcher.mu.Lock()
	_, hasKey := fakeFetcher.taken[input]
	fakeFetcher.taken[input] = true
	fakeFetcher.mu.Unlock()
	if hasKey {
		channel <- []string{}
		return
	}
	fakeResultRes, err := fetcher.Fetch(input)
	if err != nil {
		fmt.Println(err)
		channel <- []string{}
		return
	}
	fmt.Printf("found: %s %q\n", input, fakeResultRes)
	channel <- fakeResultRes
}

func coordinator(fakeFetcher FakeFetcher, channelForInput chan []string) {
	n := 1
	for val := range channelForInput {
		for _, u := range val {
			n++
			go func(u2 string) {
				worker(u2, fakeFetcher, channelForInput)
			}(u)
		}
		n--
		if n == 0 {
			break
		}
	}
}

func initCoor() {
	channelForRes := make(chan []string)
	go func() {
		channelForRes <- []string{"https://golang.org/"}
	}()
	coordinator(fetcher, channelForRes)
}

func initThreadedSol() {
	var wgInternal sync.WaitGroup
	wgInternal.Add(1)
	go func() {
		defer wgInternal.Done()
		threadedFetcher("https://golang.org/", fetcher)
	}()
}

func main() {
	initCoor()
}

// fakeFetcher is Fetcher that returns canned results.
type FakeFetcher struct {
	mu           sync.Mutex
	result       map[string]*fakeResult
	cachedResult map[string]*fakeResult
	taken        map[string]bool
}

type fakeResult struct {
	body string
	urls []string
}

func (f FakeFetcher) Fetch(url string) ([]string, error) {
	if cachedRes, okCached := f.cachedResult[url]; okCached {
		return cachedRes.urls, nil
	}
	if res, ok := f.result[url]; ok {
		f.cachedResult[url] = &fakeResult{res.body, res.urls}
		return res.urls, nil
	}

	return []string{}, fmt.Errorf("not found: %s", url)
}

var fetcher = FakeFetcher{
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
	taken:        make(map[string]bool),
}
