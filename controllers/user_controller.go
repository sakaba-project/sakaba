package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"

	db "github.com/sakaba-project/sakaba/db"
	"github.com/sakaba-project/sakaba/models"
)

type UserController struct{}

type userJson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserController() UserController {
	return UserController{}
}

func (uc UserController) Create(c *gin.Context) {
	conn := db.Connect()
	defer conn.Close()

	in := userJson{}

	if err := c.BindJSON(&in); err != nil {
		panic(err.Error())
	}
	u := models.User{
		Name:  in.Name,
		Email: in.Email,
	}

	conn.NewRecord(&u)
	if result := conn.Create(&u); result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, u)
}

func (uc UserController) Index(c *gin.Context) {
	conn := db.Connect()
	defer conn.Close()

	var u []models.User
	if result := conn.Find(&u); result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, u)
}

func (uc UserController) Show(c *gin.Context) {
	conn := db.Connect()
	defer conn.Close()

	var u []models.User
	id := c.Params.ByName("id")

	if result := conn.Where("id = ?", id).Find(&u); result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, u)
}

func (uc UserController) Delete(c *gin.Context) {
	conn := db.Connect()
	defer conn.Close()

	var u []models.User
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)

	if result := conn.Where("id = ?", idInt).Delete(&u); result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, u)
}
