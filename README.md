<!-- @format -->

# **REST API MUSIC**

Repository ini untuk test challenge **Backend Engineer Intern DOT Indonesia**.

---

**APP**

![app](https://i.pinimg.com/736x/77/a6/1a/77a61a0041a774d6754a0d8a1251bc1f.jpg)

**ERD Schema**
(user_music_favorite adalah tabel otomatis by gorm)

![erd](dot.png)

**Layering pattern**

![flow](https://i.pinimg.com/736x/c8/1d/4b/c81d4b034203100c3e9d13c8ca3d4d80.jpg)

Modular → mudah dibaca & di-maintain

Testable → Bisa test service tanpa nyentuh database (mock test).

Separation of Concerns → tiap layer fokus ke tugasnya

[Lihat Flow Chart](https://miro.com/app/board/uXjVIx8lumg=/?moveToWidget=3458764629378649791&cot=14)

[Lihat API Docs](https://app.swaggerhub.com/apis/CLAIREGLOWUP/music-api-dot/1.0.0#/User/delete_user_favorite)

## **Fitur**

- Authentication menggunakan JWT.
- User melihat list musik.
- User menambahkan musik favorit.
- User melihat list musik favoritnya.
- User menghapus musik yang sudah difavoritkan.

---

## **Teknologi yang Digunakan**

- **Golang**
- **GORM**
- **Echo Go Framework**
- **PostgreSQL**
- **Docker**
- **Docker Compose**

---

## **📦 Cara Menjalankan Proyek**

### **1. Jalankan dengan Docker Compose**

1. Clone repository ini:
   ```bash
   git clone https://github.com/claireglowup/dot-go.git
   cd dot-go
   ```
2. **UBAH FILE .ENV.EXAMPLE JADI .ENV**

3. Bangun dan jalankan container:
   ```bash
   docker-compose up --build
   ```

### **2. Alternatif: Jalankan dengan Makefile**

Jika Makefile tersedia, jalankan:

```bash
make start
```

## **📜 Lisensi**

Proyek ini menggunakan lisensi [MIT License](LICENSE).

---
