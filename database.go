package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

func main() {

	///********************** CNEXION DB *********************************

	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())

	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	///********************** INSERTAR DATOS *********************************

	/*insert, err := db.Query("INSERT INTO usuario VALUES ( 1, 'carlos', 'perez', 15 ),( null, 'jose', 'diaz', 22 ),( null, 'carlos', 'andrade', 4 ), ( null, 'andres', 'aaww', 45 ), ( 5, 'fabricio', 'dffsss', 19 )")

	  // if there is an error inserting, handle it
	  if err != nil {
	      panic(err.Error())
	  }
	  // be careful deferring Queries if you are using transactions
	  defer insert.Close()*/

	///********************** CONSULTAR TODOS LOS ROW *********************************

	results, err := db.Query("SELECT * FROM usuario")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user usuario
		// for each row, scan the result into our user composite object
		err = results.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Edad)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the user's Name attribute

		fmt.Println(user)

	}

	///********************** CONSULTA UN SOLO ROW *********************************

	/*var user usuario
	  // Execute the query
	  err = db.QueryRow("SELECT id, nombre, apellido, edad FROM usuario where id = ?", 4).Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Edad)
	  if err != nil {
	      panic(err.Error()) // proper error handling instead of panic in your app
	  }

	  log.Println(strconv.Itoa(user.ID))
	  log.Println(user.Nombre)
	  log.Println(user.Apellido)
	  log.Println(strconv.Itoa(user.Edad))*/

}
