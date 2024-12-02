package utils

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(j))
}
