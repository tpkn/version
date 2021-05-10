package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func FileExists(file_path string) bool {
	i, err := os.Stat(file_path)
	if os.IsNotExist(err) || i.IsDir() {
		return false
	}
	return true
}

func ReadFile(file_path string) string {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Println(err)
	}
	
	return string(data)
}

func WriteFile(file_path string, content string){
	f, err := os.Create(file_path)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
}




