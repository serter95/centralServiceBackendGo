FROM golang:latest
# FROM golang:alpine AS build

# RUN mkdir /home/centralServiceBackendGo
RUN mkdir /go/src/centralServiceBackendGo

# WORKDIR /home/centralServiceBackendGo/
WORKDIR /go/src/centralServiceBackendGo/

# COPY centralServiceBackendGo /home/centralServiceBackendGo/
COPY centralServiceBackendGo /go/src/centralServiceBackendGo/

EXPOSE 3000

CMD [ "./centralServiceBackendGo" ]