# M4/W4/D32 - Sat, 02 May 2026 (WIB)

## Scholastic Pemecahan Masalah - Curriculum

Kurikulum Pemecahan Masalah Tes Skolastik LPDP untuk membangun kemampuan
membaca cerita, memilih model, mengelola batasan, dan mengecek opsi. File ini
dibuat sebagai satu tempat belajar: setiap nomor berisi konsep, pola soal, pola
menjawab, jebakan umum, aturan review, dan latihan dari
`problem-solving-questions.md`.

Target audience: peserta yang sudah bisa operasi dasar, tetapi masih sering
bingung menentukan apa yang harus dikurang, dijumlah, dibandingkan, atau
dimodelkan dari soal cerita.

---

## Curriculum Table

| # | Topic | Core Concept | LPDP Usage |
| --- | --- | --- | --- |
| 01 | Himpunan dan Irisan Data Survei | Bedakan total kelompok, irisan, dan anggota "hanya" satu kelompok | Menjawab soal survei, organisasi, IPK, minat, atau kategori ganda |
| 02 | Target dan Sisa Pekerjaan | Target akhir dikurangi progres awal untuk mencari beban tersisa | Menghitung rata-rata baru, sisa dokumen, sisa produksi, atau sisa skor |
| 03 | Perubahan Bertahap | Setiap perubahan memakai nilai terbaru, bukan selalu nilai awal | Menghindari salah hitung pada kenaikan, penurunan, dan revisi target |
| 04 | Alokasi Sumber Daya | Bagikan sumber terbatas sesuai syarat, kapasitas, dan prioritas | Memilih kombinasi orang, biaya, waktu, atau kuota yang memenuhi batas |
| 05 | Optimasi Sederhana | Cari nilai minimum/maksimum yang memenuhi semua syarat | Menentukan biaya minimum, waktu tercepat, atau pilihan paling efisien |
| 06 | Jadwal dan Urutan Langkah | Susun informasi waktu dan prasyarat secara kronologis | Menentukan urutan kegiatan, konflik jadwal, dan waktu selesai |
| 07 | Laju Gabungan | Gabungkan laju kerja, cari bagian selesai, lalu hitung sisa jika satu pihak berhenti | Menjawab soal kerja bersama, antrean, produksi, atau mesin |
| 08 | Perbandingan Keputusan | Bandingkan dua skenario dengan ukuran yang sama | Memilih opsi lebih hemat, lebih cepat, atau lebih menguntungkan |
| 09 | Tabel Kasus | Ubah cerita menjadi tabel kecil agar relasi tidak tercampur | Menjawab soal kategori, status, kuota, dan kondisi bertingkat |
| 10 | Diagram Alur | Pecah proses menjadi tahap input, proses, dan output | Mengikuti aturan seleksi, distribusi, atau perubahan status |
| 11 | Kecukupan Informasi Praktis | Tentukan data mana yang benar-benar diperlukan untuk menjawab | Menghindari memakai angka yang tidak relevan dalam soal cerita |
| 12 | Eliminasi Opsi | Coret opsi yang melanggar syarat sebelum hitung detail | Mempercepat soal dengan banyak pilihan dan batasan |

---

## Grind Plan

| Urutan | Fokus | Target | Output |
| --- | --- | --- | --- |
| 01-03 | Model dasar cerita | Nomor 01-03 | Catatan irisan, sisa target, dan perubahan bertahap |
| 04-08 | Optimasi dan jadwal | Nomor 04-08 | Catatan batas, satuan, dan pilihan terbaik |
| 09-12 | Tabel, alur, dan eliminasi | Nomor 09-12 | Catatan data relevan dan opsi yang gugur |

Aturan waktu:

- Nomor 01-03: 90-120 detik per soal.
- Nomor 04-08: 2-3 menit per soal.
- Nomor 09-12: 2-3 menit per soal.
- Review wajib menulis letak salah: salah baca kata kunci, salah model, salah satuan, salah eliminasi, atau salah hitung.

---

## Num01 - Himpunan dan Irisan Data Survei

### The Problem

Soal survei sering terlihat seperti soal persentase biasa, padahal jebakan
utamanya ada di kata **dan**, **atau**, **hanya**, dan **sisanya**. Angka irisan
tidak boleh dikurangkan dari total yang salah.

### The Concept

Rumus cepat:

```text
Hanya A = A - (A dan B)
Hanya B = B - (A dan B)
A atau B = A + B - (A dan B)
Tidak keduanya = Total - (A atau B)
```

Untuk soal IPK dan organisasi:

```text
Hanya IPK = IPK - (IPK dan aktif)
```

### Pattern of Question

```text
Total responden -> jumlah A -> jumlah B -> irisan atau tidak keduanya -> tanya hanya A/B atau persen
```

### Pattern to Answer

```text
identifikasi total, A, B, irisan -> pilih rumus hanya/gabungan -> ubah ke persen jika diminta
```

Jika yang diketahui adalah **tidak keduanya**, cari dulu minimal salah satu:

```text
Minimal salah satu = Total - Tidak keduanya
Irisan = A + B - Minimal salah satu
Hanya A = A - Irisan
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Mengurangkan total kelompok lain | `150 - 180` tidak bermakna untuk mencari hanya IPK |
| Lupa bahwa irisan masuk dua kelompok | Angka "dan" dihitung di kedua kelompok |
| Salah baca "sisanya" | Sisa berarti tidak memenuhi keduanya, bukan hanya tidak aktif |

### Review Rule

Sebelum menghitung, tulis:

```text
Total:
A:
B:
A dan B:
Yang ditanya:
```

### Practice Question

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

**Latihan tambahan - Hanya A dari data tidak keduanya**

**Soal 1**  
Dari 80 peserta pelatihan, diketahui 45 peserta menguasai Excel, 38 peserta
menguasai SQL, dan 12 peserta tidak menguasai keduanya. Berapa banyak peserta
yang **hanya menguasai Excel**?

A. 19  
B. 25  
C. 30  
D. 35  
E. 45

**Soal 2**  
Dari 100 karyawan, 60 karyawan bisa berbahasa Inggris, 45 karyawan bisa
berbahasa Mandarin, dan 20 karyawan tidak bisa keduanya. Berapa banyak karyawan
yang **hanya bisa bahasa Inggris**?

A. 25  
B. 30  
C. 35  
D. 40  
E. 45

**Soal 3**  
Dalam sebuah kelas berisi 50 siswa, 28 siswa mengikuti klub Matematika, 24 siswa
mengikuti klub Sains, dan 8 siswa tidak mengikuti kedua klub tersebut. Berapa
banyak siswa yang **hanya mengikuti klub Matematika**?

A. 10  
B. 14  
C. 18  
D. 20  
E. 22

**Soal 4**  
Dari 120 pelamar, 70 pelamar memiliki sertifikat data analyst, 55 pelamar
memiliki sertifikat project management, dan 25 pelamar tidak memiliki kedua
sertifikat tersebut. Berapa banyak pelamar yang **hanya memiliki sertifikat data
analyst**?

A. 25  
B. 30  
C. 35  
D. 40  
E. 45

**Soal 5**  
Dari 90 mahasiswa, 52 mahasiswa mengikuti organisasi kampus, 40 mahasiswa
mengikuti kegiatan relawan, dan 18 mahasiswa tidak mengikuti keduanya. Berapa
banyak mahasiswa yang **hanya mengikuti organisasi kampus**?

A. 20  
B. 25  
C. 30  
D. 32  
E. 35

**Pembahasan Soal 1**  

```text
Minimal salah satu = 80 - 12 = 68
Irisan = 45 + 38 - 68 = 15
Hanya Excel = 45 - 15 = 30
```

Jawaban: **C. 30**

**Pembahasan Soal 2**  

```text
Minimal salah satu = 100 - 20 = 80
Irisan = 60 + 45 - 80 = 25
Hanya Inggris = 60 - 25 = 35
```

Jawaban: **C. 35**

**Pembahasan Soal 3**  

```text
Minimal salah satu = 50 - 8 = 42
Irisan = 28 + 24 - 42 = 10
Hanya Matematika = 28 - 10 = 18
```

Jawaban: **C. 18**

**Pembahasan Soal 4**  

```text
Minimal salah satu = 120 - 25 = 95
Irisan = 70 + 55 - 95 = 30
Hanya Data Analyst = 70 - 30 = 40
```

Jawaban: **D. 40**

**Pembahasan Soal 5**  

```text
Minimal salah satu = 90 - 18 = 72
Irisan = 52 + 40 - 72 = 20
Hanya Organisasi = 52 - 20 = 32
```

Jawaban: **D. 32**

---

## Num02 - Target dan Sisa Pekerjaan

### The Problem

Soal target sering menjebak karena peserta langsung membagi total dengan waktu
baru. Padahal sebagian pekerjaan sudah selesai, sehingga yang dibagi adalah
sisa target, bukan target awal.

### The Concept

```text
Sisa target = target akhir - progres yang sudah dicapai
Beban per periode = sisa target / sisa periode
```

Pastikan satuan periode sama: hari, minggu, jam, atau sesi.

### Pattern of Question

```text
Target total -> progres awal -> waktu tersisa -> rata-rata yang dibutuhkan
```

### Pattern to Answer

```text
target - progres -> sisa target / sisa waktu
```

### Common Traps

| Jebakan | Contoh |
| --- | --- |
| Membagi total awal | `84 / 4` padahal 30 sudah selesai |
| Lupa sisa waktu | memakai 7 hari, bukan 4 hari |
| Salah satuan | mencampur minggu dan hari tanpa konversi |

### Review Rule

Tulis kalimat ini sebelum hitung:

```text
Yang tersisa adalah ...
```

### Practice Question

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

## Num03 - Perubahan Bertahap

### The Problem

Perubahan bertahap sering salah karena peserta memakai nilai awal untuk semua
tahap. Padahal setelah tahap pertama, nilai dasar sudah berubah.

### The Concept

```text
Nilai baru = nilai sekarang x faktor perubahan
turun p% -> x (1 - p/100)
naik p% -> x (1 + p/100)
```

Setelah tiap tahap, simpan nilai terbaru sebelum lanjut ke tahap berikutnya.

### Pattern of Question

```text
Nilai awal -> perubahan tahap 1 -> perubahan tahap 2 -> nilai akhir
```

### Pattern to Answer

```text
nilai awal -> hitung tahap 1 -> pakai hasil tahap 1 untuk tahap 2
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Semua persen dihitung dari awal | Basis tahap kedua sudah berubah |
| Menjumlah persen mentah | Turun 25% lalu turun 20% bukan turun 45% dari awal |
| Tidak simpan nilai antara | Sulit mengecek tahap mana yang salah |

### Review Rule

Buat tabel kecil:

| Tahap | Operasi | Hasil |
| --- | --- | --- |
| Awal | - | ... |
| 1 | ... | ... |
| 2 | ... | ... |

### Practice Question

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

## Num04 - Alokasi Sumber Daya

### The Problem

Soal alokasi sumber daya memakai relasi antar kelompok, batas minimal, dan total
yang harus habis dibagi. Kesalahan umum adalah tidak memodelkan relasi dengan
variabel.

### The Concept

Gunakan satu variabel untuk kelompok paling dasar, lalu turunkan kelompok lain
dari relasi dalam soal.

### Pattern of Question

```text
Total sumber daya -> beberapa pos/kelompok -> relasi lebih banyak/lebih sedikit -> tanya salah satu pos
```

### Pattern to Answer

```text
pilih variabel dasar -> bentuk semua pos -> jumlahkan = total -> cari pos yang ditanya
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Tidak memakai variabel | Relasi antar pos mudah tercampur |
| Salah membaca "lebih banyak daripada" | Selisih harus ditambahkan ke pos pembanding |
| Lupa cek total | Semua pos harus menjumlah ke total relawan |

### Review Rule

Setelah dapat jawaban, jumlahkan semua kelompok untuk cek total.

### Practice Question

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

## Num05 - Optimasi Sederhana

### The Problem

Optimasi sederhana meminta nilai minimum atau maksimum yang memenuhi semua
syarat. Jebakannya adalah memilih opsi paling murah per unit tanpa mengecek
kapasitas, sisa, atau kombinasi paket.

### The Concept

Optimasi praktis biasanya tidak perlu rumus rumit. Batasi pencarian pada
kombinasi yang dekat dengan target dan cek semua syarat.

### Pattern of Question

```text
Target minimal/maksimal -> beberapa pilihan kapasitas dan biaya -> cari kombinasi terbaik
```

### Pattern to Answer

```text
cek biaya per unit -> coba kombinasi dekat target -> pilih biaya minimum yang memenuhi syarat
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Pilih paket termurah per barang saja | Bisa kurang dari target |
| Tidak cek paket penuh | Setiap paket harus terisi penuh |
| Berhenti di kombinasi pertama | Kombinasi valid belum tentu minimum |

### Review Rule

Untuk setiap kombinasi, tulis:

```text
jumlah barang:
biaya:
valid/tidak:
```

### Practice Question

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

## Num06 - Jadwal dan Urutan Langkah

### The Problem

Soal jadwal sering menjebak karena beberapa kegiatan bisa berjalan bersamaan,
sementara kegiatan lain harus menunggu prasyarat. Jangan menjumlah semua durasi
kalau ada pekerjaan paralel.

### The Concept

Jika dua kegiatan berjalan bersamaan dan kegiatan berikutnya menunggu keduanya,
pakai durasi yang lebih lama dari dua kegiatan paralel itu.

### Pattern of Question

```text
Daftar kegiatan -> prasyarat -> durasi -> waktu minimum selesai
```

### Pattern to Answer

```text
urutkan prasyarat -> kelompokkan kegiatan paralel -> ambil durasi maksimum -> jumlahkan jalur kritis
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Menjumlah B dan C | B dan C bisa berjalan bersamaan |
| Lupa D menunggu dua kegiatan | D baru mulai setelah B dan C selesai |
| Tidak buat urutan | Prasyarat mudah tertukar |

### Review Rule

Gambar alur singkat:

```text
A -> B -> D
A -> C -> D
```

### Practice Question

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

## Num07 - Laju Gabungan

### The Problem

Soal laju gabungan sering salah karena peserta memakai waktu selesai sebagai
angka biasa, bukan sebagai laju kerja. Jika satu pihak berhenti, kamu harus
mencari sisa pekerjaan dulu sebelum menghitung waktu pihak yang tersisa.

### The Concept

```text
laju A = 1 / waktu A
laju B = 1 / waktu B
laju gabungan = laju A + laju B
pekerjaan selesai = waktu kerja x laju
```

### Pattern of Question

```text
Waktu A sendiri -> waktu B sendiri -> kerja bersama -> tanya bagian selesai atau sisa waktu
```

### Pattern to Answer

```text
ubah waktu jadi laju -> jumlahkan laju -> hitung bagian selesai -> cari sisa jika perlu
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Menjumlah waktu | 12 jam + 18 jam bukan waktu gabungan |
| Lupa satuan | Jam dan hari tidak boleh dicampur |
| Tidak hitung sisa | Kalau satu pihak mundur, sisa dibagi laju pihak yang lanjut |

### Review Rule

Selalu tulis laju sebagai pecahan pekerjaan per satuan waktu.

### Practice Question

**Latihan 1 - Bagian Selesai**

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

**Latihan 2 - Satu Pihak Berhenti**

**Soal**  
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

## Num08 - Perbandingan Keputusan

### The Problem

Soal perbandingan keputusan meminta kamu membandingkan dua opsi dengan ukuran
yang sama. Jangan bandingkan harga total kalau jumlah manfaatnya berbeda.

### The Concept

Ubah setiap opsi ke satuan pembanding yang sama, misalnya biaya per sesi,
biaya per barang, waktu per unit, atau hasil per rupiah.

### Pattern of Question

```text
Dua opsi -> total biaya/manfaat berbeda -> tanya opsi yang lebih hemat/efisien
```

### Pattern to Answer

```text
nilai opsi A per unit -> nilai opsi B per unit -> bandingkan selisih
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Bandingkan total harga | Paket lebih mahal bisa lebih murah per sesi |
| Lupa satuan pembanding | Harus sama-sama per sesi |
| Salah selisih | Selisih dihitung setelah biaya per unit ditemukan |

### Review Rule

Tulis satuan di setiap angka:

```text
Rp/sesi
```

### Practice Question

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

## Num09 - Tabel Kasus

### The Problem

Soal tabel kasus sering berisi kategori yang tumpang tindih. Kalau tidak
ditabelkan, angka kategori, gabungan, dan irisan mudah tercampur.

### The Concept

Gunakan tabel atau rumus gabungan untuk memisahkan kategori:

```text
A atau B = A + B - (A dan B)
```

### Pattern of Question

```text
Total minimal satu kategori -> jumlah A -> jumlah B -> tanya keduanya
```

### Pattern to Answer

```text
gabungan = A + B - irisan -> irisan = A + B - gabungan
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Menjumlah A dan B langsung | Yang ikut keduanya terhitung dua kali |
| Lupa arti minimal satu | Minimal satu berarti A atau B |
| Salah tanda irisan | Irisan dikurangkan dari jumlah A+B |

### Review Rule

Tulis tiga kotak:

```text
hanya A | keduanya | hanya B
```

### Practice Question

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

## Num10 - Diagram Alur

### The Problem

Soal diagram alur meminta kamu mengikuti proses tahap demi tahap. Kesalahan
umum adalah memakai total awal terus, padahal setiap tahap memakai hasil tahap
sebelumnya.

### The Concept

Setiap tahap menghasilkan basis baru untuk tahap berikutnya.

### Pattern of Question

```text
Jumlah awal -> lolos tahap 1 -> ikut tahap 2 -> lolos tahap 3
```

### Pattern to Answer

```text
awal -> tahap 1 -> tahap 2 -> tahap 3 -> hasil akhir
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Semua persen dari awal | Tahap berikutnya memakai jumlah yang lolos sebelumnya |
| Lupa pecahan | `2/3` diterapkan ke hasil tahap administrasi |
| Tidak tulis alur | Angka tahap mudah tertukar |

### Review Rule

Buat panah proses:

```text
pendaftar -> administrasi -> tes tertulis -> wawancara
```

### Practice Question

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

## Num11 - Kecukupan Informasi Praktis

### The Problem

Soal kecukupan informasi praktis menguji apakah kamu tahu data mana yang
diperlukan untuk menjawab pertanyaan. Angka yang terlihat cukup belum tentu
menutup semua kemungkinan.

### The Concept

Informasi cukup jika semua bagian yang memengaruhi jawaban sudah diketahui.
Untuk target mingguan, semua hari yang mungkin berproduksi harus tertutup.

### Pattern of Question

```text
Target -> beberapa pernyataan data -> tanya pernyataan mana yang diperlukan
```

### Pattern to Answer

```text
tentukan komponen jawaban -> cek pernyataan yang menutup setiap komponen -> pilih paket info cukup
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Abaikan hari kosong | Tanpa info Minggu, masih ada kemungkinan produksi |
| Memilih data yang tampak dominan | Senin-Jumat saja belum menutup minggu |
| Tidak cocokkan ke pertanyaan | Yang ditanya target mingguan, bukan harian |

### Review Rule

Tulis:

```text
Untuk menjawab, aku perlu tahu ...
```

### Practice Question

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

## Num12 - Eliminasi Opsi

### The Problem

Soal eliminasi opsi bisa diselesaikan lebih cepat dengan membuang pilihan yang
melanggar syarat. Jangan langsung hitung semua opsi secara penuh kalau struktur
soalnya memberi batasan jelas.

### The Concept

Gunakan syarat soal untuk mengurangi ruang pencarian. Pada deret kursi, syarat
"selisih antarbaris selalu sama" berarti barisan aritmetika.

### Pattern of Question

```text
Beberapa syarat -> opsi jawaban -> cari opsi yang memenuhi semua batasan
```

### Pattern to Answer

```text
ubah syarat jadi model -> coret opsi tidak mungkin -> hitung opsi tersisa
```

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Tidak memakai syarat selisih tetap | Modelnya jadi bebas dan terlalu luas |
| Lupa jumlah baris | Ada 6 baris, bukan 5 |
| Salah jumlah deret | Total harus 57 |

### Review Rule

Setiap jawaban harus lolos semua syarat:

```text
jumlah baris, total kursi, dan selisih tetap
```

### Practice Question

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
