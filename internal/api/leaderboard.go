package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BrunoKrugel/mythic-data/internal/model"
	"github.com/labstack/echo"
	"github.com/zakaria-chahboun/cute"
)

// clientID := os.Getenv("CLIENT_ID")
// clientSecret := os.Getenv("CLIENT_SECRET")
// accessToken := os.Getenv("ACCESS_TOKEN")

func LeaderboardHorde(c echo.Context) error {

	accessToken := os.Getenv("ACCESS_TOKEN")
	url := fmt.Sprintf("https://us.api.blizzard.com/data/wow/leaderboard/hall-of-fame/uldir/horde?namespace=dynamic-us&locale=en_US&access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		cute.Println(err.Error(), cute.BrightRed)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cute.Println(err.Error(), cute.BrightRed)
	}
	lb := &model.Leaderboard{}
	_ = json.Unmarshal(body, &lb)

	return c.JSON(http.StatusOK, lb)
}
