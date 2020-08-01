package utils

import (
	"strings"
	"time"
)

type (
	XTime time.Time
)

var mapping map[string]string

func init() {
	mapping = make(map[string]string)
	mapping["%i"] = "4"
	mapping["%I"] = "04"

	mapping["%m"] = "1"
	mapping["%d"] = "2"
	mapping["%h"] = "3"
	mapping["%s"] = "5"
	mapping["%y"] = "06"
	mapping["%M"] = "01"
	mapping["%D"] = "02"
	mapping["%H"] = "15"
	mapping["%S"] = "05"
	mapping["%Y"] = "2006"
}

// 用 %y %M %d 的格式来格式化时间
func (self XTime) Format(format string) string {

	for key, val := range mapping {
		format = strings.ReplaceAll(format, key, val)
	}
	return time.Time(self).Format(format)
}
