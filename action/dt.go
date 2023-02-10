package action

import (
	"fmt"
	"time"
)

const (
	dateFlag = "--d"
	timeFlag = "--t"

	dateTimeFormat = "15:04:05 02.01.2006"
	dateFormat     = "02.01.2006"
	timeFormat     = "15:04:05"
)

func DateTime(args []string) (res string, err error) {
	f := dateTimeFormat

	if len(args) > 0 {
		switch args[0] {
		case dateFlag:
			f = dateFormat
		case timeFlag:
			f = timeFormat
		}
	}

	return fmt.Sprintf("%s\n", time.Now().Format(f)), nil
}
