package pogreb

import (
	"encoding/binary"
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
		return int(binary.LittleEndian.Uint64(vv))
	}
	bb := make([]byte, 8)
	binary.LittleEndian.PutUint64(bb, uint64(v[0]))
	a.Put([]byte(k), bb)
	return
}

func (a *DB) Int64(k string, v ...int64) (r int64) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return int64(binary.LittleEndian.Uint64(vv))
	}
	bb := make([]byte, 8)
	binary.LittleEndian.PutUint64(bb, uint64(v[0]))
	a.Put([]byte(k), bb)
	return
}

func (a *DB) Uint64(k string, v ...uint64) (r uint64) {
	if len(v) == 0 {
		vv, err := a.GetE([]byte(k))
		if err != nil || len(vv) == 0 {
			return
		}
		return binary.LittleEndian.Uint64(vv)
	}
	bb := make([]byte, 8)
	binary.LittleEndian.PutUint64(bb, uint64(v[0]))
	a.Put([]byte(k), bb)
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
