package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generates Session ID
func generateSessionID() string {
	return fmt.Sprintf("%x", rand.Int63())
}

// Middleware for sessions
func sessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			sessionID = generateSessionID()
			c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)
		}
		c.Request.Header.Set("X-Session-ID", sessionID)
		c.Next()
	}
}

// Redirects with SessionID to Cart Service
func forwardToCartService(c *gin.Context) {
	url := fmt.Sprintf("http://localhost:8081%s", c.Request.URL.Path)
	req, _ := http.NewRequest(c.Request.Method, url, c.Request.Body)
	req.Header.Set("X-Session-ID", c.GetHeader("X-Session-ID"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cart service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

// Protect with port the Origin
func originHasAllowedPort(origin string, allowedPort int) bool {
	return fmt.Sprintf(":%d", allowedPort) == origin[len(origin)-len(fmt.Sprintf(":%d", allowedPort)):]
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")

		if origin != "" && originHasAllowedPort(origin, 3005) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		}

		// Manejo de preflight OPTIONS
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.Use(sessionMiddleware())
	r.GET("/cart", forwardToCartService)
	r.GET("/cart/add/:product_id", forwardToCartService)

	r.Run(":8080")
}
