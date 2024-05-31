package main

import "com.ardafirdausr.cupid/app/http"

func main() {
	httpApp, close, err := http.InitializeApp()
	if err != nil {
		panic(err)
	}

	defer close()
	httpApp.Start()
}
