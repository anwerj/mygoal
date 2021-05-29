package main

import (
	"fmt"
	"log"
	"net/http"

	goal "github.com/anwerj/mygoal/goal"
	gin "github.com/gin-gonic/gin"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := goal.NewSelfServiceClient(conn)

	g := gin.Default()
	g.GET("/self/:action/:a/:b", func(ctx *gin.Context) {
		action, a, b := ctx.Param("action"), ctx.Param("a"), ctx.Param("b")

		req := &goal.Name{First: a, Last: b}
		response, err := client.Welcome(ctx, req)

		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"action":  action,
				"message": fmt.Sprint(response.Result),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
	})

	if err := g.Run("localhost:4041"); err != nil {
		log.Fatalf("Failed to start the client server: %v", err)
	}
}
