# EDAS API

![EDAS API](https://img.shields.io/badge/EDAS-API-blue.svg)

## Overview

EDAS (Evaluation based on Distance from Average Solution) API is a decision support system implemented in Go (Golang). This API allows users to evaluate multiple alternatives based on various criteria and rank them accordingly. It provides a robust and efficient solution for multi-criteria decision making.

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

- **URL:** `/api/edas`
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

## Example

Here's a quick example of how to use the API with `curl`:

```sh
curl -X POST http://localhost:8080/api/edas \
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

