package config

import (
	"strconv"
)

const APP_VERSION = "1.0.0"

var (
	BOT_TOKEN             = Env("BOT_TOKEN")
	ADMIN_ID, _           = strconv.ParseInt(Env("ADMIN_ID"), 10, 64)
	CHANNEL_ID            = Env("CHANNEL_ID")
	CHAPPIE_SERVER_URL    = Env("CHAPPIE_SERVER_URL")
	CHAPPIE_SERVER_BEARER = Env(("CHAPPIE_SERVER_BEARER"))
)
