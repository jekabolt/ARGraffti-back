package graffity

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jekabolt/ARGraffti-back/client"
	"github.com/jekabolt/ARGraffti-back/store"
	"github.com/jekabolt/slf"
)

var (
	log = slf.WithContext("graffity")
)

// Multy is a main struct of service
type Graffity struct {
	config     *Configuration
	userStore  store.UserStore
	restClient *client.RestClient
	route      *gin.Engine
}

// Init initializes Multy instance
func Init(conf *Configuration) (*Graffity, error) {
	graffity := &Graffity{
		config: conf,
	}
	userStore, err := store.InitUserStore(conf.Database)
	if err != nil {
		return nil, fmt.Errorf("DB initialization: %s on port %s", err.Error(), conf.Database.Address)
	}
	graffity.userStore = userStore
	log.Infof("UserStore initialization done on %s √", conf.Database)

	// REST handlers
	if err := graffity.initHttpRoutes(conf); err != nil {
		return nil, fmt.Errorf("Router initialization: %s", err.Error())
	}
	return graffity, nil
}

// initRoutes initialize client communication services
func (graffity *Graffity) initHttpRoutes(conf *Configuration) error {
	router := gin.Default()
	graffity.route = router
	//
	gin.SetMode(gin.DebugMode)

	restClient, err := client.SetRestHandlers(
		graffity.userStore,
		router,
	)
	if err != nil {
		return err
	}
	graffity.restClient = restClient

	return nil
}

// Run runs service
func (graffity *Graffity) Run() error {
	log.Info("Running server")
	graffity.route.Run(graffity.config.RestAddress)
	return nil
}
