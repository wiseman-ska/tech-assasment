package controllers

import (
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"gopkg.in/mgo.v2"
)

type Context struct {
	MongoDBSession *mgo.Session
}

func NewContext() *Context {
	session := commons.GetSession().Copy()
	context := &Context{
		MongoDBSession: session,
	}
	return context
}

func (cont *Context) Collection(name string) *mgo.Collection {
	return cont.MongoDBSession.DB(commons.AppConf.Database).C(name)
}

func (cont *Context) Close() {
	cont.MongoDBSession.Close()
}
