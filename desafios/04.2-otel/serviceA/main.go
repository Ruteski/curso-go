package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/trace"
)

type CEPRequest struct {
	CEP string `json:"cep"`
}

func initTracer() {
	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		fmt.Println("Erro ao configurar o Zipkin exporter:", err)
		return
	}

	// Log para verificar se o exporter foi criado
	fmt.Println("Zipkin exporter configurado com sucesso")

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)

	// Log para verificar se o tracer provider foi configurado
	fmt.Println("Tracer provider configurado com sucesso")
}

func main() {
	initTracer() // Inicializa o tracer do OpenTelemetry

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/cep", handleCEPRequest)
	fmt.Printf("Serviço A rodando na porta %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handleCEPRequest(w http.ResponseWriter, r *http.Request) {
	// Cria um span para rastrear a requisição
	_, span := otel.Tracer("servicoA").Start(r.Context(), "handleCEPRequest")
	defer span.End()

	// Log para verificar se o span foi criado
	fmt.Println("Span criado para handleCEPRequest")

	// Verifica se o método da requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Valida o CEP
	var cepRequest CEPRequest
	err = json.Unmarshal(body, &cepRequest)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	cep := cepRequest.CEP
	if len(cep) != 8 || !isNumeric(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Encaminha o CEP para o Serviço B
	resp, err := http.Post("http://localhost:8001/weather", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Erro ao encaminhar para o Serviço B", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Retorna a resposta do Serviço B
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
