package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "firstWeb/db"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Key    string `json:"key"`
	Status string `json:"status"`
}

func SetRedis(c *gin.Context) {

	err := db.Client.LPush("sender", "{\"key\":\"88\",\"status\":\"fail\"}").Err()
	c.JSON(http.StatusOK, gin.H{"Data sended": http.StatusOK})
	if err != nil {
		panic(err)
	}
}

func GetRedis(c *gin.Context) {

	lenIcc, err := db.Client.LLen("sender").Result()
	if err != nil {
		panic(err)
	}

	intLen := int(lenIcc)
	for i := 0; i < intLen; i++ {
		Iint64 := int64(i)
		valueIcc, err := db.Client.LIndex("sender", Iint64).Bytes()

		if err != nil {
			panic(err)
		}

		var iccList Data
		err = json.Unmarshal(valueIcc, &iccList)
		if err != nil {
			fmt.Println("JSON ERR")
			return
		} else {
			json := gin.H{
				"Ключ":            iccList.Key,
				"Статус рассылки": iccList.Status,
			}
			c.JSON(http.StatusOK, json)
		}
	}
}

func ReWriteData(c *gin.Context) {
	lenIcc, err := db.Client.LLen("sender").Result()
	if err != nil {
		panic(err)
	}

	intLen := int(lenIcc)
	for i := 0; i < intLen; i++ {
		Iint64 := int64(i)
		valueIcc, err := db.Client.LIndex("sender", Iint64).Bytes()

		if err != nil {
			panic(err)
		}

		var iccList Data
		err = json.Unmarshal(valueIcc, &iccList)
		if err != nil {
			fmt.Println("JSON ERR")
			return
		}

		if iccList.Key == "88" {

			db.Client.Set("")
			json := gin.H{
				"Статус рассылки": iccList.Status,
			}
			c.JSON(http.StatusOK, json)
		}
	}
}
