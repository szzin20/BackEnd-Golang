package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceResponse(advice *schema.Advice) web.AdviceResponse {
	return web.AdviceResponse{
		ID:          advice.ID,
		ComplaintID: advice.ComplaintID,
		Title:       advice.Title,
		Content:     advice.Content,
		Status:      advice.Status,
		CreatedAt:   advice.CreatedAt,
	}
}

func ConvertToGetAllAdvicesResponse(advices []schema.Advice) []web.AdviceResponse {
	var results []web.AdviceResponse
	for _, advice := range advices {
		adviceResponse := web.AdviceResponse{
			ID:          advice.ID,
			ComplaintID: advice.ComplaintID,
			Title:       advice.Title,
			Content:     advice.Content,
			Status:      advice.Status,
			CreatedAt:   advice.CreatedAt,
		}
		results = append(results, adviceResponse)
	}
	return results
}
