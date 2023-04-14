package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"rest/pkg/service"
)

type Person struct {
	Name      string `json:"Name"`
	Url       string `json:"Url"`
	Extension string `json:"Extension"`
}

func (h *Handler) download(c *gin.Context) {
	jsons := make([]byte, c.Request.ContentLength)
	if _, err := c.Request.Body.Read(jsons); err != nil {
		if err.Error() != "EOF" {
			return
		}
	}

	var person Person

	err := json.Unmarshal(jsons, &person)
	if err != nil {
		panic(err)
	}
	service.DownloadFile("test.zip", person.Url)

	//fmt.Printf("Name: %s\nUrl: %s\nExtension: %s\n", person.Name, person.Url, person.Extension)

	c.String(200, "Hello %s", c.Param("name"))
}
