package messages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

var (
	awsCredentials credentials.Credentials
	awsRegion      string
)

func init() {

	awsRegion = os.Getenv("AWS_REGION")

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")

	awsCredentials = *credentials.NewStaticCredentials(
		awsAccessKey, awsSecretKey, "")
}

func HandleMessage(ctx *gin.Context) {

	message := new(Message)
	bindErr := ctx.Bind(message)

	if bindErr != nil {
		log.Error(bindErr)
	}

	payload, marshallErr := json.Marshal(message)

	if marshallErr != nil {
		log.Error(marshallErr)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: &awsCredentials,
	})

	svc := lambda.New(sess, &aws.Config{Region: aws.String(awsRegion)})

	input := &lambda.InvokeInput{
		FunctionName: aws.String(viper.GetString("aws.function.name")),
		Payload:      payload,
	}

	req, res := svc.InvokeRequest(input)

	sendErr := req.Send()

	if sendErr != nil {
		log.Error(sendErr.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		log.Infof("[%s] - Message sent from %s", strconv.FormatInt(*res.StatusCode, 10), message.Name)
		ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
	}
}
