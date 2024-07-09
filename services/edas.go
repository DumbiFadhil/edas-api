package services

import (
	"DumbiFadhil/edas-api/models"
	"sort"
)

func CalculateEDAS(request models.EDASRequest) models.EDASResponse {
	alternatives := request.Alternatives
	criteria := request.Criteria

	// Step 1: Calculate average score per criterion
	criterionAverages := make(map[string]float64)
	for _, criterion := range criteria {
		totalScore := float64(0) // Use 0 for initialization
		for _, alt := range alternatives {
			totalScore += alt.Scores[criterion.Name]
		}
		criterionAverages[criterion.Name] = totalScore / float64(len(alternatives))
	}

	// Step 2: Calculate positive and negative distances
	for _, alt := range alternatives { // No need for unused index 'i'
		var positiveDistance, negativeDistance float64
		for _, criterion := range criteria {
			avg := criterionAverages[criterion.Name]
			score := alt.Scores[criterion.Name]
			if score >= avg {
				positiveDistance += (score - avg) * criterion.Weight
			} else {
				negativeDistance += (avg - score) * criterion.Weight
			}
		}
		alt.Scores["PositiveDistance"] = positiveDistance
		alt.Scores["NegativeDistance"] = negativeDistance
	}

	// Step 3: Calculate overall appraisal score
	rankedAlternatives := make([]models.RankedAlternative, len(alternatives))
	for i, alt := range alternatives { // 'i' used here for indexing
		overallScore := (alt.Scores["PositiveDistance"] - alt.Scores["NegativeDistance"]) / 2
		rankedAlternatives[i] = models.RankedAlternative{Name: alt.Name, Score: overallScore}
	}

	// Step 4: Sort alternatives by overall score
	sort.SliceStable(rankedAlternatives, func(i, j int) bool {
		return rankedAlternatives[i].Score > rankedAlternatives[j].Score
	})

	return models.EDASResponse{Ranking: rankedAlternatives}
}
