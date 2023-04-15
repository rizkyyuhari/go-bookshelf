package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyyuhari/go-bookshelf.git/db"
	"github.com/rizkyyuhari/go-bookshelf.git/entity"
)

func GetAllBook(c *gin.Context)  {
	var books []entity.Book	
	db.DB.Order("updated_at asc").Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data" : books, 
	})
}

func  CreateBook(c * gin.Context) {
	var books  entity.Book
	err := c.ShouldBindJSON(&books)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : "Bad Request",
			"desc" : err.Error(),
			"data" : "[]",
		})
		return
	}
	db.DB.Create(&books)
	c.JSON(http.StatusOK, gin.H{
		"message" : "Berhasil menambah buku!",
		"data": "[]",
	})
}

func DeleteBook(c * gin.Context){
	var book entity.Book
	id := c.Param("id")

	err := db.DB.Where("id = ?", id).Delete(&book).Error
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Berhasil Menghapus Buku",
	})
}

func UpdateBook(c *gin.Context){
	var book entity.Book
	id := c.Param("id")
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.DB.Model(&book).Where("id = ? ",id ).Updates(book)

    c.JSON(http.StatusOK, gin.H{"message": "Berhasil merubah data!"})
}