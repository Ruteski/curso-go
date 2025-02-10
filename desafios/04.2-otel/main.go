package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type TemperatureResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type CEPRequest struct {
	CEP string `json:"cep"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão caso a variável de ambiente não esteja definida
	}

	http.HandleFunc("/weather", handleWeatherRequest)
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método da requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lê o corpo da requisição
	var cepRequest CEPRequest
	err := json.NewDecoder(r.Body).Decode(&cepRequest)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	cep := cepRequest.CEP
	if len(cep) != 8 || !isNumeric(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := getCityFromCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := getTemperature(city)
	if err != nil {
		http.Error(w, "failed to get temperature: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := TemperatureResponse{
		City:  city,
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getCityFromCEP(cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var viaCEPResponse ViaCEPResponse
	err = json.Unmarshal(body, &viaCEPResponse)
	if err != nil {
		return "", err
	}

	if viaCEPResponse.Localidade == "" {
		return "", fmt.Errorf("city not found")
	}

	return viaCEPResponse.Localidade, nil
}

func getTemperature(city string) (float64, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=43f6d105117a49d698b201736250302&q=%s", encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro na requisição: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if !json.Valid(body) {
		return 0, fmt.Errorf("resposta inválida da API: %s", string(body))
	}

	var weatherAPIResponse WeatherAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		return 0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
