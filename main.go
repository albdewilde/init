package main

import (
	"fmt"

	_ "github.com/albdewilde/init/alertfetcher"
	_ "github.com/albdewilde/init/alertupdater"
	_ "github.com/albdewilde/init/classifiedfetcher"
	_ "github.com/albdewilde/init/programfetcher"
)

func main() {
	fmt.Println("Starting the main program")
}
