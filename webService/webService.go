package webService

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func webService() {
	// url parse
	var uri = "https://www.google.com/search?q=golang&rlz=1C1GCEA_enID937ID937&oq=golang&aqs=chrome..69i57j69i59l3j46j0l5.1322j0j15&sourceid=chrome&ie=UTF-8"

	url, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(url.Query())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template.html")

		var data = map[string]string{
			"Name":    "Impact Byte",
			"Message": "Welcome to web service go",
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template2.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, nil)

	})

	port := "localhost:8080"

	fmt.Println("service running on port", port)

	http.ListenAndServe(port, nil)
}
