package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// authentication and authorization DOCTOR
const roles = "doctor"

// check "role" only
func DoctorRoleAuth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		token, err := ExtractToken(c)
		if err != nil {
			return err
		}

		if token.Valid {
			userClaims := token.Claims.(jwt.MapClaims)
			userRole := userClaims["role"].(string)

			if userRole == roles {
				return next(c)
			} else {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "You are not authorized to access this resource!",
				})
			}
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid or expired token!",
			})
		}
	}
}

// check "role" and "id"
func DoctorIDRoleAuth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		token, err := ExtractToken(c)
		if err != nil {
			return err
		}

		if token.Valid {
			userClaims := token.Claims.(jwt.MapClaims)
			userRole := userClaims["role"].(string)
			userID := int(userClaims["id"].(float64))

			if userRole == roles {
				c.Set("userID", userID)
				return next(c)
			} else {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "You are not authorized to access this resource!",
				})
			}
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid or expired token!",
			})
		}
	}
}
