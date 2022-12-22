package chess

import (
	"fmt"
	"net/http"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func ChessGroup(group *gin.RouterGroup) {
	group.GET("/:user/:game/:turn", getChessMove)
}

// getChessMove godoc
// @Summary Get all suggested moves for a given #birdchess tweet
// @Tags    chess
// @Produce json
// @Param   user path     string      true "player 1's username"
// @Param   game path     string      true "game identifier, i.e. its starting instant" Format(date-time)
// @Param   turn path     int         true "second player's turn number"                minimum(1)
// @Success 200  {string} string      "The second player's move. It is recognized by the [a-h][1-8][a-h][1-8] regexp."
// @Success 204  {string} string      "No one has played yet"
// @Failure 400  {object} model.Error "integer parsing error on turn or turn <= 0"
// @Failure 404  {object} model.Error "No user or no post found"
// @Router  /chess/{user}/{game}/{turn} [get]
func getChessMove(ctx *gin.Context) {
	username := ctx.Param("username")
	date := ctx.Param("game")
	turn, err := strconv.Atoi(ctx.Param("turn"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "Integer parsing error on 'turn' parameter",
		})
		return
	}
	if turn < 1 {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "turn < 1",
		})
		return
	}
	res, error := uncheckedGetCheckMove(username, date, turn)
	if error == nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(error.Code, model.Error{
			Code:    error.Code,
			Message: error.Message,
		})
	}
}

func uncheckedGetCheckMove(username string, date string, turn int) (string, *model.Error) {
	tweets, err := twitter.GetTweetsFromUser(username, 100, date)
	if err != nil {
		return "", &model.Error{Code: http.StatusNotFound, Message: err.Error()}
	}
	length := len(tweets.Data)
	if length < turn {
		return "", &model.Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Wanted to retrieve turn #%d. Tweets not found: %d.", turn, length),
		}
	}
	return mostPopularChessMove(tweets.Data[length-turn]), nil
}

func mostPopularChessMove(tweet twitter.ProfileTweet) string {
	return "e7e5"
}
