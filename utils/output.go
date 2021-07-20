package utils

import "fmt"

// printKey prints a desired key and its values based off of the inferred type
func printKey(id string, m interface{}) {
	fmt.Printf("%s \n\n", id)
	switch m.(type) {
	case string:
		fmt.Printf("\t%s\n", m)
	case []string:
		for _, value := range m.([]string) {
			fmt.Printf("\t%s\n", value)
		}
	case map[string]string:
		for key, value := range m.(map[string]string) {
			fmt.Printf("\t%s: %s\n", key, value)
		}
	}
	fmt.Println()
}