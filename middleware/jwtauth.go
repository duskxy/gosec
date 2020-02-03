package middleware

import (
	"gosec/service"
	"log"
	"time"
    "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GinJWTMiddlewareInit() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "gosec",
		Key:        []byte("adminadmin"),
		Timeout:    time.Minute * 5,
		MaxRefresh: time.Hour,
		// IdentityKey: identityKey,
		// PayloadFunc: func(data interface{}) jwt.MapClaims {
		// 	if v, ok := data.(*models.User); ok {
		// 		//get claims from username
		// 		v.UserClaims = models.GetUserClaims(v.UserName)
		// 		jsonClaim, _ := json.Marshal(v.UserClaims)
		// 		//maps the claims in the JWT
		// 		return jwt.MapClaims{
		// 			"userName":   v.UserName,
		// 			"userClaims": string(jsonClaim),
		// 		}
		// 	}
		// 	return jwt.MapClaims{}
		// },
		// IdentityHandler: func(c *gin.Context) interface{} {
		// 	claims := jwt.ExtractClaims(c)
		// 	//extracts identity from claims
		// 	jsonClaim := claims["userClaims"].(string)
		// 	var userClaims []models.Claims
		// 	json.Unmarshal([]byte(jsonClaim), &userClaims)
		// 	//Set the identity
		// 	return &models.User{
		// 		UserName:   claims["userName"].(string),
		// 		UserClaims: userClaims,
		// 	}
		// },
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals service.UserLoginService
			if err := c.ShouldBind(&loginVals); err == nil {
				if res, ok := loginVals.Login(c); ok {
					return res,nil
				} else {
					return res,jwt.ErrFailedAuthentication
				}
				
			} else {
				return "", jwt.ErrMissingLoginValues
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		// Authorizator: jwtAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}
