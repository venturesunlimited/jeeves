package main

import (
	"github.com/gorilla/mux"
	"github.com/jplaut/jeeves/scripts"
	"fmt"
	"strings"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	router := mux.NewRouter()
	router.HandleFunc("/callback", CallbackHandler).
	 	Methods("POST").
	 	MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
      r.ParseForm()
   	  return r.PostForm["token"][0] == token
	  })

	http.Handle("/", router)
	fmt.Println("listening...")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	messageText := strings.Replace(r.PostForm["text"][0], r.PostForm["trigger_word"][0]+" ", "", 1)

	if response := scripts.ImageMe(messageText); response != nil {
		fmt.Fprint(w, response)
	}

	return
}

