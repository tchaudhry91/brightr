package backlight_test

import (
	"testing"

	"github.com/tchaudhry91/brightr/backlight"
)

// These tests are meant to pass only on my system. Change the consts accordingly
const BACKLIGHT = "amdgpu_bl0"
const MAXBRIGHTNESS = 255
const MINBRIGHTNESS = 0

func TestGetBacklights(t *testing.T) {
	backlights, err := backlight.GetBacklights()
	if err != nil {
		t.Errorf("Failed to get Backlights: %v", err)
	}
	if backlights[0] != BACKLIGHT {
		t.Errorf("Expected %s, got %s", BACKLIGHT, backlights[0])
	}
}

func TestReadCurrentBrightness(t *testing.T) {
	backlights, err := backlight.GetBacklights()
	if err != nil {
		t.Errorf("Failed to get Backlights: %v", err)
	}
	val, err := backlight.ReadCurrentBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read backlight value: %v", err)
	}
	t.Logf("Current backlight value: %d", val)
}
func TestReadMaxBrightness(t *testing.T) {
	backlights, err := backlight.GetBacklights()
	if err != nil {
		t.Errorf("Failed to get Backlights: %v", err)
	}
	val, err := backlight.ReadMaxBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read backlight value: %v", err)
	}
	t.Logf("Current max backlight value: %d", val)
}

func TestDecreaseBrightness(t *testing.T) {
	backlights, err := backlight.GetBacklights()
	if err != nil {
		t.Errorf("Failed to get Backlights: %v", err)
	}
	val, err := backlight.ReadCurrentBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read backlight value: %v", err)
	}
	t.Logf("Current backlight value: %d", val)
	t.Log("Attempting to decrease by 10")
	err = backlight.DecreaseBrightness(backlights[0], 10)
	if err != nil {
		t.Errorf("Failed to decrease brightness: %v", err)
	}

	valNew, err := backlight.ReadCurrentBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read new backlight value: %v", err)
	}
	expected := val - 10
	if expected < 0 {
		expected = 0
	}
	if valNew != expected {
		t.Errorf("Incorrect Decrease Detected!, Expected: %d, Found: %d", expected, valNew)
	}
}

func TestIncreaseBrightness(t *testing.T) {
	backlights, err := backlight.GetBacklights()
	if err != nil {
		t.Errorf("Failed to get Backlights: %v", err)
	}
	val, err := backlight.ReadCurrentBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read backlight value: %v", err)
	}
	maxVal, err := backlight.ReadMaxBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read max backlight value: %v", err)
	}
	t.Logf("Max backlight value: %d", maxVal)
	t.Logf("Current backlight value: %d", val)
	t.Log("Attempting to increase by 10")
	err = backlight.IncreaseBrightness(backlights[0], 10)
	if err != nil {
		t.Errorf("Failed to increase brightness: %v", err)
	}

	valNew, err := backlight.ReadCurrentBrightness(backlights[0])
	if err != nil {
		t.Errorf("Failed to read new backlight value: %v", err)
	}
	expected := val + 10
	if expected > maxVal {
		expected = maxVal
	}
	if valNew != expected {
		t.Errorf("Incorrect Increase Detected!, Expected: %d, Found: %d", expected, valNew)
	}
}
