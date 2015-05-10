# go-google-analytics
Google Analytics Helper for GoLang

## Install
```
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/analytics/v3
go get github.com/m-hosoi/go-google-analytics/ga
```

## Usage
```go
import (
    "fmt"
    "os"

    "github.com/m-hosoi/go-google-analytics/ga"
)

func getNewContextFromEnv() ga.Context {
    clientID := os.Getenv("GA_CLIENT_ID")
    clientSecret := os.Getenv("GA_CLIENT_SECRET")
    refreshToken := os.Getenv("GA_REFRESH_TOKEN")
    return ga.NewContext(clientID, clientSecret, refreshToken)
}

func ExampleGetEvent() {
    res := ga.GetEvent(getNewContextFromEnv(), os.Getenv("GA_ID"), "^.*Tag/.*", 30, true)
    fmt.Println(res)
    // Output:
    // map[iPhone:10 Android:17 Windows:12]
}

func ExampleGet() {
    metrics := "ga:uniquePageviews"
    dimensions := "ga:pagePath"
    sortOrder := "-ga:uniquePageviews"
    filters := "ga:pagePath=~^/path/to/site/.*"
    res := ga.Get(getNewContextFromEnv(), os.Getenv("GA_ID"), metrics, dimensions, filters, sortOrder, 1)
    fmt.Println(res)
    // Output:
    // map[/path/to/site/index.html:100 /path/to/site/about.html:10 /path/to/site/other.html:12]
}
```
