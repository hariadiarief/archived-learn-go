package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	gorm.Model
	ID   int
	Name string
}

func main() {

	r := gin.Default()
	db, err := gorm.Open("mysql", "root@/doogether_legacy_region")

	if err != nil {
		fmt.Panic(err)
	}

	defer db.Close()

	r.GET("/cuk", func(c *gin.Context) {

		rows, err := db.Raw("SELECT ID, name FROM city").Rows()
		defer rows.Close()

		cities := make([]City, 0)

		for rows.Next() {

			var city City

			rows.Scan(&city.ID, &city.Name)

			cities = append(cities, city)
		}

		c.SecureJSON(200, gin.H{
			"data": cities,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
