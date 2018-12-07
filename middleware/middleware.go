package middleware

import (
	"errors"
	"net/http"
	"strings"

	common "kreditplus/kreditplus-api/commons"
	user "kreditplus/kreditplus-api/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// AuthorizationHeaderExtractor : Extract token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// MyAuth2Extractor : Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// UpdateContextUserModel : A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, myUserID uint) {
	var myUserModel user.ModelUser
	if myUserID != 0 {
		db := common.GetPostgreSQLDB()
		db.First(&myUserModel, myUserID)
	}
	c.Set("my_user_id", myUserID)
	c.Set("my_user_model", myUserModel)
}

// AuthMiddleware : You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
				c.JSON(http.StatusUnauthorized, common.NewError("middleware", errors.New("No token present in request")))
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			myUserID := uint(claims["id"].(float64))
			UpdateContextUserModel(c, myUserID)
		}
	}
}
