package service

import (
	"encoding/json"
	"io/ioutil"
	"othonas/cmd/config"
	"slices"
)

// CalculatePacks Function to calculate the pack distribution.
// Will return a map with the pack size as the key and the count as the value. It uses a dynamic programming approach.
func CalculatePacks(order int, packSizes []int) map[int]int {
	// Define a struct to store DP entries
	type dpEntry struct {
		totalPacks int
		packCounts map[int]int // packSize -> count
	}
	// Initialize DP table
	dp := make(map[int]dpEntry)
	dp[0] = dpEntry{totalPacks: 0, packCounts: make(map[int]int)}

	// Find the maximum pack size for setting an upper limit
	maxPackSize := slices.Max(packSizes)

	// Build the DP table
	maxTotalItems := order + maxPackSize
	for totalItemsSent := 1; totalItemsSent <= maxTotalItems; totalItemsSent++ {
		for _, p := range packSizes {
			if totalItemsSent >= p {
				if prevEntry, ok := dp[totalItemsSent-p]; ok {
					newTotalPacks := prevEntry.totalPacks + 1
					// Check if this is a better (fewer packs) combination
					if currEntry, exists := dp[totalItemsSent]; !exists || currEntry.totalPacks > newTotalPacks {
						// Create a new packCounts map
						newPackCounts := make(map[int]int)
						for k, v := range prevEntry.packCounts {
							newPackCounts[k] = v
						}
						newPackCounts[p] += 1
						dp[totalItemsSent] = dpEntry{totalPacks: newTotalPacks, packCounts: newPackCounts}
					}
				}
			}
		}
	}
	// Find the minimal totalItemsSent >= orderAmount
	minTotalItemsSent := 0
	minTotalPacks := 0
	var resultPackCounts map[int]int
	found := false
	for totalItemsSent := order; totalItemsSent <= maxTotalItems; totalItemsSent++ {
		if entry, ok := dp[totalItemsSent]; ok {
			if !found || totalItemsSent < minTotalItemsSent || (totalItemsSent == minTotalItemsSent && entry.totalPacks < minTotalPacks) {
				minTotalItemsSent = totalItemsSent
				minTotalPacks = entry.totalPacks
				resultPackCounts = entry.packCounts
				found = true
			}
		}
	}
	return resultPackCounts
}

// Function to load configuration from a JSON file
func LoadPackSizesFromFile(filePath string) []int {
	var config config.Config
	// Read the file
	file, err := ioutil.ReadFile(filePath)
	// there is no need to continue if there is an error, we cannot continue without the config file
	if err != nil {
		panic(err)
	}
	// Parse the JSON
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config.PackSizes
}
