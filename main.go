package main

import (
    "fmt"
    . "github.com/d-exclaimation/gin-gonic-api/database"
    . "github.com/d-exclaimation/gin-gonic-api/models"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "strconv"
)

var port = 5000

// Main Entry Point
func main() {
    gin.SetMode(gin.ReleaseMode)
    var err error
    var db = SetupDB()
    var app = gin.Default()

    // GET Request at 0.0.0.0:5000/wishlist/:id
    app.GET("/wishlist/:id", func(context *gin.Context) {
        var id, err = strconv.ParseInt(context.Param("id"), 10, 64)
        Handle(err)
        context.JSON(200, Get(int(id), db).ToGinH())
    })

    // GET Request at 0.0.0.0:5000/wishlist
    app.GET("/wishlist", func(context *gin.Context) {
        context.JSON(200, AllGinH(GetData(db)))
    })

    // POST Request at 0.0.0.0:5000/wishlist
    app.POST("/wishlist", func(context *gin.Context) {
    	var body ItemDTO
    	err := context.BindJSON(&body)
    	Handle(err)
    	context.JSON(200, PostData(body, db).ToGinH())
    })

    // PUT Request at 0.0.0.0:5000/wishlist/:id
    app.PUT("/wishlist/:id", func(context *gin.Context) {
        // Get the id from the parameters
        var id, err = strconv.ParseInt(context.Param("id"), 10, 64)
        Handle(err)

        // Get the JSON from the request
        var item Item
        err = context.BindJSON(&item)
        Handle(err)

        if int(id) != item.Id {
            context.JSON(400, gin.H{
                "message": "Wrong id",
            })
        } else {
            context.JSON(200, UpdateData(item, db).ToGinH())
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
