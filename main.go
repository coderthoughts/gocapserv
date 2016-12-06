package main

import (
	"fmt"
	"net/http"
	"os"
)

var requests []string

func main() {
	port := "8888"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Println("Listening on port: " + port)

	http.HandleFunc("/", rootPage)
	http.HandleFunc("/record/", fooPage)
	http.ListenAndServe(":"+port, nil)

}

func rootPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Serving root")
	request.ParseForm()
	fmt.Fprintln(writer, `<h1>Requests</h1><body><ul>`)
	for _, req := range requests {
		fmt.Fprintln(writer, `<li>`+req)
	}
	fmt.Fprintln(writer, `</ul></body>`)
}

func fooPage(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	requests = append(requests, request.URL.RequestURI())
	fmt.Println("Recorded: " + request.URL.RequestURI())
	fmt.Fprintln(writer, `<h1>Recorded!</h1><body>`+request.URL.RequestURI()+`</body>`)
}
