package context

import (
	"io/ioutil"
	"log"
	"os"
)

func ListContext() {
	home := os.Getenv("HOME")
	if home == "" {
		log.Println("error while reading environment variable")
		return
	}
	if _, err := os.Stat(home + "/.kcm"); os.IsNotExist(err) {
		log.Println("kcm home not available : ", err)
		return
	}
	dirs, err := ioutil.ReadDir(home + "/.kcm")
	if err != nil {
		log.Println("error while listing context : ", err)
		return
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			log.Println(dir.Name())
		}
	}

}
