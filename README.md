
# Hacktiv8 x MSIB - Order App

Name : Yusril Ilham Kholid

No. Peserta : GLNG-KS07-04

## Entity Relationship Diagram

The relationship is one to many, one order can have many items. The DBMS used this application is PostgreSQL.

![alt text](assets/erd.png)

## Swagger UI
![alt text](assets/swagger.png)

## Setup .env file
Before you run this application please setup your .env file.
```
DB_USER=
DB_PASSWORD=
DB_NAME=
DIALECT=
PORT=
```

## Endpoints
Http method and path you can use for this application
| No. | Method |        Path        |       Description        |
|-----|--------|--------------------|--------------------------|
| 1   | POST   |  /orders           | Create new order         |
| 2   | GET    |  /orders           | Get all orders           |
| 3   | PUT    |  /orders/{id}      | Update order by order id |
| 4   | DELETE |  /orders/{id}      | Delete order by order id |

## Create Order

Endpoint : [POST] /orders

Body request :
```json
{
  "customerName": "An Yujin",
  "items": [
    {
      "description": "Indomie Goreng",
      "itemCode": "indomiegr",
      "quantity": 2
    }
  ],
  "orderedAt": "2023-09-22T09:11:00+07:00"
}
```

Response Success (201) :
```json
{
  "statusCode": 201,
  "message": "string",
  "data": "nil"
}
```

## Get Orders

Endpoint : [GET] /orders

Response Success (200) :
```json
{
  "statusCode": 200,
  "message": "string",
  "data": []
}
```

## Update Order

Endpoint : [PUT] /orders/{id}

Param required : 
```
orderId int
```

Body request : 
```json
{
  "customerName": "An Yujin",
  "items": [
    {
      "description": "Indomie Goreng",
      "itemCode": "indomiegr",
      "quantity": 3
    }
  ],
  "orderedAt": "2023-09-22T09:11:00+07:00"
}
```

Response Success (200) :
```json
{
  "statusCode": 200,
  "message": "string",
  "data": "nil"
}
```

## Delete Order

Endpoint : [DELETE] /orders/{id}

Param required : 
```
orderId int
```

Response Success (200) :
```json
{
  "statusCode": 200,
  "message": "string",
  "data": "nil"
}
```

## Run Application
```sh
  $ go run main.go
```

## Run Swagger UI
After run go program you can run in browser :
http://localhost:8080/swagger/index.html
