Health Care API

## RULES:
Untuk setiap member yang melakukan development, harap ikuti rules di bawah ini:
1. Clone terlebih dahulu repo ini menggunakan `git clone https://github.com/Health-Care-System/BackEnd-Golang.git` <br>
    atau `git clone git@github.com:Health-Care-System/BackEnd-Golang.git`<br>
    Kemudian `cd BackEnd-Golang`
2. Jalankan di terminal masing-masing:
```
git switch development
git pull origin development
go mod tidy
```
3. Setelah itu, kalian bisa buat branch baru: `git checkout -b feature/ABC`
6. Lakukan coding di branch masing2.
7. Setelah coding, jangan lupa jalankan `go mod tidy`
8. Pastikan untuk ditest terlebih dahulu. **Real Condition == Expected Condition**
9. Buat commit yang descriptive dan pastikan files changes tidak terlalu banyak agar mudah direview<br>
    Contoh: `git commit -m "add feature login for doctor"`<br>
    Jadi sesuaikan pesan commit sesuai fungsi yang telah dibuat.
10. Lakukan push ke branch masing2. <br>
    Contoh: `git push origin feature/login`
11. Buat PR (Pull Request), dan pastikan merge nya ke arah branch **development**
12. Jangan merge sendiri
13. Setiap mau coding, **pastikan jalankan langkah nomor 2** dan seterusnya.

