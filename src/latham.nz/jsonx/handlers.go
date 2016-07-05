package main

import (
	//"encoding/json"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	"net/http"
  "strings"
	"fmt"
)

const xmlHeader string = `<?xml version=\"1.0\" encoding=\"UTF-8\"?>`
const jsonObjectSchemaTag string = `<json:object xsi:schemaLocation="http://www.datapower.com/schemas/json jsonx.xsd"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:json="http://www.ibm.com/xmlns/prod/2009/jsonx">`
const jsonObjectOpenTag string = `<json:object name=\"%s\">`
const jsonObjctCloseTag string = `</json:object>`
const jsonObjectParameter string = `<json:%T name=\"%s\">%s</json:%T>`

func CreateJsonX(body []byte) string {
	var response string
	response += fmt.Sprintln(xmlHeader)
	response += fmt.Sprintln(jsonObjectSchemaTag)

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	children, _ := jsonParsed.ChildrenMap()
	for key, child := range children {
		fmt.Sprintf("%s\n<json:string name=\"%s\">%s</json:string>", response, key, child)
	}

	response += fmt.Sprintln(jsonObjctCloseTag)
	return response
}

func WalkJsonObject(children string) string {
	var child string


	return child
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
