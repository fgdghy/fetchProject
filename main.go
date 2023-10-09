package main

import "github.com/fetchProject/app"

func main() {
	r := app.InitHandler()

	r.Run(":8080")
}
