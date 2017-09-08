package teacher

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mvcapp/model/mongo"
)

type Teacher struct {
	ID       string `bson:"id"      json:"id"`
	Name     string `bson:"name"     json:"name"`
	Password string `bson:"password" json:"password"`
}

func (u *Teacher) Insert() (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("teachers")
	err = c.Insert(u)

	if err != nil {
		if mgo.IsDup(err) {
			code = 500
		} else {
			code = 501
		}
	} else {
		code = 0
	}
	return
}
func (u *Teacher) Delete() (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("teachers")
	err = c.Remove(bson.M{"id": u.ID})

	if err != nil {
		code = 500
	} else {
		code = 0
	}
	return
}
func (u *Teacher) Update(id string) (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("teachers")

	data := Teacher{}
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
	fmt.Println(u)
	data.Name = u.Name
	data.Password = u.Password
	fmt.Println(data)
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
func (u *Teacher) Get(id string) (code int, err error) {
	mConn := mongo.Conn()
	defer mConn.Close()

	c := mConn.DB("deployment").C("teachers")

	err = c.Find(bson.M{"id": id}).One(&u)

	if err != nil {
		code = 501
	} else {
		code = 0
	}
	return
}
