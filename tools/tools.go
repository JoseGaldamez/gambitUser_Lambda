package tools

import (
	"fmt"
	"time"
)

func DateMySQL() string {
	currentDate := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", currentDate.Year(), currentDate.Month(), currentDate.Day(), currentDate.Hour(), currentDate.Minute(), currentDate.Second())
}
