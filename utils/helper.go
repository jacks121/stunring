package utils

import (
	"strings"
)

func IsMobileDevice(userAgent string) bool {
	return strings.Contains(userAgent, "Mobile") || strings.Contains(userAgent, "Android") || strings.Contains(userAgent, "iPhone")
}
