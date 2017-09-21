package goul

import "time"



var _version = time.Now().Format("20060102150405")

func GoulVersion() string {
	return _version
}
