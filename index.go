package log_default

import (
	"github.com/infrago/infra"
	"github.com/infrago/log"
)

func Driver() log.Driver {
	return &defaultDriver{}
}

func init() {
	infra.Register("default", Driver())
}
