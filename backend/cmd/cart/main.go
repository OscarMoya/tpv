package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// simple placeholder for the cart table
var cartStore sync.Map

// SessionID  middleware
func getSessionID(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing session"})
		c.Abort()
		return
	}
	c.Set("session_id", sessionID)
	c.Next()
}

// Add to cart
func addToCart(c *gin.Context) {
	sessionID := c.GetString("session_id")
	productID := c.Param("product_id")

	cart, _ := cartStore.LoadOrStore(sessionID, []string{})
	items := cart.([]string)

	items = append(items, productID)
	cartStore.Store(sessionID, items)

	c.JSON(http.StatusOK, gin.H{"message": "Product added", "cart": items})
}

// Get cart from session ID
func getCart(c *gin.Context) {
	sessionID := c.GetString("session_id")

	if cart, ok := cartStore.Load(sessionID); ok {
		c.JSON(http.StatusOK, gin.H{"cart": cart})
	} else {
		c.JSON(http.StatusOK, gin.H{"cart": []string{}})
	}
}

func main() {
	r := gin.Default()

	authRoutes := r.Group("/")
	authRoutes.Use(getSessionID)
	authRoutes.GET("/cart", getCart)
	authRoutes.GET("/cart/add/:product_id", addToCart)

	r.Run(":8081")
}
