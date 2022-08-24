package main

import (
	"fmt"
	"wall-server/app"
)

func main() {
	_, err := app.Start()
	if err != nil {
		fmt.Println(err)
	}
}
