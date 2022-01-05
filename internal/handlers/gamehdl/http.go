package gamehdl

import (
	"go-hex-example/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	gameController ports.GameController
}

func NewHTTPHandler(gameController ports.GameController) *HTTPHandler {
	return &HTTPHandler{
		gameController: gameController,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	game, err := hdl.gameController.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
	}

	c.JSON(200, game)
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	game, err := hdl.gamesService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildResponseCreate(game))
}

func (hdl *HTTPHandler) RevealCell(c *gin.Context) {
	body := BodyRevealCell{}
	c.BindJSON(&body)

	game, err := hdl.gamesService.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildResponseRevealCell(game))
}
