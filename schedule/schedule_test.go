package database

import (
	"log"
	"testing"

	"egbitbucket.dtvops.net/deadline/common"

	"github.com/stretchr/testify/assert"
)

//"egbitbucket.dtvops.net/deadline/common"
//create a new database
//assert

/*
var dbdriver string

func TestGoodDB (test *testing.T) {

sql.Register(dbdriver, &sqlite3.SQLiteDriver{})
db, err := sql.Open(dbdriver,"mysqlite_3")
assert.Nil(test, err, "SQL handle could not be created")
assert.NotNil(test, db, "No database was found")

}

func TestBadDB (test *testing.T) {

sql.Register(dbdriver, &sqlite3.SQLiteDriver{})
db, err := sql.Open(dbdriver,"mysqlite_3")
assert.NotNil(test, err, "SQL could not be created when it wasn't supposed to be")
//assert.NotNil(test, db, "No database was found")



}
*/
//an open and close? main function?

//show that it can get information
var fd = NewScheduleDAO()
var s = common.Schedule{
	Name:   "sample_schedule",
	Timing: "daily",
	Handler: common.Handler{
		Name:    "email handler",
		Address: "kp755d@att.com",
	},
}

func TestSendFile(test *testing.T) {
	assert.Nil(test, fd.save(s), "Could not save the file.")

	//will put file in directory

}

func TestGetFile(test *testing.T) {
	f, err := fd.getByName("sample_schedule")
	assert.Nil(test, err, "Could not find the file.")
	assert.NotNil(test, f, "Could not find the file.")
	log.Printf("Received the following information: %#v\n", f)
	//will get sample schedule from directory

}
