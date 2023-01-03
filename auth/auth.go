package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/gdsc-ys/sprint4-jira-thread/auth/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func TokenAuthMiddleware(c *gin.Context) {
	token, err := c.Request.Cookie("access-token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication failed",
		})
		c.Abort()
		return
	}

	conn, err := grpc.Dial("grpc.jimmy0006.site:3001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	co := pb.NewUserStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.Get(ctx, &pb.Token{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppbjBtaW5AeW9uc2VpLmFjLmtyIiwiZXhwIjoxNjczNjI2NTE4fQ.sDqZOMfi-koEf19dGVFKPWQCMMIPZCUwstpW92C11xg"})
	r, err := co.Get(ctx, &pb.Token{Token: token.Value})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	c.Set("userid", r.GetId())
	c.Next()
}
