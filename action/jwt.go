package action

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func JWT(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no token provided")
	}

	parts := strings.Split(args[0], ".")
	if len(parts) != 3 {
		return "", errors.New("wrong token format")
	}

	headers, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", fmt.Errorf("base64 decode headers: %w", err)
	}

	claims, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("base64 decode claims: %w", err)
	}

	var out bytes.Buffer

	if err = json.Indent(&out, claims, "", "  "); err != nil {
		return "", fmt.Errorf("json ident claims: %w", err)
	}

	prettyClaims := out.String()

	out.Reset()

	if err = json.Indent(&out, headers, "", "  "); err != nil {
		return "", fmt.Errorf("json ident headers: %w", err)
	}

	prettyHeaders := out.String()

	dates, err := parseDates(claims)
	if err != nil {
		return "", fmt.Errorf("parse dates: %w", err)
	}

	return fmt.Sprintf(`
# headers
%s

# claims
%s

# parsed dates
%s
`, prettyHeaders, prettyClaims, dates), nil

}

func parseDates(claims []byte) (res string, err error) {
	var h map[string]interface{}

	if err = json.Unmarshal(claims, &h); err != nil {
		return res, fmt.Errorf("json unmarshal claims: %w", err)
	}

	var dates strings.Builder

	for _, k := range []string{"iat", "exp", "nbf"} {
		v, ok := h[k]
		if !ok {
			continue
		}

		if val, ok := v.(float64); ok {
			dates.WriteString(fmt.Sprintf("%s: %s\n", k, time.Unix(int64(val), 0).UTC().String()))
		}
	}

	return dates.String(), err
}
