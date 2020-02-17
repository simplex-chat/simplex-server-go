package api

import (
	"encoding/json"
	"log"
	"os"
)

type handlerSchema map[string]map[string]map[string]interface{}

func getResourceSchema(name string) handlerSchema {
	file, _ := os.Open("./api/schema/" + name + ".json")
	defer file.Close()
	var schemaObj handlerSchema
	dec := json.NewDecoder(file)
	if err := dec.Decode(&schemaObj); err != nil {
		log.Fatal(err)
	}

	jtdToJsonSchema(schemaObj, "request", "params")
	jtdToJsonSchema(schemaObj, "request", "body")
	jtdToJsonSchema(schemaObj, "request", "qs")
	jtdToJsonSchema(schemaObj, "response", "body")

	return schemaObj
}

func jtdToJsonSchema(schemaObj handlerSchema, reqRes string, loc string) {
	schema := schemaObj[reqRes][loc]
	if schema != nil {
		schema["type"] = "object"
		schema["additionalProperties"] = false
		keys := getKeys(schema["properties"])
		if keys != nil {
			schema["required"] = keys
		}
	}
}

func getKeys(object interface{}) []string {
	props, ok := object.(map[string]interface{})
	if ok {
		keys := make([]string, 0, len(props))
		for k := range props {
			keys = append(keys, k)
		}
		return keys
	}
	return nil
}
