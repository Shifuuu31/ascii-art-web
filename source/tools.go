package source

import (
	"log"
	"os"
)

func ReadTemplate(banner string) string {
	file, err := os.ReadFile(banner)
	CheckError(err)
	return string(file[1:])
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
