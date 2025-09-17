package ui

import (
	"fmt"
	"strings"
)

func TopLocationBar(location string, time string) {
    paddedLocation := PadLocation(location)
    spaces := 98 - (3 + 40 + len(time) + 4)
	fmt.Println(strings.Repeat("-", 98))
    fmt.Print("|  ")
    fmt.Print(paddedLocation)
    fmt.Print(strings.Repeat(" ", spaces))
    fmt.Print(time)
    fmt.Println("  |")
	fmt.Println(strings.Repeat("-", 98))
	fmt.Println("")
	fmt.Println("")
}

func PadLocation(location string) string {
    if len(location) >= 40 {
        return location[:40]
    }
    return location + strings.Repeat(" ", 41-len(location))
}