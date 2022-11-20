package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/georgealan")
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.printf("error: %s", err)
			os.exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	/*
		These _, below is a very special variable in Go called underscore, which can be used in a place,
		can be any type, and we cannot read or print from it.
	*/
	/*if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: can't copy - %s", err)
	}*/

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Println(r)
	fmt.Printf("%#v\n", r)
	fmt.Println()

	fmt.Println(githubInfo("georgealan"))
}

/*
url.PathEscape(login), is used in a case if a string contains any space or other character it doesn't break the url patch.
*/

// githubInfo returns name and number of public repositories for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	defer resp.Body.Close()

	var r struct { // anonymous struct
		Name         string
		Public_Repos int
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, fmt.Errorf("error: can't decode - %s", err)
	}
	return r.Name, r.Public_Repos, nil
}

type Reply struct {
	Name         string
	Public_Repos int
	//NumRepos int `json:"public_repos"` // Here you can use a different name that don't match from json file structure.
}

/*
JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
