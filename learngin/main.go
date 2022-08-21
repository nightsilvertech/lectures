package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

type GreetingsRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func RequestIDGenerator() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := uuid.New().String()
		context.Set("RequestID", requestId)
		context.Next()
	}
}

func Authentication(username string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		dataHeaders := strings.Split(authHeader, " ")
		jwtToken := dataHeaders[1]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			usernameClaims := claims["username"]
			if usernameClaims == username {
				context.Next()
			} else {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "username not recognized",
				})
				return
			}
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "error casting claims",
			})
			return
		}
	}
}

// JWT - Json Web Token
// Guna nya : checking authentication
// signing token bahwa token tersebut di generate oleh backend dan tidak bisa di generate secara sembarangan
// jwt punya umur

const secretKey = `ABCD1234`

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	// gin web framework initialization
	r := gin.New()
	r.Use(RequestIDGenerator())

	// request
	// - body (data data yang akan di berikan ke dalam request dan di prosess oleh backend)
	// - query param (data juga yang di berikan ke dalam request tetapi lewat query url)
	// - header (data tambahan yang di gunakan oleh backend untuk kebutuhan authentication dan formating)
	r.POST("/login", func(context *gin.Context) {
		token, err := GenerateJWT("isawk")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	})

	r.POST("/greetings", Authentication("isawk"), func(context *gin.Context) {
		var greetingsRequest GreetingsRequest
		err := context.ShouldBindJSON(&greetingsRequest)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s, age is %d", greetingsRequest.Name, greetingsRequest.Age),
		})
		return
	})

	// starting up gin web framework on port 8080
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
