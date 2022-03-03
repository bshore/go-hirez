package mock

import (
	"fmt"

	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

// GetGods returns all Gods and their various attributes.
func (a *APIClient) GetGods(langCode string) ([]models.God, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGods() %s", utils.NotSmiteErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	resp, err := a.makeRequest("getgods", langCode, []models.God{})
	if err != nil {
		return nil, err
	}
	var output []models.God
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetGodAltAbilities returns alt abilities for all Gods.
func (a *APIClient) GetGodAltAbilities() ([]models.GodAltAbility, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodAltAbilities(), %s", utils.NotSmiteErrMsg)
	}
	resp, err := a.makeRequest("getgodaltabilities", "", []models.GodAltAbility{})
	if err != nil {
		return nil, err
	}
	var output []models.GodAltAbility
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetGodSkins returns all available skins for a particular God.
func (a *APIClient) GetGodSkins(godID int64, langCode string) ([]models.GodSkin, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodSkins() %s", utils.NotSmiteErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	path := fmt.Sprintf("%d/%s", godID, langCode)
	resp, err := a.makeRequest("getgodskins", path, []models.GodSkin{})
	if err != nil {
		return nil, err
	}
	var output []models.GodSkin
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetGodRecommendedItems returns the recommended items for a particular God.
func (a *APIClient) GetGodRecommendedItems(godID int64, langCode string) ([]models.GodRecommendedItem, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodRecommendedItems() %s", utils.NotSmiteErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	path := fmt.Sprintf("%d/%s", godID, langCode)
	resp, err := a.makeRequest("getgodrecommendeditems", path, models.GodRecommendedItem{})
	if err != nil {
		return nil, err
	}
	var output []models.GodRecommendedItem
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetItems returns all items and their various attributes.
func (a *APIClient) GetItems(langCode string) ([]models.Item, error) {
	if langCode == "" {
		langCode = models.English
	}
	resp, err := a.makeRequest("getitems", langCode, []models.Item{})
	if err != nil {
		return nil, err
	}
	var output []models.Item
	err = a.unmarshalResponse(resp, &output)
	return output, err
}
