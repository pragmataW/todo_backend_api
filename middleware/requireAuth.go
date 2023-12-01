package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pragmataW/to-do/endpoints"
)

func RequireAuth(c *fiber.Ctx) error {
	//*Get auth cookie
	auth := c.Cookies("Authentication")
	if auth == ""{
		return c.SendStatus(http.StatusUnauthorized)
	}
	
	//* Decode / validate jwt
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(endpoints.JwtPass), nil
	})
	
	if err != nil{
		fmt.Println(err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		//* Check the exp
		if float64(time.Now().Unix()) > (claims["exp"].(float64)){
			return c.SendStatus(http.StatusUnauthorized)
		}
		//* Set important data to fiber local
		c.Locals("name", claims["Name"])
		c.Locals("surname", claims["Surname"])
		c.Locals("username", claims["Username"])
		return c.Next()
	}

	return c.SendStatus(http.StatusUnauthorized)
}