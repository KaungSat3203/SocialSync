package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func FetchFacebookPostAnalytics(postID string, accessToken string) (map[string]interface{}, error) {
	// 🛠️ Strip compound ID like "pageID_postID"
	if strings.Contains(postID, "_") {
		parts := strings.Split(postID, "_")
		if len(parts) == 2 {
			postID = parts[1]
		}
	}

	endpoint := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/insights", postID)
	params := url.Values{}

	// ✅ Use only metrics that are valid for Page Posts
	params.Set("metric", "post_impressions,post_engaged_users,post_reactions_by_type_total")
	params.Set("access_token", accessToken)

	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to Facebook Insights: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("facebook API error: %s", string(body))
	}

	var result struct {
		Data []struct {
			Name   string `json:"name"`
			Values []struct {
				Value interface{} `json:"value"`
			} `json:"values"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	metrics := make(map[string]interface{})
	for _, metric := range result.Data {
		if len(metric.Values) > 0 {
			metrics[metric.Name] = metric.Values[0].Value
		}
	}

	return metrics, nil
}
