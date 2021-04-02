package hirezapi

import (
	"fmt"
	"io/ioutil"
)

// GetGods returns all Gods and their various attributes.
func (a *APIClient) GetGods() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetGodSkins returns all available skins for a particular God.
func (a *APIClient) GetGodSkins() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetGodRecommendedItems returns the recommended items for a particular God.
func (a *APIClient) GetGodRecommendedItems() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}

// GetItems returns all items and their various attributes.
func (a *APIClient) GetItems() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil
}
