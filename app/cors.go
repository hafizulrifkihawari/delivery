package app

import (
	"github.com/gin-contrib/cors"
)

// allowCors is used to register url to user our resources
func allowCors() {
	var (
		whitelists []string = []string{"*"} // currently allow all resources
	)
	config := cors.DefaultConfig()
	config.AllowWildcard = true
	config.AllowOrigins = whitelists
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "Accept-Encoding")

	router.Use(cors.New(config))
}
