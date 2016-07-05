# jsonx
A completely trivial project based on this garbage: https://www.ibm.com/support/knowledgecenter/SS9H2Y_7.5.0/com.ibm.dp.doc/json_jsonxconversionexample.html

# Install
`go get "github.com/Jeffail/gabs"`

# Testing

Compile dependancy
`go build latham.nz/featly.common`

Compile app
`go install latham.nz/jsonx`


Then run the app (it listens on port 8080)
`$GOPATH/bin/jsonx`

Then execute some kind of post as follows in a different terminal.

`curl -d '{"first_name":"David", "last_name":"Latham"}' http://localhost:8080`
