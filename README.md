# posts-api-mysql
API for blog posts using Golang and MySQL

## Dependencies:
Using golang/dep to manager dependencies. See: https://github.com/golang/dep

## Getting started
  To create docker image: `docker build -t posts-api-mysql .`
  To start containers: `docker-compose up -d`

##### Endponts

|Method          |URI                         |
|----------------|-------------------------------|
|GET|http://localhost:3000/resources/posts|
|POST|http://localhost:3000/resources/posts            |
|PUT|http://localhost:3000/resources/posts/{number}|
|DELETE|http://localhost:3000/resources/posts/{number}|
|GET-by-id|http://localhost:3000/resources/posts/{number}|
|GET-pagination|http://localhost:3000/resources/posts/?size={number}&page={number}|

