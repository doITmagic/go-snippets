package main

import (
	"github.com/doITmagic/go-snippets/broadcast"
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
	//
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

	broadcast.Broadcast_Example()

}
