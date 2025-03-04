package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Initialize CSRF middleware
	csrfMiddleware := csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.Secure(false), // Use `true` in production for HTTPS
		csrf.Path("/"),     // CSRF cookie applies to the whole app
	)

	// Create an HTTP handler wrapped with CSRF middleware
	csrfHandler := csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Use a custom middleware to attach CSRF to Gin
	r.Use(func(c *gin.Context) {
		csrfHandler.ServeHTTP(c.Writer, c.Request)
		c.Next()
	})

	// Route to return the CSRF token
	r.GET("/", func(c *gin.Context) {
		// Retrieve the CSRF token
		csrfToken := csrf.Token(c.Request)
	
		// Return the token in the JSON response
		c.JSON(http.StatusOK, gin.H{
			"message":    "CSRF Token included",
			"csrf_token": csrfToken,
		})
	})
	
	// Protected route to validate CSRF token
	r.POST("/submit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request passed CSRF check",
		})
	})

	// Start the server
	r.Run(":8188")
}
