# M4/W4/D32 - Sat, 02 May 2026 (WIB)

## Scholastic Kuantitatif / Numerik - Curriculum

Kurikulum kuantitatif Tes Skolastik LPDP untuk membangun kemampuan membaca
pola angka, operasi dasar, pecahan, persentase, rasio, aljabar, dan soal cerita.
File ini dibuat sebagai satu tempat belajar: setiap nomor berisi konsep, pola
soal, pola menjawab, jebakan umum, aturan review, dan latihan dari
`quantitative-questions.md`.

Target audience: peserta yang ingin mulai dari basic banget, lalu naik ke pola
numerik yang sering muncul di Tes Bakat Skolastik.

---

## Curriculum Table

| #   | Topic             | Core Concept                                             | LPDP Usage                                                 |
| --- | ----------------- | -------------------------------------------------------- | ---------------------------------------------------------- |
| 01  | Deret Angka       | Pola bilangan dari selisih, rasio, atau operasi berulang | Menebak angka berikutnya dengan pola yang paling konsisten |
| 02  | Aritmetika Dasar  | Urutan operasi tambah, kurang, kali, bagi                | Menghitung cepat tanpa salah prioritas operasi             |
| 03  | Pecahan & Desimal | Konversi dan operasi pecahan/desimal                     | Menyamakan bentuk angka agar hitungan lebih ringan         |
| 04  | Persentase        | Diskon, kenaikan, penurunan, dan faktor pengali          | Menghindari jebakan persen beruntun                        |
| 05  | Perbandingan      | Rasio, proporsi, skala, dan campuran                     | Mengubah bagian rasio menjadi nilai nyata                  |
| 06  | Aljabar Sederhana | Persamaan satu/dua variabel                              | Membuat model dari informasi numerik                       |
| 07  | Soal Cerita       | Umur, jarak-waktu-kecepatan, untung-rugi, pekerjaan      | Menerjemahkan cerita menjadi persamaan sederhana           |

---

## Grind Plan

| Urutan | Fokus                  | Target      | Output                                                 |
| ------ | ---------------------- | ----------- | ------------------------------------------------------ |
| 01-03  | Pola dan operasi angka | Nomor 01-03 | Catatan pola, urutan operasi, dan konversi             |
| 04-05  | Persen dan rasio       | Nomor 04-05 | Catatan basis persen, total bagian, dan model campuran |
| 06-07  | Model numerik          | Nomor 06-07 | Catatan persamaan, timeline, dan satuan                |

Aturan waktu:

- Nomor 01-03: 60-120 detik per soal.
- Nomor 04-05: 2-4 menit per soal.
- Nomor 06-07: 2-4 menit per soal.
- Review wajib mencatat letak salah: salah pola, salah operasi, salah model, salah satuan, atau salah hitung.

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

### The Concept

Deret angka bisa memakai beberapa jenis pola. Jangan hafal satu pola saja.

| Pola               | Bentuk Cepat                               | Contoh          |
| ------------------ | ------------------------------------------ | --------------- |
| Selisih tetap      | tambah angka yang sama                     | 4, 9, 14, 19    |
| Rasio tetap        | kali angka yang sama                       | 3, 6, 12, 24    |
| Selisih bertingkat | selisihnya naik/turun berpola              | 2, 5, 10, 17    |
| Operasi berulang   | kali lalu tambah/kurang                    | 3, 7, 15, 31    |
| Pola selang-seling | posisi ganjil dan genap punya pola berbeda | 2, 10, 4, 20, 8 |

### Pattern of Question

```text
Barisan angka -> satu angka kosong/berikutnya -> opsi angka yang mirip
```

### Pattern to Answer

1. Cek selisih antar angka.
2. Kalau selisih tidak stabil, cek selisih tingkat dua.
3. Kalau masih tidak stabil, cek rasio atau operasi `x n + k`.
4. Cek apakah pola berlaku ke semua angka, bukan hanya dua angka terakhir.

```text
barisan -> pola antar angka -> verifikasi semua langkah -> angka berikutnya
```

### Common Traps

| Jebakan                    | Contoh                                                |
| -------------------------- | ----------------------------------------------------- |
| Hanya lihat angka terakhir | 63 ke opsi terdekat                                   |
| Salah pola selisih         | mengira selisihnya tambah tetap                       |
| Tidak verifikasi           | pola cocok di dua langkah, tapi gagal di langkah lain |

### Review Rule

Untuk setiap deret, wajib tulis pola dalam satu baris:

```text
Pola: setiap angka ...
```

### Practice Question

**Soal**  
Tentukan angka berikutnya:

```text
3, 7, 15, 31, 63, ...
```

A. 95  
B. 111  
C. 127  
D. 129  
E. 143

**Jawaban: C. 127**

**Pembahasan**  
Pola deret adalah dikali 2 lalu tambah 1.

```text
3 x 2 + 1 = 7
7 x 2 + 1 = 15
15 x 2 + 1 = 31
31 x 2 + 1 = 63
63 x 2 + 1 = 127
```

Jadi angka berikutnya adalah `127`.

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

Jangan hitung dari kiri secara membabi buta. Sederhanakan bagian yang mudah.

### The Concept

Urutan operasi:

```text
kurung -> kali/bagi -> tambah/kurang
```

Untuk kali dan bagi yang sejajar, kerjakan dari kiri ke kanan atau sederhanakan
faktor yang jelas aman.

Trik angka yang sering membantu:

| Bentuk  | Nilai |
| ------- | ----- |
| 25 x 4  | 100   |
| 125 x 8 | 1.000 |
| 0,5     | 1/2   |
| 0,25    | 1/4   |
| 0,75    | 3/4   |

### Pattern of Question

```text
Ekspresi hitung campuran -> angka besar tapi bisa disederhanakan
```

### Pattern to Answer

```text
pecah ekspresi -> sederhanakan faktor -> hitung akhir
```

### Common Traps

| Jebakan               | Kenapa Salah                                 |
| --------------------- | -------------------------------------------- |
| Menjumlah dulu        | Tambah/kurang kalah prioritas dari kali/bagi |
| Tidak menyederhanakan | Hitungan jadi besar dan rawan salah          |
| Salah baca ribuan     | 1.150 bisa terbaca 115 kalau terburu-buru    |

### Review Rule

Setelah mengerjakan, cek ulang ekspresi dengan tanda kurung mental:

```text
(bagian kali/bagi) + (bagian kali/bagi)
```

### Practice Question

**Soal**  
Nilai dari:

```text
48 x 125 / 24 + 36 x 25
```

adalah:

A. 1.125  
B. 1.150  
C. 1.175  
D. 1.200  
E. 1.250

**Jawaban: B. 1.150**

**Pembahasan**  
Kerjakan kali dan bagi sebelum penjumlahan.

```text
48 x 125 / 24 = 48 / 24 x 125
                 = 2 x 125
                 = 250

36 x 25 = 900

250 + 900 = 1.150
```

Jadi hasilnya adalah `1.150`.

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

### The Concept

Kunci pecahan dan desimal adalah menyamakan bentuk sebelum operasi.

Konversi penting:

| Desimal | Pecahan |
| ------- | ------- |
| 0,1     | 1/10    |
| 0,2     | 1/5     |
| 0,25    | 1/4     |
| 0,375   | 3/8     |
| 0,5     | 1/2     |
| 0,625   | 5/8     |
| 0,75    | 3/4     |

### Pattern of Question

```text
Pecahan + desimal campur -> bentuk angka perlu disamakan
```

### Pattern to Answer

1. Cari desimal yang punya pecahan mudah.
2. Samakan bentuk angka.
3. Sederhanakan bagian yang saling menghapus.
4. Baru hitung hasil akhir.

```text
desimal mudah -> pecahan setara -> sederhanakan -> hasil akhir
```

### Common Traps

| Jebakan                | Contoh                        |
| ---------------------- | ----------------------------- |
| Salah konversi         | 0,625 dianggap 6/25           |
| Semua dibuat desimal   | boleh, tapi rawan salah koma  |
| Tidak lihat pembatalan | 0,625 dan 5/8 sebenarnya sama |

### Review Rule

Sebelum menghitung, tulis bentuk yang kamu pilih:

```text
Mau dihitung sebagai pecahan atau desimal?
```

### Practice Question

**Soal**  
Nilai dari:

```text
3/4 + 0,625 - 5/8 + 0,2
```

adalah:

A. 17/20  
B. 9/10  
C. 19/20  
D. 1  
E. 21/20

**Jawaban: C. 19/20**

**Pembahasan**  
Ubah desimal yang mudah:

```text
0,625 = 5/8
0,2 = 1/5
```

Maka:

```text
3/4 + 0,625 - 5/8 + 0,2
= 3/4 + 5/8 - 5/8 + 1/5
= 3/4 + 1/5
= 15/20 + 4/20
= 19/20
```

---

## Num04 - Persentase

### The Problem

Soal persentase sering terlihat sederhana, tetapi jebakannya ada pada nilai
dasar. Kenaikan atau diskon tahap kedua biasanya dihitung dari nilai terbaru,
bukan dari nilai awal.

### The Concept

Persentase lebih aman dibaca sebagai faktor pengali.

```text
naik p% -> kali (1 + p/100)
turun p% -> kali (1 - p/100)
diskon p% -> kali (1 - p/100)
```

### Pattern of Question

```text
Nilai awal -> naik/turun/diskon bertahap -> bandingkan harga akhir
```

### Pattern to Answer

```text
ambil angka mudah -> ubah persen jadi faktor -> hitung tahap demi tahap
```

### Common Traps

| Jebakan                 | Kenapa Salah                                   |
| ----------------------- | ---------------------------------------------- |
| Menjumlah persen mentah | Naik 25% lalu turun 20% bukan otomatis naik 5% |
| Pakai nilai awal terus  | Diskon dihitung dari harga setelah kenaikan    |
| Salah faktor            | Diskon 20% berarti kali 0,8, bukan kali 0,2    |

### Review Rule

Untuk persen bertahap, tulis nilai setelah setiap tahap:

```text
awal -> tahap 1 -> tahap 2 -> akhir
```

### Practice Question

**Soal**  
Harga sebuah barang dinaikkan 25%, lalu diberi diskon 20% dari harga setelah
kenaikan. Harga akhir dibandingkan harga awal adalah:

A. Turun 5%  
B. Tetap  
C. Naik 2,5%  
D. Naik 5%  
E. Naik 10%

**Jawaban: B. Tetap**

**Pembahasan**  
Misalkan harga awal `100`.

```text
Naik 25% -> 100 x 1,25 = 125
Diskon 20% -> 125 x 0,8 = 100
```

Harga akhir kembali sama dengan harga awal, jadi jawabannya `tetap`.

---

## Num05 - Perbandingan

### The Problem

Soal perbandingan tidak selalu berbentuk `A : B`. Kadang muncul sebagai selisih
bagian, pembagian keuntungan proporsional, atau campuran yang sebagian diambil
lalu diganti bahan murni. Kuncinya adalah tahu bagian mana yang menjadi total,
bagian mana yang menjadi porsi, dan bagian mana yang berubah.

### The Concept

Rasio dibaca sebagai bagian. Kalau perbandingan `A : B : C = 2 : 5 : 8`, maka:

```text
A = 2x
B = 5x
C = 8x
```

Untuk pembagian proporsional:

```text
porsi orang = bagian orang / total bagian
nilai orang = porsi orang x total nilai
```

Untuk campuran replacement:

```text
ambil campuran -> komposisi keluar mengikuti rasio awal
ganti bahan murni -> tambahkan hanya bahan pengganti
bentuk rasio akhir -> selesaikan total
```

### Pattern of Question

```text
Rasio/saham/campuran -> total atau selisih diketahui -> cari nilai salah satu bagian
```

### Pattern to Answer

```text
ubah rasio jadi bagian -> hubungkan dengan total/selisih -> cari nilai nyata
```

Untuk campuran:

```text
Total wadah = T
X awal = bagian X / total bagian x T
Y awal = bagian Y / total bagian x T
```

### Common Traps

| Jebakan                               | Kenapa Salah                                            |
| ------------------------------------- | ------------------------------------------------------- |
| Memakai bagian sebagai nilai langsung | `2 : 5 : 8` masih bagian, bukan rupiah                  |
| Salah total bagian                    | Total bagian harus dijumlahkan dulu                     |
| Salah porsi proporsional              | Modal 300 dari total 1.000 berarti 30%, bukan 300%      |
| Salah campuran keluar                 | Yang keluar mengikuti rasio campuran, bukan bahan murni |

### Review Rule

Sebelum hitung, tulis:

```text
Total bagian:
Nilai 1 bagian:
Yang ditanya:
```

Untuk campuran, tambahkan:

```text
Yang keluar:
Yang masuk:
Rasio akhir:
```

### Practice Question

**Latihan 1 - Rasio Selisih**

**Soal**  
Perbandingan uang A : B : C adalah `2 : 5 : 8`. Jika selisih uang C dan A
adalah Rp180.000, jumlah uang mereka adalah:

A. Rp420.000  
B. Rp450.000  
C. Rp480.000  
D. Rp500.000  
E. Rp540.000

**Jawaban: B. Rp450.000**

**Pembahasan**  
Misalkan:

```text
A = 2x
B = 5x
C = 8x
```

Selisih C dan A:

```text
8x - 2x = 180.000
6x = 180.000
x = 30.000
```

Jumlah uang mereka:

```text
2x + 5x + 8x = 15x
15 x 30.000 = 450.000
```

Jadi jumlah uang mereka adalah `Rp450.000`.

**Latihan 2 - Pembagian Proporsional**

**Soal**  
Tiga investor, Pak Andi, Pak Budi, dan Pak Candra, masing-masing
menginvestasikan Rp200 juta, Rp300 juta, dan Rp500 juta dalam sebuah usaha.
Setahun kemudian, usaha menghasilkan keuntungan total Rp240 juta yang dibagi
proporsional berdasarkan modal. Berapa keuntungan yang diterima Pak Budi?

A. Rp52 juta  
B. Rp60 juta  
C. Rp72 juta  
D. Rp80 juta

**Jawaban: C. Rp72 juta**

**Pembahasan**  
Total investasi:

```text
200 + 300 + 500 = 1.000 juta
```

Porsi Pak Budi:

```text
300 / 1.000 = 30%
```

Keuntungan Pak Budi:

```text
30% x 240 juta = 72 juta
```

Jadi keuntungan yang diterima Pak Budi adalah `Rp72 juta`.

**Latihan 3 - Campuran Replacement**

**Soal**  
Sebuah wadah berisi campuran bahan kimia X dan Y dengan rasio `5 : 3`. Jika 16
liter campuran tersebut diambil dan diganti dengan 16 liter bahan kimia Y murni,
rasio bahan kimia X dan Y di dalam wadah bergeser menjadi `3 : 5`. Berapa liter
kapasitas total wadah tersebut?

A. 32 liter  
B. 40 liter  
C. 48 liter  
D. 56 liter  
E. 64 liter

**Jawaban: B. 40 liter**

**Pembahasan**  
Misalkan kapasitas total wadah adalah `T` liter.

Rasio awal `X : Y = 5 : 3`, jadi:

```text
X awal = 5/8 T
Y awal = 3/8 T
```

Saat 16 liter campuran diambil, yang keluar mengikuti rasio `5 : 3`:

```text
X yang keluar = 5/8 x 16 = 10
Y yang keluar = 3/8 x 16 = 6
```

Setelah itu, 16 liter Y murni ditambahkan. Maka:

```text
X akhir = 5/8 T - 10
Y akhir = 3/8 T - 6 + 16
        = 3/8 T + 10
```

Rasio akhir `X : Y = 3 : 5`:

```text
(5/8 T - 10) / (3/8 T + 10) = 3/5
5(5/8 T - 10) = 3(3/8 T + 10)
25/8 T - 50 = 9/8 T + 30
16/8 T = 80
2T = 80
T = 40
```

Jadi kapasitas total wadah adalah `40 liter`.

---

## Num06 - Aljabar Sederhana

### The Problem

Soal aljabar sederhana sering salah karena peserta tidak memilih variabel yang
tepat atau terlalu cepat menghilangkan informasi. Padahal modelnya biasanya
cukup satu atau dua persamaan.

### The Concept

Aljabar dipakai untuk menyimpan nilai yang belum diketahui. Setelah model
terbentuk, gunakan substitusi atau eliminasi.

### Pattern of Question

```text
Dua hubungan angka -> satu nilai ditanya -> bentuk persamaan
```

### Pattern to Answer

```text
tulis persamaan -> isolasi salah satu variabel -> substitusi -> cek hasil
```

### Common Traps

| Jebakan                         | Kenapa Salah                                 |
| ------------------------------- | -------------------------------------------- |
| Salah substitusi                | Tanda minus sering berubah                   |
| Tidak cek kembali               | Nilai x dan y harus memenuhi semua persamaan |
| Menghitung mental terlalu cepat | Persamaan pendek tetap bisa salah tanda      |

### Review Rule

Setelah dapat nilai, masukkan lagi ke persamaan awal.

### Practice Question

**Soal**  
Jika:

```text
2x + 3y = 41
x + y = 16
```

maka nilai `x` adalah:

A. 5  
B. 6  
C. 7  
D. 8  
E. 9

**Jawaban: C. 7**

**Pembahasan**  
Dari persamaan kedua:

```text
x + y = 16
x = 16 - y
```

Substitusi ke persamaan pertama:

```text
2(16 - y) + 3y = 41
32 - 2y + 3y = 41
y = 9
```

Maka:

```text
x + 9 = 16
x = 7
```

---

## Num07 - Soal Cerita

### The Problem

Soal cerita kuantitatif menjebak bukan karena rumusnya sulit, tetapi karena
informasi harus diterjemahkan dulu. Pada soal jarak-waktu-kecepatan, kesalahan
umum adalah lupa bahwa satu pihak mulai lebih dulu atau lupa memakai kecepatan
relatif.

### The Concept

Untuk soal jarak-waktu-kecepatan:

```text
jarak = kecepatan x waktu
waktu = jarak / kecepatan
```

Jika dua objek saling mendekat:

```text
kecepatan relatif = kecepatan 1 + kecepatan 2
```

### Pattern of Question

```text
Cerita jarak/waktu/kecepatan -> ada jeda mulai -> tanya waktu bertemu
```

### Pattern to Answer

```text
hitung perjalanan awal -> cari sisa jarak -> pakai kecepatan relatif -> tambah ke waktu mulai
```

### Common Traps

| Jebakan                   | Kenapa Salah                                                    |
| ------------------------- | --------------------------------------------------------------- |
| Abaikan jeda mulai        | Mobil pertama sudah menempuh jarak sebelum mobil kedua bergerak |
| Pakai satu kecepatan saja | Saat saling mendekat, kecepatannya dijumlah                     |
| Salah ubah jam            | 1,5 jam berarti 1 jam 30 menit                                  |

### Review Rule

Tulis timeline:

```text
08.00 -> 08.30 -> waktu bertemu
```

### Practice Question

**Soal**  
Jarak kota A dan kota B adalah 180 km. Mobil pertama berangkat dari kota A
pukul 08.00 menuju kota B dengan kecepatan 60 km/jam. Mobil kedua berangkat dari
kota B pukul 08.30 menuju kota A dengan kecepatan 40 km/jam. Jika keduanya
melewati jalan yang sama, pukul berapa mereka bertemu?

A. 09.30  
B. 09.45  
C. 10.00  
D. 10.15  
E. 10.30

**Jawaban: C. 10.00**

**Pembahasan**  
Dari pukul 08.00 sampai 08.30, mobil pertama sudah berjalan 0,5 jam:

```text
60 x 0,5 = 30 km
```

Sisa jarak saat mobil kedua mulai bergerak:

```text
180 - 30 = 150 km
```

Karena kedua mobil saling mendekat, kecepatan relatifnya:

```text
60 + 40 = 100 km/jam
```

Waktu untuk bertemu setelah 08.30:

```text
150 / 100 = 1,5 jam
```

Maka mereka bertemu pada:

```text
08.30 + 1 jam 30 menit = 10.00
```
