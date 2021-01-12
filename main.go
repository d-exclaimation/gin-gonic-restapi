package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "strconv"
)

// JSON Item Struct
type Item struct {
    Id int
    Name string
    Price int
}

type ItemDTO struct {
    Name string
    Price int
}

var port = 5000

// Main Entry Point
func main() {
    gin.SetMode(gin.ReleaseMode)
    var err error
    var db = setupDB()
    var app = gin.Default()

    // GET Request at 0.0.0.0:5000/wishlist/:id
    app.GET("/wishlist/:id", func(context *gin.Context) {
        var id, err = strconv.ParseInt(context.Param("id"), 10, 64)
        handle(err)
        context.JSON(200, get(int(id), db))
    })

    // GET Request at 0.0.0.0:5000/wishlist
    app.GET("/wishlist", func(context *gin.Context) {
        context.JSON(200, getData(db))
    })

    // POST Request at 0.0.0.0:5000/wishlist
    app.POST("/wishlist", func(context *gin.Context) {
    	var body ItemDTO
    	err := context.BindJSON(&body)
    	handle(err)
    	postData(body, db)
    	context.JSON(200, body)
    })

    // PUT Request at 0.0.0.0:5000/wishlist/:id
    app.PUT("/wishlist/:id", func(context *gin.Context) {
        // Get the id from the parameters
        var id, err = strconv.ParseInt(context.Param("id"), 10, 64)
        handle(err)

        // Get the JSON from the request
        var item Item
        err = context.BindJSON(&item)
        handle(err)

        if int(id) != item.Id {
            context.JSON(400, gin.H{
                "message": "Wrong id",
            })
        } else {
            updateData(item, db)
            context.JSON(200, item)
        }

    })

    // Setup the server
    var address = fmt.Sprintf(":%d", port)
    err = app.Run(address)
    if err != nil {
        fmt.Println(err)
    }
    db.Close()
}
