# posts-api-mysql
API for blog posts using Golang and MySQL

##### Dependencies:
go get -u github.com/go-sql-driver/mysql <br />
go get -u github.com/gorilla/mux <br />

##### URI`s
|                |ASCII                         |
|----------------|-------------------------------|
|GET|http://localhost:3000/resources/posts|
|POST|http://localhost:3000/resources/posts            |
|PUT|http://localhost:3000/resources/posts/{number}|
|DELETE|http://localhost:3000/resources/posts/{number}|
|GET-by-id|http://localhost:3000/resources/posts/{number}|
|GET-pagination|http://localhost:3000/resources/posts/?size={number}&page={number}|

