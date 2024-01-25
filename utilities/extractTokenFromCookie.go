package utilities

import (
	_"log"
	"net/http"
)

func ExtractTokenFromCookie(r* http.Request) string{
cookies:= r.Cookies()
for _,cookie := range cookies{
    if cookie.Name=="token"{
		return cookie.Value
	}
}
return ""
}