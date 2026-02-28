package vertexai

import (
	"context"
	"fmt"
	"cloud.google.com/go/vertexai/genai"
)

type Client struct {
	GenModel *genai.GenerativeModel
}

func NewClient(ctx context.Context, projectID, region, modelName string) (*Client, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, err
	}
	
	model := client.GenerativeModel(modelName)
	return &Client{GenModel: model}, nil
}

func (c *Client) GenerateResponse(ctx context.Context, prompt string) (string, error) {
	resp, err := c.GenModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	// Extract text from the first candidate
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
	}

	return "No diagnosis generated", nil
}
