package main

import (
	"examplesvc/service/servicea"
	"examplesvc/service/serviceb"
	"flag"
)

func main() {
	s := flag.String("service", "", "use servicea or serviceb")
	flag.Parse()
	switch *s {
	case "servicea":
		servicea.Serve()
	case "serviceb":
		serviceb.Serve()
	default:
		flag.Usage()
	}
}
