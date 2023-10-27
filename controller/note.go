package controller

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/TODO-List/database"
	"github.com/OctavianoRyan25/TODO-List/helper"
	"github.com/OctavianoRyan25/TODO-List/model"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var(
	appJSON = "application/json"
)

func Store(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Note := model.Note{}
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Note)
	} else {
		ctx.ShouldBind(&Note)
	}

	Note.UserID = userID

	err := db.Debug().Create(&Note).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Err" 		: "Bad Request",
			"Message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Status": "Success create note",
		"Note"	: Note,
	})
}

func Update(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Note := model.NoteResponse{}
	NoteId, _ := strconv.Atoi(ctx.Param("noteId"))
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Note)
	} else {
		ctx.ShouldBind(&Note)
	}

	Note.UserID = userID
	Note.ID = uint(NoteId)

	err := db.Model(&Note).Where("id = ?", NoteId).Updates(model.NoteResponse{
		Title: Note.Title, 
		Slug: Note.Slug, 
		Body: Note.Body,
}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status"	:	"Success updating data",
		"Data"		:	Note,
	})
}

func Index(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64)) 

	var AllNotes []model.Note

	//photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	//userID := uint(userData["id"].(float64)) 

	err := db.Preload(clause.Associations).Find(&AllNotes, "user_id=?", userID).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Err" 		: "Bad Request",
			"Message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status"	: "Success",
		"Data"		: AllNotes,
	})
}


func GetByID(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Note := model.Note{}
	noteId, _ := strconv.Atoi(ctx.Param("noteId"))
	userID := uint(userData["id"].(float64)) 

	Note.UserID = userID
	Note.ID = uint(noteId)

	err := db.First(&Note).Where("id = ?", noteId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status"	: "Success",
		"Data"		: Note,
	})
}


func Destroy(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Note := model.Note{}
	noteId, _ := strconv.Atoi(ctx.Param("noteId"))
	userID := uint(userData["id"].(float64)) 

	Note.UserID = userID
	Note.ID = uint(noteId)

	err := db.First(&Note).Where("id = ?", noteId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	db.Delete(&Note)

	ctx.JSON(http.StatusOK, gin.H{
		"message"	: "Data has been deleted",
		"data"		: Note,
	})
}