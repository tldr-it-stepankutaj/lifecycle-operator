package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/log"
)

type WebhookPayload struct {
	Operator string `json:"operator"`
	Current  string `json:"current"`
	Latest   string `json:"latest"`
	Message  string `json:"message"`
}

func SendWebhookNotification(ctx context.Context, webhookURL, operator, currentVersion, latestVersion, message string) error {
	logger := log.FromContext(ctx)

	payload := WebhookPayload{
		Operator: operator,
		Current:  currentVersion,
		Latest:   latestVersion,
		Message:  message,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook payload: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		logger.Error(err, "failed to send webhook notification")
		return fmt.Errorf("failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned non-success status: %d", resp.StatusCode)
	}

	logger.Info("webhook notification sent", "operator", operator, "status", resp.Status)
	return nil
}
