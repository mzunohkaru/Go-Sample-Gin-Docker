package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// 本の一覧を取得する関数
func getBooks(c *gin.Context) {
	// HTTPリクエストが来たときに、本の一覧を取得してJSON形式で返す
	c.IndentedJSON(http.StatusOK, books)
}

// 特定のIDの本を取得する関数
func bookById(c *gin.Context) {
	// クエストから本のIDを取得
	id := c.Param("id")
	
	// 指定されたIDの本を検索
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// 本を借りる処理
func checkoutBook(c *gin.Context) {
	// HTTPリクエストからidクエリパラメータを取得
	id, ok := c.GetQuery("id")

	// idが存在しない場合
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	// 指定されたIDの本を検索
	book, err := getBookById(id)

	// 本が見つからない場合
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	// 在庫がない場合
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	// 本が存在し、かつ在庫がある場合
	// 在庫を1つ減らします
	book.Quantity -= 1

	// 更新された本の情報を200ステータスコード（OK）と共にJSON形式で返します
	c.IndentedJSON(http.StatusOK, book)
}

// 本を返却する処理
func returnBook(c *gin.Context) {
	// HTTPリクエストからidクエリパラメータを取得
	id, ok := c.GetQuery("id")

	// idが存在しない場合
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	// 指定されたIDの本を検索
	book, err := getBookById(id)

	// 本が見つからない場合
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	// 本が存在し、かつ在庫がある場合
	// 在庫を1つ増やします
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// IDに対応する本を取得する関数
func getBookById(id string) (*book, error) {
	// b のIDが指定されたIDと一致するかを確認
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main_second() {
	router := gin.Default()
	// URLに対応する処理を設定
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)

	// サーバーを起動
	router.Run(":8080")
}
