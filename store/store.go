/*
Copyright 2018 Idealnaya rabota LLC
Licensed under Multy.io license.
See LICENSE for details
*/
package store

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	errType        = errors.New("wrong database type")
	errEmplyConfig = errors.New("empty configuration for datastore")
)

// Default table names
const (
	TableUsers     = "UserCollection"
	TableGraffitys = "GraffitysCollection"
)

// Conf is a struct for database configuration
type Conf struct {
	Address     string
	DBUsers     string
	DBgraffitys string
}

type UserStore interface {
	NewUser(user User) error
	GetAllUserGraffitys(userid string) ([]Graffity, error)
	GetNearbyGraffitys(longitude, latitude string) ([]Graffity, error)
	PostGraffity(Graffity) error
	CheckUser(userid string) bool
}

type MongoUserStore struct {
	config  *Conf
	session *mgo.Session

	usersData *mgo.Collection
	graffitys *mgo.Collection
}

func InitUserStore(conf Conf) (UserStore, error) {
	uStore := &MongoUserStore{
		config: &conf,
	}

	session, err := mgo.Dial(conf.Address)
	if err != nil {
		return nil, err
	}
	uStore.session = session
	uStore.usersData = uStore.session.DB(conf.DBUsers).C(TableUsers)
	uStore.graffitys = uStore.session.DB(conf.DBgraffitys).C(TableGraffitys)

	return uStore, nil
}

func (mStore *MongoUserStore) NewUser(user User) error {
	users := []User{}
	err := mStore.usersData.Find(bson.M{
		"userid": user.UserID,
	}).All(&users)

	if err == mgo.ErrNotFound {
		return mStore.usersData.Insert(user)
	} else {
		return errors.New("User already exists")
	}
	// return mStore.usersData.Insert(user)
}

func (mStore *MongoUserStore) GetAllUserGraffitys(userid string) ([]Graffity, error) {
	allGraffitys := []Graffity{}
	err := mStore.graffitys.Find(bson.M{
		"userid": userid,
	}).All(&allGraffitys)

	return allGraffitys, err
}

func (mStore *MongoUserStore) GetNearbyGraffitys(longitude, latitude string) ([]Graffity, error) {
	allGraffitys := []Graffity{}
	err := mStore.graffitys.Find(nil).All(&allGraffitys)
	return allGraffitys, err
}

func (mStore *MongoUserStore) PostGraffity(graffity Graffity) error {
	return mStore.graffitys.Insert(graffity)
}

func (mStore *MongoUserStore) CheckUser(userid string) bool {
	users := []User{}
	if err := mStore.graffitys.Find(bson.M{"userid": userid}).All(&users); err == mgo.ErrNotFound {
		return true
	}
	return false
}
func (mStore *MongoUserStore) Close() error {
	mStore.session.Close()
	return nil
}
