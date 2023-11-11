package opening

import "github.com/DamiaoCanndido/gopportunities/entities"

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type ShowOpeningsResponse struct {
	Message string                     `json:"message"`
	Data    []entities.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}
