package main

import (
	"fmt"
)

var (
	GitHash   string
	BuildTime string
	Version   string
)

func main() {
	fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
}
