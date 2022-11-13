# go-gin-gorm-mysql
Go : GORM package for MySQL and Gin based RESTful APIs

This repository contains developing of RESTful web service APIs in Go Programming Language using Gin Web Framework (Gin) and uses ORM library for Golang (GORM) package to connect with MySQL Database for DB CRUD operations.

**Prerequisite:**
- Make sure GO already installed and had working environment.<br>
while developing this code go version used,<br>
```
$go version
go version go1.19.3 windows/amd64 
```
- Male sure MySQL installed and tested with `MySQL Workbench` or `MySQL Command Line Client`<br>
while developing this code MySQL version used,<br>
```
Server version: 8.0.31 MySQL Community Server - GPL
```

**Steps to Run:**
- Run `go run main.go` in terminal
- Go to another command terminal and test the REST APIs using CURL (or use REST client like Postman).

**Test Results:**<br>
CRUD REST APIs test using CURL and it's Results:

```
$curl -X GET http://localhost:8080/ping
"Hello.. Welcome !!!"

$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:57:15 GMT
Content-Length: 2

[]
$curl -iX POST http://localhost:8080/books  -H "Content-Type: application/json" -d @book1.json
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:57:37 GMT
Content-Length: 294

{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"}
$
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   296  100   296    0     0  24734      0 --:--:-- --:--:-- --:--:-- 26909
[
    {
        "ID": 0,
        "CreatedAt": "2022-11-12T16:57:37.002-08:00",
        "UpdatedAt": "2022-11-12T16:57:37.002-08:00",
        "DeletedAt": null,
        "id": "19dd457e-42b1-4210-96fd-6157829ad013",
        "name": "C Programming Language",
        "author": "Brian W. Kernighan, Dennis M. Ritchie",
        "price": 57.09,
        "pages": 272,
        "date_published": "1988-03-22"
    }
]

$
$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:57:43 GMT
Content-Length: 296

[{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"}]
$
$
$curl -iX GET http://localhost:8080/books/19dd457e-42b1-4210-96fd-6157829ad013
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:57:59 GMT
Content-Length: 294

{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"}
$
$
$curl -iX POST http://localhost:8080/books  -H "Content-Type: application/json" -d @book2.json
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:58:13 GMT
Content-Length: 281

{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T16:58:13.342-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}
$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:58:21 GMT
Content-Length: 578

[{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"},
{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T16:58:13.342-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}]
$
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   578  100   578    0     0  34026      0 --:--:-- --:--:-- --:--:-- 38533
[
    {
        "ID": 0,
        "CreatedAt": "2022-11-12T16:57:37.002-08:00",
        "UpdatedAt": "2022-11-12T16:57:37.002-08:00",
        "DeletedAt": null,
        "id": "19dd457e-42b1-4210-96fd-6157829ad013",
        "name": "C Programming Language",
        "author": "Brian W. Kernighan, Dennis M. Ritchie",
        "price": 57.09,
        "pages": 272,
        "date_published": "1988-03-22"
    },
    {
        "ID": 0,
        "CreatedAt": "2022-11-12T17:34:06.936-08:00",
        "UpdatedAt": "2022-11-12T17:34:06.936-08:00",
        "DeletedAt": null,
        "id": "4ab40f24-7ecd-4690-a642-c13ec3ebfbe3",
        "name": "The C++ Programming Language",
        "author": "Bjarne Stroustrup",
        "price": 64.31,
        "pages": 1376,
        "date_published": "2013-05-09"
    }
]

$
$curl -iX GET http://localhost:8080/books/ca29f679-9b03-4e88-8175-4217f5293e37
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 00:58:52 GMT
Content-Length: 281

{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T16:58:13.342-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}
$
$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:02:47 GMT
Content-Length: 578

[{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"},
{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T16:58:13.342-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}]
$curl -iX GET http://localhost:8080/books/ca29f679-9b03-4e88-8175-4217f5293e37
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:03:09 GMT
Content-Length: 281

{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T16:58:13.342-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}
$curl -iX PATCH http://localhost:8080/books/ca29f679-9b03-4e88-8175-4217f5293e37 -H "Content-Type: application/json" -d @update2.json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:03:23 GMT
Content-Length: 291

{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T17:03:23.833-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language - 2nd Edition","author":"B Stroustrup ","price":65.01,"pages":1399,"date_published":"2017-07-07"}
$curl -iX GET http://localhost:8080/books/ca29f679-9b03-4e88-8175-4217f5293e37
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:03:31 GMT
Content-Length: 291

{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T17:03:23.833-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language - 2nd Edition","author":"B Stroustrup ","price":65.01,"pages":1399,"date_published":"2017-07-07"}
$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:03:35 GMT
Content-Length: 588

[{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"},
{"ID":0,"CreatedAt":"2022-11-12T16:58:13.342-08:00","UpdatedAt":"2022-11-12T17:03:23.833-08:00","DeletedAt":null,"id":"ca29f679-9b03-4e88-8175-4217f5293e37","name":"The C++ Programming Language - 2nd Edition","author":"B Stroustrup ","price":65.01,"pages":1399,"date_published":"2017-07-07"}]

$curl -iX DELETE http://localhost:8080/books/ca29f679-9b03-4e88-8175-4217f5293e37
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:04:12 GMT


$curl -iX GET http://localhost:8080/books
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 13 Nov 2022 01:04:16 GMT
Content-Length: 296

[{"ID":0,"CreatedAt":"2022-11-12T16:57:37.002-08:00","UpdatedAt":"2022-11-12T16:57:37.002-08:00","DeletedAt":null,"id":"19dd457e-42b1-4210-96fd-6157829ad013","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"}]
$
```