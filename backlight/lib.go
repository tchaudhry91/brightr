package backlight

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const BACKLIGHTDIR = "/sys/class/backlight"
const BRIGHTNESSFILE = "brightness"
const MAXBRIGHTNESSFILE = "max_brightness"

func GetBacklights() (backlights []string, err error) {
	backlights = []string{}
	backDirs, err := ioutil.ReadDir(BACKLIGHTDIR)
	if err != nil {
		return
	}
	for _, d := range backDirs {
		backlights = append(backlights, d.Name())
	}
	if len(backlights) == 0 {
		return backlights, errors.New("No Backlight Directory Found")
	}
	return backlights, nil
}

func ReadCurrentBrightness(device string) (currentBacklight int, err error) {
	valBytes, err := ioutil.ReadFile(filepath.Join(BACKLIGHTDIR, device, BRIGHTNESSFILE))
	if err != nil {
		return 0, err
	}
	valStr := strings.TrimSpace(string(valBytes))
	return strconv.Atoi(valStr)
}

func ReadMaxBrightness(device string) (currentBacklight int, err error) {
	valBytes, err := ioutil.ReadFile(filepath.Join(BACKLIGHTDIR, device, MAXBRIGHTNESSFILE))
	if err != nil {
		return 0, err
	}
	valStr := strings.TrimSpace(string(valBytes))
	return strconv.Atoi(valStr)
}

func DecreaseBrightness(device string, amount int) (err error) {
	current, err := ReadCurrentBrightness(device)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(BACKLIGHTDIR, device, BRIGHTNESSFILE), os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(strconv.Itoa(maxInt(current-amount, 0)))
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func IncreaseBrightness(device string, amount int) (err error) {
	current, err := ReadCurrentBrightness(device)
	if err != nil {
		return err
	}
	max, err := ReadMaxBrightness(device)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(BACKLIGHTDIR, device, BRIGHTNESSFILE), os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(strconv.Itoa(minInt(current+amount, max)))
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// minInt returns the smaller of x or y.
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// maxInt returns the larger of x or y.
func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
