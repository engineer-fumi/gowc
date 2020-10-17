package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/engineer-fumi/gowc"
)

func main() {
	var (
		p = flag.String("p", "", "Please specify the name of the file or the directory.")
	)
	flag.Parse()

	path := *p
	args := flag.Args()
	if *p == "" {
		if 0 == len(args) {
			flag.Usage()
			return
		}
		path = args[0]
	}

	wc := gowc.NewGoWC(path)
	total, err := wc.Counte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
