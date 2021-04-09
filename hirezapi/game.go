package hirezapi

import (
	"fmt"
	"io/ioutil"

	"github.com/bshore/go-hirez/models"
)

// GetGods returns all Gods and their various attributes.
func (a *APIClient) GetGods(langCode string) ([]models.God, error) {
	if !IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGods() %s", notSmiteErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	resp, err := a.makeRequest("getgods", langCode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.God
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetGodSkins returns all available skins for a particular God.
func (a *APIClient) GetGodSkins(godID int64, langCode string) ([]models.GodSkin, error) {
	if !IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodSkins() %s", notSmiteErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	path := fmt.Sprintf("%d/%s", godID, langCode)
	resp, err := a.makeRequest("getgodskins", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.GodSkin
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetGodRecommendedItems returns the recommended items for a particular God.
func (a *APIClient) GetGodRecommendedItems(godID int64, langCode string) ([]models.GodRecommendedItem, error) {
	if !IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodRecommendedItems() %s", notSmiteErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	path := fmt.Sprintf("%d/%s", godID, langCode)
	resp, err := a.makeRequest("getgodrecommendeditems", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.GodRecommendedItem
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetItems returns all items and their various attributes.
func (a *APIClient) GetItems(langCode string) ([]models.Item, error) {
	if langCode == "" {
		langCode = English
	}
	resp, err := a.makeRequest("getitems", langCode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Item
	err = a.unmarshalResponse(body, &output)
	return output, err
}
