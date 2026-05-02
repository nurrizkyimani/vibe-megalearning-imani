# Scholastic Pemecahan Masalah Questions

Bank soal untuk `problem-solving-curriculum.md`. Setiap bagian mengacu ke nomor
kurikulum dan nama topik, tanpa kode singkatan tambahan.

Target latihan:

- Kerjakan dulu tanpa melihat jawaban.
- Tulis model singkat sebelum membuka pembahasan.
- Setelah selesai, catat apakah salah karena kata kunci, model, satuan,
  eliminasi, atau hitung.

---

## 01 - Himpunan dan Irisan Data Survei

Mengacu ke curriculum #01: Himpunan dan Irisan Data Survei.

**Soal**  
Hasil survei terhadap 300 mahasiswa penerima beasiswa:

- 180 mahasiswa aktif di organisasi.
- 150 mahasiswa memiliki IPK >= 3,5.
- 90 mahasiswa aktif organisasi dan memiliki IPK >= 3,5.
- Sisanya tidak memenuhi keduanya.

Berapa persen mahasiswa yang hanya memiliki IPK >= 3,5 tanpa aktif organisasi?

A. 15%  
B. 20%  
C. 25%  
D. 30%  
E. 35%

**Jawaban: B. 20%**

**Pembahasan**  
Yang ditanya adalah **hanya IPK**, berarti kelompok IPK dikurangi irisan IPK
dan aktif organisasi.

```text
Hanya IPK = IPK - (IPK dan aktif)
           = 150 - 90
           = 60
```

Persentase dari total mahasiswa:

```text
60 / 300 x 100 = 20%
```

Jadi jawabannya adalah `20%`.

---

## 02 - Target dan Sisa Pekerjaan

Mengacu ke curriculum #02: Target dan Sisa Pekerjaan.

**Soal**  
Sebuah tim harus menyelesaikan 84 laporan dalam 7 hari. Dalam 3 hari pertama,
tim menyelesaikan 30 laporan. Agar target tercapai tepat waktu, berapa rata-rata
laporan per hari yang harus diselesaikan pada hari-hari tersisa?

A. 12  
B. 13  
C. 13,5  
D. 14  
E. 15

**Jawaban: C. 13,5**

**Pembahasan**  
Yang dibagi bukan 84, tetapi sisa laporan.

```text
Sisa laporan = 84 - 30 = 54
Sisa hari = 7 - 3 = 4
Rata-rata sisa = 54 / 4 = 13,5
```

Jadi rata-rata yang diperlukan adalah `13,5 laporan per hari`.

---

## 03 - Perubahan Bertahap

Mengacu ke curriculum #03: Perubahan Bertahap.

**Soal**  
Sebuah program awalnya menargetkan 240 peserta. Setelah seleksi administrasi,
jumlah peserta berkurang 25%. Dari peserta yang tersisa, 20% mengundurkan diri.
Berapa peserta yang masih tersisa?

A. 132  
B. 144  
C. 150  
D. 156  
E. 180

**Jawaban: B. 144**

**Pembahasan**  
Setiap persen memakai nilai terbaru.

```text
Setelah berkurang 25%:
240 x 0,75 = 180

Setelah 20% dari sisa mengundurkan diri:
180 x 0,8 = 144
```

Jadi peserta yang masih tersisa adalah `144`.

---

## 04 - Alokasi Sumber Daya

Mengacu ke curriculum #04: Alokasi Sumber Daya.

**Soal**  
Sebuah panitia memiliki 49 relawan. Setiap pos layanan membutuhkan minimal 5
relawan. Pos pendaftaran harus mendapat 2 relawan lebih banyak daripada pos
informasi. Pos konsumsi harus mendapat 3 relawan lebih banyak daripada pos
pendaftaran. Jika hanya ada tiga pos tersebut dan semua relawan harus dibagi,
berapa relawan di pos konsumsi?

A. 15  
B. 17  
C. 19  
D. 21  
E. 23

**Jawaban: C. 19**

**Pembahasan**  
Misalkan pos informasi mendapat `x` relawan.

```text
Informasi = x
Pendaftaran = x + 2
Konsumsi = x + 5
```

Total relawan:

```text
x + (x + 2) + (x + 5) = 49
3x + 7 = 49
3x = 42
x = 14
```

Maka:

```text
Informasi = 14
Pendaftaran = 16
Konsumsi = 19
```

Jadi relawan di pos konsumsi adalah `19`.

---

## 05 - Optimasi Sederhana

Mengacu ke curriculum #05: Optimasi Sederhana.

**Soal**  
Sebuah organisasi harus mengirim minimal 73 barang bantuan. Paket besar memuat
12 barang dengan biaya Rp45.000. Paket sedang memuat 8 barang dengan biaya
Rp32.000. Paket kecil memuat 5 barang dengan biaya Rp22.000. Setiap paket harus
terisi penuh. Berapa biaya kirim minimum?

A. Rp274.000  
B. Rp277.000  
C. Rp279.000  
D. Rp282.000  
E. Rp288.000

**Jawaban: C. Rp279.000**

**Pembahasan**  
Bandingkan biaya per barang:

```text
Besar: 45.000 / 12 = 3.750
Sedang: 32.000 / 8 = 4.000
Kecil: 22.000 / 5 = 4.400
```

Paket besar paling murah per barang, tetapi kapasitas harus mencapai minimal
73. Coba kombinasi dekat 73:

```text
6 besar = 72 barang, belum cukup
5 besar + 1 sedang + 1 kecil = 73 barang, biaya 225.000 + 32.000 + 22.000 = 279.000
6 besar + 1 kecil = 77 barang, biaya 270.000 + 22.000 = 292.000
5 besar + 2 sedang = 76 barang, biaya 225.000 + 64.000 = 289.000
4 besar + 3 sedang + 1 kecil = 77 barang, biaya 180.000 + 96.000 + 22.000 = 298.000
3 besar + 4 sedang + 1 kecil = 73 barang, biaya 135.000 + 128.000 + 22.000 = 285.000
```

Kombinasi termurah yang memenuhi minimal 73 barang adalah `Rp279.000`.

---

## 06 - Jadwal dan Urutan Langkah

Mengacu ke curriculum #06: Jadwal dan Urutan Langkah.

**Soal**  
Empat kegiatan A, B, C, dan D harus dilakukan. B hanya boleh dilakukan setelah
A selesai. C hanya boleh dilakukan setelah A selesai. D hanya boleh dilakukan
setelah B dan C selesai. Durasi A = 20 menit, B = 30 menit, C = 15 menit, dan
D = 25 menit. Jika B dan C bisa berjalan bersamaan setelah A selesai, berapa
waktu minimum untuk menyelesaikan semua kegiatan?

A. 70 menit  
B. 75 menit  
C. 80 menit  
D. 90 menit  
E. 95 menit

**Jawaban: B. 75 menit**

**Pembahasan**  
A harus selesai dulu:

```text
A = 20 menit
```

Setelah A, B dan C bisa berjalan bersamaan. Karena D menunggu keduanya selesai,
ambil durasi yang lebih lama:

```text
max(B, C) = max(30, 15) = 30 menit
```

Lalu D:

```text
Total = 20 + 30 + 25 = 75 menit
```

---

## 07 - Laju Gabungan

Mengacu ke curriculum #07: Laju Gabungan.

**Soal**  
Mesin A dapat menyelesaikan satu pekerjaan dalam 12 jam. Mesin B dapat
menyelesaikan pekerjaan yang sama dalam 18 jam. Jika keduanya bekerja bersama
selama 4 jam, berapa bagian pekerjaan yang sudah selesai?

A. 4/9  
B. 5/9  
C. 2/3  
D. 7/9  
E. 5/6

**Jawaban: B. 5/9**

**Pembahasan**  
Ubah menjadi laju per jam:

```text
A = 1/12 pekerjaan per jam
B = 1/18 pekerjaan per jam
Gabungan = 1/12 + 1/18 = 3/36 + 2/36 = 5/36
```

Dalam 4 jam:

```text
4 x 5/36 = 20/36 = 5/9
```

Jadi pekerjaan yang selesai adalah `5/9`.

**Latihan tambahan**  
Pekerja A dapat menyelesaikan satu proyek sendirian dalam 10 hari. Pekerja B
dapat menyelesaikan proyek yang sama sendirian dalam 15 hari. Mereka bekerja
bersama selama 4 hari, lalu A mundur dari proyek. Berapa hari lagi B harus
bekerja agar proyek selesai?

A. 3 hari  
B. 4 hari  
C. 5 hari  
D. 6 hari  
E. 7 hari

**Jawaban: C. 5 hari**

**Pembahasan**  
Ubah waktu kerja menjadi laju per hari:

```text
A = 1/10 proyek per hari
B = 1/15 proyek per hari
```

Laju gabungan:

```text
1/10 + 1/15 = 3/30 + 2/30 = 5/30 = 1/6
```

Mereka bekerja bersama selama 4 hari:

```text
4 x 1/6 = 4/6 = 2/3 proyek selesai
```

Sisa pekerjaan:

```text
1 - 2/3 = 1/3 proyek
```

Karena yang lanjut hanya B, bagi sisa pekerjaan dengan laju B:

```text
1/3 / 1/15 = 1/3 x 15 = 5
```

Jadi B harus bekerja `5 hari lagi`.

---

## 08 - Perbandingan Keputusan

Mengacu ke curriculum #08: Perbandingan Keputusan.

**Soal**  
Paket A memberi 18 sesi belajar dengan harga Rp720.000. Paket B memberi 24 sesi
belajar dengan harga Rp900.000. Jika kualitas sesi dianggap sama, paket mana
yang lebih murah per sesi dan berapa selisihnya?

A. Paket A lebih murah Rp2.500 per sesi  
B. Paket A lebih murah Rp5.000 per sesi  
C. Paket B lebih murah Rp2.500 per sesi  
D. Paket B lebih murah Rp5.000 per sesi  
E. Keduanya sama

**Jawaban: C. Paket B lebih murah Rp2.500 per sesi**

**Pembahasan**  
Bandingkan biaya per sesi.

```text
Paket A = 720.000 / 18 = 40.000
Paket B = 900.000 / 24 = 37.500
Selisih = 40.000 - 37.500 = 2.500
```

Jadi Paket B lebih murah `Rp2.500 per sesi`.

---

## 09 - Tabel Kasus

Mengacu ke curriculum #09: Tabel Kasus.

**Soal**  
Dalam sebuah kelas, 32 siswa mengikuti minimal satu dari dua kelas tambahan:
Bahasa Inggris dan Statistik. Sebanyak 20 siswa mengikuti Bahasa Inggris, 18
siswa mengikuti Statistik, dan beberapa siswa mengikuti keduanya. Berapa siswa
yang mengikuti keduanya?

A. 4  
B. 5  
C. 6  
D. 7  
E. 8

**Jawaban: C. 6**

**Pembahasan**  
Gunakan rumus gabungan:

```text
Inggris atau Statistik = Inggris + Statistik - Keduanya
32 = 20 + 18 - Keduanya
Keduanya = 38 - 32 = 6
```

Jadi siswa yang mengikuti keduanya adalah `6`.

---

## 10 - Diagram Alur

Mengacu ke curriculum #10: Diagram Alur.

**Soal**  
Dari 120 pendaftar, 75% lolos seleksi administrasi. Dari yang lolos administrasi,
2/3 mengikuti tes tertulis. Dari yang mengikuti tes tertulis, 40% dinyatakan
lolos ke tahap wawancara. Berapa pendaftar yang lolos ke tahap wawancara?

A. 24  
B. 30  
C. 36  
D. 40  
E. 48

**Jawaban: A. 24**

**Pembahasan**  
Ikuti alurnya tahap demi tahap.

```text
Lolos administrasi = 120 x 75% = 90
Mengikuti tes tertulis = 90 x 2/3 = 60
Lolos wawancara = 60 x 40% = 24
```

Jadi yang lolos ke tahap wawancara adalah `24`.

---

## 11 - Kecukupan Informasi Praktis

Mengacu ke curriculum #11: Kecukupan Informasi Praktis.

**Soal**  
Sebuah tim ingin mengetahui apakah target 500 paket per minggu tercapai. Data
yang tersedia:

1. Tim memproduksi 80 paket per hari pada Senin sampai Jumat.
2. Tim memproduksi 60 paket pada Sabtu.
3. Tidak ada produksi pada Minggu.

Informasi mana yang diperlukan untuk menjawab apakah target mingguan tercapai?

A. Pernyataan 1 saja cukup  
B. Pernyataan 1 dan 2 cukup  
C. Pernyataan 1 dan 3 cukup  
D. Pernyataan 2 dan 3 cukup  
E. Ketiga pernyataan diperlukan

**Jawaban: E. Ketiga pernyataan diperlukan**

**Pembahasan**  
Untuk memastikan target mingguan tercapai atau tidak, kita perlu menutup semua
hari dalam minggu tersebut.

```text
Senin-Jumat = 5 x 80 = 400
Sabtu = 60
Minggu = 0
Total mingguan = 460
```

Target 500 tidak tercapai. Tanpa pernyataan 3, masih ada kemungkinan produksi
Minggu yang belum diketahui. Jadi ketiga pernyataan diperlukan.

---

## 12 - Eliminasi Opsi

Mengacu ke curriculum #12: Eliminasi Opsi.

**Soal**  
Sebuah ruang seminar memiliki 6 baris kursi. Setiap baris harus berisi lebih
banyak kursi daripada baris sebelumnya. Total kursi adalah 57. Jika baris
pertama berisi 7 kursi dan selisih antarbaris selalu sama, berapa kursi pada
baris terakhir?

A. 10  
B. 11  
C. 12  
D. 13  
E. 14

**Jawaban: C. 12**

**Pembahasan**  
Deret kursi membentuk barisan aritmetika:

```text
Baris: 7, 7+d, 7+2d, 7+3d, 7+4d, 7+5d
Total = 42 + 15d
```

Karena total 57:

```text
42 + 15d = 57
15d = 15
d = 1
```

Baris terakhir:

```text
7 + 5d = 7 + 5 = 12
```

Jadi jawabannya adalah `12`.
