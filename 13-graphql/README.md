# site do programa para ajudar a criacao do projeto em graphql, seguir os passos indicados
https://gqlgen.com

# inicia um projeto utlizando o gqlgen
go run github.com/99designs/gqlgen init

# cria o esqueleto do graphql com base no meu schema alterado
go run github.com/99designs/gqlgen generate

# injetar meu usecase para o graphql
## - dentro da pasta graph no arquivo resolver.go graph/resolver.go