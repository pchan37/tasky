package dbManager

import (
	"github.com/globalsign/mgo"
)

type DBManager struct {
	Name     string
	session  *mgo.Session
	Database *mgo.Database
}

func New(name string, addresses ...string) (d *DBManager) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: addresses,
	})
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Strong, true)
	session.SetSafe(&mgo.Safe{})
	d = &DBManager{Name: name, session: session, Database: session.DB(name)}
	return
}

func Clear(d *DBManager) (success bool) {
	success = true
	if err := d.Database.DropDatabase(); err != nil {
		success = false
		return
	}
	if err := d.session.DB(d.Name); err != nil {
		success = false
		return
	}
	return
}

func Close(d *DBManager) {
	d.session.Close()
}
