package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Add(description string, amount float64) {
	if len(os.Args) < 5 {
		fmt.Print("provide the arguments: expense-tracker <action> --description <text> --amount <value>")
		return
	}

	lists, err := Loader()
	if err != nil {
		log.Fatal(err)
	}

	var pk int64 = 1

	if len(lists) > 0 {
		pk = int64(lists[len(lists)-1].ID) + 1
	}

	var list = Expense{
		ID:          pk,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	lists = append(lists, list)

	err = saveList(lists)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Expense added successfully (ID: %v)", pk)
}

func List() {
	lists, err := Loader()
	if err != nil {
		log.Fatal(err)
	}
	if lists == nil {
		fmt.Print("The list is empty")
		return
	}
	fmt.Printf("ID\tDate\t\tDescription\tAmount\n")

	for _, item := range lists {
		fmt.Printf("%d\t%s\t%s\t\t$%.2f\n",
			item.ID,
			item.CreatedAt.Format("2006-01-02"),
			item.Description,
			item.Amount)
	}
}

func Summary() {
	lists, err := Loader()
	if err != nil {
		log.Fatal(err)
	}
	if lists == nil {
		fmt.Print("Total expenses: $0")
		return
	}

	var total float64 = 0

	for _, item := range lists {
		total += item.Amount
	}

	fmt.Printf("Total expenses: $%v", total)
}

func Delete(id int64) {
	lists, err := Loader()

	if err != nil {
		log.Fatal(err)
	}

	if lists == nil {
		fmt.Print("The list is empty")
		return
	}

	var Newlist []Expense

	for _, item := range lists {
		if item.ID == id {
			continue
		}
		Newlist = append(Newlist, item)
	}

	saveList(Newlist)
	fmt.Print("Expense deleted successfully")
}

func SummaryParam(month int) {
	lists, err := Loader()

	if err != nil {
		log.Fatal(err)
	}

	if lists == nil {
		fmt.Print("The list is empty")
		return
	}

	var total float64 = 0
	var monthName string = ""

	for _, item := range lists {
		if int(item.CreatedAt.Month()) == month {
			total += item.Amount
			monthName = item.CreatedAt.Month().String()
		}
	}

	fmt.Printf("Total expenses for month %v: $%.2f\n", monthName, total)
}