package main

import (
	"readcommend/internal/bootstrap"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		panic(err)
	}
}
