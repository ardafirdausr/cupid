package main

import "com.ardafirdausr.cupid/app/http"

func main() {
	httpApp := http.InitializeApp()
	httpApp.Start()
}
