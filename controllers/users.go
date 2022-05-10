package controllers

import (
	"github.com/Fix-Pay/reserva-mesa-sala/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(pass)
	err = user.Create()
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	user := models.User{Id: uint(id)}
	if err := user.Get(); err != nil {
		return err
	}
	return c.JSON(user)
}

func Auth(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	triedPassword := user.Password
	user.Password = ""
	if err := user.Get(); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(triedPassword)); err != nil {
		return err
	}
	atClaims := &models.UserClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}
	return c.JSON(models.Auth{
		Token: token,
		User:  user,
	})
}
