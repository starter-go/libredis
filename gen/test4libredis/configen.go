package test4libredis

import "github.com/starter-go/application"

//starter:configen(version="4")

// ComponentsForTestLibRedis ...
func ComponentsForTestLibRedis(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
