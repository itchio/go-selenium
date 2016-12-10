package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium/src"
)

func Test_CommandSwitchToParentFrame_CorrectResponseCanBeReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst navigating.", err)
	}

	idx, _ := goselenium.ByIndex(0)
	_, err = driver.SwitchToFrame(idx)
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst switching to frame 0.", err)
	}

	resp, err := driver.SwitchToParentFrame()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was thrown or response was not a success.", err)
	}

	printObjectResult(resp)
}