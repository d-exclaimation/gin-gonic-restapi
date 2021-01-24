package models

import (
	"github.com/gin-gonic/gin"
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

func (item Item) ToGinH() gin.H {
	return gin.H{
		"id": item.Id,
		"name": item.Name,
		"price": item.Price,
	}
}

func AllGinH(items []*Item) []gin.H {
	var res []gin.H
	for i := range items {
		res = append(res, items[i].ToGinH())
	}
	return res
}
