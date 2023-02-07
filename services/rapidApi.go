package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RapidApi struct {
}

func (r RapidApi) GetNews() {

	url := "https://livescore6.p.rapidapi.com/matches/v2/list-live?Category=soccer&Timezone=-7"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("X_RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("X_RAPIDAPI_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
