package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) HandleSubscriptionPaymentSuccess(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("x-webhook-signature")
	timestamp := r.Header.Get("x-webhook-timestamp")

	body, _ := io.ReadAll(r.Body)

	webhookSignature, err := s.paymentService.VerifyWebhookSignature(signature, timestamp, string(body))
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	var webhookData map[string]interface{}

	switch v := webhookSignature.Object.(type) {
	case string:
		if err := json.Unmarshal([]byte(v), &webhookData); err != nil {
			s.HandleError(w, err, "failed to unmarshal webhook object", http.StatusBadRequest)
			return
		}
	case map[string]interface{}:
		webhookData = v
	default:
		s.HandleError(w, err, "invalid webhook object type", http.StatusInternalServerError)
		return
	}

	var orderId, userEmail string
	if data, ok := webhookData["data"].(map[string]interface{}); ok {
		if order, ok := data["order"].(map[string]interface{}); ok {
			if orderID, ok := order["order_id"].(string); ok {
				orderId = orderID
			}
		}

		if customerDetails, ok := data["customer_details"].(map[string]interface{}); ok {
			if email, ok := customerDetails["customer_email"].(string); ok {
				userEmail = email
			}
		}
	}

	go func() {
		bgCtx := context.Background()

		activatedSubscription, err := s.subscriptionService.ActivateUserSubscription(bgCtx, orderId, userEmail)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println("subscription created with id", activatedSubscription.Id)
	}()

	err = s.WriteJson(w, http.StatusOK, &Response{})
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
