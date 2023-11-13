package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceRequest(advice web.AdviceRequest, ComplaintID int) *schema.Advice {
	return &schema.Advice{
		Title:   advice.Title,
		Content: advice.Content,
	}
}
