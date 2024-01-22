package server

import (
	"fmt"
	"github.com/disism/godis/jwt"
	"github.com/disism/karma/ent"
	"github.com/disism/karma/ent/users"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

const (
	UsernameAlreadyExists      = "Username already exists"
	UserDoesNotExist           = "User does not exist"
	PasswordVerificationFailed = "Password verification failed"
)

// CreateUserForm create user form.
type CreateUserForm struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=32"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

func (s *Server) CreateUser() error {
	defer s.client.Close()

	var f CreateUserForm
	if err := s.ctx.ShouldBind(&f); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	// check username is exist.
	exist, err := s.client.Users.Query().Where(users.UsernameEQ(f.Username)).Exist(s.ctx)
	if err != nil {
		return fmt.Errorf("check username is exist error: %w", err)
	}
	if exist {
		ErrorConflict(s.ctx, UsernameAlreadyExists)
		return nil
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	create, err := s.client.Users.
		Create().
		SetUsername(f.Username).
		SetPassword(string(hash)).
		Save(s.ctx)
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}

	device, err := s.client.Devices.Create().
		SetUser(create).
		SetIP(s.ctx.ClientIP()).
		SetDevice(s.ctx.Request.UserAgent()).
		Save(s.ctx)
	if err != nil {
		return fmt.Errorf("create user create device error: %w", err)
	}

	k, err := jwt.New(
		JWTIssues(),
		strconv.FormatUint(create.ID, 10),
		jwt.WithID(strconv.FormatUint(device.ID, 10)),
	).Generate(JWTSecret())
	if err != nil {
		return fmt.Errorf("create user jwt generate token error: %w", err)
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"access_token": k,
	})
	return nil
}

// LoginForm login form.
type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=32"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

func (s *Server) Login() error {
	defer s.client.Close()

	var f LoginForm
	if err := s.ctx.ShouldBind(&f); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Users.Query().Where(users.UsernameEQ(f.Username)).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorUnauthorized(s.ctx, UserDoesNotExist)
			return nil
		}
		return fmt.Errorf("query user error: %w", err)
	}

	// CompareHashAndPassword to compare whether the password is correct or not.
	if err := bcrypt.CompareHashAndPassword([]byte(query.Password), []byte(f.Password)); err != nil {
		ErrorUnauthorized(s.ctx, PasswordVerificationFailed)
		return nil
	}

	device, err := s.client.Devices.Create().
		SetUser(query).
		SetIP(s.ctx.ClientIP()).
		SetDevice(s.ctx.Request.UserAgent()).
		Save(s.ctx)
	if err != nil {
		return fmt.Errorf("create user create device error: %w", err)
	}

	k, err := jwt.New(
		JWTIssues(),
		strconv.FormatUint(query.ID, 10),
		jwt.WithID(strconv.FormatUint(device.ID, 10)),
	).Generate(JWTSecret())
	if err != nil {
		return fmt.Errorf("login jwt generate token error: %w", err)
	}

	s.ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"access_token": k,
	})
	return nil
}
