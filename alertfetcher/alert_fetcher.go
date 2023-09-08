package alertfetcher

import (
	"fmt"

	_ "github.com/albdewilde/init/locationfetcher"
	_ "github.com/albdewilde/init/userfetcher"
)

func init() {
	fmt.Println("initialize the alert fetcher package")
}
