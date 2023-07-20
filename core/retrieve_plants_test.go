package core

import (
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetPlantData(t *testing.T) {
	// Mock HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockUrl := "http://mock.url"
	mockResponse := `<html>
		<h1>Plant name</h1>
		<div id="plantinfoouter">
			<ul>
				<li>Some data</li>
				<li>Latin Name: Latin name</li>
				<li>Family Name: Family name</li>
			</ul>
		</div>
		<div id="description">
			<p>Description</p>
		</div>
	</html>`
	httpmock.RegisterResponder("GET", mockUrl, httpmock.NewStringResponder(200, mockResponse))

	// Call the function we're testing
	plant, err := GetPlantData(mockUrl)

	// Assert function results
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if plant.Name != "Plant name" {
		t.Errorf("expected 'Plant name', got '%v'", plant.Name)
	}

	if plant.LatinName != " Latin name" {
		t.Errorf("expected 'Latin name', got '%v'", plant.LatinName)
	}

	if plant.FamilyName != " Family name" {
		t.Errorf("expected 'Family name', got '%v'", plant.FamilyName)
	}

	if !strings.Contains(plant.Description, "Description") {
		t.Errorf("expected description to contain 'Description', got '%v'", plant.Description)
	}
}

func TestGetPlants(t *testing.T) {
	// Mock HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockUrl := "http://mock.url"
	mockResponse := `<html>
		<div class="plantbox">
			<div class="dbseemore" href="http://mock.url/plant"></div>
		</div>
	</html>`
	httpmock.RegisterResponder("GET", mockUrl, httpmock.NewStringResponder(200, mockResponse))

	mockPlantUrl := "http://mock.url/plant"
	mockPlantResponse := `<html>
		<h1>Plant name</h1>
		<div id="plantinfoouter">
			<ul>
				<li>Some data</li>
				<li>Latin Name: Latin name</li>
				<li>Family Name: Family name</li>
			</ul>
		</div>
		<div id="description">
			<p>Description</p>
		</div>
	</html>`
	httpmock.RegisterResponder("GET", mockPlantUrl, httpmock.NewStringResponder(200, mockPlantResponse))

	// Call the function we're testing
	plants, err := GetPlants(mockUrl)

	// Assert function results
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if len(plants) != 1 {
		t.Errorf("expected 1 plant, got %v", len(plants))
	}

	plant := plants[0]
	if plant.Name != "Plant name" {
		t.Errorf("expected 'Plant name', got '%v'", plant.Name)
	}

	if plant.LatinName != " Latin name" {
		t.Errorf("expected 'Latin name', got '%v'", plant.LatinName)
	}

	if plant.FamilyName != " Family name" {
		t.Errorf("expected 'Family name', got '%v'", plant.FamilyName)
	}

	if !strings.Contains(plant.Description, "Description") {
		t.Errorf("expected description to contain 'Description', got '%v'", plant.Description)
	}
}
