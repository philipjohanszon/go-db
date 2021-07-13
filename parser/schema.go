package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Schema struct {
	schemaString string
}

/*
	SCHEMA FORMAT (SAVED TO FILE)

	//user table

	users::firstName|string,lastName|string,email|string,family|object,family.siblings|number,family.parents|number

*/

func (schema *Schema) Parse(schemaJSON *[]byte) error {
	var decodedJSON map[string]interface{}

	json.Unmarshal(*schemaJSON, &decodedJSON)

	//check for keys inside keys

	keysError := schema.getSchemaKeys("", decodedJSON)

	if keysError != nil {
		fmt.Println(keysError.Error())
		return keysError
	}

	schema.trimComma()

	return nil
}

//Goes recursively through all the keys and if they have the value of map then they will be looped through recursively
//if the key has a value except map then it will be added to the schemaString
func (schema *Schema) getSchemaKeys(prefix string, decodedJSON map[string]interface{}) error {
	keys := reflect.ValueOf(decodedJSON).MapKeys()

	//Goes through keys
	for index, _ := range keys {

		currentKey := keys[index].Interface().(string)

		//checks if the value in the key is an index
		if isMap(decodedJSON[currentKey]) {
			schema.addToSchemaString(prefix+currentKey, "map")

			//selecting only the map that is inside the map so that it can run recursively
			cutMap, ok := decodedJSON[currentKey].(map[string]interface{})

			if !ok {
				return fmt.Errorf("Error converting interface to map: " + fmt.Sprint(decodedJSON[currentKey]))
			}

			keysError := schema.getSchemaKeys(prefix+currentKey+".", cutMap)

			if keysError != nil {
				return keysError
			}

		} else {
			//Gets variable type in string format
			varType := reflect.TypeOf(decodedJSON[currentKey]).String()

			schema.addToSchemaString(prefix+currentKey, varType)
		}
	}

	return nil
}

func (schema *Schema) addToSchemaString(key string, varType string) {
	schema.schemaString += key + "|" + varType + ","
}

//checks if variable is map, will be used to check if an object has a object inside when parsin json schema
func isMap(x interface{}) bool {
	t := fmt.Sprintf("%T", x)
	return strings.HasPrefix(t, "map[")
}

//sometimes the schemastring will get a comma at the end, this removes it
func (schema *Schema) trimComma() {
	lastChar := schema.schemaString[len(schema.schemaString)-1:]

	if lastChar == "," {
		schema.schemaString = schema.schemaString[:len(schema.schemaString)-1]
	}
}

func (schema *Schema) Get() string {
	return schema.schemaString
}
