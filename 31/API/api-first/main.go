package main

import (
	"fmt"
	"github.com/x86nick/go100/31/API/api-first/handlers"
	"net/http"
	"os"
)

func main() {
	// we need to define antoher handler for users and it is going to be part of handlers package
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	// 2. add a default route
	http.HandleFunc("/", handlers.RootHandler) // we will create rootHandler function
	// 1. start a server
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		//log.Fatal(err)
		// panic(err) - it will terminate the program
		/*
				panic: listen tcp 127.0.0.1:11111: bind: address already in use
			goroutine 1 [running]:
			main.main()
			/Users/nkumar17/GO100/go100/31/API/api-first/main.go:9 +0xc9
			exit status 2
		*/
		fmt.Println(err) // if error -> listen tcp 127.0.0.1:11111: bind: address already in use
		os.Exit(1)       // exit status 1
	}

}
