# go-bookshelf

#### This is a simple web api of bookshelf
## using go gin, gorm, and postgres

type Book struct {
	Title string
	Description string
	Price int
	Rating int
}

This Api has 4 http method, which is CREATE, Read, UPDATE, and DELETE  
  - GET with endpoint "/books"
  - POST with endpoint "/books" and  field (title string, description string, rating int, price int) as a request body.
  - DELETE with endpoint "/books/:id" (this endpoint need url parameter id to delete row from DB).
  - PUT with endpoint "/books/:id" (this endpoint take one url parameter and also accept request body from client to update book data. for req body field you can see struct Book above).
