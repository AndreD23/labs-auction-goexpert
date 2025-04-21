# # Sistema de Leilões

## 🚀 Como Iniciar o Projeto

### Pré-requisitos
- Docker
- Docker Compose
- Git

---

### 1. Configuração do Ambiente

1. Clone o repositório:
```bash
git clone <url-do-repositório>
cd fullcycle-auction_go
```

2. Crie o arquivo de configuração `.env`:
```bash
cp .env.example .env
```
---

### 2. Executando a Aplicação

1. Inicie os containers com Docker Compose:
```bash
docker-compose up -d --build
```
2. Verifique se os containers estão rodando:

---

## 📡 Testando as Rotas

### Criar um Leilão (POST)
```bash
curl -X POST http://localhost:8080/auction \
-H "Content-Type: application/json" \
-d '{
    "product_name": "iPhone 15",
    "category": "Eletrônicos",
    "description": "iPhone 15 Pro Max 256GB",
    "condition": 0
}'
```

### Buscar Leilões (GET)
```bash
curl http://localhost:8080/auction
```

### Buscar Leilão por ID (GET)
```bash
curl http://localhost:8080/auction/{auctionId}
```

### Criar um Lance (POST)
```bash
curl -X POST http://localhost:8080/bid \
-H "Content-Type: application/json" \
-d '{
    "user_id": "1",
    "auction_id": "123",
    "amount": 50000.00
}'
```

### Buscar Lances de um Leilão (GET)
```bash
curl http://localhost:8080/bid/{auctionId}
```

---

## 🗄️ Verificando Dados no MongoDB

1. Acesse o container do MongoDB:
```bash
docker exec -it mongodb mongosh
```

2. Comandos úteis:
```
use auction
db.auctions.find()
db.bids.find()
```

---

## 🧪 Executando Testes

1. Execute os testes
```
go test ./... -v
```
