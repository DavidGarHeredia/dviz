package reader

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJson(file string) map[string]json.RawMessage {
	reader := getJsonReader(file)
	return getJsonContent(reader)
}

func getJsonReader(file string) io.Reader {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	return jsonFile
}

func getJsonContent(reader io.Reader) map[string]json.RawMessage {
	byteValue, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	var result map[string]json.RawMessage
	json.Unmarshal([]byte(byteValue), &result)
	return result
}
