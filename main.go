package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Filename = "expenses.json"

func main() {

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Print("use: expense-tracker <action> --description <text> --amount <value>")
			return
		}
		addcmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addcmd.String("description", "", "description of an expense")
		amount := addcmd.Float64("amount", 0, "amount of an expense")

		addcmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			log.Fatal("The flags --description and --amount are required, and amount must be > 0")
		}
		Add(*description, *amount)

	case "list":
		List()

	case "summary":
		if len(os.Args) == 3 {
			Summary()
		} else if len(os.Args) == 4 {
			var addcmd = flag.NewFlagSet("summary", flag.ExitOnError)
			month := addcmd.Int("month", 0, "month number (1-12)")

			addcmd.Parse(os.Args[2:])

			if *month <= 0 || *month > 12 {
				log.Fatal("The --month flag must be between 1 and 12")
			}
			SummaryParam(*month)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Print("use: expense-tracker delete --id <expense id>")
			return
		}
		addcmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := addcmd.Int64("id", 0, "ID of an expense")

		addcmd.Parse(os.Args[2:])

		Delete(*id)

	default:
		fmt.Println("Invalid command")
	}
}