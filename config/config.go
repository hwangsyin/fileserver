package config

import (
	"os"
	"path/filepath"
	"errors"
	"io/ioutil"
	"encoding/xml"
)

var (
	Configs = make(map[string]string, 16)
)

type property struct {
	XMLName xml.Name `xml:"property"`
	Name string `xml:"name"`
	Value string `xml:"value"`
}

type configuration struct {
	XMLName xml.Name `xml:"configuration"`
	Properties []property `xml:"property"`
}

func Parse() error {
	home, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	
	configFileName := filepath.Join(home, "conf", "server.xml")
	fileInfo, err := os.Stat(configFileName)
	if os.IsNotExist(err) {
		return errors.New(configFileName + " not exists")
	}
	
	if !fileInfo.Mode().IsRegular() {
		return errors.New(configFileName + " is not regular file")
	}
	
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}
	
	var result configuration
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return err
	}
	
	for _, p := range result.Properties {
		Configs[p.Name] = p.Value
	}
	
	return nil
}