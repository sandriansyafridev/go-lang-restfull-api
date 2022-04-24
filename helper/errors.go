package helper

import "log"

func FatalIfError(message string, err error) {
	if err != nil {
		log.Fatal(message+": ", err.Error())
	}
}
