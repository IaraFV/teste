package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	/*dbConfig "./dbconfig"*/
	_ "github.com/lib/pq"
)

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "123456"

const DbName = "API-Golang"

const TableName = "pessoas"
//banco
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//type
type projeto struct {
	ID           string `json:"id"`
	Nome_projeto string `json:"nome_projeto"`
	Equipe_resp  string `json:"equipe_resp"`
}
type pessoa struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Id_Projeto string `json:"Id_Projeto"`
	Id_equipe  string `json:"id_equipe"`
}

type equipe struct {
	ID           string `json:"id"`
	NOME_EQUIPE  string `json:"nome_equipe"`
	Scrum_master string `json:"scrum_master"`
	Id_projeto   string `json:"id_projeto"`
}
type tarefa struct {
	ID          string `json:"id_tar"`
	Nome_tar    string `json:"Nome_tar"`
	Projeto_res string `json:"Projeto_res"`
}

//var
var tarefas = []tarefa{
	{ID: "1", Nome_tar: "tentativa um", Projeto_res: "1"},
}

//banco
var db *sql.DB
var err error

//
var equipes = []equipe{
	{ID: "1", NOME_EQUIPE: "tentando dnv", Scrum_master: "lucas", Id_projeto: "1"},
}
var projetos = []projeto{
	{ID: "1", Nome_projeto: "beta", Equipe_resp: "1"},
	{ID: "2", Nome_projeto: "alpha", Equipe_resp: "2"},
}
var pessoas = []pessoa{
	{ID: "1", Nome: "lolo", Id_Projeto: "5", Id_equipe: "1"},
	{ID: "2", Nome: "lili", Id_Projeto: "4", Id_equipe: "3"},
	{ID: "3", Nome: "lulu", Id_Projeto: "3", Id_equipe: "2"},
	{ID: "4", Nome: "lucas", Id_Projeto: "2", Id_equipe: "3"},
	{ID: "5", Nome: "iara", Id_Projeto: "2", Id_equipe: "8"},
}

//router
func main() {
	router := gin.Default()
	router.GET("/projetos", getProjetos)
	router.GET("/projetos/:id", getProjetoByID)
	router.POST("/projetos", postProjeto)
	router.PUT("/projetos/:id", updateProjetosById)

	router.GET("/equipes", getEquipes)
	router.GET("/equipes/:id", getEquipeByID)
	router.POST("/equipes", postEquipe)

	router.GET("/pessoas", getPessoas)
	router.GET("/pessoas/:id", getpessoaByID)
	router.POST("/pessoas", postpessoas)
	router.DELETE("/pessoas/:id", deletePessoaById)
	router.PUT("/pessoas/:id", updatePessoaById)

	router.GET("/tarefas/", getTarefas)
	router.POST("/tarefas", postTarefa)
	router.PUT("/tarefas", updateTarefas)
	router.GET("/tarefas/:id", getTarefasByID)

	router.Run("localhost:8000")


	
	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
	sqlSelectID()
	sqlInsert()
	sqlUpdate()
	sqlDelete()
}

func sqlSelect() {

	sqlStatement, err := db.Query("SELECT id, nome, id_projeto,id_equipe FROM " + TableName)
	checkErr(err)

	for sqlStatement.Next() {

		var Pessoas dbConfig.Pessoas

		err = sqlStatement.Scan(&Pessoas.ID, &Pessoas.nome, &Pessoas.id_projeto, &Pessoas.id_equipe)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", Pessoas.ID, Pessoas.nome, Pessoas.id_projeto, Pessoas.id_equipe)

	}
}

func sqlSelectID() {

	var pessoas dbConfig.Pessoas

	sqlStatement := fmt.Sprintf("SELECT id, nome, id_projeto,id_equipe FROM %s where id = $1", TableName)

	err = db.QueryRow(sqlStatement, 1).Scan(&pessoas.ID, &pessoas.Title, &pessoas.Body)
	checkErr(err)

	fmt.Printf("%d\t%s\t%s \n", pessoas.ID, pessoas.Title, pessoas.Body)
}

func sqlInsert() {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2, $3)", TableName)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := insert.Exec(5, "Maps in Golang", "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium")
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func sqlUpdate() {

	sqlStatement := fmt.Sprintf("update %s set body=$1 where id=$2", TableName)

	update, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := update.Exec("But I must explain to you how all this mistaken idea", 5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
func sqlDelete() {

	sqlStatement := fmt.Sprintf("delete from %s where id=$1", TableName)

	delete, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := delete.Exec(5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func getEquipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, equipes)
}
func getProjetos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projetos)
}
func getPessoas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pessoas)
}

func gettarefas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tarefas)
}

func postTarefas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tarefas)
}

func postEquipe(c *gin.Context) {
	var newEquipe equipe

	if err := c.BindJSON(&newEquipe); err != nil {
		return
	}

	equipes = append(equipes, newEquipe)
	c.IndentedJSON(http.StatusCreated, newEquipe)
}
func postpessoas(c *gin.Context) {
	var newpessoa pessoa

	if err := c.BindJSON(&newpessoa); err != nil {
		return
	}

	pessoas = append(pessoas, newpessoa)
	c.IndentedJSON(http.StatusCreated, newpessoa)
}

func postProjeto(c *gin.Context) {
	var newProjeto projeto

	if err := c.BindJSON(&newProjeto); err != nil {
		return
	}

	projetos = append(projetos, newProjeto)
	c.IndentedJSON(http.StatusCreated, newProjeto)
}

func getEquipeByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range equipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipes not found"})
}
func getpessoaByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range pessoas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
}
func getProjetoByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range projetos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func updateProjetosById(c *gin.Context) {
	id := c.Param("id")
	for i := range projetos {
		if projetos[i].ID == id {
			c.BindJSON(&projetos[i])
			c.IndentedJSON(http.StatusOK, projetos[i])
			return
		}
	}
}
func deletePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range pessoas {
		if a.ID == id {
			pessoas = append(pessoas[:i], pessoas[i+1:]...)
			return
		}
	}
}

func updatePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i := range pessoas {
		if pessoas[i].ID == id {
			c.BindJSON(&pessoas[i])
			c.IndentedJSON(http.StatusOK, pessoas[i])
			return
		}
	}
}

func getTarefas(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tarefas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func getTarefasByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tarefas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipes not found"})
}

func postTarefa(c *gin.Context) {
	var newTarefa tarefa
	if err := c.BindJSON(&newTarefa); err != nil {
		return
	}

	tarefas = append(tarefas, newTarefa)
	c.IndentedJSON(http.StatusCreated, newTarefa)
}

func updateTarefas(c *gin.Context) {
	id := c.Param("id")
	for i := range tarefas {
		if tarefas[i].ID == id {
			c.BindJSON(&tarefas[i])
			c.IndentedJSON(http.StatusOK, tarefas[i])
			return
		}
	}
}
