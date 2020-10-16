package main

import (
	"flag"
	"fmt"

	"github.com/techxmind/location2ip"
)

func main() {
	var showHelp bool
	var count int
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.IntVar(&count, "n", 1, "number of generations")

	flag.Parse()

	usage := `
Generate ip address from given geo location

loc2ip [-n COUNT] LOCATION

example
  loc2ip -n 5 北京
	`

	args := flag.Args()

	if showHelp || len(args) == 0 {
		fmt.Println(usage)
		return
	}

	for i := 0; i < count; i++ {
		if ip, err := location2ip.Generate(args[0]); err != nil {
			fmt.Printf(
				"\033[0;31m%s\033[0m\n",
				err,
			)
			break
		} else {
			fmt.Println(ip)
		}
	}
}
