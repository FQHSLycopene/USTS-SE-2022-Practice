package utils

import "github.com/golang-jwt/jwt/v4"

var mySigningKey = []byte("Go-Bill")

type MyClaims struct {
	jwt.RegisteredClaims
	Identity string
	Status   int
}

func GenerateToken(identity string, status int) (string, error) {

	c := MyClaims{
		Identity: identity,
		Status:   status,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析Token
func AnalyseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*MyClaims), err
}
