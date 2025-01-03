// TUGAS BESAR ALPRO 2 KELOMPOK 10 S1IF-11-01
// DIVA OCTAVIANI      (2311102006)
// ARJUN WERDHO KUMORO (2311102009)
// RICO ADE PRATAMA    (2311102138)
// PROGRAM JUAL BELI BARANG UNTUK PEGAWAI TOKO
Program Jual Beli
Kamus
type Item <
ID : int
Name : string
Harga_Beli : float
Harga_Jual : float
Stock : int
> 
type Transaction <
ItemID : int
ItemName : string
Quantity : int
TotalPrice : float
SaleDate : time.Time
>
type History <
Action : string
Category : string
ItemID : int
Details : string
Date : time.Time
>
global items : array of Item
global transactions : array of Transaction
global riwayat : array of History
global historyCount : integer

//Fungsi untuk menambahkan data barang
program addItem
kamus
    ID : integer
    Nama : string
    Harga_Beli, Harga_Jual : real
    Stock : integer
    i : integer
algoritma
// Loop melalui array items untuk mencari slot kosong
    for each i in items do
// Validasi input - memastikan nilai harga dan stok positif
        if Harga_Beli <= 0 or Harga_Jual <= 0 or Stock <= 0 then
            output("## Input Salah!! Harga beli, harga jual, dan stok harus lebih besar dari 0.")
            return
        endif
// Menambahkan item baru ke array items pada indeks i
        items[i] <- Item(ID, Nama, Harga_Beli, Harga_Jual, Stock)
// Mencatat ke history penambahan barang
        riwayat[historyCount] <- History(Action: "add", Category: "Data", ItemID: ID,
            Details: format("Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d", Nama, Harga_Beli, Harga_Jual, Stock),
            Date: current time)
        historyCount <- historyCount + 1
// Menampilkan konfirmasi penambahan berhasil
        output("- %s (ID: %d, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d) pada %s",
            Nama, ID, Harga_Beli, Harga_Jual, Stock, current time)
        output("Data Barang berhasil ditambahkan.")
        return
    endfor
// Jika loop selesai tanpa return, berarti array penuh
    output("## Tidak dapat menambahkan, kapasitas array penuh.")
endprogram


// Fungsi untuk mengubah (edit) Data Barang
program editItem
kamus
    ID : integer
    Nama_Baru : string
    Harga_Beli_Baru, Harga_Jual_Baru : real
    Stock_Baru : integer
    i : integer
    found : boolean
algoritma
// Loop mencari item dengan ID yang sesuai
    found <- false
    for each i in items do
        if items[i].ID = ID then
            found <- true
// Validasi input: memastikan semua nilai numerik > 0
            if Harga_Beli_Baru <= 0 or Harga_Jual_Baru <= 0 or Stock_Baru <= 0 then
                output("## Input Salah!! Harga beli, harga jual, dan stok harus lebih besar dari 0.")
                return
            endif
// Update data barang dengan informasi baru
            items[i] <- Item(ID, Nama_Baru, Harga_Beli_Baru, Harga_Jual_Baru, Stock_Baru)
// Catat riwayat perubahan data
            riwayat[historyCount] <- History(Action: "edit", Category: "Data", ItemID: ID,
                Details: format("Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d", Nama_Baru, Harga_Beli_Baru, Harga_Jual_Baru, Stock_Baru),
                Date: current time)
// Increment jumlah riwayat
            historyCount <- historyCount + 1
// Tampilan pesan bahwa data berhasil diubah
            output("Data Barang berhasil diubah.")
            return
        endif
    endfor
 // Tampilan pesan jika barang tidak ditemukan
    if not found then
        output("## Tidak ditemukan data barang dengan ID: %d.", ID)
    endif
endprogram


//  Fungsi untuk menghapus Data Barang
program deleteItem
kamus
    ID : integer
    i : integer
    found : boolean
algoritma
// Inisialisasi variabel found sebagai false
    found <- false
// Loop mencari item dengan ID yang sesuai
    for each i in items do
// Cek apakah ID barang pada iterasi saat ini sama dengan ID yang dicari
        if items[i].ID = ID then
// Jika ditemukan, set found menjadi true
            found <- true
// Catat riwayat penghapusan data
            riwayat[historyCount] <- History(Action: "delete", Category: "Data", ItemID: ID,
                Details: format("Nama: %s", items[i].Nama),
                Date: current time)
// Increment jumlah riwayat
            historyCount <- historyCount + 1
// Hapus item dari daftar barang
            remove items[i]
// Tampilan jika barang berhasil dihapus
            output("Data Barang dengan ID %d berhasil dihapus.", ID)
            return
        endif
    endfor
// Tampilan jika barang tidak ditemukan
    if not found then
        output("## Tidak ditemukan data barang dengan ID: %d.", ID)
    endif
endprogram


//Fungsi untuk menambahkan transaksi penjualan
procedure addTransaction(itemID, quantity)
    // Cek apakah jumlah barang yang dijual lebih besar dari 0
    if quantity <= 0 then
        print("## Input Salah!! Jumlah Barang yang dijual harus lebih besar dari 0.")
        return
    endif

    // Cari item berdasarkan itemID
    for i <- 0 to jumlahBarang - 1 do
        if items[i].ID == itemID then
            // Cek apakah stok cukup
            if items[i].Stock < quantity then
                print("## Maaf, Stok tidak mencukupi!")
                return
            endif

            // Hitung total harga transaksi
            totalPrice <- items[i].Price * quantity

            // Buat transaksi baru
            transaksi[transactionCount] <- Transaction{
                ItemID:     items[i].ID,
                ItemName:   items[i].Name,
                Quantity:   quantity,
                TotalPrice: totalPrice,
                SaleDate:   currentDate(),
            }

            // Update stok barang
            items[i].Stock <- items[i].Stock - quantity

            // Update jumlah transaksi
            transactionCount <- transactionCount + 1

            print("Transaksi berhasil ditambahkan.")
            return
        endif
    endfor

    print("## Item dengan ID tersebut tidak ditemukan.")
endprocedure





// Fungsi untuk mengubah transaksi penjualan
procedure editTransaksiPenjualan()
    input(indexTransaksi)  // Mengambil index transaksi yang ingin diedit
    
    if indexTransaksi >= 0 and indexTransaksi < jumlahTransaksi then
        // Jika index transaksi valid, lanjutkan untuk edit transaksi
        print("Transaksi ditemukan, silakan input jumlah barang yang terjual baru")

        // Input jumlah baru barang yang terjual
        input(jumlahBarangBaru)

        // Cari barang yang sesuai dengan transaksi
        for i <- 0 to jumlahBarang - 1 do
            if items[i].ID == daftarTransaksi[indexTransaksi].ItemID then
                // Hitung perubahan jumlah barang
                deltaQuantity <- jumlahBarangBaru - daftarTransaksi[indexTransaksi].Quantity
                
                // Jika perubahan jumlah barang lebih besar, cek apakah stok mencukupi
                if deltaQuantity > 0 and items[i].Stock < deltaQuantity then
                    print("## Data transaksi tidak bisa diedit, stok tidak mencukupi.")
                    return
                endif

                // Simpan detail lama transaksi
                oldDetails <- "Jumlah: " + daftarTransaksi[indexTransaksi].Quantity + ", Total Harga: Rp" + daftarTransaksi[indexTransaksi].TotalPrice

                // Update stok barang dengan perubahan jumlah
                items[i].Stock <- items[i].Stock + daftarTransaksi[indexTransaksi].Quantity - jumlahBarangBaru
                
                // Update jumlah barang yang terjual pada transaksi yang dipilih
                daftarTransaksi[indexTransaksi].Quantity <- jumlahBarangBaru

                // Update total harga sesuai dengan jumlah barang baru
                daftarTransaksi[indexTransaksi].TotalPrice <- jumlahBarangBaru * items[i].Price

                // Menyimpan perubahan transaksi
                print("Jumlah barang yang terjual berhasil diubah")
                return
            endif
        endfor

    else
        print("Index transaksi tidak valid")
    endif
endprocedure




//Fungsi untuk menghapus transaksi penjualan
procedure deleteTransaction()
    input(indexTransaksi)  // Mengambil index transaksi yang ingin dihapus
    if indexTransaksi >= 0 and indexTransaksi < jumlahTransaksi then
        // Jika index transaksi valid, lanjutkan untuk menghapus transaksi
        
        // Geser transaksi setelah transaksi yang dihapus
        for j <- indexTransaksi to jumlahTransaksi-2 do
            daftarTransaksi[j] <- daftarTransaksi[j+1]  // Geser data transaksi setelah index yang dihapus
        endfor
        
        jumlahTransaksi--  // Kurangi jumlah transaksi
        print("Transaksi berhasil dihapus")
    else
        print("Index transaksi tidak valid")
    endif
endprocedure



//Fungsi untuk menampilkan Data Barang yang tersedia
program displayItem
kamus
    i : integer
algoritma
    // Cek apakah daftar barang kosong
    if items is empty then
        output("## Tidak ada barang untuk ditampilkan.")
        return
    endif
    output("Data Barang:")
    // Loop untuk menampilkan setiap item dalam daftar barang
    for each i in items do
// Tampilkan informasi tentang setiap barang
        output("- ID: %d, Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d", 
            items[i].ID, items[i].Nama, items[i].Harga_Beli, items[i].Harga_Jual, items[i].Stock)
    endfor
endprogram

// Fungsi untuk mengurutkan Data Barang berdasarkan kategori
PROGRAM sortItems(by: STRING, order: STRING)
{ I.S.: Array items berisi data barang dengan atribut ID, Nama, Harga_Beli, Harga_Jual, dan Stock }
{ F.S.: Array items terurut berdasarkan kriteria dan urutan yang dipilih }

KAMUS:
    n, i, j, minOrMax : FLOAT64     
    temp : RECORD { ID, Nama, Harga_Beli, Harga_Jual, Stock }

ALGORITMA:
    n <- LENGTH(items) 

    SWITCH by DO // Memilih atribut untuk pengurutan
        CASE "ID": // Jika ingin mengurutkan berdasarkan ID
            FOR i <- 1 TO n-1 DO 
                minOrMax <- i //
                FOR j <- i+1 TO n DO 
                    IF (order = "asc" AND items[j].ID < items[minOrMax].ID) OR
                       (order = "desc" AND items[j].ID > items[minOrMax].ID) THEN
                        minOrMax <- j // Perbarui indeks elemen minimum/maksimum
                    END IF
                END FOR
                // Tukar elemen
                temp <- items[i] 
                items[i] <- items[minOrMax] ke posisi i 
                items[minOrMax] <- temp // Pindahkan elemen i ke posisi
            END FOR

        CASE "Nama": // Jika ingin mengurutkan berdasarkan Nama
            FOR i <- 1 TO n-1 DO
                minOrMax <- i
                FOR j <- i+1 TO n DO
                    IF (order = "asc" AND items[j].Nama < items[minOrMax].Nama) OR
                       (order = "desc" AND items[j].Nama > items[minOrMax].Nama) THEN
                        minOrMax <- j
                    END IF
                END FOR
                // Tukar elemen
                temp <- items[i]
                items[i] <- items[minOrMax]
                items[minOrMax] <- temp
            END FOR

        CASE "Harga_Beli": // Jika ingin mengurutkan berdasarkan Harga_Beli
            FOR i <- 1 TO n-1 DO
                minOrMax <- i
                FOR j <- i+1 TO n DO
                    IF (order = "asc" AND items[j].Harga_Beli < items[minOrMax].Harga_Beli) OR
                       (order = "desc" AND items[j].Harga_Beli > items[minOrMax].Harga_Beli) THEN
                        minOrMax <- j
                    END IF
                END FOR
                // Tukar elemen
                temp <- items[i]
                items[i] <- items[minOrMax]
                items[minOrMax] <- temp
            END FOR

        CASE "Harga_Jual": // Jika ingin mengurutkan berdasarkan Harga_Jual
            FOR i <- 1 TO n-1 DO
                minOrMax <- i
                FOR j <- i+1 TO n DO
                    IF (order = "asc" AND items[j].Harga_Jual < items[minOrMax].Harga_Jual) OR
                       (order = "desc" AND items[j].Harga_Jual > items[minOrMax].Harga_Jual) THEN
                        minOrMax <- j
                    END IF
                END FOR
                // Tukar elemen
                temp <- items[i]
                items[i] <- items[minOrMax]
                items[minOrMax] <- temp
            END FOR

        CASE "Stock": // Jika ingin mengurutkan berdasarkan Stock
            FOR i <- 1 TO n-1 DO
                minOrMax <- i
                FOR j <- i+1 TO n DO
                    IF (order = "asc" AND items[j].Stock < items[minOrMax].Stock) OR
                       (order = "desc" AND items[j].Stock > items[minOrMax].Stock) THEN
                        minOrMax <- j
                    END IF
                END FOR
                // Tukar elemen
                temp <- items[i]
                items[i] <- items[minOrMax]
                items[minOrMax] <- temp
            END FOR

        DEFAULT: // Jika kriteria pengurutan tidak valid
            PRINT "Kriteria urutkan tidak valid." // Pesan kesalahan
            RETURN // Keluar dari program
    END SWITCH

// Fungsi untuk mencari Data Barang berdasarkan kata kunci
PROGRAM searchItems(keyword : STRING)
{ I.S.: Array items berisi data barang dengan atribut ID, Nama, Harga_Beli, Harga_Jual, dan Stock }
{ F.S.: Menampilkan data barang yang sesuai dengan kata kunci }

KAMUS:
    found : BOOLEAN
    i : FLOAT64
ALGORITMA:
    // Tampilkan header tabel hasil pencarian
    PRINT "Hasil Pencarian:"
    PRINT "+-----+--------------------+------------------+------------------+-----------------+"
    PRINT "| ID  | Nama Barang        | Harga Beli (Rp)  | Harga Jual (Rp)  | Stok Tersedia   |"
    PRINT "+-----+--------------------+------------------+------------------+-----------------+"
    
    found <- FALSE // Inisialisasi tidak ada data yang ditemukan

    // Iterasi untuk mencari data yang sesuai dengan kata kunci
    FOR i <- 1 TO LENGTH(items) DO
        // Cek apakah atribut barang mengandung kata kunci (case-insensitive)
        IF CONTAINS(LOWERCASE(STRING(items[i].ID)), LOWERCASE(keyword)) OR
           CONTAINS(LOWERCASE(items[i].Nama), LOWERCASE(keyword)) OR
           CONTAINS(LOWERCASE(STRING(items[i].Harga_Beli)), LOWERCASE(keyword)) OR
           CONTAINS(LOWERCASE(STRING(items[i].Harga_Jual)), LOWERCASE(keyword)) OR
           CONTAINS(LOWERCASE(STRING(items[i].Stock)), LOWERCASE(keyword)) THEN

            // Tampilkan data barang yang sesuai
            PRINT "| ", items[i].ID, " | ", items[i].Nama, " | ",
                  FORMAT_CURRENCY(items[i].Harga_Beli), " | ",
                  FORMAT_CURRENCY(items[i].Harga_Jual), " | ",
                  items[i].Stock, " |"
            
            found <- TRUE // Tandai bahwa data ditemukan
        END IF
    END FOR

    // Jika tidak ada data yang ditemukan
    IF found = FALSE THEN
        PRINT "|                     Tidak ada Data yang sesuai.                                  |"
    END IF

    // Tampilkan footer tabel
    PRINT "+-----+--------------------+------------------+------------------+-----------------+"
END PROGRAM

// Fungsi untuk menampilkan laporan modal dan pendapatan
PROGRAM displayReport()
{ I.S.: Array transactions berisi data transaksi, Array items berisi data barang }
{ F.S.: Menampilkan Total Modal, Total Pendapatan, dan Total Keuntungan }

KAMUS:
    totalHarga_Beli, totalHarga_Jual, totalProfit : FLOAT64
    i, j : INTEGER

ALGORITMA:
    totalHarga_Beli <- 0
    totalHarga_Jual <- 0
    totalProfit <- 0

    FOR i <- 1 TO LENGTH(transactions) DO
        FOR j <- 1 TO LENGTH(items) DO
            IF items[j].ID = transactions[i].ItemID THEN
// Modal
                totalHarga_Beli <- totalHarga_Beli + (transactions[i].Quantity * items[j].Harga_Beli)
// Pendapatan
                totalHarga_Jual <- totalHarga_Jual + (transactions[i].Quantity * items[j].Harga_Jual)
// Keuntungan                totalProfit <- totalProfit + ((items[j].Harga_Jual - items[j].Harga_Beli) * transactions[i].Quantity)
            END IF
        END FOR
    END FOR

    PRINT "Total Modal      : Rp ", FORMAT_CURRENCY(totalHarga_Beli)
    PRINT "Total Pendapatan : Rp ", FORMAT_CURRENCY(totalHarga_Jual)
    PRINT "Total Keuntungan : Rp ", FORMAT_CURRENCY(totalProfit)
END PROGRAM

// Fungsi untuk menampilkan 5 Barang yang paling banyak terjual
PROCEDURE displayTopSoldItems()
{ I.S.: Array transactions berisi data transaksi, Array items berisi data barang }
{ F.S.: Menampilkan 5 barang dengan jumlah penjualan terbanyak }

KAMUS:
    TYPE SoldItem IS RECORD
        ItemID : FLOAT64         
        TotalSold : FLOAT64 
    END RECORD

    soldItemsMap : MAP<FLOAT64, FLOAT64>
    soldItems : ARRAY OF SoldItem
    i, j : FLOAT64 


ALGORITMA:
    // Menghitung total barang yang terjual berdasarkan ItemID
    FOR EACH transaction IN transactions DO
        soldItemsMap[transaction.ItemID] <- soldItemsMap[transaction.ItemID] + transaction.Quantity
    END FOR

    // Memindahkan data dari Map ke Array untuk diurutkan
    FOR EACH itemID, totalSold IN soldItemsMap DO
        APPEND(SoldItem(itemID, totalSold)) TO soldItems
    END FOR

    // Mengurutkan array berdasarkan TotalSold secara menurun (descending)
    FOR i <- 1 TO LENGTH(soldItems) - 1 DO
        FOR j <- i + 1 TO LENGTH(soldItems) DO
// Jika jumlah barang pada indeks i lebih kecil daripada pada indeks j, tukar posisinya
            IF soldItems[i].TotalSold < soldItems[j].TotalSold THEN
                SWAP(soldItems[i], soldItems[j]) // Menukar elemen array
            END IF
        END FOR
    END FOR

    PRINT "Total Modal      : Rp ", FORMAT(totalHarga_Beli, ".2f")
    PRINT "Total Pendapatan : Rp ", FORMAT(totalHarga_Jual, ".2f")
    PRINT "Total Keuntungan : Rp ", FORMAT(totalProfit, ".2f")
END PROGRAM

// Fungsi untuk menampilkan Riwayat/History
PROGRAM showHistory()
{ I.S.: Array riwayat berisi data riwayat aksi/tindakan }
{ F.S.: Menampilkan semua riwayat tindakan yang telah dilakukan }
KAMUS
  i : FLOAT64
ALGORITMA
  PROGRAM showHistory()
    IF LENGTH(riwayat) = 0 THEN
        PRINT "## Tidak ada riwayat tindakan."
        RETURN
    END IF

    PRINT "### Riwayat Aksi:"
    FOR EACH history IN riwayat DO
        PRINT history.Date, " - ", history.Action, ": ", history.Category, 
              " (ID: ", history.ItemID, ", ", history.Details, ") pada ", history.Date
    END FOR
END PROGRAM

//Fungsi utama untuk menjalankan aplikasi
program aplikasiJualBeli
kamus
    pilihan : integer
    ID, qty, index : integer
    Nama, sortBy, order, keyword : string
    Harga_Beli, Harga_Jual : real
    Stock : integer
algoritma
    while true do
        output("===================================================================")
        output(">>>>>>>>>>> Menu Aplikasi Jual Beli Untuk Pegawai Toko <<<<<<<<<<<")
        output("==================================================================")
     output("1.  Tampilkan Data                       (•⊙ω⊙•)")
     output("2.  Tambah Data                          (＾∀＾)")
     output("3.  Edit Data                            (｡◕‿◕｡)")
     output("4.  Hapus Data                           (⇀‸↼‶)ᕗ")
     output("5.  Tambah Transaksi Penjualan           (๑>ᴗ<๑)")
     output("6.  Edit Transaksi Penjualan             ( ͡° ͜ʖ ͡°)")
     output("7.  Hapus Transaksi Penjualan            (╥﹏╥)")
     output("8.  Urutkan Data                         (•̀‿⊹ )")
     output("9.  Pencarian Data                       q(◉ᴥ◉)p")
     output("10. Lihat Laporan Modal dan Pendapatan   [̲̅$̲̅(̲̅5̲̅)̲̅$̲̅]")
     output("11. TOP 5 Data Terlaris                  (⌐▨_▨)")
     output("12. History Data dan Transaksi           (☞ﾟヮﾟ)☞")
     output("13. Keluar                               (✖╭╮✖)")
        output("===================================================================")
        output("Pilih menu (1-13): ")
        input(pilihan)

        case pilihan of
            1:
                call displayItems()
            2:
                output("Masukkan ID Barang   : ")
                input(ID)
                output("Masukkan nama Barang : ")
                input(Nama)
                output("Masukkan harga beli  : ")
                input(Harga_Beli)
                output("Masukkan harga jual  : ")
                input(Harga_Jual)
                output("Masukkan jumlah stok : ")
                input(Stock)
                call addItem(ID, Nama, Harga_Beli, Harga_Jual, Stock)
            3:
                output("Masukkan ID Barang yang ingin diedit : ")
                input(ID)
                output("Masukkan nama barang baru : ")
                input(Nama)
                output("Masukkan harga beli baru  : ")
                input(Harga_Beli)
                output("Masukkan harga jual baru  : ")
                input(Harga_Jual)
                output("Masukkan jumlah stok baru : ")
                input(Stock)
                call editItem(ID, Nama, Harga_Beli, Harga_Jual, Stock)
            4:
                output("Masukkan ID Barang yang ingin dihapus : ")
                input(ID)
                call deleteItem(ID)
            5:
                output("Masukkan ID Barang yang terjual : ")
                input(ID)
                output("Masukkan jumlah Barang yang terjual : ")
                input(qty)
                call addTransaction(ID, qty)
            6:
                output("Masukkan indeks transaksi yang ingin diedit : ")
                input(index)
                output("Masukkan jumlah baru Barang yang terjual : ")
                input(qty)
                call editTransaction(index, qty)
            7:
                output("Masukkan indeks transaksi yang ingin dihapus : ")
                input(index)
                call deleteTransaction(index)
            8:
                output("Urutkan berdasarkan (pilih: ID/Nama/Harga_Beli/Harga_Jual/Stock): ")
                input(sortBy)
                output("Urutkan (asc/desc): ")
                input(order)
                call sortItems(sortBy, order)
            9:
                output("Masukkan Kata Kunci pencarian : ")
                input(keyword)
                call searchItems(keyword)
            10:
                call displayReport()
            11:
                call displayTopSoldItems()
            12:
                call showHistory()
            13:
                output("Terima kasih telah menggunakan aplikasi ini! (^‿^)")
                break
            otherwise:
                output("Pilihan tidak valid.")
        endcase
    endwhile
endprogram