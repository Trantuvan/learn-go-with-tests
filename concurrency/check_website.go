package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChan <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChan
		results[result.string] = result.bool
	}

	return results
}
