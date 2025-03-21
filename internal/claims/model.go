package claims

import "time"

type Claim struct {
    ClaimID     string    `json:"claim_id" bson:"claim_id"`
    PolicyID    string    `json:"policy_id" bson:"policy_id"`
    Amount      float64   `json:"amount" bson:"amount"`
    Incident    string    `json:"incident_type" bson:"incident_type"`
    AIDecision  AIDecision    `json:"ai_decision" bson:"ai_decision"`
    AINotes     string    `json:"ai_notes" bson:"ai_notes"`
    CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at"`
}

type AIDecision string

const (
    DecisionApproved AIDecision = "APPROVED"
    DecisionRejected AIDecision = "REJECTED"
    DecisionReview   AIDecision = "REVIEW"
	DecisionPending  AIDecision = "PENDING"
)

