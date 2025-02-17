package common

import (
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v3"
)

func SetCookieStr(ck *fiber.Cookie) string{
	
	cookieStr := fmt.Sprintf("%s=%s", ck.Name, ck.Value)

	if !ck.Expires.IsZero() {
		cookieStr += "; Expires=" + ck.Expires.UTC().Format(http.TimeFormat)
	}
	if ck.MaxAge > 0 {
		cookieStr += fmt.Sprintf("; Max-Age=%d", ck.MaxAge)
	}
	if ck.Path != "" {
		cookieStr += "; Path=" + ck.Path
	}
	if ck.Domain != "" {
		cookieStr += "; Domain=" + ck.Domain
	}
	if ck.SameSite != "" {
		cookieStr += "; SameSite=" + ck.SameSite
	}
	if ck.Secure {
		cookieStr += "; Secure"
	}
	if ck.HTTPOnly {
		cookieStr += "; HttpOnly"
	}
	if ck.Partitioned {
		cookieStr += "; Partitioned"
	}
	return cookieStr
}