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
    git clone https://github.com/your-username/edas-api.git
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

## Usage

### Running the Server

Start the server with:

```sh
./edas-api

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
            {"name": "Alt1", "scores": {"Criteria1": 3.5, "Criteria2": 7.0}},
            {"name": "Alt2", "scores": {"Criteria1": 4.0, "Criteria2": 6.5}},
            {"name": "Alt3", "scores": {"Criteria1": 3.0, "Criteria2": 8.0}}
        ],
        "criteria": [
            {"name": "Criteria1", "weight": 0.5},
            {"name": "Criteria2", "weight": 0.5}
        ]
    }
    ```

- **Response:**

    ```json
    {
        "ranking": [
            {"name": "Alt3", "score": 0.75},
            {"name": "Alt2", "score": 0.25},
            {"name": "Alt1", "score": -0.25}
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
            {"name": "Alt1", "scores": {"Criteria1": 3.5, "Criteria2": 7.0}},
            {"name": "Alt2", "scores": {"Criteria1": 4.0, "Criteria2": 6.5}},
            {"name": "Alt3", "scores": {"Criteria1": 3.0, "Criteria2": 8.0}}
        ],
        "criteria": [
            {"name": "Criteria1", "weight": 0.5},
            {"name": "Criteria2", "weight": 0.5}
        ]
    }'

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

