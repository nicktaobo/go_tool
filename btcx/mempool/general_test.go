package mempool_test

import (
	"github.com/nicktaobo/go_tool/btcx"
	"github.com/nicktaobo/go_tool/btcx/mempool"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client = mempool.NewClient(btcx.MainNet)

func TestGetPrices(t *testing.T) {
	s, err := client.GetPrices()
	if err != nil {
		t.Fatalf("test failed: %v", err)
	}
	t.Log("result:", s)
	assert.True(t, s.Time > 0)
	assert.True(t, s.USD > 0)
}

func TestGetDifficultyAdjustment(t *testing.T) {
	s, err := client.GetDifficultyAdjustment()
	if err != nil {
		t.Fatalf("test failed: %v", err)
	}
	t.Log("result:", s)
	assert.True(t, s.RemainingBlocks > 0)
}
