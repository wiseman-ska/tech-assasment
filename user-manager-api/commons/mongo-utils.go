package commons


import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session

func addIndexesToDataBase()  {
	var err error
	session := GetSession().Copy()
	defer session.Close()
	userIndex := mgo.Index{
		Key: []string{"email"},
		Unique: true,
		Background: true,
		Sparse: true,
	}
	userCol := session.DB(AppConf.Database).C("users")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addUserIndexes]: %s\n", err)
	}
}

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConf.MongoDBHost},
			Username: AppConf.DBUser,
			Password: AppConf.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConf.MongoDBHost},
		Username: AppConf.DBUser,
		Password: AppConf.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

