package common

import (
	"log"
	"os"
)

// Logger is used for CLI only
var Logger = log.New(os.Stdout, "", 0)
