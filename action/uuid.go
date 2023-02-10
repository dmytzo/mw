package action

import (
	"github.com/google/uuid"
)

func UUID(_ []string) (string, error) {
	return uuid.NewString(), nil
}
