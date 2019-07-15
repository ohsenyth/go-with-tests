package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Anonymous functions have a number of features which make them useful, two of which we're using above.
	// Firstly, they can be executed at the same time that the're declared - this is what the () at the end of the anonymous function is doing.
	// Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point
	// when you declare the anonymous function are also available in the body of the function.
	// for _, url := range urls {
	// 	go func(u string) {
	// 		results[u] = wc(u)
	// 	}(url)
	// }

	// time.Sleep(2 * time.Second)

	// return results

	// using channels
	for _, url := range urls {
		go func(u string) {

			// Send statement
			// <- operator, taking a channel on the left and a value on the right
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {

		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
