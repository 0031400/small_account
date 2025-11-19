package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Panicln(err)
	}
}
func Login(email, password string) (string, string, error) {
	id := 0
	salt := ""
	passwordHash := ""
	token := ""
	err := DB.QueryRow("SELECT id,password_salt,password_hash FROM users WHERE email = ?", email).Scan(&id, &salt, &passwordHash)
	if err == sql.ErrNoRows {
		return "", "the email doesn't exist", nil
	} else if err != nil {
		return "", "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password+salt))
	if err != nil {
		return "", "password is wrong", nil
	}
	err = DB.QueryRow("SELECT token FROM tokens WHERE user_id = ?", id).Scan(&token)
	if err != nil && err != sql.ErrNoRows {
		return "", "", err
	}
	if token != "" {
		return token, "", nil
	}
	token, err = GenerateRandom(128)
	if err != nil {
		return "", "", err
	}
	now := time.Now().Unix()
	_, err = DB.Exec("INSERT INTO tokens(user_id,token,created_at) VALUES (?,?,?)", id, token, now)
	if err != nil {
		return "", "", err
	}
	return token, "", err
}
func Register(email, password string) (string, error) {
	exist := 0
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&exist)
	if err != nil {
		return "", err
	}
	if exist > 0 {
		return "the email exists", nil
	}
	salt, err := GenerateRandom(16)
	if err != nil {
		return "", err
	}
	passwordHash, err := HashPassword(password, salt)
	if err != nil {
		return "", err
	}
	now := time.Now().Unix()
	_, err = DB.Exec("INSERT INTO users(email,username,password_salt,password_hash,created_at) VALUES (?,?,?,?,?)", email, email, salt, passwordHash, now)
	if err != nil {
		return "", err
	}
	return "", nil
}

func HashPassword(password, salt string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func GenerateRandom(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
