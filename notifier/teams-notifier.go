package notifier
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)
type Config struct {
	WebhookURL string `json:"webhook_url"`
}
type teamsMessage struct {
	Type string `json:"@type"`
	Context string `json:"@context"`
	ThemeColor string `json:"themeColor"`
	Summary string `json:"summary"`
	Sections []teamsSection `json:"sections"`
}
type teamsSection struct {
	ActivityTitle string `json:"activityTitle"`
	Facts []teamsFact `json:"facts"`
}
type teamsFact struct {
	Name string `json:"name"`
	Value string `json:"value"`
}
func Notify(config Config, alert string) error {
	message := teamsMessage{
		Type: "MessageCard",
		Context: "http://schema.org/extensions",
		ThemeColor: "0078D7",
		Summary: "Consul Alert",
		Sections: []teamsSection{
			{
				ActivityTitle: "Consul Alert",
				Facts: []teamsFact{
					{
						Name: "Alert",
						Value: alert,
					},
				},
			},
		},
	}
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}
	resp, err := http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response status: %s", resp.Status)
	}
	return nil
}
func Validate(config Config) error {
	if config.WebhookURL == "" {
		return fmt.Errorf("Webhook URL must be set")
	}
	return nil
}