package pogreb

import (
	"github.com/monopolly/numbers"
)

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
