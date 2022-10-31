package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/BrunoKrugel/mythic-data/internal/model"
	"github.com/labstack/echo"
	"github.com/zakaria-chahboun/cute"
)

func CharacterData(c echo.Context) error {
	resp, err := http.Get("https://raider.io/api/v1/characters/profile?region=us&realm=azralon&name=porend")
	if err != nil {
		cute.Println(err.Error(), cute.BrightRed)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cute.Println(err.Error(), cute.BrightRed)
	}
	char := &model.Characters{}
	_ = json.Unmarshal(body, &char)

	return c.JSON(http.StatusOK, char)
}
