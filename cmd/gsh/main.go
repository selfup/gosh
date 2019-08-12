package main

import "flag"

func main() {
	var cmd string
	flag.StringVar(&cmd, "cmd", "", "REQUIRED TO RUN ANYTHING")

	flag.Parse()
}
