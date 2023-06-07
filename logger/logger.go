package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func NewLogger(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)

	go func(ww io.Writer) {
		defer l.wg.Done()
		for v := range l.ch {
			_, err := ww.Write([]byte(v + "\n"))
			if err != nil {
				fmt.Println("Error writing to file")
			}
		}
	}(w)

	return &l
}

func (l *Logger) Log(s string) {
	select {
	case l.ch <- s:
		//fmt.Println("Logger channel is not full, send message")
	default:
		//fmt.Println("Logger channel is full")
	}
}

func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}
