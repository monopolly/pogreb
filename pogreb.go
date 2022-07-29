package pogreb

import (
	"os"
	"os/signal"

	"github.com/akrylysov/pogreb"
	"github.com/monopolly/numbers"
)

//volume, "db", "nice" > volume/db/nice
func New(volume string) (a *DB, err error) {
	a = new(DB)
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

// запускает горутину чтобы от сигнала системы закрыть базу
func (a *DB) CloseSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			a.db.Close()
			return
		}
	}()
}

func (a *DB) Count() int {
	return int(a.db.Count())
}

func (a *DB) String(k string, v ...string) (r string) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return string(vv)
	}
	a.Put([]byte(k), []byte(v[0]))
	return
}

func (a *DB) Int(k string, v ...int) (r int) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return numbers.BytesInt(vv)
	}
	a.Put([]byte(k), numbers.IntBytes(v[0]))
	return
}

func (a *DB) Int64(k string, v ...int64) (r int64) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return numbers.BytesInt64(vv)
	}
	a.Put([]byte(k), numbers.Int64Bytes(v[0]))
	return
}

func (a *DB) Uint64(k string, v ...uint64) (r uint64) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return numbers.BytesUint64(vv)
	}
	a.Put([]byte(k), numbers.Uint64Bytes(v[0]))
	return
}

func (a *DB) Bool(k string, v ...bool) (r bool) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return vv[0] == 1
	}
	switch v[0] {
	case true:
		a.Put([]byte(k), []byte{1})
	default:
		a.db.Delete([]byte(k))
	}
	return
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
