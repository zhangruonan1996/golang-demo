package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type idol struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	Birthday string   `json:"birthday"`
	Works    []string `json:"works"`
}

var idols = []idol{
	{ID: "00001", Name: "章若楠", Gender: "女", Birthday: "1996-11-14", Works: []string{"悲伤逆流成河", "你的婚礼", "如果声音不记得", "难哄"}},
	{ID: "00002", Name: "白敬亭", Gender: "男", Birthday: "1993-10-15", Works: []string{"匆匆那年", "开端", "你是我的城池营垒", "难哄"}},
	{ID: "00003", Name: "杨洋", Gender: "男", Birthday: "1991-09-09", Works: []string{"微微一笑很倾城", "全职高手", "左耳", "你是我的荣耀"}},
}

func getIdols(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, idols)
}

func postIdol(c *gin.Context) {
	var newIdol idol

	if err := c.BindJSON(&newIdol); err != nil {
		return
	}

	idols = append(idols, newIdol)
	c.IndentedJSON(http.StatusCreated, newIdol)
}

func getIdolById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range idols {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/idols", getIdols)
	router.GET("/idols/:id", getIdolById)
	router.POST("/idols", postIdol)

	router.Run("localhost:8080")
}