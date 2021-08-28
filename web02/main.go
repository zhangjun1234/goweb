package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("pages/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := template.Must(template.ParseFiles("web02/index.html"))
	t.Execute(file, " my data")

}
func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8081", nil)
}
