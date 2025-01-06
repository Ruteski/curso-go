# Sempre instalar o compilador do protobuf
## https://grpc.io/docs/protoc-installation/


# Instalar plugins para o Go
## https://grpc.io/docs/languages/go/quickstart/



# Após ter o protofile pronto rodar comando
## - protoc --go_out=. --go-grpc_out=. proto/course_category.proto
## - go mod tidy


# Instalar o evans grpc client, para trabalhar chamadas grpc
# https://github.com/ktr0731/evans
### go install github.com/ktr0731/evans@latest
### rodar o evans(tem que estar na porta 50051) -> evans -r repl
### sempre verificar onde esta o serviço >> "pb.OrderService@127.0.0.1:50051>" tem que ter o pg.<service>@