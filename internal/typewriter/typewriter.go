package typewriter

import (
	"fmt"
	"time"
)

// Type fn
func Type(text string, speed time.Duration) {
	sleep := int(speed) / len(text)
	for _, ch := range text {
		fmt.Printf(string(ch))
		time.Sleep(time.Duration(sleep))
	}
}
