package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type GreetingsRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

type GreetingsHeader struct {
	APIKey    string `header:"API-KEY" binding:"required"`
	ClientKey string `header:"CLIENT-KEY" binding:"required"`
	RequestID string `header:"REQUEST-ID" binding:"required"`
}

type GreetingsQueryParams struct {
	Email string `form:"email" binding:"required"`
}

//func BasicAuth() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		data := context.GetHeader("Authorization")
//
//		split := strings.Split(data, ":")
//		username := split[0]
//		password := split[1]
//
//		fmt.Println(username, password)
//		context.Next()
//	}
//}

func RequestIDGenerator() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := uuid.New().String()
		context.Set("RequestID", requestId)
		context.Next()
	}
}

func main() {
	// gin web framework initialization
	r := gin.New()
	r.Use(RequestIDGenerator())

	// request
	// - body (data data yang akan di berikan ke dalam request dan di prosess oleh backend)
	// - query param (data juga yang di berikan ke dalam request tetapi lewat query url)
	// - header (data tambahan yang di gunakan oleh backend untuk kebutuhan authentication dan formating)

	// handler creation for handling request to specific endpoint
	r.POST("/greetings", func(context *gin.Context) {
		// query
		var greetingQueryParam GreetingsQueryParams
		err := context.ShouldBindQuery(&greetingQueryParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// body
		var greetingRequest GreetingsRequest
		err = context.ShouldBindJSON(&greetingRequest)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// header
		var greetingHeader GreetingsHeader
		err = context.ShouldBindHeader(&greetingHeader)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// get request id from context
		requestId, _ := context.Get("RequestID")
		stringRequestId := requestId.(string)

		context.JSON(http.StatusOK, gin.H{
			"requestId": stringRequestId,
			"query":     greetingQueryParam,
			"data":      greetingRequest,
			"header":    greetingHeader,
		})
		return
	})

	// starting up gin web framework on port 8080
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
