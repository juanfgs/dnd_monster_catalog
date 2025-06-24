package encounter

type CreateRequest struct {
	MinChallengeRating float64 `json:"min_challenge_rating"`
	MaxChallengeRating float64 `json:"max_challenge_rating"`
	Quantity int64 `json:"quantity"`
}

