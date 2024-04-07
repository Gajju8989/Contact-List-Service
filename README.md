Contact App Task
Author: Gajendra Singh

Description
This project is a contact application designed to manage contacts. It provides APIs for user registration, contact creation, updating spam status, and searching for contacts by name or phone number.

Dependencies
- Go 1.21.6
- github.com/dgrijalva/jwt-go v3.2.0+incompatible
- github.com/gin-gonic/gin v1.9.1
- github.com/go-sql-driver/mysql v1.8.1
- github.com/google/wire v0.6.0
- github.com/joho/godotenv v1.5.1
- golang.org/x/crypto v0.22.0
- gorm.io/driver/mysql v1.5.6
- gorm.io/gorm v1.25.9
See go.mod file for additional dependencies

Installation and Usage
1. Navigate to the project directory:

```bash
cd contact-app-task



Ensure you have Go installed and configured properly.

Installing Dependencies

To install all dependencies listed in the go.mod file for your Go project, follow these steps:

1.Open your terminal or command prompt.

2.Navigate to the root directory of your Go project.

3.Run the following command to download and install all dependencies listed in the go.mod file:


go mod tidy


This command reads your go.mod file, downloads any missing dependencies, and removes any dependencies that are no longer used by your project. It also updates the go.sum file to reflect the changes.

After running go mod tidy, all dependencies listed in your go.mod file should be installed in your project.

Local Environment Setup
Before running the application, you need to set up your local environment variables by creating a .env file in the project root directory. Add the following configuration to the .env file:

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD="xyz"
DB_NAME=db

Replace the values with your database configuration.

1. Run the following command to start the server:

 go run main.go



The server will start running on port 8080 by default.

API Documentation
The API documentation is generated using Swagger. To view the API documentation, follow these steps:

1.Start the server (if not already running) by following step 4 in the Installation and Usage section.

2.Open a web browser and navigate to the following URL:

http://localhost:8080/swagger/index.html

This will open the Swagger UI where you can explore and test the available APIs.


