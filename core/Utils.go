package core

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"os"
	"path/filepath"
)

func Console(message ...string) {
	if len(message) < 2 {
		fmt.Println(fmt.Sprintf("[Core]: %s", message))
		return
	}
	fmt.Println(fmt.Sprintf("[%s]: %s", message[1], message[0]))
	return
}

// Checkers&Matchers

func CheckForFatal(err ...interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error, " + err.Error())
		os.Exit(1)
	}
}

func CheckErrorNotPanic(err error) {
	if err != nil {
		Console(err.Error(), "ERROR")
	}
}

// CheckKeyExists checks if map has the key or not.
func CheckKeyExists(key string, toCheck map[string]interface{}) bool {
	if toCheck[key] != nil {
		return true
	}
	return false
}

// ExpectResult Expect result from given string.
func ExpectResult(expectation string, result interface{}) bool {
	if expectation == result {
		return true
	}
	return false
}

// String Parsing

// PosString get string position
func PosString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// ContainsString returns true if slice contains element
func ContainsString(slice []string, element string) bool {
	return !(PosString(slice, element) == -1)
}

// File IO

// RemoveDirectory removes directory with contents
func RemoveDirectory(path string) error {
	d, err := os.Open(path)
	CheckError(err)
	defer func(d *os.File) {
		err := d.Close()
		CheckError(err)
	}(d)
	names, err := d.Readdirnames(-1)
	CheckError(err)
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		CheckError(err)
	}
	err = os.Remove(path)
	CheckError(err)
	return nil
}

// FileExists checks for file or directory existence
func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}

func GetUser(session *sessions.Session) UserData {
	val := session.Values["user"]
	var user = UserData{}
	user, ok := val.(UserData)
	if !ok {
		return UserData{Authenticated: false}
	}
	return user
}
