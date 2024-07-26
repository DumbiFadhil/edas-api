package services

import (
	"DumbiFadhil/edas-api/models"
	"context"
	"log"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllHistory() ([]models.HistoryListItem, error) {
	collection := db.Collection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Project only the necessary fields (UUID and rankings)
	projection := bson.M{"uuid": 1, "rankings": 1}

	// Find options to control batch size
	findOptions := options.Find().SetProjection(projection).SetBatchSize(100)

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Printf("GetAllHistory: Failed to find history: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var historyItems []models.HistoryListItem

	// Iterate and decode into HistoryListItem
	for cursor.Next(ctx) {
		var result bson.M // Raw MongoDB document
		if err := cursor.Decode(&result); err != nil {
			log.Printf("GetAllHistory: Error decoding history item: %v", err)
			continue
		}

		var item models.HistoryListItem

		// Extract UUID if it exists
		if uuidValue, ok := result["uuid"].(uuid.UUID); ok {
			item.UUID = uuidValue
		}

		// Extract rankings
		if rankingsValue, ok := result["rankings"].(bson.A); ok { // bson.A for array
			for _, rankRaw := range rankingsValue {
				if rankMap, ok := rankRaw.(bson.M); ok {
					var rankedAlternative models.RankedAlternative
					bsonBytes, _ := bson.Marshal(rankMap)
					bson.Unmarshal(bsonBytes, &rankedAlternative)
					item.Rankings = append(item.Rankings, rankedAlternative)
				}
			}
		}

		historyItems = append(historyItems, item)

		filteredRankings := make([]models.RankedAlternative, 0, len(item.Rankings))
		for _, rank := range item.Rankings {
			if !math.IsNaN(rank.Score) {
				filteredRankings = append(filteredRankings, rank)
			} else {
				log.Printf("GetAllHistory: Filtering out RankedAlternative with NaN score")
			}
		}
		item.Rankings = filteredRankings // Update item with filtered rankings

		historyItems = append(historyItems, item) // Append the filtered item
	}

	// Check for errors after iteration is complete
	if err := cursor.Err(); err != nil {
		log.Printf("GetAllHistory: Cursor error: %v", err)
		return nil, err
	}

	log.Printf("GetAllHistory: Successfully retrieved %d history items", len(historyItems))
	return historyItems, nil

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

// DeleteHistoryByUUID deletes a history record by its UUID
func DeleteHistoryByUUID(uuidStr string) error {
	collection := db.Collection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert the UUID string to a uuid.UUID type
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		return err
	}

	log.Println("Deleting history with UUID:", uuid)

	// Delete the document with the specified UUID
	result, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		log.Println("Failed to delete history by UUID:", err)
		return err
	}

	if result.DeletedCount == 0 {
		log.Println("No history found with UUID:", uuid)
		return mongo.ErrNoDocuments // No document found
	}

	log.Println("Deleted history with UUID:", uuid)
	return nil
}
