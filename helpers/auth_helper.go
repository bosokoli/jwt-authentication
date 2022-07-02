package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, userType string) (err error) {
	userType = c.GetString("user_type")
	err = nil
	if userType != "ROLE" {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}
