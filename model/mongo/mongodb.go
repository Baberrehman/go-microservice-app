package mongo

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/mgo.v2"
	"net"
)

var session *mgo.Session

// Conns return mongodb session.
func Conn() *mgo.Session {
	return session.Copy()
}

/*
func Close() {
	session.Close()
}
*/

func init() {
	url := "**.**.**.**:port"
	auth := true
	username := "***"
	password := "***"
	fmt.Println(url)
	if !auth {
		sess, err := mgo.Dial(url)
		if err != nil {
			return
		}
		session = sess
		session.SetMode(mgo.Monotonic, true)
	} else {

		sess, err := mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{url},
			Username: username,
			Password: password,
			DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {

				conf := &tls.Config{
					InsecureSkipVerify: true,
				}

				return tls.Dial("tcp", addr.String(), conf)
			},
		})
		if err != nil {
			return
		}
		session = sess
		session.SetMode(mgo.Monotonic, true)
	}

}
