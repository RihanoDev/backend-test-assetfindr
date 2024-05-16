# Backend Test - AssetFindr

## Nama: Rizky Haffiyan Roseno

## Project Test Post API

### Deskripsi
API untuk mengelola postingan dan tag dengan menggunakan Go, Gin, GORM, dan PostgreSQL.

### Instalasi

1. **Clone repository**:
    ```bash
    git clone https://github.com/RihanoDev/backend-test-assetfindr.git
    cd backend-test-assetfindr/post-api
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Konfigurasi Environment**:
    Buat file `.env` berdasarkan template berikut dan isi dengan konfigurasi yang dimiliki.
    ```
    DB_USERNAME=your_db_username
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    ```

4. **Setup Database**:
    Pastikan database PostgreSQL berjalan dan sesuai dengan konfigurasi di `.env`.

5. **Run the Application**:
    ```bash
    go run main.go
    ```
    atau dapat menggunakan
   
    ```bash
    air
    ```
