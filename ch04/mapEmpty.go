package main

import (
	"encoding/json"
	"fmt"
)

var JSONrecord = `{
	"Flag": true,
	"Array": ["a", "b", "c"],
	"Entiry": {
		"a1": "b1",
		"a2": "b2",
		"a3": "b3",
		"Value": -456,
		"Null": null
	},
	"Message": "Hello Go!"
}`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Its a string:", k, c)
		case float64:
			fmt.Println("Its a float64:", k, c)
		case bool:
			fmt.Println("Its a boolean:", k, c)
		case map[string]interface{}:
			fmt.Println("Its a map!")
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("...Is %v: %T!\n", k, c)
		}
	}
}

func exploreMap(m map[string]interface{}) {
	for k, v := range m {
	  embMap, ok := v.(map[string]interface{})
		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}
func main() {
	amap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &amap)
	if err != nil {
		fmt.Println(err)
		return
	}

	typeSwitch(amap)
	fmt.Println("-------------")
	exploreMap(amap)
}


