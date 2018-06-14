package main

import (
	"log"

	"github.com/alvaradojl/test/pkg/service1"
)

func main() {

	res, _ := service1.Add(1, 2)

	log.Printf("\nthis is a test 1 + 2 = %d\n", res)
}
