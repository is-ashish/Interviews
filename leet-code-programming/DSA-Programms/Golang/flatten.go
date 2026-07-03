package main

import (
	"encoding/json"
	"fmt"
)

// inputJSON := `{
// 	"name": "John Doe",
// 	"email": "jd@email.io",
// 	"password": "xxxxx",
// 	"phone_number": "+91907xxxxxx",
// 	"test": {
// 		"id": "1234",
// 		"tel": {
// 			"tags": ["hello", "world"]
// 		}
// 	}
// }`

// flatten recursively flattens a nested JSON structure into a flat map.
// Keys are concatenated with underscores for nested objects.
// Array elements are appended with their index.
func flatten(data interface{}, prefix string) map[string]interface{} {

	result := make(map[string]interface{})

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newKey := key
			if prefix != "" {
				newKey = prefix + "_" + key
			}
			if value != nil {
				flattened := flatten(value, newKey)
				for k, val := range flattened {
					result[k] = val
				}
			}
		}
	case []interface{}:
		for index, item := range v {
			newKey := fmt.Sprintf("%s_%d", prefix, index)
			if item != nil {
				flattened := flatten(item, newKey)
				for k, val := range flattened {
					result[k] = val
				}
			}
		}
	default:
		if prefix != "" {
			result[prefix] = v
		}
	}
	fmt.Println(result)
	return result
}

func main_() {
	inputJSON := `{
		"name": "John Doe",
		"email": "jd@email.io",
		"password": "xxxxx",
		"phone_number": "+91907xxxxxx",
		"test": {
			"id": "1234",
			"tel": {
				"tags": ["hello", "world"]
			}
		}
	}`

	var data interface{}
	err := json.Unmarshal([]byte(inputJSON), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	flattened := flatten(data, "")

	outputJSON, err := json.MarshalIndent(flattened, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(outputJSON))
}
