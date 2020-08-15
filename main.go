package main

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
  )


type City struct {
	gorm.Model
	ID int
	name string
	governingDistrictID int
}

func main() {
	// r := gin.Default()
	// r.GET("/cuk", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	/**/
	// db, err := gorm.Open("mysql", "root@/doogether_legacy_region")

	// if err != nil {
	// 	fmt.Print(err)
	// }

	// rows, err := db.Raw("SELECT ID, name FROM city").Rows()
	// defer rows.Close()

	// for rows.Next() {
	// 	var ID int
	// 	var name string

	// 	rows.Scan(&ID, &name)
	
	// 	fmt.Printf("ID : %v , name : %v \n", ID, name)
	// }



	/**/
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