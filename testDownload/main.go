package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(res http.ResponseWriter, r *http.Request) {
	fileName := "1.zip"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Header().Add("Content-Type", "application/octet-stream")
	res.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")
	_, err = io.Copy(res, file)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
}
func main() {
	http.HandleFunc("/testDownload", Download)
	http.ListenAndServe(":8080", nil)
}
