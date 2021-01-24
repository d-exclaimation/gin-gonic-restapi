package database

import (
	"database/sql"
	"fmt"
	. "github.com/d-exclaimation/gin-gonic-api/models"
	_ "github.com/lib/pq"
)


func SetupDB() *sql.DB {
	// Create a string address to the database
	var psqlInfo = fmt.Sprintf("postgres://postgres:password@localhost/restful?sslmode=disable")

	// Try to connect using the default sql libraries
	db, err := sql.Open("postgres", psqlInfo)
	Handle(err)

	// Ping the database due to the fact that connecting doesn't actually make sure database is available
	err = db.Ping()
	Handle(err)

	return db
}


func GetData(db *sql.DB) []*Item {
	// Creating a container and making a select query
	var results = make([]*Item, 0)
	rows, err := db.Query("SELECT * FROM wishlist")
	Handle(err)

	// For each item in the result rows, append as Item into container
	for rows.Next() {
		var curr = &Item{
			Id: 0,
			Name: "",
			Price: 0,
		}
		err = rows.Scan(&curr.Id, &curr.Name, &curr.Price)
		Handle(err)
		results = append(results, curr)
	}

	// In case traversing the rows causes an error
	err = rows.Err()
	Handle(err)

	return results
}

func Get(id int, db *sql.DB) *Item {
	// Create an empty item pointer, and send a query with the id
	var result = &Item{}
	row := db.QueryRow("SELECT * FROM wishlist WHERE list_id = $1", id)

	// Get the result into Item
	err := row.Scan(&result.Id, &result.Name, &result.Price)
	Handle(err)
	return result
}


func PostData(item ItemDTO , db *sql.DB) *Item {
	// Create an empty item pointer, and send a query with the new item
	var result = &Item{}
	row := db.QueryRow("INSERT INTO wishlist (name, price) VALUES ($1, $2) RETURNING *", item.Name, item.Price)

	// Get the result back and return it
	err := row.Scan(&result.Id, &result.Name, &result.Price)
	Handle(err)
	return result
}

func UpdateData(item Item, db *sql.DB) *Item {
	// Create an empty item pointer, and send a query with the updated item
	var result = &Item{}
	row := db.QueryRow("UPDATE wishlist SET name = $1, price = $2 WHERE list_id = $3 RETURNING *", item.Name, item.Price, item.Id)

	// Get the result back and return it
	err := row.Scan(&result.Id, &result.Name, &result.Price)
	Handle(err)
	return result
}