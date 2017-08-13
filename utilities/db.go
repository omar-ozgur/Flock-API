package utilities

import (
	"os"
)

// Database configuration values
var DB_INFO = os.Getenv("DB_INFO")
var DB_USER = "postgres"
var DB_PASSWORD = "postgres"
var DB_NAME = "flock_api"
var DB_HOST = "localhost"

// Table names
var USERS_TABLE = "users"
var EVENTS_TABLE = "events"
var ATTENDEES_TABLE = "attendees"
