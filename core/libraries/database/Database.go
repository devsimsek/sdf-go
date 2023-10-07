package database

import (
	"SDF/core"
	"encoding/base64"
	"encoding/json"
	"os"
  "io"
	"path/filepath"
)

type Database interface {
    Open(fileName string, path ...string) bool
    Save() bool
    Set(key string, value interface{})
    Get(key string) interface{}
    Update(key string, value interface{})
    Delete(key string)
    Create(fileName string, path ...string) error
    Read() map[string]interface{}
}

type Db struct {
	Dir      string
	File     string
	Contents map[string]interface{}
}

func Initialize() *Db {
  return &Db{}
}

func (tempDatabase *Db) Open(fileName string, path ...string) bool {
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

func (tempDatabase *Db) Save() bool {
	content, err := json.Marshal(tempDatabase.Contents)
	core.CheckError(err)
	content = []byte(base64.StdEncoding.EncodeToString(content))
	err = os.WriteFile(tempDatabase.Dir+"/"+tempDatabase.File, content, 0644)
	core.CheckError(err)
	return true
}

func (tempDatabase *Db) Set(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func (tempDatabase *Db) Get(key string) interface{} {
	return tempDatabase.Contents[key]
}

func (tempDatabase *Db) Update(key string, value interface{}) {
	tempDatabase.Contents[key] = value
}

func (tempDatabase *Db) Delete(key string) {
	delete(tempDatabase.Contents, key)
}

func (tempDatabase *Db) Create(fileName string, path ...string) (*os.File, error) {
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

func (tempDatabase *Db) Read() map[string]interface{} {
	return tempDatabase.Contents
}
