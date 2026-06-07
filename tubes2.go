package main
import "fmt"
 
const NMAX = 100
const SMAX = 500
 
type tabWarga [NMAX]Warga
type tabSetoran [SMAX]Setoran
 
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
 
var dataWarga tabWarga
var dataSetoran tabSetoran
var jumlahWarga, jumlahSetoran, idAuto int
 
func tambahWarga(data *tabWarga, n *int) {
	var nama, alamat string
	fmt.Print("Nama   : ")
	fmt.Scan(&nama)
	fmt.Print("Alamat : ")
	fmt.Scan(&alamat)
	idAuto = idAuto + 1
	data[*n].id = idAuto
	data[*n].nama = nama
	data[*n].alamat = alamat
	data[*n].aktif = true
	*n = *n + 1
	fmt.Println("ID:", idAuto)
}

func cariIdx(id int) int {
	var i int
	for i = 0; i < jumlahWarga; i++ {
		if dataWarga[i].id == id && dataWarga[i].aktif {
			return i
		}
	}
	return -1
} 
 
func ubahWarga(data *tabWarga, n int) {
	var id, idx int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)
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
	data[idx].nama = nama
	data[idx].alamat = alamat
	fmt.Println("Data berhasil diubah.")
}
 
func hapusWarga(data *tabWarga, n int) {
	var id, idx int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)
	idx = cariIdx(id)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}
	data[idx].aktif = false
	fmt.Println("Warga berhasil dihapus.")
}
 
func lihatWarga(data tabWarga, n int) {
	var i int
	var ada bool
	for i = 0; i < n; i++ {
		if data[i].aktif {
			fmt.Printf("ID: %d | Nama: %-15s | Alamat: %s\n",
				data[i].id, data[i].nama, data[i].alamat)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Belum ada data warga.")
	}
}
 
func catatSetoran(data *tabSetoran, n *int) {
	var id, idx int
	fmt.Print("ID warga : ")
	fmt.Scan(&id)
	idx = cariIdx(id)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}
	var tanggal, jenis string
	var berat float64
	fmt.Print("Tanggal (DD-MM-YYYY) : ")
	fmt.Scan(&tanggal)
	fmt.Print("Jenis sampah         : ")
	fmt.Scan(&jenis)
	fmt.Print("Berat (kg)           : ")
	fmt.Scan(&berat)
	data[*n].idWarga = id
	data[*n].tanggal = tanggal
	data[*n].jenis = jenis
	data[*n].berat = berat
	*n = *n + 1
}
 
func lihatSetoran(data tabSetoran, n int) {
	if n == 0 {
		fmt.Println("Belum ada data setoran.")
		return
	}
	var i, idx int
	var nama string
	for i = 0; i < n; i++ {
		idx = cariIdx(data[i].idWarga)
		if idx != -1 {
			nama = dataWarga[idx].nama
		} else {
			nama = "-"
		}
		fmt.Printf("[%d] %-15s | %s | %-10s | %.2f kg\n",
			i+1, nama, data[i].tanggal, data[i].jenis, data[i].berat)
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
			tambahWarga(&dataWarga, &jumlahWarga)
		} else if pilih == 2 {
			ubahWarga(&dataWarga, jumlahWarga)
		} else if pilih == 3 {
			hapusWarga(&dataWarga, jumlahWarga)
		} else if pilih == 4 {
			lihatWarga(dataWarga, jumlahWarga)
		} else if pilih == 5 {
			catatSetoran(&dataSetoran, &jumlahSetoran)
		} else if pilih == 6 {
			lihatSetoran(dataSetoran, jumlahSetoran)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
 