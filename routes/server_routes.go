package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server represents a server in the response
type Server struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

// ServersResponse represents the response format for the /servers/list endpoint
type ServersResponse struct {
	Proxy Server   `json:"proxy"`
	Hubs  []Server `json:"hubs"`
}

// APIEndpoint represents the API endpoint for updating servers
const APIEndpoint = "http://127.0.0.1:5712/servers/updateservers"

// ProxyReady handles POST requests to /servers/proxyready
func ProxyReady(c *gin.Context) {
	// Send POST request to the predefined API endpoint
	if err := sendPOSTRequest(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send POST request"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "POST request sent successfully"})
}

// sendPOSTRequest sends a POST request to the predefined API endpoint with a header containing the API key
func sendPOSTRequest() error {
	// Create a new HTTP request with an empty body
	req, err := http.NewRequest("POST", APIEndpoint, nil)
	if err != nil {
		return err
	}

	// Add the API key header
	req.Header.Set("apiKey", "f6998720-87d5-43df-aca7-764475c35e90")

	// Send the POST request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// ListServers handles GET requests to /servers/list
func ListServers(c *gin.Context) {
	// Define proxy and hub servers
	proxy := Server{Name: "Proxy", Port: 25577}
	hubs := []Server{
		{Name: "Hub1", Port: 25565},
		// Add more hub servers if needed
	}

	// Create the ServersResponse object
	response := ServersResponse{
		Proxy: proxy,
		Hubs:  hubs,
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, response)
}
