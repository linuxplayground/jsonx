package main

import (
	//"encoding/json"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	"net/http"

	"fmt"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var response string

	response = `<?xml version="1.0" encoding="UTF-8"?>
<json:object xsi:schemaLocation="http://www.datapower.com/schemas/json jsonx.xsd"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:json="http://www.ibm.com/xmlns/prod/2009/jsonx">`

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	children, _ := jsonParsed.ChildrenMap()
	for key, child := range children {
		response = fmt.Sprintf("%s\n<json:string name=\"%s\">%s</json:string>", response, key, child)
	}

	response = fmt.Sprintf("%s\n</json:object>\n", response)

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, response)
}
