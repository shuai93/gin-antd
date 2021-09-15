package ws

import (
	"backend/utils/logging"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Message struct {
	Id   int    `json: id`
	From string `json: from`
	Body string `json: body`
	To   string `json to`
}

func PushMessage(c *gin.Context, hub *Hub) {
	var message Message

	_ = c.BindJSON(&message)

	jsonBytes, _ := json.Marshal(message)
	logging.Info("Push message is %s " + string(jsonBytes))

	//body := []byte(message.Body)

	h := gin.H{
		"author": gin.H{
			"username":  message.From,
			"id":        message.Id,
			"avatarUrl": "https://image.flaticon.com/icons/svg/2446/2446032.svg",
		},
		"text":      message.Body,
		"type":      "text",
		"timestamp": time.Now().UnixNano() / 1e6,
	}
	bodyBytes := new(bytes.Buffer)
	json.NewEncoder(bodyBytes).Encode(h)

	for client := range hub.clients {
		if client.username == message.To {
			client.send <- bodyBytes.Bytes()
		}
	}

	c.JSON(http.StatusOK, message)

}

func GetMessage(c *gin.Context, hub *Hub) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "OK",
	})

}
