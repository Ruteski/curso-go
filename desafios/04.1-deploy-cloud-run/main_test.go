package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Estrutura do serviço para injeção de dependência
type WeatherService struct {
	GetCityFromCEP func(cep string) (string, error)
	GetTemperature func(city string) (float64, error)
}

var (
	// Serviço padrão com as implementações reais das funções
	defaultService = WeatherService{
		GetCityFromCEP: getCityFromCEP,
		GetTemperature: getTemperature,
	}
)

// TestCEPValido testa o cenário de um CEP válido e encontrado.
func TestCEPValido(t *testing.T) {
	// Mock da função GetCityFromCEP
	mockService := WeatherService{
		GetCityFromCEP: func(cep string) (string, error) {
			return "São Paulo", nil
		},
		GetTemperature: func(city string) (float64, error) {
			return 25.0, nil // Temperatura simulada
		},
	}

	// Substitui o serviço padrão pelo mock
	originalService := defaultService
	defaultService = mockService
	defer func() { defaultService = originalService }()

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handleWeatherRequest))
	defer server.Close()

	// Faz uma requisição GET para o endpoint /weather com um CEP válido
	resp, err := http.Get(server.URL + "/weather?cep=01001000")
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code esperado: %d, obtido: %d", http.StatusOK, resp.StatusCode)
	}

	// Decodifica a resposta JSON
	var result TemperatureResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Erro ao decodificar a resposta JSON: %v", err)
	}
}

// TestCEPInvalido testa o cenário de um CEP inválido.
func TestCEPInvalido(t *testing.T) {
	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handleWeatherRequest))
	defer server.Close()

	// Faz uma requisição GET com um CEP inválido
	resp, err := http.Get(server.URL + "/weather?cep=123")
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o status code
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Status code esperado: %d, obtido: %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}

	// Verifica a mensagem de erro
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}
	expectedMessage := "invalid zipcode"
	if !bytes.Contains(body, []byte(expectedMessage)) {
		t.Errorf("Mensagem esperada: %s, obtida: %s", expectedMessage, string(body))
	}
}

// TestCEPNotFound testa o cenário de um CEP não encontrado.
func TestCEPNotFound(t *testing.T) {
	// Mock da função GetCityFromCEP para simular CEP não encontrado
	mockService := WeatherService{
		GetCityFromCEP: func(cep string) (string, error) {
			return "", fmt.Errorf("city not found")
		},
	}

	// Substitui o serviço padrão pelo mock
	originalService := defaultService
	defaultService = mockService
	defer func() { defaultService = originalService }()

	// Cria um servidor de teste
	server := httptest.NewServer(http.HandlerFunc(handleWeatherRequest))
	defer server.Close()

	// Faz uma requisição GET com um CEP que não existe
	resp, err := http.Get(server.URL + "/weather?cep=99999999")
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o status code
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Status code esperado: %d, obtido: %d", http.StatusNotFound, resp.StatusCode)
	}

	// Verifica a mensagem de erro
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}
	expectedMessage := "can not find zipcode"
	if !bytes.Contains(body, []byte(expectedMessage)) {
		t.Errorf("Mensagem esperada: %s, obtida: %s", expectedMessage, string(body))
	}
}
