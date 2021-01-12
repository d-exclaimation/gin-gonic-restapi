package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


func setupDB() *sql.DB {
	// Create a string address to the database
	var psqlInfo = fmt.Sprintf("postgres://postgres:password@localhost/restful?sslmode=disable")

	// Try to connect using the default sql libraries
	db, err := sql.Open("postgres", psqlInfo)
	handle(err)

	// Ping the database due to the fact that connecting doesn't actually make sure database is available
	err = db.Ping()
	handle(err)

	return db
}


func getData(db *sql.DB) []*Item {
	// Creating a container and making a select query
	var results = make([]*Item, 0)
	rows, err := db.Query("SELECT * FROM wishlist;")
	handle(err)

	// For each item in the result rows, append as Item into container
	for rows.Next() {
		var curr = &Item{
			Id: 0,
			Name: "",
			Price: 0,
		}
		err = rows.Scan(&curr.Id, &curr.Name, &curr.Price)
		handle(err)
		results = append(results, curr)
	}

	// In case traversing the rows causes an error
	err = rows.Err()
	handle(err)

	return results
}

func get(id int, db *sql.DB) *Item {
	// Create an empty item pointer, and send a query with the id
	var result = &Item{
		Id:    0,
		Name:  "",
		Price: 0,
	}
	var query = fmt.Sprintf("SELECT * FROM wishlist WHERE list_id = %d;", id)
	row := db.QueryRow(query)

	// Get the result into Item
	err := row.Scan(&result.Id, &result.Name, &result.Price)
	handle(err)
	return result
}


func postData(item ItemDTO , db *sql.DB) {
	// Send a insert query
	var query = fmt.Sprintf("('%s', %d);", item.Name, item.Price)
	_, err := db.Exec("INSERT INTO wishlist (name, price) VALUES " + query)
	handle(err)
}

func updateData(item Item, db *sql.DB) {
	// Send an update query
	_, err := db.Exec("UPDATE wishlist SET name = $1, price = $2 WHERE list_id = $3", item.Name, item.Price, item.Id)
	handle(err)
}