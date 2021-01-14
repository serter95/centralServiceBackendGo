# creo la imagen que servira para compilar el proyecto
FROM golang:latest AS central_service_builder
# creo la carpeta que tendra el proyecto a compilar
RUN mkdir /go/src/centralServiceBackendGo
# seteo las variales de entorno
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# me posiciono en el directorio de trabajo
WORKDIR /go/src/centralServiceBackendGo/
# copio todos mis archivos a la carpeta de compilacion
COPY * /go/src/centralServiceBackendGo/
# ejecuto la instalacion de paquetes
RUN go install
# compilo la app para tener los binarios
RUN go build

# tomo la version alpine:3.9.4 de go que es mas ligera
FROM alpine:3.9.4
# Algunas configuraciones de ambiente
ENV GOROOT /usr/local/go
# Creo la carpeta
RUN mkdir -p /go/src/centralServiceBackendGo
# Copio de la imagen anterior el compilado
COPY --from=central_service_builder /go/src/centralServiceBackendGo/centralServiceBackendGo /go/src/centralServiceBackendGo
# Expongo el puerto 3000
EXPOSE 3000
# Le indico cual sera el directorio de trabajo
WORKDIR /go/src/centralServiceBackendGo/
# Al instanciar la imagen lo que se ejecutara es el comando
CMD [ "./centralServiceBackendGo" ]