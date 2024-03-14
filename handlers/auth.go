// handlers/auth_handler.go

package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/models"
)

func SignIn(c *gin.Context) {
	var signInRequest struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここでは単純なユーザー名とパスワードでの認証を行います
	// 本番環境では、データベースや他の認証機構を使用することが推奨されます

	// 仮のユーザー名とパスワードを定義
	validUsername := validUsername
	validPassword := validPassword

	// リクエストされたユーザー名とパスワードを検証
	if signInRequest.Name == validUsername && signInRequest.Password == validPassword {
		// 認証成功時の処理
		// トークンを生成して返す
		token, _ := generateToken(signInRequest.Name, signInRequest.Password)

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		// 認証失敗時の処理
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

var (
	validUsername = os.Getenv("VALID_USERNAME")
	validPassword = os.Getenv("VALID_PASSWORD")
	jwtKey        = []byte(os.Getenv("JWT_KEY"))
)

func generateToken(name, password string) (string, error) {
	// トークンの有効期限を設定（例: 24時間）
	expirationTime := time.Now().Add(24 * time.Hour)

	// トークンのペイロードを作成
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   name,
	}

	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を追加
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SignUp(c *gin.Context) {
	// リクエストから新しいユーザーの情報を取得
	var signUpRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// データベースに新しいユーザーを登録（仮の登録ロジック）
	// ここでは単純に成功したものとして、登録されたユーザー情報を返す
	newUser := saveUserToDatabase(signUpRequest)

	// ユーザーの登録に成功した場合のレスポンス
	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": newUser})
}

// saveUserToDatabase は新しいユーザーをデータベースに登録する仮の関数です
func saveUserToDatabase(signUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}) models.User {
	// ここでデータベースに新しいユーザーを登録するロジックを実装する
	// 今回は仮のロジックなので、新しいユーザーを作成して返すだけとします
	return models.User{
		Name:     signUpRequest.Name,
		Password: signUpRequest.Password,
		// 他のユーザー情報も適切に設定する
	}
}
