# E-Commerce Alta

## Requirements
For running this project please install go. Go:
```
go version
go1.22.3 darwin/arm64
```

## Getting Started
Init folder name:
```
go mod init e-commerce-alta
```

## Requirement:
install several libraries needed for the project, for example:
```
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo-jwt/v4
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

## Setup supabase
```
Login supabase
Insert project
Select menu table editor
Create new schema
Rename "e-commerce-alta"
```

## Setup database:
```
Add .env
env value example:
poshost= your host name
posuser= your username
pospw= your password
posport= port your database
dbname= your db name

JWT_SECRET= your pass jwt key
```

## Run the app:
```
go run main.go
```

## Testing api:
```
open insomnia
add htpp request
select method POST 
testing api, for example:
http://localhost:5000/login
click send
```

## OpenApi Documentation
This is link open api e-commerce-alta:
```
https://app.swaggerhub.com/apis-docs/FarahRaihanunnisa/E-CommerceManagementAPI/1.0.0
```