package utils

import (
	"log"
)

// printKey prints a desired key and its values based off of the inferred type
func printKey(id string, m interface{}) {
	log.Printf("%s \t➝️", id)
	switch m.(type) {
	case string:
		log.Printf("\t%s", m)
	case []string:
		for _, value := range m.([]string) {
			log.Printf("\t%s ", value)
		}
	case map[string]string:
		for key, value := range m.(map[string]string) {
			log.Printf("\t%s: %s ", key, value)
		}
	}
	log.Printf("\n")
}
