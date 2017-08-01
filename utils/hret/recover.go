package hret

import "github.com/asofdate/sso-jwt-auth/utils/logger"

type httpPanicFunc func()

// HttpPanic user for stop panic up.
func HttpPanic(f ...httpPanicFunc) {
	if r := recover(); r != nil {
		logger.Error("system generator panic.", r)
		for _, val := range f {
			val()
		}
	}
}
