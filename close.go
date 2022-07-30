package pogreb

import (
	"os"
	"os/signal"
)

// goroutine to catch os signal
func Closer(add chan *DB) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		var dblist []*DB
		select {
		case db := <-add:
			dblist = append(dblist, db)
		case <-c:
			for _, x := range dblist {
				x.Close()
			}
			os.Exit(1)
		}
	}()
}
