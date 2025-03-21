package claims

type CreateClaimRequest struct {
    ClaimID    string  `json:"claim_id" binding:"required"`
    PolicyID   string  `json:"policy_id" binding:"required"`
    Amount     float64 `json:"amount" binding:"required"`
    Incident   string  `json:"incident_type" binding:"required"`
}

type ClaimResponse struct {
    ClaimID    string  `json:"claim_id"`
    PolicyID   string  `json:"policy_id"`
    Amount     float64 `json:"amount"`
    Incident   string  `json:"incident_type"`
    AIDecision string  `json:"ai_decision"`
    AINotes    string  `json:"ai_notes"`
    CreatedAt  string  `json:"created_at"`
}
