package kcmcontext

import (
	"io/ioutil"
	"log"
	"os"
)

func ListContext() {
	if _, err := os.Stat("/home/shovan/.kcm"); os.IsNotExist(err) {
		log.Println("kcm home not available : ", err)
		return
	}
	dirs, err := ioutil.ReadDir("/home/shovan/.kcm")
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
