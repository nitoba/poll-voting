package controllers

import "github.com/gin-gonic/gin"

type RegisterVoterController struct{}

func (*RegisterVoterController) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register voter",
	})
}

func NewRegisterVoterController() *RegisterVoterController {
	return &RegisterVoterController{}
}
