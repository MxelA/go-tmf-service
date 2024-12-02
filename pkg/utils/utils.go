package utils

import (
	"encoding/json"
	"fmt"
)

var PrettyPrint = func(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(j))
}
