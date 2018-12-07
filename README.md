# kreditplus-api
API with Golang for Training Kreditplus

Prepare plugin before code :

- go get -u -v github.com/gin-gonic/gin
- go get -u -v github.com/denisenkom/go-mssqldb
- go get -u -v github.com/lib/pq
- go get -u -v github.com/Sirupsen/logrus
- go get -u -v github.com/dgrijalva/jwt-go
- go get -u -v golang.org/x/crypto/bcrypt
- go get -u -v github.com/jinzhu/gorm

Step-by-step running :

Database :
1. PostgreSQL :
   - Create Database KreditPlus (script included)
   - Create Table tbl_user (script included)

2. SQL Server :
   - Create Database KreditPlus (script included)
   - Create Table tbl_employee (script included)

App :
1. Run : go run main.go
