# Berkontribusi Untuk Negeri

Terima kasih atas ketertarikan Anda untuk berkontribusi pada pengembangan **Neurasita**. Kami menyambut segala bentuk kontribusi, mulai dari penulisan kode, dokumentasi, pelaporan bug, pengelolaan dataset, hingga ide pengembangan.

## Jenis Kontribusi yang Dapat Diberikan

- Penambahan atau perbaikan fitur
- Perbaikan bug dan peningkatan stabilitas
- Penambahan atau pengelolaan dataset
- Dokumentasi dan tutorial
- Optimalisasi performa sistem
- Pengujian dan validasi model
- Donasi untuk keberlangsungan server

## Teknologi yang Digunakan

- Backend: Go (Golang), PostgreSQL
- Frontend: Vanilla JavaScript
- DevOps: Docker, GitHub Actions

## Panduan Kontribusi

### 1. Fork dan Clone Repositori

Silakan fork repositori ini, lalu clone ke komputer lokal Anda.

```bash
git clone https://github.com/nurfianqodar/neurasita.git
cd neurasita
```

### 2. Buat Branch Fitur

Gunakan nama branch yang deskriptif, misalnya:

```bash
git checkout -b fitur/upload-dataset-api
```

### 3. Lakukan Perubahan

- Ikuti standar penulisan kode masing-masing bahasa
- Tambahkan dokumentasi bila perlu
- Jalankan pengujian sebelum mengajukan pull request

### 4. Commit Perubahan

Gunakan pesan commit yang singkat dan jelas:

```bash
git commit -m "feat: implementasi endpoint unggah dataset"
```

### 5. Push dan Ajukan Pull Request

```bash
git push origin fitur/upload-dataset-api
```

Ajukan pull request ke branch `main` disertai deskripsi yang menjelaskan perubahan Anda.

## Gaya Penulisan Kode

- Go: gunakan `gofmt` dan struktur idiomatik
- JavaScript/TypeScript: gunakan `eslint` dan `prettier` bila tersedia
- Gunakan nama variabel dan fungsi yang bermakna (duck typing)
- Selalu sertakan komentar pada deklarasi fungsi atau variable global

## Pengujian

Pastikan semua pengujian dapat dijalankan dan lolos:

```bash
# Untuk backend Go
go test ./...

# Untuk frontend
npm run test
```

## Etika Berkontribusi

Kami menjunjung tinggi sikap terbuka, saling menghargai, dan kolaboratif. Harap menjaga komunikasi yang baik dan profesional dalam setiap interaksi. Neurasita mengikuti pedoman perilaku [Contributor Covenant](https://www.contributor-covenant.org/).

## Bantuan

Jika mengalami kendala teknis atau memiliki pertanyaan sebelum berkontribusi, Anda dapat membuka _issue_ dengan label `pertanyaan` atau menghubungi tim pengembang melalui email: **[nurfianqodar@gmail.com](mailto:nurfianqodar@gmail.com)**

Terima kasih telah mendukung dan membangun Neurasita bersama kami.
