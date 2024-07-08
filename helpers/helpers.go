package helpers

import (
	"fmt"
	"time"
)

func GetGreetings(time time.Time, name string) string {
	return fmt.Sprintf("%s: Hello there, %s", time, name)
}
