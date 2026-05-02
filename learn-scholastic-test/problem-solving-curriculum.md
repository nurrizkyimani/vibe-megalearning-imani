# M4/W4/D31 - Fri, 01 May 2026 (WIB)

## Scholastic Pemecahan Masalah - Curriculum

Kurikulum singkat untuk membangun kemampuan Pemecahan Masalah Tes Skolastik
LPDP: mulai dari membaca data survei, irisan kelompok, target-sisa pekerjaan,
optimasi sederhana, jadwal, sampai eliminasi opsi. Fokus bagian ini bukan
sekadar menghitung, tetapi memilih model yang tepat dari informasi cerita.

Setiap topik dipetakan ke soal latihan di `problem-solving-questions.md`.

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

## Detailed Topic Offset

Detail penjelasan per nomor saat ini sudah dibuat sampai:

```text
01 -> 03
```

Nomor berikutnya yang belum dibedah detail adalah `04 - Alokasi Sumber Daya`.

---

## Num01 - Himpunan dan Irisan Data Survei

### The Problem

Soal survei sering terlihat seperti soal persentase biasa, padahal jebakan
utamanya ada di kata **dan**, **atau**, **hanya**, dan **sisanya**. Angka irisan
tidak boleh dikurangkan dari total yang salah.

Contoh:

```text
Aktif organisasi = 180
IPK >= 3,5 = 150
Aktif dan IPK >= 3,5 = 90
```

Angka `90` adalah irisan. Artinya, `90` orang itu sudah termasuk di dalam
kelompok aktif dan juga di dalam kelompok IPK.

### The Concept

Rumus cepat:

```text
Hanya A = A - (A dan B)
Hanya B = B - (A dan B)
A atau B = A + B - (A dan B)
Tidak keduanya = Total - (A atau B)
```

Untuk soal seperti IPK dan organisasi:

```text
Hanya IPK = IPK - (IPK dan aktif)
```

Jangan mengurangkan `IPK - aktif`, karena aktif adalah total kelompok lain,
bukan bagian yang harus dikeluarkan dari IPK.

### What the practice shows

Soal terkait: nomor `01 - Himpunan dan Irisan Data Survei` di
`problem-solving-questions.md`.

Nomor 01 melatih membedakan total kelompok dan irisan. Yang dikurang dari
kelompok IPK adalah bagian yang juga aktif, bukan total semua mahasiswa aktif.

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Mengurangkan total kelompok lain | `150 - 180` tidak bermakna untuk mencari hanya IPK |
| Lupa bahwa irisan masuk dua kelompok | Angka "dan" dihitung di kedua kelompok |
| Salah baca "sisanya" | Sisa berarti tidak memenuhi keduanya, bukan hanya tidak aktif |

### Review Rule

Sebelum menghitung, tulis dulu:

```text
Total:
A:
B:
A dan B:
Yang ditanya:
```

Kalau yang ditanya memakai kata **hanya**, kurangi kelompok itu dengan
irisannya.

---

## Num02 - Target dan Sisa Pekerjaan

### The Problem

Soal target sering menjebak karena peserta langsung membagi total dengan waktu
baru. Padahal sebagian pekerjaan sudah selesai, sehingga yang dibagi adalah
sisa target, bukan target awal.

Contoh:

```text
Target = 60 dokumen
Sudah selesai = 18 dokumen
Sisa hari = 3
```

Yang dicari adalah rata-rata untuk sisa pekerjaan:

```text
(60 - 18) / 3
```

### The Concept

Pola cepat:

```text
Sisa target = target akhir - progres yang sudah dicapai
Beban per periode = sisa target / sisa periode
```

Pastikan satuan periode sama: hari, minggu, jam, atau sesi.

### What the practice shows

Soal terkait: nomor `02 - Target dan Sisa Pekerjaan` di
`problem-solving-questions.md`.

Nomor 02 melatih mencari target harian baru setelah progres awal diketahui.
Fokusnya adalah tidak membagi target awal dengan sisa hari.

### Common Traps

| Jebakan | Contoh |
| --- | --- |
| Membagi total awal | `60 / 3` padahal 18 sudah selesai |
| Lupa sisa waktu | memakai 5 hari, bukan 3 hari |
| Salah satuan | mencampur minggu dan hari tanpa konversi |

### Review Rule

Tulis kalimat ini sebelum hitung:

```text
Yang tersisa adalah ...
```

Kalau belum jelas yang tersisa apa, modelnya belum aman.

---

## Num03 - Perubahan Bertahap

### The Problem

Perubahan bertahap sering salah karena peserta memakai nilai awal untuk semua
tahap. Padahal setelah tahap pertama, nilai dasar sudah berubah.

Contoh:

```text
Kuota awal = 200
Tahap 1 turun 10%
Tahap 2 naik 25% dari hasil tahap 1
```

Tahap kedua harus memakai hasil setelah turun 10%, bukan `200` lagi.

### The Concept

Pola cepat:

```text
Nilai baru = nilai sekarang x faktor perubahan
```

Faktor perubahan:

```text
turun p% -> x (1 - p/100)
naik p% -> x (1 + p/100)
```

Setelah tiap tahap, simpan nilai terbaru sebelum lanjut ke tahap berikutnya.

### What the practice shows

Soal terkait: nomor `03 - Perubahan Bertahap` di
`problem-solving-questions.md`.

Nomor 03 melatih membaca frasa "dari sisa", "dari hasil tahap sebelumnya", atau
"dari jumlah terbaru" sebagai tanda bahwa basis hitung berubah.

### Common Traps

| Jebakan | Kenapa Salah |
| --- | --- |
| Semua persen dihitung dari awal | Basis tahap kedua sudah berubah |
| Menjumlah persen mentah | Turun 10% lalu naik 25% bukan otomatis naik 15% |
| Tidak simpan nilai antara | Sulit mengecek tahap mana yang salah |

### Review Rule

Buat tabel kecil:

| Tahap | Operasi | Hasil |
| --- | --- | --- |
| Awal | - | ... |
| 1 | ... | ... |
| 2 | ... | ... |

---

## Num04-Num12 - Strategi Model dan Eliminasi

### The Problem

Nomor 04-12 biasanya tidak sulit karena rumusnya, tetapi karena informasi
soalnya ramai. Ada angka kapasitas, batas waktu, prioritas, syarat minimal,
atau pilihan yang mirip-mirip. Kalau semua angka langsung dihitung, waktu habis
sebelum modelnya jelas.

### The Concept

Untuk nomor 04-12, fokusnya adalah memilih struktur penyelesaian:

1. Alokasi sumber daya memakai batas dan prioritas.
2. Optimasi sederhana mencari pilihan minimum atau maksimum yang memenuhi syarat.
3. Jadwal dan urutan langkah harus disusun kronologis.
4. Laju gabungan harus disamakan satuannya; untuk variasi lanjutan, cari sisa
   pekerjaan lalu bagi dengan laju pihak yang tersisa.
5. Perbandingan keputusan harus memakai ukuran pembanding yang sama.
6. Tabel kasus membantu memisahkan kategori.
7. Diagram alur membantu mengikuti perubahan status.
8. Kecukupan informasi melatih memilih data yang relevan.
9. Eliminasi opsi mempercepat soal yang banyak batasan.

### What the practice shows

1. Nomor 04 melatih alokasi kuota dengan batas minimal.
2. Nomor 05 melatih biaya minimum dari beberapa kombinasi paket.
3. Nomor 06 melatih jadwal dengan prasyarat.
4. Nomor 07 melatih laju kerja gabungan, lalu latihan tambahannya melatih kerja
   bareng dulu, mencari sisa pekerjaan, dan membagi sisa itu dengan laju pekerja
   yang masih lanjut.
5. Nomor 08 melatih membandingkan dua skenario keputusan.
6. Nomor 09 melatih tabel kategori.
7. Nomor 10 melatih proses bertahap.
8. Nomor 11 melatih memilih informasi yang cukup.
9. Nomor 12 melatih eliminasi opsi.

### Review rule

Untuk setiap soal nomor 04-12, tulis dulu:

```text
Tujuan:
Batas/syarat:
Model:
Opsi yang gugur:
```

Kalau tujuan dan batas belum jelas, jangan mulai hitung detail.

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
- Review wajib menulis letak salah: salah baca kata kunci, salah model, salah
  satuan, salah eliminasi, atau salah hitung.
