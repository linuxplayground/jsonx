package main

import (
	//"encoding/json"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	"net/http"
  "strings"
	"fmt"
	"strconv"
)

const xmlHeader string = `<?xml version=\"1.0\" encoding=\"UTF-8\"?>`
const jsonObjectSchemaTag string = `<json:object xsi:schemaLocation="http://www.datapower.com/schemas/json jsonx.xsd"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:json="http://www.ibm.com/xmlns/prod/2009/jsonx">`
const jsonObjectOpenTag string = `<json:object name="%s">`
const jsonObjctCloseTag string = `</json:object>`
const jsonArrayOpenTag string = `<json:array name="%s">`
const jsonArrayCloseTag string = `</json:array>`
const jsonObjectParameter string = `<json:%s%s>%s</json:%s>`
const jsonNullValueParameter string = `<json:null name="%s" />`
const jsonKeyName string = ` name="%s"`

func CreateJsonX(body []byte) string {
	var response string
	response += fmt.Sprintln(xmlHeader)
	response += fmt.Sprintln(jsonObjectSchemaTag)

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	children, _ := jsonParsed.ChildrenMap()
	response += WalkJsonObject(children, 1)
	// for key, child := range children {
	// 	fmt.Sprintf("%s\n<json:string name=\"%s\">%s</json:string>", response, key, child)
	// }

	response += fmt.Sprintln(jsonObjctCloseTag)
	return response
}

func WalkJsonObject(children map[string]*gabs.Container, depth int) string {
	var response string
	for i := 0; i < depth; i++ {
		response += fmt.Sprintf("\t")
	}
	for key, child := range children {
		sChild := child.String()
		if strings.HasPrefix(sChild, "{}") {
			response += fmt.Sprintf(jsonNullValueParameter, key)
		} else if strings.HasPrefix(sChild,"{") {
			response += fmt.Sprintf(jsonObjectOpenTag, key)
			x, err := gabs.ParseJSON(child.Bytes())
			if err != nil {
				panic(err)
			}
			newChild, _ := x.ChildrenMap()
			response += WalkJsonObject(newChild, depth+1)
			response += fmt.Sprintf(jsonObjctCloseTag)
		} else if strings.HasPrefix(sChild,"[") {
			response += fmt.Sprintf(jsonArrayOpenTag, key)
			count, err := child.ArrayCount()
			if err != nil {
				panic(err)
			}
			for i := 0; i < count; i++ {
					arrayElement, err := child.ArrayElement(i)
					if err != nil {
						panic(err)
					}
					response += FormatObjectToXml("", arrayElement.String())
			}
			response += fmt.Sprintf(jsonArrayCloseTag)
		} else {
			response += FormatObjectToXml(key, sChild)
		}
		response += fmt.Sprintln()
	}
	return response
}

func _WalkJsonObject(child *gabs.Container, depth int) {

}

func FormatObjectToXml(key, sChild string) string {
	if key != "" {
		key = fmt.Sprintf(jsonKeyName, key)
	}
	if IsNumber(sChild) {
		return fmt.Sprintf(jsonObjectParameter, "number",key,sChild,"number")
	} else if IsBoolean(sChild) {
		return fmt.Sprintf(jsonObjectParameter, "boolean",key,sChild,"boolean")
	} else {
		return fmt.Sprintf(jsonObjectParameter, "string",key,sChild,"string")
	}
}

func IsNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return true
}

func IsBoolean(value string) bool {
	_, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return true
}

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	response := CreateJsonX(body)

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, response)
}
