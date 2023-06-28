package main

import (
	"fmt"

	"github.com/doITmagic/go-snippets/errors"
)

func main() {
	////using logger package
	//myLogger := logger.NewLogger(os.Stdout, 10)
	//defer myLogger.Close()
	//
	//for i := 0; i < 10; i++ {
	//	myLogger.Log(fmt.Sprintf("Hello %d", i))
	//	time.Sleep(100 * time.Millisecond)
	//}
	////
	////using hashtable package
	//testHashTable := hashtable.InitHashTable()
	//testHashTable.Insert("test")
	//testHashTable.Insert("test1")
	//testHashTable.Insert("test1")
	//testHashTable.Insert("test2")
	//testHashTable.Insert("test3")
	//testHashTable.Insert("test4")
	//found := testHashTable.Search("test")
	//testHashTable.Delete("test")
	//fmt.Println(testHashTable, found)
	//
	////using broadcast package
	//broadcast.ExampleBroadcast()
	//
	////using fanOut package
	//fanOut.Example()
	//
	////using prepend package
	//s := []interface{}{2, 3, 4}
	//var mySlice prepend.Prepender
	//s = mySlice.Prepend(1, s)
	//s = append(s, 11)
	//fmt.Println(s)
	//
	////using decorator package
	//decorator.Example()
	//
	////slice tricks with experimental packages
	//slicesexp.Example()

	//er1 := errors.New("error 1")
	//er2 := errors.New("error 2")
	//er3 := errorsWrap.New("error 3")
	//err4 := errorsWrap.New("error 4")

	//fmt.Errorf("error 6: %w", err5)
	//fmt.Println(errorsWrap.Unwrap(err6))
	customError1 := errors.New("second custom error")

	fmt.Println(errors.StackTrace(customError1))

}
