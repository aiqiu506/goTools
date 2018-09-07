package grqTime

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

const (
	YYYY = "2006"
	YY   = "06"
	MM   = "01"
	M    = "1"
	DD   = "02"
	D    = "2"
	HH   = "15"
	H    = "03"
	II   = "04"
	SS   = "06"
)

func FormaterTime(fstr string, timestamp timestamp.Timestamp) string {

	return timestamp.Format(YY + "/" + MM + "/" + D)
}
