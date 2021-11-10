package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        string
	Uuid      string
	Name      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) CreateUser() (err error) {

	_, err = UserCollection.InsertOne(context.Background(), bson.D{
		{Key: "Uuid", Value: createUUID()},
		{Key: "Name", Value: user.Name},
		{Key: "Email", Value: user.Email},
		{Key: "Password", Value: user.Password},
	})
	if err != nil {
		return
	}

	return
}

func UserByEmail(email string) (user User, err error) { /*really messy, but FindOne doesn't seem to work*/

	var foundUuid string
	var foundUser string
	var foundEmail string
	var foundPassword string

	cursor, err := UserCollection.Find(context.Background(), bson.M{"Email": email})
	if err != nil {
		return
	}

	var filteredCursor []bson.M

	if err = cursor.All(context.Background(), &filteredCursor); err != nil {
		return
	}

	for _, singleUser := range filteredCursor {

		foundUuid = fmt.Sprint(singleUser["Uuid"])
		foundUser = fmt.Sprint(singleUser["Name"])
		foundEmail = fmt.Sprint(singleUser["Email"])
		foundPassword = fmt.Sprint(singleUser["Password"])
	}

	user = User{
		Uuid:     foundUuid,
		Name:     foundUser,
		Email:    foundEmail,
		Password: foundPassword,
	}

	return user, nil

}

func UserByUuid(Uuid string) (user User, err error) { /*really messy, but FindOne doesn't seem to work*/

	var foundUuid string
	var foundUser string
	var foundEmail string
	var foundPassword string

	cursor, err := UserCollection.Find(context.Background(), bson.M{"Uuid": Uuid})
	if err != nil {
		return
	}

	var filteredCursor []bson.M

	if err = cursor.All(context.Background(), &filteredCursor); err != nil {
		return
	}

	for _, singleUser := range filteredCursor {

		foundUuid = fmt.Sprint(singleUser["Uuid"])
		foundUser = fmt.Sprint(singleUser["Name"])
		foundEmail = fmt.Sprint(singleUser["Email"])
		foundPassword = fmt.Sprint(singleUser["Password"])
	}

	user = User{
		Uuid:     foundUuid,
		Name:     foundUser,
		Email:    foundEmail,
		Password: foundPassword,
	}

	return user, err

}

func (user *User) CreateSession() (session Session, err error) {

	session = Session{
		Name:  user.Name,
		Uuid:  user.Uuid,
		Email: user.Email,
	}

	if ok, _ := session.CheckSession(); !ok {

		_, err = SessionCollection.InsertOne(context.Background(), bson.D{
			{Key: "Name", Value: session.Name},
			{Key: "Uuid", Value: session.Uuid},
			{Key: "Email", Value: session.Email},
		})
		if err != nil {
			return
		}
	}

	return session, nil

}

func (session *Session) CheckSession() (valid bool, err error) {

	_, err = SessionCollection.Find(context.Background(), bson.M{"Uuid": session.Uuid})
	if err != nil {
		return false, err
	}
	// fmt.Println("Session already exists! Handing back the same")
	return true, err
}

func (user *User) CheckUserExistance() (valid bool, err error) {

	var checker bson.M

	if err = UserCollection.FindOne(context.Background(), bson.M{"Email": user.Email}).Decode(&checker); err != nil {
		return false, err /*false means not ok, so user wasnt found, so no such user exists*/
	}

	return true, err
}
