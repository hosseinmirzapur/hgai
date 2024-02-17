package utils

import "log"

func HandleError(err error) {
	log.Printf("error happened: %+v\n", err)
}
