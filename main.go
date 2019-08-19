package main

import (
	"encoding/json"
	"fmt"
)

func ConvertToHCL(policies map[string][]string) (string, error) {
	str := ""
	for key, value := range policies {
		capabilities, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		str += fmt.Sprintf(`path "%s" {
   capabilities = %s
}
`, key, string(capabilities))
	}
	return str, nil
}

func main() {
	//policies := map[string]map[string]map[string][]string{
	//	"path": map[string]map[string][]string{
	//		"secret/foo": map[string][]string{
	//			"capabilities": []string{"read", "list"},
	//		},
	//		"secret/bar/*": map[string][]string{
	//			"capabilities": []string{"list", "update", "read"},
	//		},
	//	},
	//}
	//value, err := json.Marshal(policies)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//value2, err := json.Marshal(runtime.RawExtension{Raw: value})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(value2))

	policies := map[string][]string{
		"secret/bar/*": []string{"list", "update", "read"},
		"secret/foo":   []string{"read", "list"},
	}

	doc, err := ConvertToHCL(policies)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc)
}
