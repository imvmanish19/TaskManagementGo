package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define the User struct in task-service based on what user-service returns
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

const userServiceURL = "http://user-service:8081/users/"

// UserServiceClient manages HTTP communication with the user-service
type UserServiceClient struct {
	Client *http.Client
}

// NewUserServiceClient creates a new instance of UserServiceClient
func NewUserServiceClient() *UserServiceClient {
	return &UserServiceClient{
		Client: &http.Client{},
	}
}

// GetUser makes a GET request to the user-service to fetch a user by ID
func (c *UserServiceClient) GetUser(userID string) (*User, error) {
	// Make the GET request to user-service
	resp, err := c.Client.Get(fmt.Sprintf("%s%s", userServiceURL, userID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the response into the User struct
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	return &user, nil
}
