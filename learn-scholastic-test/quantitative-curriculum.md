# M4/W4/D30 - Thu, 30 Apr 2026 (WIB)

## Scholastic Kuantitatif / Numerik - Curriculum

Kurikulum singkat untuk membangun fondasi kuantitatif Tes Skolastik LPDP:
mulai dari pola bilangan, operasi dasar, pecahan, persentase, rasio, aljabar,
sampai soal cerita. Setiap topik tidak dipetakan ke fungsi seperti folder Go
lain, tetapi dipetakan ke soal latihan di `quantitative-questions.md`.

Target audience: peserta yang ingin mulai dari basic banget, lalu naik ke pola
numerik yang sering muncul di Tes Bakat Skolastik.

---

## Curriculum Table

| # | Topic | Core Concept | LPDP Usage |
| --- | --- | --- | --- |
| 01 | Deret Angka | Pola bilangan dari selisih, rasio, atau operasi berulang | Menebak angka berikutnya dengan pola yang paling konsisten |
| 02 | Aritmetika Dasar | Urutan operasi tambah, kurang, kali, bagi | Menghitung cepat tanpa salah prioritas operasi |
| 03 | Pecahan & Desimal | Konversi dan operasi pecahan/desimal | Menyamakan bentuk angka agar hitungan lebih ringan |
| 04 | Persentase | Diskon, kenaikan, penurunan, dan faktor pengali | Menghindari jebakan persen beruntun |
| 05 | Perbandingan | Rasio, proporsi, dan skala | Mengubah bagian rasio menjadi nilai nyata |
| 06 | Aljabar Sederhana | Persamaan satu/dua variabel | Membuat model dari informasi numerik |
| 07 | Soal Cerita | Umur, jarak-waktu-kecepatan, untung-rugi, pekerjaan | Menerjemahkan cerita menjadi persamaan sederhana |

---

## Detailed Topic Offset

Detail penjelasan per nomor saat ini sudah dibuat sampai:

```text
01 -> 03
```

Nomor berikutnya yang belum dibedah detail adalah `04 - Persentase`.

---

## Num01 - Deret Angka

### The Problem

Soal deret angka terlihat seperti tebak-tebakan, padahal yang diuji adalah
kemampuan menemukan pola paling sederhana dan paling konsisten. Kesalahan umum:
melihat satu selisih cocok, lalu langsung memilih jawaban tanpa mengecek seluruh
barisan.

Contoh pola dasar:

```text
3, 7, 15, 31, 63, ...
```

Setiap angka dikali 2 lalu tambah 1:

```text
3 x 2 + 1 = 7
7 x 2 + 1 = 15
15 x 2 + 1 = 31
31 x 2 + 1 = 63
```

Maka angka berikutnya adalah:

```text
63 x 2 + 1 = 127
```

### The Concept

Deret angka bisa memakai beberapa jenis pola. Jangan hafal satu pola saja.

| Pola | Bentuk Cepat | Contoh |
| --- | --- | --- |
| Selisih tetap | tambah angka yang sama | 4, 9, 14, 19 |
| Rasio tetap | kali angka yang sama | 3, 6, 12, 24 |
| Selisih bertingkat | selisihnya naik/turun berpola | 2, 5, 10, 17 |
| Operasi berulang | kali lalu tambah/kurang | 3, 7, 15, 31 |
| Pola selang-seling | posisi ganjil dan genap punya pola berbeda | 2, 10, 4, 20, 8 |

Langkah tetap:

1. Cek selisih antar angka.
2. Kalau selisih tidak stabil, cek selisih tingkat dua.
3. Kalau masih tidak stabil, cek rasio atau operasi `x n + k`.
4. Cek apakah pola berlaku ke semua angka, bukan hanya dua angka terakhir.

### What the practice shows

Soal terkait: nomor `01 - Deret Angka` di `quantitative-questions.md`.

Nomor 01 melatih pola operasi berulang `x 2 + 1`. Fokusnya bukan cuma menemukan
jawaban, tapi membuktikan pola berlaku dari awal sampai akhir.

### Common Traps

| Jebakan | Contoh |
| --- | --- |
| Hanya lihat angka terakhir | 63 ke opsi terdekat |
| Salah pola selisih | mengira selisihnya tambah tetap |
| Tidak verifikasi | pola cocok di dua langkah, tapi gagal di langkah lain |

### Review Rule

Untuk setiap deret, wajib tulis pola dalam satu baris:

```text
Pola: setiap angka ...
```

Kalau pola tidak bisa ditulis jelas, jangan pilih jawaban dulu.

---

## Num02 - Aritmetika Dasar

### The Problem

Soal aritmetika dasar sering salah bukan karena konsepnya sulit, tetapi karena
urutan operasi keliru atau hitungan mental terlalu cepat. Di Tes Skolastik,
angka sering dibuat agar bisa disederhanakan kalau urutannya benar.

Contoh:

```text
48 x 125 / 24 + 36 x 25
```

Jangan hitung dari kiri secara membabi buta. Sederhanakan bagian yang mudah:

```text
48 / 24 = 2
2 x 125 = 250
36 x 25 = 900
250 + 900 = 1.150
```

### The Concept

Urutan operasi:

```text
kurung -> kali/bagi -> tambah/kurang
```

Untuk kali dan bagi yang sejajar, kerjakan dari kiri ke kanan atau sederhanakan
faktor yang jelas aman.

Pola cepat:

```text
pecah ekspresi -> sederhanakan faktor -> hitung akhir
```

Trik angka yang sering membantu:

| Bentuk | Nilai |
| --- | --- |
| 25 x 4 | 100 |
| 125 x 8 | 1.000 |
| 0,5 | 1/2 |
| 0,25 | 1/4 |
| 0,75 | 3/4 |

### What the practice shows

Soal terkait: nomor `02 - Aritmetika Dasar` di `quantitative-questions.md`.

Nomor 02 melatih operasi campuran kali, bagi, dan tambah. Kuncinya adalah
menyederhanakan `48 / 24` sebelum mengalikan dengan `125`.

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Menjumlah dulu | Tambah/kurang kalah prioritas dari kali/bagi |
| Tidak menyederhanakan | Hitungan jadi besar dan rawan salah |
| Salah baca ribuan | 1.150 bisa terbaca 115 kalau terburu-buru |

### Review Rule

Setelah mengerjakan, cek ulang ekspresi dengan tanda kurung mental:

```text
(bagian kali/bagi) + (bagian kali/bagi)
```

Kalau struktur ekspresinya belum jelas, jangan hitung final dulu.

---

## Num03 - Pecahan & Desimal

### The Problem

Pecahan dan desimal sering menjebak karena bentuk angkanya campur. Peserta
sering mencoba menghitung semuanya sebagai desimal panjang, padahal beberapa
angka punya bentuk pecahan yang mudah.

Contoh:

```text
3/4 + 0,625 - 5/8 + 0,2
```

Karena `0,625 = 5/8`, bagian `0,625 - 5/8` saling menghapus.

Sisa:

```text
3/4 + 0,2 = 0,75 + 0,2 = 0,95 = 19/20
```

### The Concept

Kunci pecahan dan desimal adalah menyamakan bentuk sebelum operasi.

Konversi penting:

| Desimal | Pecahan |
| --- | --- |
| 0,1 | 1/10 |
| 0,2 | 1/5 |
| 0,25 | 1/4 |
| 0,375 | 3/8 |
| 0,5 | 1/2 |
| 0,625 | 5/8 |
| 0,75 | 3/4 |

Langkah tetap:

1. Cari desimal yang punya pecahan mudah.
2. Samakan bentuk angka.
3. Sederhanakan bagian yang saling menghapus.
4. Baru hitung hasil akhir.

### What the practice shows

Soal terkait: nomor `03 - Pecahan & Desimal` di `quantitative-questions.md`.

Nomor 03 melatih konversi `0,625 = 5/8` dan `0,2 = 1/5`. Tujuannya agar kamu tidak
terjebak menghitung desimal panjang ketika ada pembatalan sederhana.

### Common Traps

| Jebakan | Contoh |
| --- | --- |
| Salah konversi | 0,625 dianggap 6/25 |
| Semua dibuat desimal | boleh, tapi rawan salah koma |
| Tidak lihat pembatalan | 0,625 dan 5/8 sebenarnya sama |

### Review Rule

Sebelum menghitung, tulis bentuk yang kamu pilih:

```text
Mau dihitung sebagai pecahan atau desimal?
```

Campur bentuk hanya boleh kalau konversinya sangat jelas.

---

## Num04-Num07 - Persen, Rasio, Aljabar, dan Soal Cerita

### The Problem

Topik nomor 04-07 biasanya lebih dekat dengan soal cerita. Angka tidak lagi berdiri
sendiri; kamu harus menerjemahkan kalimat menjadi operasi. Jebakan utamanya:
menghitung terlalu cepat sebelum tahu "nilai awal", "nilai akhir", dan relasi
antar variabel.

### The Concept

Untuk nomor 04-07, fokusnya naik dari hitungan tunggal ke model sederhana:

1. Persentase dibaca sebagai faktor pengali.
2. Rasio dibaca sebagai bagian.
3. Aljabar dipakai untuk menyimpan nilai yang belum diketahui.
4. Soal cerita diterjemahkan menjadi jarak, waktu, kecepatan, harga, umur, atau pekerjaan.

Pola persentase minimum:

```text
naik p% -> kali (1 + p/100)
turun p% -> kali (1 - p/100)
```

Pola rasio minimum:

```text
A : B : C = 2 : 5 : 8
nilai A = 2x
nilai B = 5x
nilai C = 8x
```

Varian campuran replacement:

```text
Total wadah = T
Rasio awal X : Y = 5 : 3
X awal = 5/8 T
Y awal = 3/8 T
```

Kalau sebagian campuran diambil, komposisi yang keluar mengikuti rasio campuran
saat itu. Setelah diganti bahan murni, hanya bahan yang ditambahkan yang berubah
sesuai penggantinya.

Pola cepat:

```text
ambil campuran -> kurangi X dan Y sesuai rasio awal
ganti Y murni -> tambahkan hanya ke Y
bentuk rasio akhir -> selesaikan T
```

### What the practice shows

1. Nomor 04 melatih persentase beruntun.
2. Nomor 05 melatih rasio menjadi nilai nyata.
3. Nomor 06 melatih persamaan dua variabel.
4. Nomor 07 melatih soal cerita jarak-waktu-kecepatan.

Latihan tambahan nomor 05 juga melatih pembagian keuntungan proporsional dan
campuran replacement, yaitu saat sebagian campuran diambil lalu diganti bahan
murni.

### Review rule

Untuk setiap soal nomor 04-07, tulis dulu:

```text
Yang diketahui:
Yang ditanya:
Model:
```

Kalau model belum ada, jangan buru-buru hitung.

---

## Grind Plan

| Urutan | Fokus | Target | Output |
| --- | --- | --- | --- |
| 01-03 | Pola dan operasi angka | Nomor 01-03 | Catatan pola, urutan operasi, dan konversi |
| 04-07 | Model numerik | Nomor 04-07 | Catatan persen, rasio, persamaan, dan cerita |

Aturan waktu:

- Nomor 01-03: 60-120 detik per soal.
- Nomor 04-07: 2-4 menit per soal.
- Review wajib mencatat letak salah: salah pola, salah operasi, salah model, atau salah hitung.
