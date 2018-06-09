package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jekabolt/slf"
)

const (
	msgErrHeaderError = "wrong authorization headers"
)

type RestClient struct {
	// userStore store.UserStore
	log slf.StructuredLogger
}

func SetRestHandlers(
	// userDB store.UserStore,
	r *gin.Engine,
) (*RestClient, error) {
	restClient := &RestClient{
		// userStore: userDB,
		log: slf.WithContext("rest-client"),
	}
	r.POST("/graffity", restClient.postGraffity())
	r.POST("/near/graffity/", restClient.nearGraffitys())
	r.GET("/map/zones/", restClient.getMapZones())

	return restClient, nil
}

func (restClient *RestClient) postGraffity() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getIdenity(c)
		if err != nil {
			restClient.log.Errorf("addWallet: getIdenity: %s\t[addr=%s]", err.Error(), c.Request.RemoteAddr)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrHeaderError,
			})
			return
		}
		fmt.Println("token ", token)

	}
}

func (restClient *RestClient) nearGraffitys() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (restClient *RestClient) getMapZones() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func getIdenity(c *gin.Context) (string, error) {
	authHeader := strings.Split(c.GetHeader("Authorization"), " ")
	if len(authHeader) < 2 {
		return "", errors.New(msgErrHeaderError)
	}
	return authHeader[1], nil
}
