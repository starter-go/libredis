package lib4libredis

import "github.com/starter-go/application"

//starter:configen(version="4")

// ComponentsForLibRedis ...
func ComponentsForLibRedis(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
