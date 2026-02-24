# Expense Tracker CLI

A simple command-line Expense Tracker application that helps you manage your finances. You can add, list, delete, and summarize expenses using straightforward CLI commands.

This project is based on the challenge from:
https://roadmap.sh/projects/expense-tracker

## Features

- Add expenses with description and amount
- List all expenses
- Delete expenses by ID
- View total summary of all expenses
- View summary for a specific month of the current year

## Usage

### Add an expense

go run . add --description "Lunch" --amount 20

Output:
Expense added successfully (ID: 1)

go run . add --description "Dinner" --amount 10

Output:
Expense added successfully (ID: 2)

### List all expenses

go run . list

Output:
ID   Date         Description   Amount
1    2024-08-06   Lunch         $20
2    2024-08-06   Dinner        $10

### View total summary

go run . summary

Output:
Total expenses: $30

### Delete an expense

go run . delete --id 2

Output:
Expense deleted successfully

### View summary after deletion

go run . summary

Output:
Total expenses: $20

### View summary for a specific month

go run . summary --month 8

Output:
Total expenses for August: $20

## File Storage

All expenses are stored locally in a JSON file:

expenses.json

If the file does not exist, it is created automatically.

## Commands Summary

| Command | Description | Example |
|---------|-------------|---------|
| add | Add a new expense | go run . add --description "Lunch" --amount 20 |
| list | List all expenses | go run . list |
| summary | Show expense summary | go run . summary |
| summary --month | Show summary for specific month | go run . summary --month 8 |
| delete | Delete an expense | go run . delete --id 1 |