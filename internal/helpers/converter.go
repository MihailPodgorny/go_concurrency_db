package helpers

import "github.com/dustin/go-humanize"

func ConvertFromString(s string) (uint64, error) {
	return humanize.ParseBytes(s)
}
