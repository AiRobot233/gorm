package main

import "gin/routes"

func main() {
	router := routes.Routes()
	_ = router.Run(":9501")
}
