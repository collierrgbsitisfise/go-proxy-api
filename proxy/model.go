package proxy

import "gopkg.in/mgo.v2/bson"

type Proxy struct {
	ID      bson.ObjectId `bson:"_id"`
	Time    int           `json:"time"`
	Ip      string        `json:"ip"`
	Port    int           `json:"port"`
	Country string        `json: "country"`
}

type Proxys []Proxy
