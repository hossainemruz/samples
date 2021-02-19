package json

import (
	"encoding/json"
	"strings"
)

type Data struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func ReadJsonFromArbitraryString(input string, output interface{}) error {
	return json.NewDecoder(strings.NewReader(input)).Decode(output)
}
