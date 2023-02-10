package action

import (
	"encoding/base64"
	"errors"
	"fmt"
)

func Base64Encode(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("not enough arguments")
	}

	content := args[0]

	if len(args) > 1 {
		for _, a := range args[1:] {
			content += fmt.Sprintf(" %s", a)
		}
	}

	return base64.StdEncoding.EncodeToString([]byte(content)), nil
}

func Base64Decode(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("not enough arguments")
	}

	res, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		return "", fmt.Errorf("base64 decode string: %w", err)
	}

	return string(res), nil
}
