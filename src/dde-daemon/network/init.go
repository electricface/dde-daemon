package network

import "dde-daemon"

func init() {
	loader.Register(&loader.Module{"network", Start, Stop, true})
}