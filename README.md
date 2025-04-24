# Account Service

## Project Overview
Account Service using `echo` (`Golang`) with `PostgreSQL`.

## Prerequisites
- Docker
- Docker Compose
- Golang
- PostgreSQL

## Getting Started

### Running the Application
Clone the repository:
```bash
git clone https://github.com/dedebadru/go-account-service.git
cd go-account-service
```

### Using Docker Compose
Build and start the application:
```bash
docker-compose up --build
```

Access the application:
- Account Service: `http://localhost:8080`

### Stopping the Application
```bash
docker-compose down
```

## Technologies
- Backend: Golang, echo
- Database: PostgreSQL
- Containerization: Docker, Docker Compose

## API
### Account Registration
```bash
curl --request POST \
  --url http://localhost:8080/daftar \
  --header 'Content-Type: application/json' \
  --data '{
	"nama": "Fulan",
	"nik": "1234567890123452",
	"no_hp": "083456781234"
}'
```

### Credit Balance
```bash
curl --request POST \
  --url http://localhost:8080/tabung \
  --header 'Content-Type: application/json' \
  --data '{
	"no_rekening": "100156231795",
	"nominal": 500000
}'
```

### Debit Balance
```bash
curl --request POST \
  --url http://localhost:8080/tarik \
  --header 'Content-Type: application/json' \
  --data '{
	"no_rekening": "100156231795",
	"nominal": 300000
}'
```

### Get Balance
```bash
curl --request GET \
  --url http://localhost:8080/saldo/100156231795
```