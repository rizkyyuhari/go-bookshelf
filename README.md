# go-bookshelf

#### This is a simple web api of bookshelf
## using go gin, gorm, and postgres
This Api has 4 http method, which is CREATE, Read, UPDATE, and DELETE  
  - GET with endpoint "v1/books"
  - POST with endpoint "v1/books" and  field (title string, description string, rating int, price int) as a request body.
  - DELETE with endpoint "v1/books/:id" (this endpoint need url parameter id to delete row from DB).
  - PUT with endpoint "v1/books" (this endpoint only update title from specific row. need 2 query parameter, first is "id" and second "title") 
