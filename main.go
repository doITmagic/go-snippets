package main

import (
	"fmt"
	"os"
	"time"

	"github.com/doITmagic/go-snippets/broadcast"
	"github.com/doITmagic/go-snippets/hashtable"
	"github.com/doITmagic/go-snippets/logger"
	"github.com/doITmagic/go-snippets/patterns/messaging/fanOut"
	"github.com/doITmagic/go-snippets/patterns/structural/decorator"
	"github.com/doITmagic/go-snippets/prepend"
	slicesexp "github.com/doITmagic/go-snippets/slices"
)

func main() {
	//using logger package
	myLogger := logger.NewLogger(os.Stdout, 10)
	defer myLogger.Close()

	for i := 0; i < 10; i++ {
		myLogger.Log(fmt.Sprintf("Hello %d", i))
		time.Sleep(100 * time.Millisecond)
	}
	//
	//using hashtable package
	testHashTable := hashtable.InitHashTable()
	testHashTable.Insert("test")
	testHashTable.Insert("test1")
	testHashTable.Insert("test1")
	testHashTable.Insert("test2")
	testHashTable.Insert("test3")
	testHashTable.Insert("test4")
	found := testHashTable.Search("test")
	testHashTable.Delete("test")
	fmt.Println(testHashTable, found)

	//using broadcast package
	broadcast.ExampleBroadcast()

	//using fanOut package
	fanOut.Example()

	//using prepend package
	s := []interface{}{2, 3, 4}
	var mySlice prepend.Prepender
	s = mySlice.Prepend(1, s)
	s = append(s, 11)
	fmt.Println(s)

	//using decorator package
	decorator.Example()

	//slice tricks with experimental packages
	slicesexp.Example()

}
