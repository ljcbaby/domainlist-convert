package convert

import "github.com/dn-11/provider2domainset/conf"

type Task struct {
	conf.File
	Source string
	Target string
}
