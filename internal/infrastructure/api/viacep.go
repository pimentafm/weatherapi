package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCEPAPI struct{}

type viaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
	Erro        string `json:"erro"`
}

func NewViaCEPAPI() *ViaCEPAPI {
	return &ViaCEPAPI{}
}

func (v *ViaCEPAPI) GetCity(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return "", fmt.Errorf("invalid zipcode")
	}

	var viaCEPResp viaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return "", err
	}

	if viaCEPResp.Erro == "true" {
		return "", fmt.Errorf("can not find zipcode")
	}

	return viaCEPResp.Localidade, nil
}
