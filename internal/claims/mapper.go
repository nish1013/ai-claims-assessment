package claims

func ToClaimEntity(dto CreateClaimRequest) Claim {
    return Claim{
        ClaimID:    dto.ClaimID,
        PolicyID:   dto.PolicyID,
        Amount:     dto.Amount,
        Incident:   dto.Incident,
        AIDecision: DecisionPending,
		AINotes:    "",
        CreatedAt:  time.Now(),
    }
}

func ToClaimResponse(entity Claim) ClaimResponse {
    return ClaimResponse{
        ClaimID:    entity.ClaimID,
        PolicyID:   entity.PolicyID,
        Amount:     entity.Amount,
        Incident:   entity.Incident,
        AIDecision: entity.AIDecision,
        AINotes:    entity.AINotes,
        CreatedAt:  entity.CreatedAt.Format(time.RFC3339),
    }
}
