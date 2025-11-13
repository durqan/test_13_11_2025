package services

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

func CheckLinks(links []string) map[string]bool {
	client := &http.Client{Timeout: 10 * time.Second}
	results := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, url := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			fullURL := ensureProtocol(url)
			resp, err := client.Get(fullURL)
			if err != nil {
				mu.Lock()
				results[url] = false
				mu.Unlock()
				return
			}
			resp.Body.Close()

			mu.Lock()
			results[url] = resp.StatusCode == 200
			mu.Unlock()
		}(url)
	}

	wg.Wait()
	return results
}

func ensureProtocol(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}
