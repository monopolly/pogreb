package pogreb

type loggs struct {
	h func(s string)
}

func (a *loggs) Write(p []byte) (n int, err error) {
	if a.h == nil {
		return
	}
	a.h(string(p))
	return
}
