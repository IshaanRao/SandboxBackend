package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sandboxbackend/models"
)

// Ranks - Enum-like structure for player ranks
type Ranks string

const (
	DEFAULT Ranks = "DEFAULT"
	VIP     Ranks = "VIP"
	HELPER  Ranks = "HELPER"
	MOD     Ranks = "MOD"
	ADMIN   Ranks = "ADMIN"
)

// SandboxPlayer represents the player data
type SandboxPlayer struct {
	UUID       string `json:"uuid"`
	PlayerRank Ranks  `json:"playerRank"`
	StaffRank  Ranks  `json:"staffRank"`
}

// SetRankPayload represents the JSON payload for setting player rank
type SetRankPayload struct {
	Rank string `json:"rank"`
}

// GetPlayer handles GET requests to /players/:uuid
func GetPlayer(c *gin.Context) {
	uuid := c.Param("uuid")

	// Load players from the mock database
	players, err := models.LoadPlayers()
	if err != nil {
		fmt.Println("Failed to load player data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load player data"})
		return
	}

	var foundPlayer *models.Player
	for i, p := range players {
		if p.UUID == uuid {
			foundPlayer = &players[i]
			break
		}
	}

	// If player not found, create a default player
	if foundPlayer == nil {
		// Create a default player
		defaultPlayer := models.Player{
			UUID:       uuid,
			PlayerRank: "DEFAULT",
			StaffRank:  "DEFAULT",
			Inventory:  &models.Inventory{},
		}

		// Append the default player to the players list
		players = append(players, defaultPlayer)
		foundPlayer = &defaultPlayer

		// Save players back to the JSON file
		if err := models.SavePlayers(players); err != nil {
			fmt.Println("Failed to save player data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save player data"})
			return
		}
	}

	c.JSON(http.StatusOK, foundPlayer)
}

// SetRank handles POST requests to /players/setrank/:uuid
func SetRank(c *gin.Context) {
	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")

	// Load players from the mock database
	players, err := models.LoadPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load player data"})
		return
	}

	// Find the player by UUID
	var foundPlayer *models.Player
	for i, p := range players {
		if p.UUID == uuid {
			foundPlayer = &players[i]
			break
		}
	}

	// If player not found, return error
	if foundPlayer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Parse JSON payload
	var payload SetRankPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update player's rank
	foundPlayer.PlayerRank = payload.Rank

	// Save players back to the mock database
	err = models.SavePlayers(players)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save player data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player rank updated successfully"})
}

// GetInvContents handles GET requests to /players/inv/:uuid
func GetInvContents(c *gin.Context) {
	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")

	// Load players from the mock database
	players, err := models.LoadPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load player data"})
		return
	}

	// Find the player by UUID
	var foundPlayer *models.Player
	for i, p := range players {
		if p.UUID == uuid {
			foundPlayer = &players[i]
			break
		}
	}

	// If player not found, return error
	if foundPlayer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Create SandboxInventory object from player data
	invContents := &models.SandboxInventory{
		UUID: uuid,
	}
	if foundPlayer.Inventory != nil {
		invContents.InvContents = foundPlayer.Inventory.Contents
		invContents.ArmorContents = foundPlayer.Inventory.ArmorContents
	}

	// Return SandboxInventory data
	c.JSON(http.StatusOK, invContents)
}

// SetInvContents handles POST requests to /players/setinvcontents/:uuid
func SetInvContents(c *gin.Context) {
	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")

	// Parse JSON payload
	var payload map[string]string
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Load players from the mock database
	players, err := models.LoadPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load player data"})
		return
	}

	// Find the player by UUID
	var foundPlayer *models.Player
	for i, p := range players {
		if p.UUID == uuid {
			foundPlayer = &players[i]
			break
		}
	}

	// If player not found, return error
	if foundPlayer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Update player's inventory contents
	foundPlayer.Inventory.Contents = payload["invContents"]

	// Save players back to the mock database
	if err := models.SavePlayers(players); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player data"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Inventory contents updated successfully"})
}

// SetArmorContents handles POST requests to /players/setarmorcontents/:uuid
func SetArmorContents(c *gin.Context) {
	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")

	// Parse JSON payload
	var payload map[string]string
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Load players from the mock database
	players, err := models.LoadPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load player data"})
		return
	}

	// Find the player by UUID
	var foundPlayer *models.Player
	for i, p := range players {
		if p.UUID == uuid {
			foundPlayer = &players[i]
			break
		}
	}

	// If player not found, return error
	if foundPlayer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Update player's armor contents
	foundPlayer.Inventory.ArmorContents = payload["armorContents"]

	// Save players back to the mock database
	if err := models.SavePlayers(players); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player data"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Armor contents updated successfully"})
}
