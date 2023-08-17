package auth

import (
	"fmt"
	"time"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/Daaaai0809/grpc-auth-practice/data"
	jwt "github.com/dgrijalva/jwt-go"
	proto "github.com/Daaaai0809/grpc-auth-practice/proto"
)

func GenerateToken(userName string, role data.Role) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := token.Claims.(jwt.MapClaims)
	claims["userName"] = userName
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte("secret"))
}

func CheckPassword(userName string, password string) (data.Role, error) {
	for _, user := range data.Users {
		if user.UserName != userName {
			continue
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return data.Guest, err
		}

		return user.Role, nil
	}

	return data.Guest, errors.New("user not found")
}

func CheckAdminToken(req *proto.AdminRequest) error {
	token := req.GetToken()

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	exp := claims["exp"].(float64)
	now := float64(time.Now().Unix())
	if exp < now {
		return errors.New("token expired")
	}

	role := claims["role"].(data.Role)

	if role != data.Admin {
		return errors.New("permission denied")
	}

	return nil
}

func CheckUserToken(req *proto.RequiredAuthRequest) error {
	token := req.GetToken()

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	exp := claims["exp"].(float64)
	now := float64(time.Now().Unix())
	if exp < now {
		return errors.New("token expired")
	}

	role := claims["role"].(data.Role)

	if role != data.Admin && role != data.User {
		return errors.New("permission denied")
	}

	return nil
}
