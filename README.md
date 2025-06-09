
📦 Backend Summarizer

Project ini menyediakan layanan untuk meringkas teks panjang dan mengelola data ringkasan melalui GraphQL API serta dikemas dengan Docker Compose untuk deployment yang mudah.

🚀 Quickstart

Clone repository:
```
git clone https://github.com/username/backend-summarizer.git
cd backend-summarizer
```
Pastikan Docker & Docker Compose terinstal.

Jalankan semua service:
```
docker-compose up --build
```
Akses GraphQL Playground di http://localhost:8080/

🗂 Layanan di Docker Compose
```
version: "3.8"
services:
  db:
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 123
      POSTGRES_DB: summary_db
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  api:
    build: .
    env_file:
      - .env
    environment:
      POSTGRES_DSN: "postgresql://postgres:123@db:5432/summary_db?sslmode=disable"
      ML_SERVICE_URL: "http://ml:5000/generate"
    ports:
      - "8080:8080"
    depends_on:
      - db
      - ml

  ml:
    build:
      context: ./ml_dummy
      dockerfile: Dockerfile
    ports:
      - "5000:5000"

volumes:
  pgdata:
```
db: PostgreSQL menyimpan tabel summaries (otomatis membuat DB summary_db).

api: Service Go (Gin + gqlgen) untuk GraphQL endpoint.

ml: Dummy Flask service yang mensimulasikan endpoint ML.

🔧 Konfigurasi Environment

Buat file .env di root project:
```
POSTGRES_DSN=postgresql://postgres:123@db:5432/summary_db?sslmode=disable
ML_SERVICE_URL=http://ml:5000/generate
```
Docker Compose akan memuat env ini untuk service api.

📄 GraphQL Schema & Endpoints
```
schema {
  query: Query
  mutation: Mutation
}

type Summary {
  id: ID!
  input: String!
  summary: String!
  createdAt: String!
}

type Query {
  # Ambil semua ringkasan
  summaries: [Summary!]!
  # Ambil ringkasan berdasarkan ID
  summaryByID(id: ID!): Summary
}

type Mutation {
  # Buat ringkasan baru dari teks input
  createSummary(input: String!): Summary!
}
```
Endpoint utama: http://localhost:8080/query

Buka GraphQL Playground untuk langsung mencoba query & mutation.

Contoh Query: Ambil Semua Ringkasan
```
query {
  summaries {
    id
    input
    summary
    createdAt
  }
}
```
Contoh Mutation: Buat Ringkasan Baru
```
mutation {
  createSummary(input: "Teks panjang untuk diringkas...") {
    id
    input
    summary
    createdAt
  }
}
```
