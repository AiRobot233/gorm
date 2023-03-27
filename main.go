package main

import "gin/routes"

func main() {
	r := routes.Routes()
	_ = r.Run(":9501")
}
