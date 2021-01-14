# centralServiceBackendGo
API to consume three external search services 

1. itunes
2. TV maze
3. Crcind

Central Service API

# Install using local
```
git clone https://github.com/serter95/centralServiceBackendGo.git
cd centralServiceBackendGo
go install
go run main.go
```

# Install using docker
```
git clone https://github.com/serter95/centralServiceBackendGo.git
cd centralServiceBackendGo
docker pull golang:latest
docker build -t central_service_backend:latest .
docker run -p 3000:3000 central_service_backend:latest
```

# Documentation
http://localhost:3000/swagger/index.html