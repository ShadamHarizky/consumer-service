# consumer-service - Technical Test PT.EDOT

## Persyaratan

Sebelum menjalankan service, pastikan point2 di bawah terpenuhi:

- [Go](https://golang.org/) (Golang) - Versi 1.18 atau lebih 
- [RabbitMQ](https://www.rabbitmq.com/)
- [Redis](https://redis.com/)
- [MongoDB](https://www.mongodb.com/)

## Memulai

1. clone repositori:

    ```bash
    git clone https://github.com/ShadamHarizky/consumer-service
    cd consumer-service
    ```

2. install dependensi:

    ```bash
    go mod init github.com/ShadamHarizky/consumer-service
    go mod tidy
    ```

3. Collection MongoDB

    - buat collection di mongodb dengan nama message_consumer

4. Buatkan .env sesuai dengan .env-example:

    - Rubah detail koneksi RabbitMQ, MongoDB dan Redis dalam file env.
    - Pastikan RabbitMQ diinstal dan berjalan.
    - Pastikan Redis diinstal dan berjalan.
    - Pastikan Mongodb diinstal dan berjalan.

5. Jalankan Project:

    ```bash
    go run .
    ```

## Konfigurasi

Update env pada file `.env` dan sesuaikan dengan konfigurasi pada komputer masing2:
- `RABBITMQ_URL`: URL server RabbitMQ (default: `amqp://guest:guest@localhost:5672/`)
- `REDIS_ADDRESS`: URL server Redis (default: `localhost:6379`)
- `MONGODB_URL`: URL Connection MongoDB (default: `mongodb://localhost:27017`)
- `DB_NAME`: Database Name (default: `technical-test`)
- `REDIS_CHANNEL`: Redis key pub/sub atau channel (default: `testing`)
- `RABBITMQ_QUEUE_NAME`: Queue Name RabbitMQ (default: `your-queue-name`)
