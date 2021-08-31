package main

import (
	"fmt"
	"goWeb/DBChain"
	"goWeb/copyDirectory"
	"goWeb/testZip"
	"io"
	"net/http"
	"os"
	"text/template"
)

var dataBaseName string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!", r.URL.Path)
	fmt.Fprintln(w, "请求url user 的值为：", r.FormValue("user"))
	fmt.Fprintln(w, "请求体post username 的值为：", r.PostFormValue("username"))
}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "https://www.baidu.com")
	w.WriteHeader(302)
}

//func testMap(w http.ResponseWriter, r *http.Request) {
//	//myMap := MakeTestData()
//	myMap := make(map[string]interface{})
//	DataMaps :=GetDBData("student")
//	myMap["stus"] = DataMaps
//	t := template.Must(template.ParseFiles("index.html"))
//	t.Execute(w, myMap)
//}
func testMap(w http.ResponseWriter, r *http.Request) {
	myMap := make(map[string]interface{})

	arr := [5]int32{1, 2, 3, 4, 5}
	myMap["arr"] = arr
	t:= template.Must(template.New("test.vue").Delims("{[","]}").ParseFiles("test.vue"))
	os.RemoveAll("pages/test2.vue")
	file, err := os.OpenFile("pages/test2.vue", os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			return
		}
	t.Execute(file, myMap)
}


func testTableMap(w http.ResponseWriter, r *http.Request) {
	dataMap := DBChain.GetFinalMap(dataBaseName)
	t:= template.Must(template.New("testTableMap.vue").Delims("{[","]}").ParseFiles("testTableMap.vue"))
	os.RemoveAll("pages/testTableMap.vue")
	file, err := os.OpenFile("pages/testTableMap.vue", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(file, dataMap)
}

func ToDownload(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("ToDownload.html"))
	t.Execute(w, nil)
}

//func testWholeProcess(w http.ResponseWriter, r *http.Request) {
//	//get data from dbchain and add it to template
//	myMap := make(map[string]interface{})
//	DataMaps :=GetDBData("student")
//	myMap["stus"] = DataMaps
//	t := template.Must(template.ParseFiles("index.html"))
//
//    //write template to specific directory
//	os.RemoveAll("pages/index.html")
//	file, err := os.OpenFile("pages/index.html", os.O_WRONLY|os.O_CREATE, 0600)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	t.Execute(file, myMap)
//
//	//compress to zip
//	testZip.ToZip("pages","./1.zip")
//
//	//redirect
//	w.Header().Set("location", "http://localhost:8080/ToDownload")
//	w.WriteHeader(302)
//}
func testWholeProcess(w http.ResponseWriter, r *http.Request) {
	//copy template to pages
	copyDirectory.Dir("templates", "pages")
	//get tableStruct from DBChain
	dataJson := DBChain.GetFinalJson(dataBaseName)
	//write data to data.json
	os.RemoveAll("pages/data.json")
	file, err := os.OpenFile("pages/data.json", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte(dataJson))
	file.Close()

	//compress to zip
	testZip.ToZip("pages", "./code.zip")

	//redirect
	w.Header().Set("location", "http://localhost:8080/ToDownload")
	w.WriteHeader(302)
}

func download(res http.ResponseWriter, r *http.Request) {
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

func getDatabaseName(res http.ResponseWriter, req *http.Request) {
	receiveParam := req.URL.RawQuery
	fmt.Println(receiveParam)
	dataBaseName = receiveParam
}

//func GetDBData(tableName string) []map[string]string {
//	resp :=DBChain.QueryTableData(tableName)
//	DataMaps :=DBChain.ExtractData(resp)
//	//for _,v :=range DataMaps{
//	//	for k2,v2:=range v{
//	//		fmt.Println("submap key :",k2,",value :",v2)
//	//	}
//	//	fmt.Println("===================================")
//	//}
//	return DataMaps
//}

//func MakeTestData() map[string]interface{} {
//	myMap := make(map[string]interface{})
//	stuMap1 := make(map[string]string)
//	stuMap1["name"] = "zhangsan"
//	stuMap1["age"] = "11"
//	stuMap2 := make(map[string]string)
//	stuMap2["name"] = "lisi"
//	stuMap2["age"] = "18"
//	var stus []map[string]string
//	stus = append(stus, stuMap1)
//	stus = append(stus, stuMap2)
//	myMap["stus"] = stus
//	return myMap
//}
func main() {
	//http.HandleFunc("/hello", handler)
	//http.HandleFunc("/testRedirect", testRedirect)
	http.HandleFunc("/testTableMap", testTableMap)
	http.HandleFunc("/testMap", testMap)
	//http.HandleFunc("/testWholeProcess", testWholeProcess)
	http.HandleFunc("/ToDownload", ToDownload)
	http.HandleFunc("/download", download)
	http.HandleFunc("/transDatabaseName", getDatabaseName)
	http.HandleFunc("/testWholeProcess", testWholeProcess)
	http.ListenAndServe(":8080", nil)
}
