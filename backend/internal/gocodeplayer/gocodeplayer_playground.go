package gocodeplayer

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"backend/internal/gocode"

	"github.com/go-playground/validator"
	"github.com/imroc/req/v3"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// options for the client
type Options func(*PlaygroundClient)

// WithLogger sets the logger for the playgroundClient
func WithLogger(logger *slog.Logger) func(*PlaygroundClient) {
	return func(c *PlaygroundClient) {
		c.logger = logger
	}
}

type PlaygroundClient struct {
	url       string
	logger    *slog.Logger
	reqClient *req.Client
}

// NewPlaygroundClient creates a new PlaygroundClient instance with the given URL
func NewPlaygroundClient(url string, opts ...Options) (*PlaygroundClient, error) {
	if url == "" {
		return nil, fmt.Errorf("url cannot be empty")
	}

	client := &PlaygroundClient{
		url: url,
	}

	for _, o := range opts {
		o(client)
	}

	// set req client
	if client.reqClient == nil {
		client.reqClient = req.NewClient()
		client.reqClient.BaseURL = url
	}

	// set default logger
	if client.logger == nil {
		client.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelError,
		}))
	}
	return client, nil
}

type PlayCodeResponse struct {
	CompileErrors string `json:"compile_errors"`
	Output        string `json:"output"`
}

// Play sends the given code to the playground and returns the output
func (c *PlaygroundClient) Play(code *gocode.GoCode) (*gocode.GoCodePlayOutput, error) {
	logger := c.logger.With("method", "Play")

	// validate the code
	if code == nil {
		return nil, fmt.Errorf("code cannot be nil")
	}

	if code.Code == "" {
		return nil, fmt.Errorf("code cannot be empty")
	}

	// URL encode the code
	urlEncodedCode := url.QueryEscape(code.Code)

	rawTextBody := fmt.Sprintf("version=&body=%s&withVet=true", urlEncodedCode)

	// send the request
	var result PlayCodeResponse

	resp, err := c.reqClient.
		R().
		SetSuccessResult(&result).
		SetContentType("application/x-www-form-urlencoded; charset=UTF-8").
		SetBodyString(rawTextBody).
		Post(c.reqClient.BaseURL + "/_/compile")
	if err != nil {
		logger.Error("error getting output", "error", err)
		return nil, err
	}
	if resp.Response.StatusCode != 200 {
		logger.Error("error getting output", "status_code",
			resp.Response.StatusCode)
		return nil, fmt.Errorf(
			"error getting output: status code: %v",
			resp.Response.StatusCode)
	}

	logger.Info(fmt.Sprintf(
		"response output: %+v\n", result))

	return &gocode.GoCodePlayOutput{
		Output: result.Output,
		Error:  result.CompileErrors,
	}, nil
}
