package database

import (
	"context"
	"errors"
	"log"

	utils "github.com/ItsMeSamey/go_utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password []byte `json:"password" bson:"password"`
}

type Email struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Sender     string        `json:"sender" bson:"sender"`
	Recipient  string        `json:"recipient" bson:"recipient"`
	Subject    string        `json:"subject" bson:"subject"`
	TextBody   string        `json:"text-body" bson:"text-body"`
	HTMLBody   string        `json:"html-body" bson:"html-body"`
	Recipients []string      `json:"recipients" bson:"recipients"`
}

type Client struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Email        string        `json:"email" bson:"email"`
	SerialNumber string        `json:"serialNumber" bson:"serialNumber"`
	ClientSecret string        `json:"clientSecret" bson:"clientSecret"`
	IsRevoked    bool          `json:"isRevoked" bson:"isRevoked"`
}


type Collection[T any] struct {
	*mongo.Collection
}

func (c *Collection[T]) GetExists(filter any) (out T, exists bool, err error) {
	result := c.FindOne(context.Background(), filter)
	err = result.Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return out, false, nil
		}
		log.Printf("Error finding document: %v\n", err)
		return out, false, utils.WithStack(err)
	}

	if err := result.Decode(&out); err != nil {
		log.Printf("Error decoding document: %v\n", err)
		return out, false, utils.WithStack(err)
	}

	return out, true, nil
}
