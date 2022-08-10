# Routine ToDo
A simple todolist application written in Go


## Requirements
* MySQL installed
* Go installed

## Installation
* Clone this repo 
```bash
git clone https://github.com/arcenme/routine_todo
```
* Change Directory
```bash
cd routine_todo
```
* Initiate `.env` file
```bash
cp .env.example .env
```
* Modify `.env` file with your correct database credentials

## Packages Requirements
- gin-gonic/gin package for web framework `go get -u github.com/gin-gonic/gin`
- jinzhu/gorm package  for ORM (Object Relational Mapping) library `go get -u github.com/jinzhu/gorm`
- go-sql-driver package to connect with MySql `go get -u github.com/go-sql-driver/mysql`
- joho/godotenv package to access the environment variable `go get -u github.com/joho/godotenv`
- go-playground/validator package to validate user input value `go get -u github.com/go-playground/validator/v10`


## Usage
To run this application, execute:

```bash
go run main.go
```
You should be able to access this application at `http://127.0.0.1:4121`

## References
https://gowebexamples.com/  
https://kotakode.com/blogs/5479/Membuat-crud-api-dengan-menggunakan-Go%2C-Gin-dan-Mysql  
https://github.com/joho/godotenv  
https://github.com/go-playground/validator/blob/master/_examples/translations/main.go    
https://codepen.io/paulj05hua/pen/LYGLJYQ