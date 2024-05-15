package models

import (
	"encoding/json"
	"os"
)

// Player represents the player data
type Player struct {
	UUID       string     `json:"uuid"`
	PlayerRank string     `json:"playerRank"`
	StaffRank  string     `json:"staffRank"`
	Inventory  *Inventory `json:"inventory,omitempty"`
}

// Inventory represents the player's inventory contents
type Inventory struct {
	Contents      string `json:"contents,omitempty"`
	ArmorContents string `json:"armorContents,omitempty"`
}

// SandboxInventory represents the inventory and armor contents for a player
type SandboxInventory struct {
	UUID          string `json:"uuid"`
	InvContents   string `json:"invContents,omitempty"`
	ArmorContents string `json:"armorContents,omitempty"`
}

// LoadPlayers loads player data from the JSON file
func LoadPlayers() ([]Player, error) {
	data, err := os.ReadFile("players.json")
	if err != nil {
		return nil, err
	}

	var players []Player
	if err := json.Unmarshal(data, &players); err != nil {
		return nil, err
	}

	return players, nil
}

// SavePlayers saves player data to the JSON file
func SavePlayers(players []Player) error {
	data, err := json.MarshalIndent(players, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("players.json", data, 0644)
}
