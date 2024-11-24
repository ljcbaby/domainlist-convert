package main

import (
	"github.com/ljcbaby/domainlist-convert/cmd"
	"github.com/ljcbaby/domainlist-convert/log"
)

func main() {
	log.Init()
	cmd.Execute()
}
