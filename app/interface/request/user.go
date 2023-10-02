package request

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserScoreUpdateRequest struct {
	Score string `json:"score"`
}
