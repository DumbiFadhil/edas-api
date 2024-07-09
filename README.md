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

