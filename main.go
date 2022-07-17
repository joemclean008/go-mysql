package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	router.POST("/login", getLogin)
	router.GET("/data", data)

	router.Run("localhost:8080")

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "go"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type login struct {
	Uname    string `json:"uname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type account struct {
	ID       string
	Uname    string
	Password string
}

var accounts = []account{
	{ID: "1", Uname: "joe", Password: "test"},
	{ID: "2", Uname: "mclean", Password: "alma"},
}

func getLogin(c *gin.Context) {
	var request login
	if err := c.BindJSON(&request); err != nil {
		return
	}
	if request.Uname == accounts[0].Uname && request.Password == accounts[0].Password {
		c.IndentedJSON(http.StatusOK, "Success!")
	} else {
		c.IndentedJSON(http.StatusOK, "Login Failed!")
	}

}

type User struct {
	id       int
	name     string
	uname    string
	password string
}
type Errormessage struct {
	// export the field, i.e. change it to start with an upper case letter
	Errormessage string `json:"errormessage"`
}

func data(c *gin.Context) {
	db := dbConn()
	var user User
	rows, err := db.Query("SELECT * FROM users WHERE id='1'")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	} else {
		if rows.Next() {
			err = rows.Scan(&user.id, &user.name, &user.uname, &user.password)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad Request",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully retrieved!",
				"user":    user,
			})
		}
	}
}
