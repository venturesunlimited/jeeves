package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"regexp"
	"io/ioutil"
	"os"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func main() {
	token := "uaBDHFoe1ve9oNJPCKX6a350"

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
	// http.ListenAndServe(":4567", nil)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	messageText := strings.Replace(r.PostForm["text"][0], r.PostForm["trigger_word"][0]+" ", "", 1)

	if response := ImageMe(messageText); response != nil {
		fmt.Fprint(w, response)
	}

	return
}

type Image struct {
	UnescapedUrl string `json:"unescapedUrl"`
}

type Result struct {
	Images []Image `json:"results"`
}

type ResponseData struct {
	Results Result `json:"responseData"`
}

func ImageMe(text string) Response {
	r, _ := regexp.Compile("^image me (.*)")
	if match := r.FindStringSubmatch(text); match != nil {
		query := match[1]
		url := "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&rsz=8&q="+query+"&safe=active"
		resp, _ := http.Get(url)

		defer resp.Body.Close()

		respString, _ := ioutil.ReadAll(resp.Body)
		var data ResponseData
		err := json.Unmarshal(respString, &data)
		if err != nil {
			panic(err)
		}

		imageUrl := data.Results.Images[0].UnescapedUrl
		return Response{"text": imageUrl}
	} else {
		return nil
	}
}
