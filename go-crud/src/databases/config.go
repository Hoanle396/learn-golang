package databases

import (
	"os"

	"gitlab.com/avarf/getenvs"
)

var HOST = getenvs.GetEnvString("POSTGRES_HOST", "localhost")
var PORT, _ = getenvs.GetEnvInt("POSTGRES_PORT", 5432)
var USER = getenvs.GetEnvString("POSTGRES_USER", "postgre")
var PASSWORD = getenvs.GetEnvString("POSTGRES_PASSWORD", "postgre")
var DATABASE = getenvs.GetEnvString("POSTGRES_DATABASE", "golang")
var SSLMODE = getenvs.GetEnvString("POSTGRES_SSLMODE", "disable")
var POSTGRES_CONNECT_URL = os.Getenv("POSTGRES_CONNECT_URL")
