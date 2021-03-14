package internal

import "fmt"

// Output ...
func Output(secretValues map[string]interface{}) {
	for k, v := range secretValues {
		fmt.Printf("%s=%s\n", k, v)
	}
}
