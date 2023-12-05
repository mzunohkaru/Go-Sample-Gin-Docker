package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// 各フィールド( Id と Name )が、どのフィールド( リクエストのJson )に対応するか記述
}

func postUsers(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// ShouldBindJSON : フロントエンドから送られてきたデータの中身を取得する
		// User構造体にリクエストに入っているJsonオブフェクトをバインドさせる
		// バインドが失敗した時、エラーを返す
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": err.Error(),
		})
		return
	}

	//データ　登録　Postmanで確認

	//データ　取得
	c.JSON(http.StatusMovedPermanently, gin.H{
		"message": "User created successfully", "data": user,
	})
}

func getSamples(c *gin.Context) {
	c.JSON(http.StatusMovedPermanently, gin.H{
		"message": "サンプルAPI",
		"int":     1,
		"map": map[string]int{
			"太郎":  21,
			"みどり": 34,
		},
	})
}

func getVersionProduct(c *gin.Context) {
	// category := c.Query("category")
	// Queryメソッド : 値がない場合 nil

	category := c.DefaultQuery("category", "デフォルトの値")
	// DefaultQueryメソッド : 値がない場合 デフォルト値を設定できる

	c.JSON(http.StatusMovedPermanently, gin.H{"category": category})
}

func putUsers(c *gin.Context) {
	// :userID : パスパラメータ( URLの値の一部を可変にできる )で一意に識別する情報を受け取る
	userID := c.Param("userID")
	// パスに入って来た ID を受け取る

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": err.Error(),
		})
		return
	}

	//データ　更新

	//データ　取得
	c.JSON(http.StatusMovedPermanently, gin.H{
		"message": "User updated successfully", "id": userID,
	})
}

func deleteUsers(c *gin.Context) {
	userID := c.Param("userID")

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": err.Error(),
		})
		return
	}
	//データ　取得
	c.JSON(http.StatusMovedPermanently, gin.H{
		"message": "User updated successfully", "id": userID,
	})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	// URLグルーピング
	// グルーピング : バージョンごとに管理する（APIの使用が簡単にかわらないように）
	// GoogleドライブのAPI : drive/v3/about

	// http://localhost:8080/users
	router.POST("/users", postUsers)
	// http://localhost:8080/sample
	router.GET("/sample", getSamples)

	// http://localhost:8080/v1/products?category=book  <- グルーピングしたので、　/v1/　が入る
	v1.GET("/products", getVersionProduct)

	// http://localhost:8080/redirect
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// 更新のメソッドは2つ PUTメソッド & PATCHメソッド
	// PUT リソースの全てを更新
	// PATCH リソースの一部を更新
	// この場合は、リソース( User )の( Name )を更新する
	// http://localhost:8080/users/1  <- /1 は userID のパラメーター
	router.PUT("/users/:userID", putUsers)

	//データ　削除
	// http://localhost:8080/users/1
	router.DELETE("/users/:userID", deleteUsers)

	router.Run(":8000")
}

// http://localhost:8080/
