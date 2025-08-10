package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// สร้าง secret key สำหรับเซ็น token (ใน production ควรเก็บใน environment variable)
var jwtSecret = []byte("your-secret-key-here")

// Claims โครงสร้างข้อมูลที่จะเก็บใน JWT
type Claims struct {
	StaffID  uint   `json:"staff_id"`
	Hospital string `json:"hospital"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWTToken สร้าง JWT token
func GenerateJWTToken(staffID uint, hospital string, username string) (string, error) {
	// ตั้งค่าการหมดอายุของ token (ที่นี่ตั้งไว้ 24 ชั่วโมง)
	expirationTime := time.Now().Add(24 * time.Hour)

	// สร้าง claims
	claims := &Claims{
		StaffID:  staffID,
		Hospital: hospital,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "his-middleware",
		},
	}

	// สร้าง token ด้วย algorithm HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// เซ็น token ด้วย secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken ตรวจสอบและดึงข้อมูลจาก JWT token
func ValidateJWTToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
