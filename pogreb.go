package pogreb

import (
	"log"

	"github.com/akrylysov/pogreb"
)

//volume, "db", "nice" > volume/db/nice
func New(volume string, logger ...func(s string)) (a *DB, err error) {
	a = new(DB)
	logs := log.Default()
	if len(logger) > 0 {
		b := new(loggs)
		b.h = logger[0]
		logs.SetOutput(b)
	}
	logs.SetFlags(0)
	pogreb.SetLogger(logs)
	a.db, err = pogreb.Open(volume, nil)
	return
}

//DB ok
type DB struct {
	db *pogreb.DB
}

// has
func (a *DB) Has(k, v []byte) (has bool) {
	has, _ = a.db.Has(k)
	return
}

// put
func (a *DB) Put(k, v []byte) {
	a.db.Put(k, v)
}

// put
func (a *DB) Add(k string, v []byte) {
	a.db.Put([]byte(k), v)
}

// get
func (a *DB) Get(k []byte) (v []byte) {
	v, _ = a.db.Get(k)
	return
}

// get with error
func (a *DB) GetE(k []byte) (v []byte, err error) {
	return a.db.Get(k)
}

func (a *DB) Delete(k []byte) {
	a.db.Delete(k)
}

func (a *DB) Close() {
	a.db.Close()
}

func (a *DB) Count() int {
	return int(a.db.Count())
}

//Iterate проходит по всей базе, если обратно пришло value в базе заменяется
func (a *DB) IterateAndUpdate(f func(k, v []byte) (value []byte)) {
	it := a.db.Items()
	for {
		k, v, err := it.Next()
		if err != nil {
			return
		}
		v = f(k, v)
		if len(v) == 0 {
			continue
		}
		a.db.Put(k, v)
	}
}

//Iterates проходит по всей базе
func (a *DB) Iterate(f func(k, v []byte)) {
	it := a.db.Items()
	for {
		k, v, err := it.Next()
		if err != nil {
			return
		}
		f(k, v)
	}
}

// put
func (a *DB) Items() *pogreb.ItemIterator {
	return a.db.Items()
}
