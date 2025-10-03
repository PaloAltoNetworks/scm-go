// Helper functions if you want to cleanup a bunch of objects - runs as a test but can also be coded up as a script.
// Usually used when tf state is messed up or lost and we want to get back to a clean slate programmatically.

package objects

import (
	"strings"
	"testing"
)

// CONFIGURATION: Add the UUIDs of the objects you want to delete here.
var cleanupConfig = struct {
	Addresses     []string
	AddressGroups []string
	Tags          []string
}{
	Addresses: []string{
		"73608b14-0360-4ed1-a7b4-b9082f9d18b6",
		"abe33f5b-8040-48c3-ab6a-8de679ba35b3",
		"0ecab520-b598-48af-b3d7-f044c2c2c61e",
		"060debc5-336d-4293-884d-d334a8a1cde8",
		"551b293b-e4a6-4e18-9502-e00a953337b0",
	},
	AddressGroups: []string{
		"2f2d635b-8b18-4993-92d3-28d62d70b606",
	},
	Tags: []string{
		"f2ba5e3e-8791-42c6-8e00-8b109b1afbc9",
		"fda2cc59-9714-446f-aa9c-fa72575b4d7a",
		"c169a27e-05bf-4d78-9f4e-312d308df07b",
	},
}

// Test_CleanupObjects now uses the shared helper functions.
func Test_CleanupObjects(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// Sub-test for cleaning up Addresses.
	t.Run("cleanup addresses", func(t *testing.T) {
		for _, id := range cleanupConfig.Addresses {
			if id == "" || strings.HasPrefix(id, "replace-with-") {
				continue
			}
			// Use the shared helper function.
			deleteTestAddress(t, client, id, "")
		}
	})

	// Sub-test for cleaning up Address Groups.
	t.Run("cleanup address groups", func(t *testing.T) {
		for _, id := range cleanupConfig.AddressGroups {
			if id == "" || strings.HasPrefix(id, "replace-with-") {
				continue
			}
			// Use the shared helper function.
			deleteTestAddressGroup(t, client, id, "")
		}
	})

	// Sub-test for cleaning up Tags.
	t.Run("cleanup tags", func(t *testing.T) {
		for _, id := range cleanupConfig.Tags {
			if id == "" || strings.HasPrefix(id, "replace-with-") {
				continue
			}
			// Use the new shared helper function.
			deleteTestTag(t, client, id, "")
		}
	})
}
