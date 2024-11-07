package main

import (
	api "api/internal/newspaper-api"
)

func main() {

	api.NewAPIService().Start()
}
