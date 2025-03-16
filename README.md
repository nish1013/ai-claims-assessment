# AI-Powered Claims Assessment Service

## Overview

This repository contains a Golang-based AI microservice designed to assist insurance claim assessment by providing risk scoring, fraud detection, and claim summarization. The service can be integrated into existing claims and underwriting systems to enhance automation and decision-making.

## Features

- **Risk Scoring**: Assigns a risk score (1-10) based on claim details.
- **Fraud Detection**: Analyzes claims for potential fraud based on historical data.
- **Claim Summarization**: Generates a concise summary of claim reports.
- **REST API**: Exposes endpoints for claims applications to interact with AI models.
- **Test-Driven Development**: Includes unit tests to ensure reliability.
- **Docker Support**: Easily deployable via containerization.

## Technology Stack

- **Language**: Golang
- **Web Framework**: Gin
- **Database**: MongoDB
- **Cache**: Redis (optional)
- **AI Integration**: OpenAI API or local ML models
- **Containerization**: Docker

## API Endpoints

### 1. Claim Risk Scoring

**Endpoint:** `POST /risk-score`

```json
{
  "claim_id": "CLAIM123",
  "policy_id": "POL456",
  "amount": 10000,
  "incident_type": "Cargo Theft",
  "claim_history": ["CLAIM001", "CLAIM002"]
}
```

**Response:**

```json
{ "risk_score": 8 }
```

### 2. Fraud Detection

**Endpoint:** `POST /fraud-check`

```json
{
  "claim_id": "CLAIM123",
  "policy_id": "POL456",
  "amount": 50000,
  "incident_description": "Stolen shipment but no police report",
  "claimant_history": ["CLAIM001"]
}
```

**Response:**

```json
{ "fraud_risk": "HIGH" }
```

### 3. Claim Summarization

**Endpoint:** `POST /summarize-claim`

```json
{
  "claim_report": "The insured reported a total loss of cargo due to an accident on 5th March 2025. No injuries, police report filed."
}
```

**Response:**

```json
{ "summary": "Cargo lost in accident on 5th March 2025. No injuries, police report filed." }
```

## c

1. Clone the repository:
   ```sh
   git clone https://github.com/nish1013/ai-claims-assessment.git
   cd ai-claims-assessment
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the service:
   ```sh
   go run main.go
   ```
4. Run tests:
   ```sh
   go test ./...
   ```

## Contribution Guidelines

- Follow clean code principles.
- Ensure all features are covered with unit tests.
- Use pull requests for changes and improvements.

## License

This project is licensed under the MIT License.

