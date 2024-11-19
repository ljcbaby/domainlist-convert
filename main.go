package main

import (
	"github.com/dn-11/provider2domainset/cmd"
	"github.com/dn-11/provider2domainset/log"
)

func main() {
	log.Init()
	cmd.Execute()
}
