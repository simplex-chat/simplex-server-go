package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qri-io/jsonschema"
)

type Context struct {
	Req    *http.Request
	Params httprouter.Params
	Body   map[string]interface{}
	Resp   http.ResponseWriter
}

type apiHandler = func(http.ResponseWriter, *http.Request, httprouter.Params)

func marshal(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func handler(name string, handler func(Context)) apiHandler {
	resSchema := getResourceSchema(name)
	schema := jsonschema.RootSchema{}
	schema.UnmarshalJSON(marshal(resSchema["request"]["body"]))

	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var body map[string]interface{}
		dec := json.NewDecoder(req.Body)
		if err := dec.Decode(&body); err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad Request")
			return
		}

		errors := []jsonschema.ValError{}
		schema.Validate("", body, &errors)

		if len(errors) > 0 {
			log.Println("Invalid data:", errors)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad Request")
			return
		}

		log.Println(body["recipient"])

		handler(Context{Resp: w, Req: req, Params: params, Body: body})
	}
}
