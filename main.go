package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyyuhari/go-bookshelf.git/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	dsn := "host=localhost user=postgres password=260999 dbname=GolangDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}


	router := gin.Default()
	
	v1:= router.Group("/v1")

	v1.GET("books", func(c *gin.Context){
		books,err:= getBooks(db)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data" : books,
		})
	})
	v1.POST("books", func(c *gin.Context){
		var data book.Book
		err := c.ShouldBindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}
		
		createBook(&data,db)

		c.JSON(http.StatusOK, gin.H{
			"message" : "Berhasil menambahkan buku baru!",
		})
	})

	v1.DELETE("books/:id", func(c *gin.Context){
		id,_ := strconv.Atoi(c.Param("id"))
		
		deleteBook(id,db)

		c.JSON(http.StatusOK, gin.H{
			"message" : "Berhasil menghapus!",
		})
	})

	v1.PUT("books/", func(c *gin.Context){
		id,_ := strconv.Atoi(c.Query("id"))
		title := c.Query("title")

		 updateDescriptionBook(id,title,db)

		c.JSON(http.StatusOK, gin.H{
			"message" : "Berhasil Update Column Description!",
			"id" : id,
			"desc" : title,
		})
	})

	router.Run()

}


func getBooks(db *gorm.DB) ([]book.Book,error){
	var books []book.Book
	err := db.Order("id asc").Find(&books).Error
	return books,err
}

func createBook(data *book.Book,db *gorm.DB) error {
	err := db.Create(data).Error
	return err
}

func deleteBook(id int,db *gorm.DB) error {
	var book book.Book
	err := db.Where("id = ?", id).Delete(&book).Error
	return err
}

func updateDescriptionBook(id int, title string,db *gorm.DB) error{
	var book book.Book
	err := db.Where("id = ?", id).Find(&book).Error
	if err!= nil{
		log.Fatal(err)
	}
	book.Title = title
	err = db.Save(&book).Error
	return err
}

