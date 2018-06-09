package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jekabolt/ARGraffti-back/store"
	"github.com/jekabolt/slf"
)

const (
	msgErrHeaderError      = "wrong authorization headers"
	msgErrRequestBodyError = "missing request body params"
	msgErrNoSuchUser       = "missing request body params"
)

type RestClient struct {
	userStore store.UserStore
	log       slf.StructuredLogger
}

func SetRestHandlers(
	userDB store.UserStore,
	r *gin.Engine,
) (*RestClient, error) {
	restClient := &RestClient{
		userStore: userDB,
		log:       slf.WithContext("rest-client"),
	}
	r.POST("/auth", restClient.auth())
	r.GET("/user/graffity", restClient.getUserGraffity())
	r.POST("/graffity", restClient.postGraffity())
	r.POST("/near/graffity/", restClient.nearGraffitys())
	r.GET("/map/zones/", restClient.getMapZones())

	return restClient, nil
}

func (restClient *RestClient) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds := store.Login{}

		if err := decodeBody(c, &creds); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrRequestBodyError,
			})
			return
		}

		user := store.User{
			UserID:     creds.UserID,
			Gang:       creds.Gang,
			Graffities: []string{},
		}

		err := restClient.userStore.NewUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "created",
		})
	}
}

func (restClient *RestClient) getUserGraffity() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := getIdenity(c)
		if err != nil {
			restClient.log.Errorf("addWallet: getIdenity: %s\t[addr=%s]", err.Error(), c.Request.RemoteAddr)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrHeaderError,
			})
			return
		}
		g, err := restClient.userStore.GetAllUserGraffitys(userid)
		if err != nil {
			restClient.log.Errorf("addWallet: getIdenity: %s\t[addr=%s]", err.Error(), c.Request.RemoteAddr)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": g,
		})
	}
}

func (restClient *RestClient) postGraffity() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := getIdenity(c)
		if err != nil {
			restClient.log.Errorf("addWallet: getIdenity: %s\t[addr=%s]", err.Error(), c.Request.RemoteAddr)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrHeaderError,
			})
			return
		}

		if restClient.userStore.CheckUser(userid) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrNoSuchUser,
			})
			return
		}

		gr := store.Graffity{}
		if err := decodeBody(c, &gr); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": msgErrRequestBodyError,
			})
			return
		}
		err = restClient.userStore.PostGraffity(gr)
		if err := decodeBody(c, &gr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "created",
		})

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
