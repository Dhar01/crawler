package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int  // keep track of pages
	baseURL            *url.URL        // keep track of original base URL
	mu                 *sync.Mutex     // ensures pages map is thread safe
	concurrencyControl chan struct{}   // ensures control on the channel
	wg                 *sync.WaitGroup // ensures the main func waits
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
