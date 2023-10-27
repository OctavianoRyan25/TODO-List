package auth

import (
	"log"
	"net/http"

	"github.com/OctavianoRyan25/TODO-List/database"
	"github.com/OctavianoRyan25/TODO-List/helper"
	"github.com/OctavianoRyan25/TODO-List/model"

	"github.com/gin-gonic/gin"
)

// var(
// 	appJSON = "application/json"
// )

func UserRegister(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)
	_, _ = db, contentType
	User := model.User{}
	err := ctx.ShouldBindJSON(&User)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	err = db.Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status"	: "Account succesfully created",
		"username"	: User.Username,
		"email"		: User.Email,
		"password"	: User.Password,
	})
}

func UserLogin(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)
	_, _ = db, contentType
	User := model.User{}
	password := ""
	err := ctx.ShouldBindJSON(&User)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}
	
	password = User.Password

	err = db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error"		: "Unauthorized",
			"message"	: "Invalid email",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid password",
		})
		return
	} 
	token := helper.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusCreated, gin.H{
		"status"	: "Login succesfully",
		"id"		: User.ID,
		"token"		: token,
	})
}