# Library Backend API

## Overview
Backend untuk pengelolaan perpustakaan, dengan fitur login, CRUD buku, dan peminjaman.

## Cara Menjalankan
1. **Local**: Jalankan `go run main.go`.
2. **Docker**: 
   ```bash
   docker build -t library-api .
   docker run -p 8080:8080 library-api
