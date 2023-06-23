package detectClosedChannel

func IsClosed(ch chan interface{}) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
			return
		}
	}()
	ch <- struct{}{}
	return
}
