package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func Loader() ([]Expense, error){
	if _, err := os.Stat(Filename); errors.Is(err, os.ErrNotExist){
		return []Expense{}, nil
	}

	data, err := os.ReadFile(Filename)
	if err != nil{
		return nil, err
	}

	var list []Expense
	err = json.Unmarshal(data, &list)
	if err != nil{
		return nil, err
	}
	return list, nil
}

func saveList(list []Expense) error{
	data, err := json.Marshal(list)
	if err != nil{
		log.Fatal(err)
	}
	return os.WriteFile(Filename, data, 0644)
}