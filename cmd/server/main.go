package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
	"github.com/timkelleher.com/discord-webhooks/pkg/integrations/cronicle"
	"github.com/timkelleher.com/discord-webhooks/pkg/integrations/ferry"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "pong",
		})
	})
	r.POST("/cronicle", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		var data cronicle.CronicleRequest
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		payload := cronicle.NewPayload(data)
		// todo: split up conf/paylod structs
		discord.NewBackupWebHook(payload, payload)

		c.JSON(200, gin.H{
			"result": "success",
		})
	})

	r.POST("/ferry", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		var data ferry.FerryRequest
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		payload := ferry.NewPayload(data)
		// todo: split up conf/paylod structs
		discord.NewGenericWebHook(payload, payload)

		c.JSON(200, gin.H{
			"result": "success",
		})
	})

	r.Run("0.0.0.0:3000")
}
