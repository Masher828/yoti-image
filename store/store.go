package store

import (
	"encoding/json"
	"log"
	"os"
	"yoti/constants"
)

var dataStore map[string]string = make(map[string]string)

func Get(key string) (string, error) {
	value, isPresent := dataStore[key]
	if !isPresent {
		return "", constants.InvalidKey
	}

	return value, nil
}

func Add(key, value string) error {

	dataStore[key] = value
	return nil

}

func Delete(key string) error {
	if _, ok := dataStore[key]; !ok {
		return constants.InvalidKey
	}
	delete(dataStore, key)
	return nil
}

func LoadStore() error {
	data, err := os.ReadFile("store.txt")
	if err != nil {
		if err != os.ErrNotExist {
			log.Println(err.Error())
		}
		return err
	}

	err = json.Unmarshal(data, &dataStore)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func StoreInFile() error {
	f, err := os.Create("store.txt")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer f.Close()

	data, err := json.Marshal(dataStore)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
