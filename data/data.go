package data

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientGorum *mongo.Client
var GorumDb *mongo.Database
var UserCollection *mongo.Collection
var SessionCollection *mongo.Collection
var ThreadCollection *mongo.Collection
var PostCollection *mongo.Collection

func init() {

	var err error
	ClientGorum, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://srijan:srijandb@cluster0.xdk3r.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = ClientGorum.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	GorumDb = ClientGorum.Database("gorum")
	UserCollection = GorumDb.Collection("user")
	SessionCollection = GorumDb.Collection("session")
	ThreadCollection = GorumDb.Collection("thread")
	PostCollection = GorumDb.Collection("post")
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
