Go Pack Size Calculator API

The API can calculate the number of packs we need to ship to the customer. 
I used gorilla mux for routing, swagger ui for documentation and testing 

Description:
 Order to given pack sizes api behaves order by following rules:
    1.	Only whole packs can be sent. Packs cannot be broken open. 
    2.	Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order. 
    3.	Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order. 
    4.  Pack sizes can be changed without having to change code


Installation 
1. Make sure you have Go installed
2. Clone repo
3. Install dependencies with command : go mod tidy

For usage just run command: go run main.go 

API Endpoints
The API provides the following endpoints for calculating pack sizes:

GET /api/v1/pack-sizes: To get existing pack sizes 
POST /api/v1/pack-sizes: To set new pack sizes. Body have to be array of int. For example : [50,100,150,200]
GET /api/v1/calculate-packs: To calculate order packs by quantity. Body have to be json object.For example: {"OrderQuantity":504812}. Response example: {"300":1,"100":1}

To run unit tests, in project directory run: go test -v

Swagger ui link: http://localhost:8080/swagger/index.html
