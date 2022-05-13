package controllers

import (
	"fixpay/reserva-mesa-sala/data"
	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}
type LoginRequest struct {
	Email    string
	Password string
}

func SignIn(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadGateway, "Valores inseridos invalidos")
	}

	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	user := new(data.User)
	has, err := engine.Where("email = ?", req.Email).Desc("id").Get(user)
	if err != nil {
		return err
	}
	if !has {
		return fiber.NewError(fiber.StatusBadRequest, "Login invalido")
	}
	token, exp, err := creatJWTToken(*user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func SignUp(c *fiber.Ctx) error {
	req := new(SignupRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if req.Email == "" || req.Name == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadGateway, "Valores inseridos invalidos")
	}

	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	// Encriptando senha
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &data.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	_, err = engine.Insert(&user)
	if err != nil {
		return err
	}
	token, exp, err := creatJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func creatJWTToken(user data.User) (string, int64, error) {
	// Tempo válido do Token
	exp := time.Now().Add(time.Minute * 30).Unix()

	//Criando Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
