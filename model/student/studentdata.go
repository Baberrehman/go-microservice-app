package student

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mvcapp/model/mongo"
)

type Student struct {
	ID       string `bson:"id"      json:"id,omitempty"`
	Name     string `bson:"name"     json:"name,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
}

const (
	ErrDupRows  = 500
	ErrDatabase = 501
)

func (u *Student) Insert() (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("students")
	err = c.Insert(u)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}
	return
}
func (u *Student) Delete() (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("students")
	err = c.Remove(bson.M{"id": u.ID})

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}
	return
}
func (u *Student) Update(id string) (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("students")

	data := Student{}
	query := c.Find(bson.M{
		"id": id,
	})

	err1 := query.One(&data)

	if err1 != nil {
		fmt.Println(err1)
		return 500, err1
	}
	/*err1 := query.One(&doc)
	if err1 != nil {
		utils.Error.Println(err1)
		return err1
	}*/
	data.Name = u.Name
	data.Password = u.Password
	changeDoc := mgo.Change{
		Update:    data,
		ReturnNew: true,
	}

	_, err2 := query.Apply(changeDoc, &data)
	if err2 != nil {
		fmt.Println(err2)
		return 500, err2
	}

	return 200, nil
}
func (u *Student) Get(id string) (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("students")

	err = c.Find(bson.M{"id": id}).One(&u)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}
	return
}
