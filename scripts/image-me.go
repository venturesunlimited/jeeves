package scripts

import (
	"regexp"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type GoogleImageSearchResponse struct {
	Results struct {
		Images []struct {
			UnescapedUrl string `json:"unescapedUrl"`
		} `json:"results"`
	} `json:"responseData"`
}

func ImageMe(text string) Response {
	r, _ := regexp.Compile("^image me (.*)")
	if match := r.FindStringSubmatch(text); match != nil {
		query := match[1]
		url := "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&rsz=8&q="+query+"&safe=active"
		resp, _ := http.Get(url)

		defer resp.Body.Close()

		respString, _ := ioutil.ReadAll(resp.Body)
		var data GoogleImageSearchResponse
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
