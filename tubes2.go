package main
import "fmt"

type Warga struct {
	id     int
	nama   string
	alamat string
	aktif  bool
}

type Setoran struct {
	idWarga int
	tanggal string
	jenis   string
	berat   float64
}

var dataWarga [100]Warga
var dataSetoran [500]Setoran
var jumlahWarga, jumlahSetoran, idAuto int

func cariIdx(id int) int {
	var i int
	for i = 0; i < jumlahWarga; i++ {
		if dataWarga[i].id == id && dataWarga[i].aktif {
			return i
		}
	}
	return -1
}

func tambahWarga() {
	var nama, alamat string
	fmt.Print("Nama   : ")
	fmt.Scan(&nama)
	fmt.Print("Alamat : ")
	fmt.Scan(&alamat)

	idAuto = idAuto + 1
	dataWarga[jumlahWarga].id = idAuto
	dataWarga[jumlahWarga].nama = nama
	dataWarga[jumlahWarga].alamat = alamat
	dataWarga[jumlahWarga].aktif = true
	jumlahWarga = jumlahWarga + 1

	fmt.Println("ID:", idAuto)
}

func ubahWarga() {
	var id int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)

	var idx int
	idx = cariIdx(id)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	var nama, alamat string
	fmt.Print("Nama baru   : ")
	fmt.Scan(&nama)
	fmt.Print("Alamat baru : ")
	fmt.Scan(&alamat)

	dataWarga[idx].nama = nama
	dataWarga[idx].alamat = alamat
	fmt.Println("Data berhasil diubah.")
}

func hapusWarga() {
	var id int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)

	var idx int
	idx = cariIdx(id)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	dataWarga[idx].aktif = false
	fmt.Println("Warga berhasil dihapus.")
}

func lihatWarga() {
	var i int
	var ada bool
	for i = 0; i < jumlahWarga; i++ {
		if dataWarga[i].aktif {
			fmt.Printf("ID: %d | Nama: %-15s | Alamat: %s\n",
				dataWarga[i].id, dataWarga[i].nama, dataWarga[i].alamat)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Belum ada data warga.")
	}
}

func catatSetoran() {
	var id int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)

	var idx int
	idx = cariIdx(id)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	var tanggal, jenis string
	var berat float64
	fmt.Print("Tanggal              : ")
	fmt.Scan(&tanggal)
	fmt.Print("Jenis sampah         : ")
	fmt.Scan(&jenis)
	fmt.Print("Berat (kg)           : ")
	fmt.Scan(&berat)

	dataSetoran[jumlahSetoran].idWarga = id
	dataSetoran[jumlahSetoran].tanggal = tanggal
	dataSetoran[jumlahSetoran].jenis = jenis
	dataSetoran[jumlahSetoran].berat = berat
	jumlahSetoran = jumlahSetoran + 1
}

func lihatSetoran() {
	if jumlahSetoran == 0 {
		fmt.Println("Belum ada data setoran.")
		return
	}
	var i int
	for i = 0; i < jumlahSetoran; i++ {
		var idx int
		idx = cariIdx(dataSetoran[i].idWarga)
		var nama string
		if idx != -1 {
			nama = dataWarga[idx].nama
		} else {
			nama = "-"
		}
		fmt.Printf("[%d] %-15s | %s | %-10s | %.2f kg\n",
			i+1, nama, dataSetoran[i].tanggal, dataSetoran[i].jenis, dataSetoran[i].berat)
	}
}


func main() {
	var pilih int
	var jalan bool
	jalan = true

	for jalan {
		fmt.Println()
		fmt.Println("1. Tambah warga")
		fmt.Println("2. Ubah warga")
		fmt.Println("3. Hapus warga")
		fmt.Println("4. Lihat warga")
		fmt.Println("5. Catat setoran")
		fmt.Println("6. Lihat setoran")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		fmt.Println()

		if pilih == 1 {
			tambahWarga()
		} else if pilih == 2 {
			ubahWarga()
		} else if pilih == 3 {
			hapusWarga()
		} else if pilih == 4 {
			lihatWarga()
		} else if pilih == 5 {
			catatSetoran()
		} else if pilih == 6 {
			lihatSetoran()
		} else if pilih == 0 {
			jalan = false
			fmt.Println("Terima Kasih")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}