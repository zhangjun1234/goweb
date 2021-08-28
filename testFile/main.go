package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("pages/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//str := "hello"
	content := []byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<body>\n<form action=\"http://localhost:8080/hello?user=admin&pwd=123456\" method=\"POST\">\n    用 户 名 ： <input type=\"text\" name=\"username\" ><br/>\n    密 码 ： <input type=\"password\" name=\"password\" ><br/>\n    <input type=\"submit\">\n</form>\n</body>\n</html>")
	if _, err = file.Write(content); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write file successful")

}
