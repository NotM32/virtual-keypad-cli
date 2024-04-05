package query

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	endpoint = "https://api.securecomwireless.com/"
)

type Door struct {
	Number    string `json:"number"`
	Name      string `json:"name"`
	Panel     string
	AuthToken string
}

type DoorRequest struct {
	Status string `json:"status"`
	Type   string `json:"type"`
}

// Access method triggers the access funciton for the door and opens the strike for the duration
func (door *Door) Access() {
	client := resty.New()

	url := fmt.Sprintf("%s/1/panels/%s/output_statuses/%s?auth_token=%s", endpoint, door.Panel, door.Number, door.AuthToken)

	var job JobResponse
	_, err := client.R().
		EnableTrace().
		SetResult(&job).
		SetBody(DoorRequest{
			Status: "ACCESS",
			Type:   "door",
		}).
		Post(url)
	if err != nil {
		fmt.Println("Could not request to trigger access: ", err)
	}
}

// Unlock method forces the door to stay unlocked until relocked
func (door *Door) Unlock() {
	client := resty.New()

	url := fmt.Sprintf("%s/1/panels/%s/output_statuses/%s?auth_token=%s", endpoint, door.Panel, door.Number, door.AuthToken)

	var job JobResponse
	_, err := client.R().
		EnableTrace().
		SetResult(&job).
		SetBody(DoorRequest{
			Status: "UNLOCK",
			Type:   "door",
		}).
		Post(url)
	if err != nil {
		fmt.Println("Could not request to unlock door: ", err)
	}
}

// Lock method forces locks the door
func (door *Door) Lock() {
	client := resty.New()

	url := fmt.Sprintf("%s/1/panels/%s/output_statuses/%s?auth_token=%s", endpoint, door.Panel, door.Number, door.AuthToken)

	var job JobResponse
	_, err := client.R().
		EnableTrace().
		SetResult(&job).
		SetBody(DoorRequest{
			Status: "LOCK",
			Type:   "door",
		}).
		Post(url)
	if err != nil {
		fmt.Println("Could not request to lock door: ", err)
	}
}

// GetDoors function returns a list of Door objects that can be interacted with
func GetDoors(authToken string, userCode string, panel string, panelSecondary string) []Door {
	client := resty.New()

	url := fmt.Sprintf("%s/v2/panels/%s/doors?page=1&page_size=100&auth_token=%s&auth_user_code=%s", endpoint, panel, authToken, userCode)

	var deviceinfo []struct {
		Door Door `json:"device_information"`
	}

	_, err := client.R().
		EnableTrace().
		SetResult(&deviceinfo).
		Get(url)
	if err != nil {
		fmt.Println("Could not fetch doors: ", err)
	}

	var doors []Door
	for _, device := range deviceinfo {
		device.Door.Panel = panelSecondary
		device.Door.AuthToken = authToken
		doors = append(doors, device.Door)
	}
	return doors
}
