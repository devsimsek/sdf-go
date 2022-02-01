package database

import (
	"SDF/core"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Database struct {
	Dir      string
	File     string
	Contents map[string]interface{}
}

var tempDatabase Database

func Open(fileName string, path ...string) bool {
	var dir string
	if len(path) < 1 {
		directory, err := os.Getwd()
		core.CheckError(err)
		dir = directory
	} else {
		dir = path[0]
	}
	// Open database file
	file, _ := os.Open(dir + "/" + fileName)
	defer func(file *os.File) {
		err := file.Close()
		core.CheckError(err)
	}(file)
	tempDatabase.Dir = dir
	tempDatabase.File = fileName
	// Parse contents of database
	contents, _ := ioutil.ReadAll(file)
	contents, err := base64.StdEncoding.DecodeString(string(contents))
	err = json.Unmarshal(contents, &tempDatabase.Contents)
	core.CheckError(err)
	if tempDatabase.Contents != nil {
		return true
	}
	return false
}

func Save() bool {
	content, err := json.Marshal(tempDatabase.Contents)
	core.CheckError(err)
	content = []byte(base64.StdEncoding.EncodeToString(content))
	err = os.WriteFile(tempDatabase.Dir+"/"+tempDatabase.File, content, 0644)
	core.CheckError(err)
	return true
}

func Set(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func Get(key string) interface{} {
	return tempDatabase.Contents[key]
}

func Update(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func Delete(key string) {
	delete(tempDatabase.Contents, key)
}

func Create(fileName string, path ...string) (*os.File, error) {
	var dir string
	if len(path) < 1 {
		directory, err := os.Getwd()
		core.CheckError(err)
		dir = directory
	} else {
		dir = path[0]
	}
	if err := os.MkdirAll(filepath.Dir(dir+"/"+fileName), 0770); err != nil {
		return nil, err
	}
	file, err := os.Create(dir + "/" + fileName)
	core.CheckError(err)
	_, err = file.WriteString(base64.StdEncoding.EncodeToString([]byte("{}")))
	core.CheckError(err)
	// Open database
	Open(fileName, dir)
	return file, nil
}

func Read() map[string]interface{} {
	return tempDatabase.Contents
}
