package link

import (
	"encoding/json"
	"io/ioutil"
)

type fileProvider struct {
}

func (fileProvider) Get(fileName string, v interface{}) error {
	content, e := ioutil.ReadFile(fileName)
	if e != nil {
		return e
	}
	var err = json.Unmarshal(content, &v)
	return err
}

func (fileProvider) Save(fileName string, v interface{}) error {
	jsonString, e := json.Marshal(v)
	if e != nil {
		return e
	}
	e = ioutil.WriteFile(fileName, jsonString, 0644)
	return e
}
