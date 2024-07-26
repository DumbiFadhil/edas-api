# EDAS API

[![EDAS API](https://img.shields.io/badge/EDAS-API-blue.svg)](https://github.com/DumbiFadhil/edas-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/DumbiFadhil/edas-api)](https://goreportcard.com/report/github.com/DumbiFadhil/edas-api) 
[![GoDoc](https://godoc.org/github.com/DumbiFadhil/edas-api?status.svg)](https://godoc.org/github.com/DumbiFadhil/edas-api)

## Overview

EDAS (Evaluation based on Distance from Average Solution) API is a decision support system implemented in Go (Golang). This API allows users to evaluate multiple alternatives based on various criteria and rank them accordingly. It provides a robust and efficient solution for multi-criteria decision making.

This algorithm provides a structured way to evaluate and rank multiple alternatives based on their performance across various criteria. It's particularly useful when criteria have varying importance (weights) and can be either beneficial (higher scores are better) or cost-oriented (lower scores are better).

**How EDAS Works**

1. **Average Score Calculation:**
   - For each criterion, calculate the average score across all alternatives.

2. **Distance Calculation:**
   - For each alternative and criterion:
      - Determine the positive distance (how much better the alternative is than average).
      - Determine the negative distance (how much worse the alternative is than average).
      - Adjust distances based on the criterion's type (benefit or cost) and weight.

3. **Normalization:**
   - Normalize positive and negative distances to a 0-1 scale for easier comparison.

4. **Final Score:**
   - Calculate a final score for each alternative by combining normalized distances. This score considers both how much better the alternative is on some criteria and how much worse it might be on others.

5. **Ranking:**
   - Sort alternatives based on their final scores, from highest to lowest.

## Key Features

- **Multi-Criteria Evaluation:**  Handles complex decisions with multiple factors.
- **Weighted Criteria:**  Prioritize criteria based on their importance.
- **Benefit/Cost Flexibility:**  Works with criteria where higher or lower scores are preferred.
- **Accurate Ranking:**  Provides a clear, data-driven ranking of alternatives.
- **API-Driven:** Easy to integrate into your applications.
- **History Tracking (MongoDB):** (Optional) Logs evaluation requests and results.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Example Request/Response](#example-requestresponse) 
- [Contributing](#contributing)
- [License](#license)

## Code Breakdown

1. **Initialization:**
    ```Go
    criterionAverages := make(map[string]float64)  // Map to store average scores
    rankedAlternatives := make([]models.RankedAlternative, len(alternatives)) // Slice for ranked results
    ```
Two data structures are created:

criterionAverages stores the average score for each criterion.
rankedAlternatives will hold the final results (alternatives with their calculated scores and ranks).

2. **Calculate Average Score per Criterion:**
    ```Go
    for _, criterion := range criteria {
        totalScore := float64(0)
        for _, alt := range alternatives {
            totalScore += alt.Scores[criterion.Name]
        }
        criterionAverages[criterion.Name] = totalScore / float64(len(alternatives))
    }
    ```
- We iterate through each criterion.
- For each criterion, we sum the scores across all alternatives.
- The total score is divided by the number of alternatives to get the average.
- The average score is stored in the criterionAverages map.

3. **Calculate Average Score per Criterion:**
    ```Go
    for _, alt := range alternatives {
        var positiveDistance, negativeDistance float64
        for _, criterion := range criteria {
            avg := criterionAverages[criterion.Name]
            score := alt.Scores[criterion.Name]
            // ... (calculations for positive and negative distances based on criterion type)
        }
        alt.Scores["PositiveDistance"] = positiveDistance
        alt.Scores["NegativeDistance"] = negativeDistance
    }
    ```
- We iterate through each alternative.
- For each criterion, we calculate:
  - The positiveDistance (how much better the alternative is than average).
  - The negativeDistance (how much worse the alternative is than average).
- These distances are adjusted based on whether the criterion is a "benefit" (higher score is better) or "cost" (lower score is better).
- The calculated distances are stored in the alt.Scores map.

4. **Normalize Distances:**
    ```Go
    // ... (find maxPositiveDistance and maxNegativeDistance)

    for _, alt := range alternatives {
        alt.Scores["NormalizedPositiveDistance"] = alt.Scores["PositiveDistance"] / maxPositiveDistance
        alt.Scores["NormalizedNegativeDistance"] = alt.Scores["NegativeDistance"] / maxNegativeDistance
    }
    ```
- After finding the maximum positive and negative distances, we normalize them to the range of 0 to 1.
- Each alternative's positive and negative distances are divided by the respective maximum values.

5. **Calculate Final Score:**
    ```Go
    for i, alt := range alternatives {
        finalScore := (alt.Scores["NormalizedPositiveDistance"] + (1 - alt.Scores["NormalizedNegativeDistance"])) / 2
        rankedAlternatives[i] = models.RankedAlternative{Name: alt.Name, Score: finalScore}
    }
    ```
- For each alternative, the final score is calculated as the average of the normalized positive distance and the complement of the normalized negative distance (1 minus the normalized negative distance).
- RankedAlternative structs are created to store the Name and Score of each alternative.

6. **Rank Alternatives:**
    ```Go
    sort.SliceStable(rankedAlternatives, func(i, j int) bool {
        return rankedAlternatives[i].Score > rankedAlternatives[j].Score
    })

    // Assign ranks
    for i := range rankedAlternatives {
      rankedAlternatives[i].Rank = i + 1
    }
    ```
- The rankedAlternatives slice is sorted in descending order based on the Score of each RankedAlternative.

7. **Create Response and Save History (Optional & Require MongoDB Connection)**:
    ```Go
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
    ```
- An EDASResponse is created, containing the ranked alternatives.
- Optionally, the request, response, and rankings can be saved to a history.

## Features

- **Multi-criteria Decision Analysis:** Evaluate alternatives based on multiple criteria.
- **Flexible Weights:** Assign weights to each criterion to reflect their importance.
- **Accurate Ranking:** Get accurate rankings of alternatives using the EDAS method.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- Go 1.16 or higher
- Git
- MongoDB (if using history tracking)

### Steps

1. Clone the repository:

    ```sh
    git clone https://github.com/DumbiFadhil/edas-api.git
    cd edas-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Build the project:

    ```sh
    go build
    ```

4. Run the server:

    ```sh
    ./edas-api
    ```

**OR**

3. Immediately start the server without building executable file:

    ```sh
    go run main.go
    ```

## Usage

### Running the Server

- Start the server with:

    ```sh
    ./edas-api
    ```

The server will start on `http://localhost:8080`.

## API Endpoints

### Calculate EDAS

- **URL:** `/api/v1/edas`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Request Body:**

    ```json
    {
        "alternatives": [
            {"name": "Alt1", "scores": {"Criteria1": 20, "Criteria2": 10, "Criteria3": 30, "Criteria4": 25, "Criteria5": 5}},
            {"name": "Alt2", "scores": {"Criteria1": 15, "Criteria2": 7, "Criteria3": 21, "Criteria4": 60, "Criteria5": 3}},
            {"name": "Alt3", "scores": {"Criteria1": 18, "Criteria2": 8, "Criteria3": 22, "Criteria4": 40, "Criteria5": 4}},
            {"name": "Alt4", "scores": {"Criteria1": 23, "Criteria2": 6, "Criteria3": 31, "Criteria4": 40, "Criteria5": 8}},
            {"name": "Alt5", "scores": {"Criteria1": 17, "Criteria2": 9, "Criteria3": 27, "Criteria4": 20, "Criteria5": 4}}

        ],
        "criteria": [
            {"name": "Criteria1", "weight": 0.3, "type": "benefit"},
            {"name": "Criteria2", "weight": 0.1, "type": "benefit"},
            {"name": "Criteria3", "weight": 0.2, "type": "cost"},
            {"name": "Criteria4", "weight": 0.2, "type": "cost"},
            {"name": "Criteria5", "weight": 0.2, "type": "cost"}
        ]
    }
    ```

- **Response:**

    ```json
    {
        "ranking": [
            {
                "name": "Alt5",
                "score": 0.9244443265623181
            },
            {
                "name": "Alt1",
                "score": 0.8198183184953796
            },
            {
                "name": "Alt3",
                "score": 0.6761050498859149
            },
            {
                "name": "Alt2",
                "score": 0.45498470350972064
            },
            {
                "name": "Alt4",
                "score": 0.25764249729644173
            }
        ]
    }
    ```

## Things to note
- Every variable is Case-Sensitive, variables such as name, scores, weight, and type should be fully lowercased
- Every variable but name and type should be float64 OR integer
- Variable type is Case-Sensitive, and any values other than "benefit" will be treated as cost

**VERY IMPORTANT**
- The API will **NOT** run unless the .env file contains the correct MongoDB configuration, such as URI and database name.

## Example

Here's a quick example of how to use the API with `curl`:

```sh
curl -X POST http://localhost:8080/api/v1/edas \
    -H "Content-Type: application/json" \
    -d '{
        "alternatives": [
            {"name": "Alt1", "scores": {"Criteria1": 20, "Criteria2": 10, "Criteria3": 30, "Criteria4": 25, "Criteria5": 5}},
            {"name": "Alt2", "scores": {"Criteria1": 15, "Criteria2": 7, "Criteria3": 21, "Criteria4": 60, "Criteria5": 3}},
            {"name": "Alt3", "scores": {"Criteria1": 18, "Criteria2": 8, "Criteria3": 22, "Criteria4": 40, "Criteria5": 4}},
            {"name": "Alt4", "scores": {"Criteria1": 23, "Criteria2": 6, "Criteria3": 31, "Criteria4": 40, "Criteria5": 8}},
            {"name": "Alt5", "scores": {"Criteria1": 17, "Criteria2": 9, "Criteria3": 27, "Criteria4": 20, "Criteria5": 4}}

        ],
        "criteria": [
            {"name": "Criteria1", "weight": 0.3, "type": "benefit"},
            {"name": "Criteria2", "weight": 0.1, "type": "benefit"},
            {"name": "Criteria3", "weight": 0.2, "type": "cost"},
            {"name": "Criteria4", "weight": 0.2, "type": "cost"},
            {"name": "Criteria5", "weight": 0.2, "type": "cost"}
        ]
    }'
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure that your code follows the project's coding standards and includes appropriate tests.

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

