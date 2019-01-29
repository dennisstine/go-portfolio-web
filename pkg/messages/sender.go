package messages

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleMessage(ctx *gin.Context) {

	message := new(Message)
	bindErr := ctx.Bind(message)

	if bindErr != nil {

		log.Error(bindErr)
	}

	log.Info(message)
}
