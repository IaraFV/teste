package main

import "fmt"

type Pessoas struct{
	ID 	string
	Nome       string 
	Id_equipe  string
}
const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "123456"

const DbName = "API-Golang"

const TableName = "pessoas"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

