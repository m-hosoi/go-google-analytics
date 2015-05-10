package ga

import (
	"strconv"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/analytics/v3"
)

// NewContext is constractor
func NewContext(clientID, clientSecret, refreshToken string) Context {
	return Context{clientID: clientID, clientSecret: clientSecret, refreshToken: refreshToken}
}

// Context .
type Context struct {
	clientID     string
	clientSecret string
	refreshToken string
}

// CreateAnalyticsService : create service
func (ac Context) CreateAnalyticsService() *analytics.Service {
	config := &oauth2.Config{
		ClientID:     ac.clientID,
		ClientSecret: ac.clientSecret,
		Endpoint:     google.Endpoint,
	}
	t := &oauth2.Token{
		RefreshToken: ac.refreshToken,
	}
	ctx := context.Background()
	c := config.Client(ctx, t)
	svc, err := analytics.New(c)
	checkErr(err)
	return svc
}

// GetEvent is ...
func GetEvent(c Context, id string, category string, days int, useRegex bool) map[string]int {
	filters := ""
	if useRegex {
		filters = "ga:eventCategory=~" + category
	} else {
		filters = "ga:eventCategory==" + category
	}
	return Get(c, id, "ga:uniqueEvents", "ga:eventAction", filters, "-ga:uniqueEvents", days)
}

// Get is ...
func Get(c Context, id, metrics, dimensions, filters, sortOrder string, days int) map[string]int {
	svc := c.CreateAnalyticsService()
	ds := analytics.NewDataGaService(svc)
	now := time.Now()
	startTime := now.AddDate(0, 0, -1*days)
	query := ds.Get("ga:"+id, startTime.Format("2006-01-02"), now.Format("2006-01-02"), metrics).
		Dimensions(dimensions).
		Sort(sortOrder)
	if filters != "" {
		query = query.Filters(filters)
	}
	data, err := query.Do()
	checkErr(err)
	res := map[string]int{}
	for _, v := range data.Rows {
		i, _ := strconv.Atoi(v[1])
		res[v[0]] = i
	}
	return res

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
