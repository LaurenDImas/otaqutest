TEST SCRAPPING http://115.85.80.33/test-scrapping/avail.html

- Buat Database mysql dengan nama **otaqu_test** jalankan script sql membuat table terdapat di folder database
- Sesuaikan konfigurasi koneksi database di file .env
- Jalankan server golang menggunakan script go run main.go

Buatlah API menggunakan Golang (GoFiber Framework) dengan ketentuan sebagai berikut : 

1. Buatlah script untuk mendapatkan list Hotel dari link di bawah ini http://115.85.80.33/test-scrapping/avail.html
=> script scrapping terdapat di folder handler/hotel.go -> di function CreateHotel

 2. Kemudian Parsing pada hasil scrapping di atas dan insert ke dalam table "hotel".
=> POST http://localhot:8000/hotel/create
Data sebelumnya akan ter truncate jika melakukan request create

 3. Buatlah satu endpoint untuk menampilkan list hotel tersebut dengan format response JSON. 
=> GET http://localhot:8000/hotel

4. Gunakan script sql di bawah ini untuk membuat table "hotel".
=> Terdapat di folder database/otaqutest.sql