package service

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"insert_DM/cnf"
	"insert_DM/domain"
	"insert_DM/domain/dto"
	"insert_DM/repository"
	"insert_DM/utils"
	"net/http"
	"time"
)

type UserServiceImpl struct {
	DB             *sql.DB
	userRepository repository.UserRepository
}

func NewUserService(DB *sql.DB, userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{DB: DB, userRepository: userRepository}
}

func (service UserServiceImpl) Register(ctx *fiber.Ctx, register dto.UserRegister) dto.WebRes {
	var responses dto.WebRes

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	//# HASH PASSWORD
	hashPsw, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil { // <-- response ketika gagal hash
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data: map[string]interface{}{
				"Error": "Error Hash Password",
			},
		}
		return responses
	}
	register.Password = string(hashPsw)

	//#MAPPING REQUEST
	createAcc := domain.User{
		Username: register.Username,
		Password: register.Password,
	}
	err = service.userRepository.Register(ctx.Context(), tx, createAcc)
	if err != nil {
		if err.Error() == fmt.Sprintf("username %s already exists", register.Username) {
			responses = dto.WebRes{
				Code:   http.StatusConflict, // Status code untuk konflik
				Status: "Conflict",
				Data:   "Username already exists",
			}
			return responses
		} else {
			responses = dto.WebRes{
				Code:   http.StatusInternalServerError,
				Status: "InternalServerError",
				Data:   err.Error(),
			}
			return responses
		}
	}

	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	//IF SUCCESS
	responses = dto.WebRes{
		Code:   http.StatusCreated,
		Status: "StatusCreated",
		Data:   "success created account",
	}

	return responses
}

func (service UserServiceImpl) Login(ctx *fiber.Ctx, login dto.UserRequestLogin) dto.WebRes {
	var responses dto.WebRes

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	//#MAPPING DATA
	ReqLogin := domain.User{
		Username: login.Username,
		Password: login.Password,
	}
	err = service.userRepository.Login(ctx.Context(), tx, ReqLogin)
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "BadRequest",
			Data:   err.Error(),
		}
		return responses
	}
	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	//#RESPONSE SUCCESS
	responses = dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Login Success",
	}

	//#GENERATE JWT TOKEN
	expTime := time.Now().Add(time.Minute * 2) // <-- token kadaluarsa dalam 2min
	claimToken := &cnf.JWTClaim{
		UserName: login.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "king_jamal",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//#ALGORITM FOR JWT
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	//# SIGN TOKEN
	token, err := tokenAlgo.SignedString([]byte(cnf.JWT_KEY))
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   "Failed to generate token",
		}
		return responses
	}

	//#SET TOKEN KE COOKIE
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
	})

	return responses
}

func (service UserServiceImpl) Logout(ctx *fiber.Ctx) dto.WebRes {
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	responses := dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Logout Success",
	}

	return responses
}
