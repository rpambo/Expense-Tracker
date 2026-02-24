# Expense Tracker CLI
A simple command-line Expense Tracker application that helps you manage your finances.
You can add, list, delete, and summarize expenses using straightforward CLI commands.

This project is based on the challenge from:

https://roadmap.sh/projects/expense-tracker

Show the total amount for a specific month of the current year.

$ go run . add --description "Lunch" --amount 20
    Expense added successfully (ID: 1)

$ go run . add --description "Dinner" --amount 10
    Expense added successfully (ID: 2)

$ go run . list
    ID   Date         Description   Amount
    1    2024-08-06   Lunch         $20
    2    2024-08-06   Dinner        $10

$ go run . summary
    Total expenses: $30

$ go run . delete --id 2
    Expense deleted successfully

$ go run . summary
    Total expenses: $20

$ go run . summary --month 8
    Total expenses for August: $20

# File Storage

All expenses are stored locally in a JSON file:
    expenses.json

If the file does not exist, it is created automatically.