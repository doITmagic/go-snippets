package main

import (
	"fmt"
	"os"
	"time"

	"github.com/doITmagic/go-snippets/logger"
)

func main() {
	//using logger package
	myLogger := logger.NewLogger(os.Stdout, 10)
	defer myLogger.Close()

	for i := 0; i < 10; i++ {
		myLogger.Log(fmt.Sprintf("Hello %d", i))
		time.Sleep(100 * time.Millisecond)
	}

}
