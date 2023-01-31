package sub2config

import (
	"fmt"

	"github.com/BaymaxRice/sub2config/utils/ping"
)

type Link interface {
	fmt.Stringer
	Ping(round int, dst string) (ping.Status, error)
	DestAddr() string
	Description() string
	Safe() string
	Config
}

type Config interface {
	Config() map[string]string
}
