package middlewares
import 
func BasicAuth() gin.HandlerFunc {
return gin.BasicAuth(gin.Accounts{
	"pragmatic" ; "reviews"
})
}