package utils

import "time"

// Delay pauses the execution for the specified number of seconds
// to create more natural dialogue timing.
func Delay(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
