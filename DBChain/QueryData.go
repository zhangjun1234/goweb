package DBChain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	databaseName = "HJM9XONGDY"
	baseUrl      = "https://controlpanel.dbchain.cloud/relay/dbchain/"
)

func QueryTables(token string, databaseCode string) []string {
	QueryDataUrl := baseUrl + "/tables/" + token + "/" + databaseCode
	resp := Get(QueryDataUrl)
	tmpMap := JSONToMap(resp)
	resultsInterfaces := tmpMap["result"].([]interface{})
	var tablesSlice []string
	for _, v := range resultsInterfaces {
		//fmt.Println("value :",v)
		tablesSlice = append(tablesSlice, v.(string))
	}
	return tablesSlice
}

func JSONToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func QueryTableColOpts(token string, databaseCode string, tableName string, field string) []string {
	QueryColOptsUrl := baseUrl + "/column-options/" + token + "/" + databaseCode + "/" + tableName + "/" + field
	resp := Get(QueryColOptsUrl)
	tmpMap := JSONToMap(resp)
	tmpinterfaces := tmpMap["result"].([]interface{})
	var colOptsArray []string
	for _, v := range tmpinterfaces {
		colOptsArray = append(colOptsArray, v.(string))
	}
	return colOptsArray
}

func QueryTableColDataType(token string, databaseCode string, tableName string, field string) string {
	QueryColOptsUrl := baseUrl + "/column-data-type/" + token + "/" + databaseCode + "/" + tableName + "/" + field
	resp := Get(QueryColOptsUrl)
	tmpMap := JSONToMap(resp)
	return tmpMap["result"].(string)
}

func QueryTableFields(token string, databaseCode string, tableName string) []string {
	QueryDataUrl := baseUrl + "/tables/" + token + "/" + databaseCode + "/" + tableName
	resp := Get(QueryDataUrl)
	return ExtractFieldsSlice(resp)
}

func ExtractFieldsSlice(resp string) []string {
	myMap := JSONToMap(resp)
	resultinterface := myMap["result"]
	resultMap, _ := resultinterface.(map[string]interface{})
	fieldsInterfaces := resultMap["fields"].([]interface{})
	var fieldSlices []string
	for _, v := range fieldsInterfaces {
		fieldSlices = append(fieldSlices, v.(string))
	}
	return fieldSlices
}

func GetFinalJson(databaseCode string) string {
	token := MakeAccessCode()
	tablesName := QueryTables(token, databaseCode)
	bigMap := make(map[string]interface{})
	var tablesMap []map[string]interface{}

	for _, tableName := range tablesName {
		tableMap := make(map[string]interface{})
		fields := QueryTableFields(token, databaseCode, tableName)
		//tableMap := make(map[string]interface{})
		var fieldMap []map[string]interface{}
		colMap := make(map[string]interface{})

		for _, field := range fields {
			colMap["name"] = field
			colopts := QueryTableColOpts(token, databaseCode, tableName, field)
			dataType := QueryTableColDataType(token, databaseCode, tableName, field)
			if len(colopts) > 0 {
				colMap["propertyArr"] = colopts
			}
			if len(dataType) > 0 {
				colMap["fieldType"] = dataType
			}
			fieldMap = append(fieldMap, colMap)
		}
		tableMap["name"] = tableName
		tableMap["field"] = fieldMap

		tablesMap = append(tablesMap, tableMap)
	}
	bigMap["table"] = tablesMap
	return MapToJson(bigMap)
}

func MapToJson(myMap map[string]interface{}) string {
	bytes, err := json.Marshal(myMap)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
