package auth

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err

}
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	log.Printf("User Type: %s", userType)
	uid := c.GetString("uid")
	log.Printf("User id: %s", uid)
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("User not matched")
		return err
	}

	if userType != "USER" {
		if err := CheckUserType(c, userType); err != nil {
			return err
		}
	}
	return err
}
