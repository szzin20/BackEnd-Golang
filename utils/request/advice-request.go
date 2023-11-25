package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceRequest(advice web.AdviceRequest, TransactionID uint) *schema.Advice {
	return &schema.Advice{
		TransactionID: TransactionID,
		Message:       advice.Message,
		Image:         advice.Image,
		Audio:         advice.Audio,
	}
}
