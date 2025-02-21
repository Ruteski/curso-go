# API de Consulta de Temperatura por CEP

Esta é uma API em Go que recebe um CEP válido via query parameter, consulta a localização correspondente e retorna as temperaturas atuais em Celsius, Fahrenheit e Kelvin.

---

## Como Executar a Aplicação

### Requisitos

- Docker
- Docker Compose

## Passo 1: Construir e Rodar a Aplicação com Docker Compose
### Use o Docker Compose para construir a imagem e rodar a aplicação:
```bash
docker-compose up --build app
```
### A aplicação estará disponível em http://localhost:8000

## Passo 2: Testar a API
### Você pode testar a API usando curl, Postman ou diretamente no navegador.
### Exemplo de Requisição:
```bash
curl "http://localhost:8000/weather?cep=01001000"
```

### Resposta Esperada:
```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.15
}
```

## Passo 3: Rodar os Testes Automatizados
### Para rodar os testes automatizados, execute:
```bash
docker-compose up tests
``` 
### Isso vai executar todos os testes e exibir o resultado no terminal.

## Passo 4: Parar os Contêineres
### Para parar os contêineres, execute:
```bash
docker-compose down
```

## Passo 5: Deploy Google Cloud

Este é o passo a passo para criar uma conta no Google Cloud e realizar o deploy do projeto.

### 5.1: Criar uma Conta no Google Cloud

1. Acesse o [Google Cloud Console](https://console.cloud.google.com/).

2. Se você não tiver uma conta, clique em **Criar Conta**.

3. Complete o cadastro com seu e-mail e informações solicitadas.

4. Após a criação, você terá acesso a um **Crédito de $300** por 90 dias, o que é suficiente para completar este desafio.

### 5.2: Ativar o Faturamento no Google Cloud

1. No Console do Google Cloud, vá para **Faturamento**.

2. Clique em **Ativar faturamento** e siga as instruções para configurar a conta de faturamento.