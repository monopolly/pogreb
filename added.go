package pogreb

import (
	"fmt"
	"path"
)

//New engine — volume, dbpath1, dbpath2...
func New(volume ...string) (a *DB, err error) {
	a, err = Open(path.Join(volume...), nil)
	return
}

//Iterate проходит по всей базе, если обратно пришло value в базе заменяется
func (a *DB) Iterate(f func(k, v []byte) (value []byte)) {
	it := a.Items()
	for {
		k, v, err := it.Next()
		if err == ErrIterationDone {
			return
		}
		v = f(k, v)
		if v == nil {
			continue
		}
		a.Put(k, v)
	}
}

//Iterates проходит по всей базе
func (a *DB) Iterates(f func(k, v []byte)) {
	it := a.Items()
	for {
		k, v, err := it.Next()
		if err == ErrIterationDone {
			break
		}
		if err != nil {
			fmt.Println("iterate error", err)
		}
		f(k, v)
	}
}
