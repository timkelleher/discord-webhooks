package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/timkelleher.com/discord-webhooks/pkg/cronicle"
	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
)

func main() {
	r := gin.Default()
	r.POST("/cronicle", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		var data cronicle.WebHookRequest
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		payload := cronicle.NewPayload(data)
		// todo: split up conf/paylod structs
		discord.NewWebHook(payload, payload)

		c.JSON(200, gin.H{
			"result": "success",
		})
	})
	r.Run("0.0.0.0:9999")
}
