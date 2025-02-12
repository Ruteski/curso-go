# Projeto de Exemplo de comunição entre Serviço A e Serviço B e Observabilidade com Zipkin e OpenTelemetry

Este projeto consiste em dois serviços (`serviceA` e `serviceB`) que se comunicam entre si e utilizam o Zipkin para tracing distribuído. O `serviceA` recebe requisições com um CEP e encaminha para o `serviceB`, que consulta a temperatura da cidade correspondente ao CEP.

## Pré-requisitos

Antes de rodar o projeto, certifique-se de que você tem os seguintes requisitos instalados:

- **Docker**: [Instale o Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Instale o Docker Compose](https://docs.docker.com/compose/install/)
- **Git** (opcional, para clonar o repositório): [Instale o Git](https://git-scm.com/)

## Estrutura do Projeto

O projeto está organizado da seguinte forma:
├── serviceA
│ ├── Dockerfile
│ ├── go.mod
│ ├── go.sum
│ ├── main.go
│ └── ...
├── serviceB
│ ├── Dockerfile
│ ├── go.mod
│ ├── go.sum
│ ├── main.go
│ └── ...
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md


## Como Rodar o Projeto

Siga os passos abaixo para rodar o projeto em um ambiente de desenvolvimento.

### 1. Clone o Repositório (se necessário)

Se você ainda não clonou o repositório, faça isso agora:

```bash
git clone https://github.com/Ruteski/curso-go.git
cd seu-repositorio
```

### 2. Construa e Suba os Contêineres
Use o Docker Compose para construir e subir os contêineres:
```bash
docker-compose up --build
ou
docker-compose up -d --build
```

Isso irá:
- Construir as imagens para serviceA e serviceB.
- Subir os serviços serviceA, serviceB e o Zipkin.

### 3. Acesse os Serviços
- Serviço A: Estará disponível em http://localhost:8000.
- Serviço B: Estará disponível em http://localhost:8001.
- Zipkin: Estará disponível em http://localhost:9411 para visualizar os traces.

### 4. Teste a Aplicação
Envie uma requisição POST para o serviceA usando o comando curl:

```bash
curl -X POST http://localhost:8000/cep -d '{"cep": "01001000"}' -H "Content-Type: application/json"
```

Isso deve retornar a temperatura da cidade correspondente ao CEP fornecido.

### 5. Visualize os Traces no Zipkin
Acesse o Zipkin em http://localhost:9411 para visualizar os traces gerados pelos serviços.

### 6. Parar os Serviços
Para parar os serviços, execute:
```bash
docker-compose down
```
Isso encerrará todos os contêineres e removerá a rede criada pelo Docker Compose.

## Contribuição
Se você quiser contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.