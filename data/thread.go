package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	Body      string
	CreatedBy string
	CreatedAt string
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func CreateThread(topic, body, user string) (err error) {

	_, err = ThreadCollection.InsertOne(context.Background(), bson.D{
		{Key: "Uuid", Value: createUUID()},
		{Key: "Topic", Value: topic},
		{Key: "Body", Value: body},
		{Key: "CreatedBy", Value: user},
		{Key: "CreatedAt", Value: time.Now().Format("02-01-2006 15:04:05")},
	})
	if err != nil {
		return
	}

	return
}

func GetThreads() (th []Thread) {

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := ThreadCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return
	}

	var thread Thread

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var singleThread bson.M
		if err = cursor.Decode(&singleThread); err != nil {
			log.Fatal(err)
		}
		thread = Thread{
			Uuid:      fmt.Sprint(singleThread["Uuid"]),
			Topic:     fmt.Sprint(singleThread["Topic"]),
			Body:      fmt.Sprint(singleThread["Body"]),
			CreatedBy: fmt.Sprint(singleThread["CreatedBy"]),
			CreatedAt: fmt.Sprint(singleThread["CreatedAt"]),
		}

		th = append(th, thread)

	}

	return th
}

func ThreadByUuid(Uuid string) (thread Thread, err error) { /*really messy, but FindOne doesn't seem to work*/

	var foundUuid string
	var foundTopic string
	var foundBody string
	var foundCreatedBy string
	var foundCreatedAt string

	cursor, err := ThreadCollection.Find(context.Background(), bson.M{"Uuid": Uuid})
	if err != nil {
		return
	}

	var filteredCursor []bson.M

	if err = cursor.All(context.Background(), &filteredCursor); err != nil {
		return
	}

	for _, singleUser := range filteredCursor {

		foundUuid = fmt.Sprint(singleUser["Uuid"])
		foundTopic = fmt.Sprint(singleUser["Topic"])
		foundBody = fmt.Sprint(singleUser["Body"])
		foundCreatedAt = fmt.Sprint(singleUser["CreatedAt"])
		foundCreatedBy = fmt.Sprint(singleUser["CreatedBy"])
	}

	thread = Thread{
		Uuid:      foundUuid,
		Topic:     foundTopic,
		Body:      foundBody,
		CreatedAt: foundCreatedAt,
		CreatedBy: foundCreatedBy,
	}

	return thread, err

}

func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {

	_, err = PostCollection.InsertOne(context.Background(), bson.D{
		{Key: "Uuid", Value: createUUID()},
		{Key: "Body", Value: body},
		{Key: "CreatedBy", Value: user.Name},
		{Key: "CreatedAt", Value: time.Now().Format("02-01-2006 15:04:05")},
	})
	if err != nil {
		return
	}

	return
}
