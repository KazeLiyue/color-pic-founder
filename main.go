package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Hen struct {
	Err  string `json:"error"`
	Data []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		Uid        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		R18        bool     `json:"r18"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Tags       []string `json:"tags"`
		Ext        string   `json:"ext"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original,omitempty"`
			Regular  string `json:"regular,omitempty"`
			Small    string `json:"small,omitempty"`
			Thumb    string `json:"thumb,omitempty"`
			Mini     string `json:"mini,omitempty"`
		} `json:"urls"`
	} `json:"data"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/h", search)
	e.Static("/", "html")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func search(c echo.Context) error {
	resp, err := http.Get("https://api.lolicon.app/setu/v2?" + c.QueryString())
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	h, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	unjson := Hen{}
	err = json.Unmarshal(h, &unjson)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, unjson)
}
