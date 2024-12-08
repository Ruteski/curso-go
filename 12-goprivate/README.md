# traz a variavel onde informa o local dos repos privados
go env | grep PRIVATE

# informa o repositorio que é privado, para adicionar mais repositorios privados, separar eles por ","
export GOPRIVATE=github.com/ruteski/fcutils-secret,github.com/ruteski/fcutils-secret

# Autenticacao via http

## configurar credenciais de acesso ao github
vim ~/netrc

## adiconar isso
machine github.com
login Ruteski
password <gerar token no github>

# Autenticacao via ssh
vim ~/.gitconfig
[url "ssh://git@github.com/"]
        instedOf https://github.com/