package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect to postgres DB by "jackc/pgx" package
	//default port of postgres is 5432
	con, err := sql.Open("pgx", "host=localhost port=5432 dbname=test user=ehsan.d password=")
	if err != nil {
		log.Fatalf("Can not connect to postgres DB : %v/n", err)
	}

	//close connection to postgres DB
	defer con.Close()

	//log connection status
	log.Println("Connected to postgres DB")

	//test connection by send ping to postgres DB and set back pong or error
	err = con.Ping()
	if err != nil {
		log.Fatal("Can not ping to postgres DB !!")
	}

	log.Println("pinged postgres DB ")

	fmt.Println("------------Show All Rows-------------")

	//get rows from DB table
	err = getAllrows(con)
	if err != nil {
		log.Fatal(err)
	}

	//insert a row

	query := `insert into users (name , last_name , phone) values ($1, $2, $3)`
	_, err = con.Exec(query, "ehsan", "dastras", "09121234567")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted a row")

	// //get rows from DB table

	fmt.Println("------------Show All Rows-------------")
	err = getAllrows(con)
	if err != nil {
		log.Fatal(err)
	}

	//update a row

	query_update := `update users set phone = $1 where id = $2`
	_, err = con.Exec(query_update, "09100000000", 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated a row")
	fmt.Println("------------Show All Rows-------------")
	err = getAllrows(con)
	if err != nil {
		log.Fatal(err)
	}

	//get a row by id

	select_query := `select id , name ,last_name ,phone from users where id = $1`
	var Name, LastName, Phone string
	var Id int
	row := con.QueryRow(select_query, 1)
	err = row.Scan(&Id, &Name, &LastName, &Phone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("record is", Id, Name, LastName, Phone)
	
	fmt.Println("show only one row")


	//delete a row by id

	delete_query := `delete from users where id = $1`
	_, err = con.Exec(delete_query, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted a row")

}

func getAllrows(con *sql.DB) error {
	rows, err := con.Query("select id , name , last_name , phone from users")
	if err != nil {
		log.Println(err)
		return err
	}
	//close connection
	defer rows.Close()

	var Name, LastName, Phone string
	var Id int

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &LastName, &Phone)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("record is", Id, Name, LastName, Phone)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error scaning rows", err)
	}

	fmt.Println("--------------------------------------")

	return nil
}
