package main

import (
	"fmt"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
)

func main() {
	fmt.Printf(`%#v`, config.GlobalConfig)
}
