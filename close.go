package pogreb

import (
	"os"
	"os/signal"
)

// goroutine to catch os signal
func Closer(db ...*DB) {
	if len(db) == 0 {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			for _, x := range db {
				x.Close()
			}
			os.Exit(1)
		}
	}()
}
