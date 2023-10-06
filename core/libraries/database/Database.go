package database

import (
	"SDF/core"
	"encoding/base64"
	"encoding/json"
	"os"
  "io"
	"path/filepath"
)

type Database struct {
	Dir      string
	File     string
	Contents map[string]interface{}
}

func (tempDatabase *Database) Open(fileName string, path ...string) bool {
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
	contents, _ := io.ReadAll(file)
	contents, err := base64.StdEncoding.DecodeString(string(contents))
	err = json.Unmarshal(contents, &tempDatabase.Contents)
	core.CheckError(err)
	if tempDatabase.Contents != nil {
		return true
	}
	return false
}

func (tempDatabase *Database) Save() bool {
	content, err := json.Marshal(tempDatabase.Contents)
	core.CheckError(err)
	content = []byte(base64.StdEncoding.EncodeToString(content))
	err = os.WriteFile(tempDatabase.Dir+"/"+tempDatabase.File, content, 0644)
	core.CheckError(err)
	return true
}

func (tempDatabase *Database) Set(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func (tempDatabase *Database) Get(key string) interface{} {
	return tempDatabase.Contents[key]
}

func (tempDatabase *Database) Update(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func (tempDatabase *Database) Delete(key string) {
	delete(tempDatabase.Contents, key)
}

func (tempDatabase *Database) Create(fileName string, path ...string) (*os.File, error) {
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
	return file, nil
}

func (tempDatabase *Database) Read() map[string]interface{} {
	return tempDatabase.Contents
}
