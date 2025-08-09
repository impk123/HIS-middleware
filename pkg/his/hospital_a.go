package his

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HospitalAClient struct {
	BaseURL string
}

func NewHospitalAClient(baseURL string) *HospitalAClient {
	return &HospitalAClient{BaseURL: baseURL}
}

func (c *HospitalAClient) SearchPatient(id string) (*PatientResponse, error) {
	url := fmt.Sprintf("%s/patient/search/%s", c.BaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("hospital A API returned non-200 status code")
	}

	var patient PatientResponse
	if err := json.NewDecoder(resp.Body).Decode(&patient); err != nil {
		return nil, err
	}

	return &patient, nil
}

type PatientResponse struct {
	FirstNameTH  string `json:"first_name_th"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th"`
	FirstNameEN  string `json:"first_name_en"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en"`
	DateOfBirth  string `json:"date_of_birth"`
	PatientHN    string `json:"patient_hn"`
	NationalID   string `json:"national_id"`
	PassportID   string `json:"passport_id"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
}
