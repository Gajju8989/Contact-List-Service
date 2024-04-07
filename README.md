# Contact App Task

Author: Gajendra Singh

## Description
This project is a contact application designed to manage contacts. It provides APIs for user registration, contact creation, updating spam status, and searching for contacts by name or phone number.

## Dependencies
- Go 1.21.6
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) v3.2.0+incompatible
- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) v1.9.1
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) v1.8.1
- [github.com/google/wire](https://github.com/google/wire) v0.6.0
- [github.com/joho/godotenv](https://github.com/joho/godotenv) v1.5.1
- [golang.org/x/crypto](https://golang.org/x/crypto) v0.22.0
- [gorm.io/driver/mysql](https://gorm.io/driver/mysql) v1.5.6
- [gorm.io/gorm](https://gorm.io/gorm) v1.25.9
- See go.mod file for additional dependencies

## Installation and Usage
1.Navigate to the project directory.

cd contact-app-task

2. Ensure you have Go installed and configured properly.

3.Run the following command to start the server:'
go run main.go
4.The server will start running on port 8080 by default.


API Documentation
The API documentation is generated using Swagger. To view the API documentation, follow these steps:
1.
Start the server (if not already running) by following step 4 in the Installation and Usage section.
2.
Open a web browser and navigate to the following URL:

http://localhost:8080/swagger/index.html
This will open the Swagger UI where you can explore and test the available APIs.
