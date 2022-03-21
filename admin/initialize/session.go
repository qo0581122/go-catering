package initialize

import (
	"catering/global"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func InitSession() {
	session := sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))
	session.Options.MaxAge = 0
	global.SESSION = session
}
