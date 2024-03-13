// handlers/user.go

package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/models"
)

var users []models.User

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	// URL パラメータから ID を取得
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		// パラメータが不正な場合のエラーレスポンス
		c.JSON(400, gin.H{"message": "Invalid user ID"})
		return
	}

	// ユーザーを検索
	user, found := findUserByID(userID)
	if found {
		// ユーザーが見つかった場合
		c.JSON(200, user)
	} else {
		// ユーザーが見つからなかった場合
		c.JSON(404, gin.H{"message": "User not found"})
	}
}

// findUserByID は指定された ID のユーザーを users スライスから検索するヘルパー関数
func findUserByID(userID int) (models.User, bool) {
	for _, user := range users {
		if user.ID == userID {
			return user, true
		}
	}
	return models.User{}, false
}

func CreateUser(c *gin.Context) {
	// リクエストからユーザー情報を取得
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		// JSON ボディのパースエラーがある場合のエラーレスポンス
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	// ユーザーの作成ロジックを実行
	createdUser := createUser(newUser)

	// 作成されたユーザーを返す
	c.JSON(201, createdUser)
}

// createUser は新しいユーザーを作成し、users スライスに追加するヘルパー関数
func createUser(newUser models.User) models.User {
	// 例: ユーザーに一意の ID を割り当てる（実際のアプリケーションでは ID の生成方法を適切に考慮する必要があります）
	newUser.ID = len(users) + 1

	// ユーザーを users スライスに追加
	users = append(users, newUser)

	return newUser
}

func UpdateUser(c *gin.Context) {
	// URL パラメータから ID を取得
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		// パラメータが不正な場合のエラーレスポンス
		c.JSON(400, gin.H{"message": "Invalid user ID"})
		return
	}

	// ユーザーを検索
	existingUser, found := findUserByID(userID)
	if !found {
		// ユーザーが見つからない場合のエラーレスポンス
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	// リクエストからユーザー情報を取得
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		// JSON ボディのパースエラーがある場合のエラーレスポンス
		c.JSON(400, gin.H{"message": "Invalid JSON format"})
		return
	}

	// ユーザー情報を更新
	updatedUser.ID = existingUser.ID
	updateUser(existingUser, updatedUser)

	// 更新されたユーザーを返す
	c.JSON(200, updatedUser)
}

// updateUser は既存のユーザー情報を更新するヘルパー関数
func updateUser(existingUser models.User, updatedUser models.User) {
	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
}

func DeleteUser(c *gin.Context) {
	// URL パラメータから ID を取得
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		// パラメータが不正な場合のエラーレスポンス
		c.JSON(400, gin.H{"message": "Invalid user ID"})
		return
	}

	// ユーザーを検索
	_, found := findUserByID(userID)
	if !found {
		// ユーザーが見つからない場合のエラーレスポンス
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	// ユーザーを削除
	deleteUser(userID)

	// ユーザー削除が成功した場合のレスポンス
	c.JSON(200, gin.H{"message": "User deleted"})
}

// deleteUser は指定された ID のユーザーを users スライスから削除するヘルパー関数
func deleteUser(userID int) {
	for i, user := range users {
		if user.ID == userID {
			// ユーザーを削除
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}
