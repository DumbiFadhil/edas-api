package services

import (
	"DumbiFadhil/edas-api/models"
	"context"
	"log"
	"sort"
	"time"

	"github.com/google/uuid"
)

// Helper function for max calculation
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func SaveHistory(history models.History) error {
	// Generate a new UUID
	history.UUID = uuid.New()

	collection := db.Collection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, history)
	if err != nil {
		log.Println("Failed to save history:", err)
		return err
	}
	return nil
}

func CalculateEDAS(request models.EDASRequest) models.EDASResponse {
	alternatives := request.Alternatives
	criteria := request.Criteria

	// Step 1: Calculate average score per criterion
	criterionAverages := make(map[string]float64)
	for _, criterion := range criteria {
		totalScore := float64(0)
		for _, alt := range alternatives {
			totalScore += alt.Scores[criterion.Name]
		}
		criterionAverages[criterion.Name] = totalScore / float64(len(alternatives))
	}

	// Step 2: Calculate positive and negative distances
	for _, alt := range alternatives {
		var positiveDistance, negativeDistance float64
		for _, criterion := range criteria {
			avg := criterionAverages[criterion.Name]
			score := alt.Scores[criterion.Name]
			if criterion.Type == "benefit" {
				positiveDistance += (max(0, score-avg) / avg) * criterion.Weight
				negativeDistance += (max(0, avg-score) / avg) * criterion.Weight
			} else { // "cost"
				positiveDistance += (max(0, avg-score) / avg) * criterion.Weight
				negativeDistance += (max(0, score-avg) / avg) * criterion.Weight
			}
		}
		alt.Scores["PositiveDistance"] = positiveDistance
		alt.Scores["NegativeDistance"] = negativeDistance
	}

	// Step 3: Normalize positive and negative distances
	var maxPositiveDistance, maxNegativeDistance float64
	for _, alt := range alternatives {
		if alt.Scores["PositiveDistance"] > maxPositiveDistance {
			maxPositiveDistance = alt.Scores["PositiveDistance"]
		}
		if alt.Scores["NegativeDistance"] > maxNegativeDistance {
			maxNegativeDistance = alt.Scores["NegativeDistance"]
		}
	}

	for _, alt := range alternatives {
		alt.Scores["NormalizedPositiveDistance"] = alt.Scores["PositiveDistance"] / maxPositiveDistance
		alt.Scores["NormalizedNegativeDistance"] = alt.Scores["NegativeDistance"] / maxNegativeDistance
	}

	// Step 4: Calculate final score
	rankedAlternatives := make([]models.RankedAlternative, len(alternatives))
	for i, alt := range alternatives {
		finalScore := (alt.Scores["NormalizedPositiveDistance"] + (1 - alt.Scores["NormalizedNegativeDistance"])) / 2
		rankedAlternatives[i] = models.RankedAlternative{Name: alt.Name, Score: finalScore}
	}

	// Step 5: Rank alternatives
	sort.SliceStable(rankedAlternatives, func(i, j int) bool {
		return rankedAlternatives[i].Score > rankedAlternatives[j].Score
	})

	// Assign ranks
	for i := range rankedAlternatives {
		rankedAlternatives[i].Rank = i + 1
	}

	// Create EDASResponse
	edasResponse := models.EDASResponse{Ranking: rankedAlternatives}

	// Save to history
	history := models.History{
		EDASRequests:  []models.EDASRequest{request},
		EDASResponses: []models.EDASResponse{edasResponse},
		Rankings:      rankedAlternatives,
	}
	err := SaveHistory(history)
	if err != nil {
		log.Println("Failed to save history:", err)
	}

	return edasResponse
}
