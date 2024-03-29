package ga_test

import (
	"fmt"
	"os"

	"../ga"
)

func getNewAccountTokenFromEnv() ga.AccountToken {
	clientID := os.Getenv("GA_CLIENT_ID")
	clientSecret := os.Getenv("GA_CLIENT_SECRET")
	refreshToken := os.Getenv("GA_REFRESH_TOKEN")
	return ga.NewAccountToken(clientID, clientSecret, refreshToken)
}

func ExampleGetEvent() {
	res, _ := ga.GetEvent(getNewAccountTokenFromEnv(), os.Getenv("GA_ID"), "^.*Tag/.*", 30, true)
	fmt.Println(res)
	// Output:
	// map[iPhone:10 Android:17 Windows:12]
}

func ExampleGet() {
	metrics := "ga:uniquePageviews"
	dimensions := "ga:pagePath"
	sortOrder := "-ga:uniquePageviews"
	filters := "ga:pagePath=~^/path/to/site/.*"
	res, _ := ga.Get(getNewAccountTokenFromEnv(), os.Getenv("GA_ID"), metrics, dimensions, filters, sortOrder, 1)
	fmt.Println(res)
	// Output:
	// map[/path/to/site/index.html:100 /path/to/site/about.html:10 /path/to/site/other.html:12]
}
