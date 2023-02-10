package action

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"mw/std"
)

var ticker = time.NewTicker(time.Second)

func Counter(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("not enough arguments")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("parse int argument: %w", err)
	}

	if count <= 0 {
		return "", errors.New("counter have to be greater than 0")
	}

	currentCount := count

	for currentCount > 0 {
		<-ticker.C

		std.WriteResult(strconv.Itoa(count-currentCount+1), false)
		currentCount -= 1
	}

	return "done", nil
}
