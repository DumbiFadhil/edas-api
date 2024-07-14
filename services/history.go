package services

import (
	"DumbiFadhil/edas-api/models"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllHistory() ([]models.History, error) {
	collection := db.Collection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Failed to find history:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var history []models.History
	if err := cursor.All(ctx, &history); err != nil {
		log.Println("Failed to decode history:", err)
		return nil, err
	}
	return history, nil
}

func GetHistoryByUUID(uuidStr string) (*models.History, error) {
	collection := db.Collection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert the UUID string to a uuid.UUID type
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		return nil, err
	}

	log.Println("Searching for history with UUID:", uuid)

	var history models.History
	err = collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&history)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No history found with UUID:", uuid)
			return nil, nil // No document found
		}
		log.Println("Failed to find history by UUID:", err)
		return nil, err
	}

	log.Println("Found history with UUID:", history.UUID)
	return &history, nil
}
