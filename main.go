package main

import "github.com/gin-gonic/gin"f

func main() {
	router := gin.Default()
	// usersGroup = router.Group("users")
	router.Run(":8000")
}
