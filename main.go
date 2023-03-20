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
	Usage *ShowUsage
	Input string
}

type ShowUsage struct {
	Amiibo []struct {
		AmiiboSeries string `json:"amiiboSeries"`
		Character    string `json:"character"`
		GameSeries   string `json:"gameSeries"`
		Games3DS     []struct {
			AmiiboUsage []struct {
				Usage string `json:"Usage"`
				Write bool   `json:"write"`
			} `json:"amiiboUsage"`
			GameID   []string `json:"gameID"`
			GameName string   `json:"gameName"`
		} `json:"games3DS"`
		GamesSwitch []struct {
			AmiiboUsage []struct {
				Usage string `json:"Usage"`
				Write bool   `json:"write"`
			} `json:"amiiboUsage"`
			GameID   []string `json:"gameID"`
			GameName string   `json:"gameName"`
		} `json:"gamesSwitch"`
		GamesWiiU []struct {
			AmiiboUsage []struct {
				Usage string `json:"Usage"`
				Write bool   `json:"write"`
			} `json:"amiiboUsage"`
			GameID   []string `json:"gameID"`
			GameName string   `json:"gameName"`
		} `json:"gamesWiiU"`
		Head    string `json:"head"`
		Image   string `json:"image"`
		Name    string `json:"name"`
		Release struct {
			Au string `json:"au"`
			Eu string `json:"eu"`
			Jp string `json:"jp"`
			Na string `json:"na"`
		} `json:"release"`
		Tail string `json:"tail"`
		Type string `json:"type"`
	} `json:"amiibo"`
}

func main() {
	var a AmiiboStruct
	a.Input = ""
	fmt.Println("server is running on port 8080 : http://localhost:8080")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("script"))))
	http.HandleFunc("/", a.Index)
	http.HandleFunc("/name", a.Name)
	http.HandleFunc("/character", a.Character)
	http.ListenAndServe(":8080", nil)
}

func (a *AmiiboStruct) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp := template.Must(template.ParseFiles("index.html"))
	details := AmiiboStruct{}
	tmp.Execute(w, details)
}

func (a *AmiiboStruct) Name(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/name" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp := template.Must(template.ParseFiles("pages/name.html"))
	if r.Method != http.MethodPost {
		tmp.Execute(w, nil)
		return
	}
	details := AmiiboStruct{
		Input: r.FormValue("input"),
		Data:  getNameData(r.FormValue("input")),
		Usage: getUsage(r.FormValue("input")),
	}
	tmp.Execute(w, details)
}

func getNameData(request string) *AmiiboStruct {
	var tmp string
	var url string = "https://www.amiiboapi.com/api/amiibo/"
	if request != "" {
		for i := 0; i != len(request); i++ {
			if request[i] == ' ' {
				tmp += "%20"
				i++
			}
			tmp += string(request[i])
		}
		request = tmp
		url += "?name=" + request
	} else {
		return &AmiiboStruct{}
	}
	return MyUnmarshal(url)
}

func (a *AmiiboStruct) Character(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/character" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp := template.Must(template.ParseFiles("pages/character.html"))
	details := AmiiboStruct{
		Input: r.FormValue("input"),
		Data:  getCharacterData(r.FormValue("input")),
		Usage: getUsage(r.FormValue("input")),
	}
	tmp.Execute(w, details)
}

func getCharacterData(request string) *AmiiboStruct {
	var tmp string
	var url string = "https://www.amiiboapi.com/api/amiibo/"
	if request != "" {
		for i := 0; i != len(request); i++ {
			if request[i] == ' ' {
				tmp += "%20"
				i++
			}
			tmp += string(request[i])
		}
		request = tmp
		url += "?character=" + request
	} else {
		return &AmiiboStruct{}
	}
	return MyUnmarshal(url)
}

func MyUnmarshal(url string) *AmiiboStruct {
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

func getUsage(request string) *ShowUsage {
	var tmp string
	var url string = "https://www.amiiboapi.com/api/amiibo/"
	if request != "" {
		for i := 0; i != len(request); i++ {
			if request[i] == ' ' {
				tmp += "%20"
				i++
			}
			tmp += string(request[i])
		}
		request = tmp
		url += "?character=" + request + "&showusage"
	} else {
		return &ShowUsage{}
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
	newDataStruct := ShowUsage{}
	jsonErr := json.Unmarshal(body, &newDataStruct)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return &newDataStruct
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 not found")
	}
}
