package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type AmiiboStruct struct {
	Amiibo []struct {
		AmiiboSeries string `json:"amiiboSeries"`
		Character    string `json:"character"`
		GameSeries   string `json:"gameSeries"`
		Head         string `json:"head"`
		Image        string `json:"image"`
		Name         string `json:"name"`
		Release      struct {
			Au string `json:"au"`
			Eu string `json:"eu"`
			Jp string `json:"jp"`
			Na string `json:"na"`
		} `json:"release"`
		Tail string `json:"tail"`
		Type string `json:"type"`
	} `json:"amiibo"`
	Data  *AmiiboStruct
	Input string
}

func main() {
	var a AmiiboStruct
	a.Input = ""
	fmt.Println("server is running on port 8080")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", a.Index)
	fmt.Println(getData("link"))
	http.ListenAndServe(":8080", nil)
}

func (a *AmiiboStruct) Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := AmiiboStruct{
		Input: r.FormValue("input"),
		Data:  getData(r.FormValue("input")),
	}
	tmp.Execute(w, details)
}

func getData(request string) *AmiiboStruct {
	var url string = "https://www.amiiboapi.com/api/amiibo/"
	fmt.Println(request)
	if request != "" {
		url += "?name=" + request
	}
	timeClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	newDataStruct := AmiiboStruct{}
	jsonErr := json.Unmarshal(body, &newDataStruct)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return &newDataStruct
}
