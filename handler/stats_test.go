package handler

import (
	"fmt"
	"golang-api/model"
	"testing"
)

func TestProcess(t *testing.T) {
	expected := model.HashRequestStats{1, 5123, 5123}
	actual := model.HashRequestStats{}
	actual.Process(5123)
	if expected != actual {
		t.Errorf("The expected and actual HashRequestStats structs for the input string '%v' should be equal.", 5123)
	}
	fmt.Println(expected)
	fmt.Println(actual)
}

func TestProcessFail(t *testing.T) {
	expected := model.HashRequestStats{1, 5123, 5123}
	actual := model.HashRequestStats{}
	if expected == actual {
		t.Errorf("The expected and actual HashRequestStats structs for the input string '%v' should not be equal.", 5123)
	}
	fmt.Println(expected)
	fmt.Println(actual)
}
