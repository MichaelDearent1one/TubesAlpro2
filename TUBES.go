package main

import "fmt"

type milestone struct {
	id              int
	NamaTugas       string
	Deskripsi       string
	TanggalMulai    string
	TanggalSelesai  string
	TanggalDeadline string
	progress        int
	Mingguke        int
	Kepentingan     int
	Selesai         bool
	Mood            Mood
	tugas           jadwalTugas
}

type jadwalTugas struct {
	jadwaltugas string
}

type Mood struct {
	SkorStres         int
	SkorMoodMilestone int
	SkorMoodTugas     int
	CatatanRasa       string
}

var TotalData [1000]milestone
var JumlahData int
var ID int = 1

func main() {
	for {
		var pilihan int
		fmt.Println("\n========== SELAMAT DATANG DI MINDSTONE ==========")
		fmt.Println("Aplikasi Pemantau Capaian Kerja & Kesehatan Mental Anda")
		fmt.Println("Silahkan pilih opsi yang anda inginkan : ")
		fmt.Println("1. Tambah Milestone")
		fmt.Println("2. Hapus Milestone")
		fmt.Println("3. Ubah Milestone")
		fmt.Println("4. Cari Milestone")
		fmt.Println("5. Lihat Semua Milestone")
		fmt.Println("6. Urutkan Milestone")
		fmt.Println("7. Statistik")
		fmt.Println("8. Jadwal Tugas")
		fmt.Println("0. Keluar")
		fmt.Println("===================================================")

		fmt.Println("Tentukan pilihan anda : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			Tambah()
		case 2:
			Hapus()
		case 3:
			Ubah()
		case 4:
			Cari()
		case 5:
			Tampilkan()
		case 6:
			Urut()
		case 7:
			Statistik()
		case 8:
			JadwalTugas()
		case 0:
			fmt.Println("Sampai jumpa! Jaga kesehatan mentalmu selalu ya!!")
			return
		default:
			fmt.Println("Pilihan anda tidak valid")
			fmt.Println("Silahkan masukan pilihan kembali dengan benar ya!!")

		}
	}
}

func Tambah() {
	var pilihan int
	fmt.Println("\n========== TAMBAH MILESTONE ==========")
	fmt.Println("1. Tambah Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		TambahMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func TambahMilestone() {
	var m milestone

	fmt.Println("Masukan Nama Kegiatan : ")
	fmt.Scan(&m.NamaTugas)
	fmt.Println("Masukan Deskripsi : ")
	fmt.Scan(&m.Deskripsi)
	fmt.Println("Masukan Tanggal Kegiatan Dimulai (YYYY/MM/DD): ")
	fmt.Scan(&m.TanggalMulai)
	fmt.Println("Masukan Tanggal Deadline (YYYY/MM/DD) : ")
	fmt.Scan(&m.TanggalDeadline)
	fmt.Println("Minggu ke- : ")
	fmt.Scan(&m.Mingguke)
	fmt.Println("Progress tugas anda saat ini (0-100) : ")
	fmt.Scan(&m.progress)
	if m.progress < 0 && m.progress > 100 {
		fmt.Println("Masukan tidak valid")
		fmt.Println("Masukan kembali pilihan dengan benar!!")
		return
	}
	fmt.Println("Tingkat Kepentingan (1-10): ")
	fmt.Scan(&m.Kepentingan)
	if m.Kepentingan < 1 || m.Kepentingan > 10 {
		fmt.Println("Masukan tidak valid")
		fmt.Println("Masukan kembali pilihan dengan benar!!")
		return
	}
	fmt.Println("Tingkat Stres Anda (1-10): ")
	fmt.Scan(&m.Mood.SkorStres)
	if m.Mood.SkorStres == 10 {
		fmt.Println("Sekedar mengingatkan saja, segera pergi ke psikolog ya")
	} else if m.Mood.SkorStres == 1 {
		fmt.Println(" Sepertinya anda memiliki sedikit beban ya!")
	}
	if m.Mood.SkorStres < 1 || m.Mood.SkorStres > 10 {
		fmt.Println("Masukan tidak valid")
		fmt.Println("Masukan kembali pilihan dengan benar!")
		return
	}
	fmt.Println("Tingkat Mood Anda (1-10) : ")
	fmt.Scan(&m.Mood.SkorMoodMilestone)
	if m.Mood.SkorMoodMilestone == 10 {
		fmt.Println(" Sepertinya suasana hati anda sangat baik ya")
	} else if m.Mood.SkorMoodMilestone == 1 {
		fmt.Println("Segera perbaiki mood anda ya")
	}
	if m.Mood.SkorMoodMilestone < 1 || m.Mood.SkorMoodMilestone > 10 {
		fmt.Println("Masukan tidak valid")
		fmt.Println("Masukan kembali pilihan dengan benar!")
		return
	}
	fmt.Println("Catatan Perasaan : ")
	fmt.Scan(&m.Mood.CatatanRasa)
	m.Selesai = false
	m.id = ID
	ID++
	TotalData[JumlahData] = m
	JumlahData++
}

func Hapus() {
	var pilihan int
	fmt.Println("\n========== HAPUS MILESTONE ==========")
	fmt.Println("1. Hapus Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		HapusMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func HapusMilestone() {
	var idHapus int

	if JumlahData == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Print("Masukkan ID Milestone yang ingin dihapus: ")
	fmt.Scan(&idHapus)

	for i := 0; i < JumlahData; i++ {

		if TotalData[i].id == idHapus {

			for j := i; j < JumlahData-1; j++ {
				TotalData[j] = TotalData[j+1]
			}

			JumlahData--

			fmt.Println("Milestone berhasil dihapus.")
			return
		}
	}

	fmt.Println("Maaf ID anda tidak ditemukan.")
	fmt.Println("Silahkan masukan ID anda kembali")
}

func Ubah() {
	var pilihan int
	fmt.Println("\n========== UBAH MILESTONE ==========")
	fmt.Println("1. Ubah Milestone")
	fmt.Println("2. Update Jadwal Tugas")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		UbahMilestone()
	} else if pilihan == 2 {
		UbahJadwalTugas()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func UbahJadwalTugas() {
	var pilihanID int

	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	fmt.Print("Masukan ID Milestone yang ingin diubah jadwal tugasnya: ")
	fmt.Scan(&pilihanID)
	cariIdMilestone(pilihanID)

	fmt.Print("Masukan Jadwal Tugas pada Milestone: ")
	fmt.Scan(&TotalData[pilihanID].tugas.jadwaltugas)
}

func UbahMilestone() {
	var pilihan string
	indeks := cariIdMilestone(ID)

	if indeks == -1 {
		fmt.Println("Milestone tidak ditemukan.")
		return
	}

	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	TampilkanMilestoneSaja()

	fmt.Println(" Masukan ID Milestone yang ingin diubah: ")
	fmt.Scan(&indeks)
	fmt.Println("Bagian mana yang ingin anda ubah")
	fmt.Println("1. Nama Tugas")
	fmt.Println("2. Deskripsi")
	fmt.Println("3. Tanggal Mulai")
	fmt.Println("4. Minggu ke-")
	fmt.Println("5. Tanggal Selesai ")
	fmt.Println("6. Tingkat Kepentingan")
	fmt.Println("7. Skor Stres")
	fmt.Println("8. Skor Mood")
	fmt.Println("9. Catatan Perasaan")
	fmt.Println("0. kembali")
	fmt.Scan(&pilihan)

	if pilihan == "1" {
		fmt.Print("Masukan Nama Kegiatan yang baru: ")
		fmt.Scan(&TotalData[indeks].NamaTugas)
	} else if pilihan == "2" {
		fmt.Print("Masukan Deskripsi yang baru: ")
		fmt.Scan(&TotalData[indeks].Deskripsi)
	} else if pilihan == "3" {
		fmt.Print("Masukan Tanggal Mulai yang baru : ")
		fmt.Scan(&TotalData[indeks].TanggalMulai)
	} else if pilihan == "4" {
		fmt.Print("Masukan Minggu ke- : ")
		fmt.Scan(&TotalData[indeks].Mingguke)
	} else if pilihan == "5" {
		fmt.Print("Masukan Tanggal Selesai yang baru: ")
		fmt.Scan(&TotalData[indeks].TanggalSelesai)
	} else if pilihan == "6" {
		fmt.Print("Masukan Tingkat Kepentingan (1-10): ")
		fmt.Scan(&TotalData[indeks].Kepentingan)
	} else if pilihan == "7" {
		fmt.Print("Masukan Tingkat Stres Anda (1-10): ")
		fmt.Scan(&TotalData[indeks].Mood.SkorStres)
	} else if pilihan == "8" {
		fmt.Print("Masukan Tingkat Mood Anda (1-10): ")
		fmt.Scan(&TotalData[indeks].Mood.SkorMoodMilestone)
	} else  if pilihan == "9" {
		fmt.Print("Masukan Catatan Perasaan: ")
		fmt.Scan(&TotalData[indeks].Mood.CatatanRasa)
	} else if pilihan == "0" {
		return
	} else {
	fmt.Println("Pilihan tidak valid.")
	}
}

func cariIdMilestone(id int) int {
	ketemu := -1
	for i := 0; i < JumlahData; i++ {
		if TotalData[i].id == id {
			ketemu = i
			return ketemu
		}
	}
	return ketemu
}

func SortingTanggal() {
	for i := 0; i < JumlahData-1; i++ {
		IndeksBawah := i
		for j := i + 1; j < JumlahData; j++ {
			if TotalData[j].TanggalSelesai < TotalData[IndeksBawah].TanggalSelesai {
				IndeksBawah = j
			}
		}
		TotalData[i], TotalData[IndeksBawah] = TotalData[IndeksBawah], TotalData[i]
	}
}

func cariTanggalMilestone(tanggal string) {
	if tanggal == "" {
		fmt.Println("Belum ada tugas yang selesai.")
		fmt.Println("Masukan kembali pilihan dengan benar!!")
		return
	}

	SortingTanggal()

	low := 0
	high := JumlahData - 1

	fmt.Println("\n--- Hasil Pencarian by Tanggal Selesai ---")
	for low <= high {
		mid := (low + high) / 2

		if TotalData[mid].TanggalSelesai == tanggal {
			m := TotalData[mid]
			status := "Belum Selesai"
			if m.Selesai {
				status = "Selesai"
			}
			fmt.Printf("\n[ID:%d] %s | Minggu ke-%d | Penting: %d | %s\n",
				m.id, m.NamaTugas, m.Mingguke, m.Kepentingan, status)
			fmt.Printf("Deskripsi : %s\n", m.Deskripsi)
			fmt.Printf("Mulai     : %s\n", m.TanggalMulai)
			fmt.Printf("Deadline  : %s\n", m.TanggalDeadline)
			fmt.Printf("Selesai   : %s\n", m.TanggalSelesai)
			fmt.Printf("Stres     : %d | Mood: %d | Perasaan: %s\n",
				m.Mood.SkorStres, m.Mood.SkorMoodMilestone, m.Mood.CatatanRasa)
			return

		} else if tanggal < TotalData[mid].TanggalSelesai {
			high = mid - 1

		} else {
			low = mid + 1
		}
	}

	fmt.Println("Milestone tidak ditemukan.")
}

func cariNamaMilestone(nama string) {
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan.")
		return
	}

	ketemu := false

	for i := 0; i < JumlahData; i++ {
		if TotalData[i].NamaTugas == nama {
			m := TotalData[i]
			status := "Belum Selesai"
			if m.Selesai {
				status = "Selesai"
			}
			fmt.Println("\n========== HASIL PENCARIAN ==========")
			fmt.Printf("[%d] | Nama: %s | Deskripsi: %s | Mulai: %s | Deadline: %s | Minggu ke-%d | Kepentingan: %d | Progress: %d%% | Stres: %d | Mood: %d | Perasaan: %s | Status: %s\n",
				m.id, m.NamaTugas, m.Deskripsi, m.TanggalMulai, m.TanggalDeadline, m.Mingguke, m.Kepentingan, m.progress, m.Mood.SkorStres, m.Mood.SkorMoodMilestone, m.Mood.CatatanRasa, status)
			fmt.Println("======================================")
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Milestone dengan nama tersebut tidak ditemukan.")
	}
}

func CariMilestone() bool {
	var pilihan int
	var nama, tanggal string
	ketemu := false
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return ketemu
	}

	fmt.Println(" Masukan pilihan yang ingin anda pilih: ")
	fmt.Println("1. Cari berdasarkan Nama Tugas")
	fmt.Println("2. Cari berdasarkan Tanggal Selesai")
	fmt.Scan(&pilihan)
	if pilihan < 1 || pilihan > 2 {
		fmt.Println("Pilihan yang anda masukan tidak valid.")
		return ketemu
	}
	if pilihan == 1 {
		fmt.Print("Masukan Nama Tugas yang ingin dicari: ")
		fmt.Scan(&nama)
		cariNamaMilestone(nama)
		ketemu = true
	} else if pilihan == 2 {
		fmt.Print("Masukan Tanggal Selesai yang ingin dicari (YYYY/MM/DD): ")
		fmt.Scan(&tanggal)
		cariTanggalMilestone(tanggal)
		ketemu = true
	}
	return ketemu
}

func Cari() {
	var pilihan int
	fmt.Println("\n========== Cari MILESTONE ==========")
	fmt.Println("1. Cari Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		CariMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func Tampilkan() {
	var pilihan int
	fmt.Println("\n========== TAMPILKAN MILESTONE ==========")
	fmt.Println("1. Tampilkan Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		TampilkanMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func TampilkanMilestone() {
	var pilihan int
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	fmt.Println(" ")
	fmt.Println("Pilih menu untuk melanjutkan")
	fmt.Println("1. Tampilkan Semua Milestone Saja")
	fmt.Println("2. Tampilkan Semua Milestone dan Jadwal Tugas")
	fmt.Println("Pilihan anda adalah : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		TampilkanMilestoneSaja()
	case 2:
		TampilkanSemua()
	}
}

func TampilkanMilestoneSaja() {
	if JumlahData == 0 {
		fmt.Println("Belum ada data milestone yang ditambahkan.")
		fmt.Println("Silahkan kembali ke halaman utama untak menambahkan data anda terlebih dahulu ya!!")
		return
	}
	for i := 0; i < JumlahData; i++ {
		status := "Belum Selesai"
		if TotalData[i].Selesai {
			status = "Selesai"
		}
		fmt.Printf("[%d] Nama Milestone: %s | Tanggal selesai: %s | Kepentingan: %d | Stres: %d | Mood: %d | %s | Deskripsi: %s | Perasaan : %s\n",
			TotalData[i].id, TotalData[i].NamaTugas, TotalData[i].TanggalSelesai, TotalData[i].Kepentingan, TotalData[i].Mood.SkorStres, TotalData[i].Mood.SkorMoodMilestone, status, TotalData[i].Deskripsi, TotalData[i].Mood.CatatanRasa)
	}
}

func TampilkanSemua() {
	if JumlahData == 0 {
		fmt.Println("Belum ada data milestone yang ditambahkan.")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}
	for i := 0; i < JumlahData; i++ {
		status := "Belum Selesai"
		if TotalData[i].Selesai {
			status = "Selesai"
		}

		fmt.Printf("[%d] %s | Tgl: %s | Penting: %d | Stres: %d | Mood: %d | %s | Deskripsi: %s | Perasaan: %s | Jadwal Tugas: %s\n",
			TotalData[i].id, TotalData[i].NamaTugas, TotalData[i].TanggalSelesai, TotalData[i].Kepentingan, TotalData[i].Mood.SkorStres, TotalData[i].Mood.SkorMoodMilestone, status, TotalData[i].Deskripsi, TotalData[i].Mood.CatatanRasa, TotalData[i].tugas.jadwaltugas)
	}
}

func Urut() {
	var pilihan int
	fmt.Println("\n========== URUTKAN MILESTONE ==========")
	fmt.Println("1. Urutkan Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		UrutkanMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func UrutKepentingan() {
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		return
	}

	for i := 0; i < JumlahData-1; i++ {
		idxMax := i
		for j := i + 1; j < JumlahData; j++ {
			if TotalData[j].Kepentingan > TotalData[idxMax].Kepentingan {
				idxMax = j
			}
		}
		TotalData[i], TotalData[idxMax] = TotalData[idxMax], TotalData[i]
	}
}

func UrutMood() {
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		return
	}

	for i := 1; i < JumlahData; i++ {
		temp := TotalData[i]
		j := i - 1
		for j >= 0 && TotalData[j].Mood.SkorMoodMilestone < temp.Mood.SkorMoodMilestone {
			TotalData[j+1] = TotalData[j]
			j--
		}
		TotalData[j+1] = temp
	}
}

func UrutkanMilestone() {
	var pilihan int
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	fmt.Println("\n========== Urutkan Milestone ==========  ")
	fmt.Println(" Masukan pilihan yang ingin anda pilih: ")
	fmt.Println("1. Urutkan berdasarkan tingkat kepentingan")
	fmt.Println("2. Urutkan berdasarkan skor mood")
	fmt.Println("Pilhan anda adalah : ")
	fmt.Scan(&pilihan)
	if pilihan < 1 || pilihan > 2 {
		fmt.Println("Pilihan yang anda masukan tidak valid.")
		return
	}

	if pilihan == 1 {
		UrutKepentingan()
	}

	if pilihan == 2 {
		UrutMood()
	}
}

func Statistik() {
	var pilihan int
	fmt.Println("\n========== STATISTIK MILESTONE ==========")
	fmt.Println("1. Lihat Statistik")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		StatistikMilestone()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func StatistikMilestone() {
	var DaftarMinggu [1000]int
	var TotalStres [1000]int
	var TotalMood [1000]int
	var TotalMoodTugas [1000]int
	var TotalProgress [1000]int
	var JumlahPerMinggu [1000]int

	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	TugasSelesai := 0
	persenSelesai := float64(TugasSelesai) / float64(JumlahData) * 100
	JumlahMinggu := 0

	for i := 0; i < JumlahData; i++ {
		if TotalData[i].Selesai {
			TugasSelesai++
		}
		Minggu := TotalData[i].Mingguke
		ketemu := false

		for m := 0; m < JumlahMinggu; m++ {
			if DaftarMinggu[m] == Minggu {
				TotalStres[m] += TotalData[i].Mood.SkorStres
				TotalMood[m] += TotalData[i].Mood.SkorMoodMilestone
				TotalMoodTugas[m] += TotalData[i].Mood.SkorMoodTugas
				TotalProgress[m] += TotalData[i].progress
				JumlahPerMinggu[m]++
				ketemu = true
				break
			}
		}
		if ketemu == false {
			DaftarMinggu[JumlahMinggu] = Minggu
			TotalStres[JumlahMinggu] = TotalData[i].Mood.SkorStres
			TotalMood[JumlahMinggu] = TotalData[i].Mood.SkorMoodMilestone
			TotalMoodTugas[JumlahMinggu] = TotalData[i].Mood.SkorMoodTugas
			TotalProgress[JumlahMinggu] = TotalData[i].progress
			JumlahPerMinggu[JumlahMinggu] = 1
			JumlahMinggu++
		}
	}

	fmt.Println("\n===== STATISTIK MINDSTONE =====")
	fmt.Printf("Total Milestone : %d\n", JumlahData)
	fmt.Printf("Target Selesai  : %d dari %d (%.1f%%)\n", TugasSelesai, JumlahData, persenSelesai)
	fmt.Println("\n--- Rata-rata per Minggu ---")

	for m := 0; m < JumlahMinggu; m++ {
		rataStres := float64(TotalStres[m]) / float64(JumlahPerMinggu[m])
		rataMood := float64(TotalMood[m]) / float64(JumlahPerMinggu[m])
		rataMoodTugas := float64(TotalMoodTugas[m]) / float64(JumlahPerMinggu[m])
		rataProgress := float64(TotalProgress[m]) / float64(JumlahPerMinggu[m])

		fmt.Printf("\nMinggu ke-%d (%d milestone)\n", DaftarMinggu[m], JumlahPerMinggu[m])
		fmt.Printf("  Rata-rata Stres       : %.2f / 10\n", rataStres)
		fmt.Printf("  Rata-rata Mood        : %.2f / 10\n", rataMood)
		fmt.Printf("  Rata-rata Mood Tugas  : %.2f / 10\n", rataMoodTugas)
		fmt.Printf("  Rata-rata Progress    : %.1f%%\n", rataProgress)

		if rataStres >= 9 {
			fmt.Println("  Tingkat stres anda sangat tinggi!!")
			fmt.Println("  Segera kurangi aktivitas yang terlalu membebani dan luangkan waktu untuk beristirahat!!")
		} else if rataStres >= 7 {
			fmt.Println("  Tingkat stres anda tinggi!!")
			fmt.Println("  Gawat! Anda perlu segera mengurangi aktivitas yang terlalu membebani.")
		} else if rataStres >= 4 {
			fmt.Println("  Tingkat stres anda sedang.")
			fmt.Println("  Cobalah mengatur waktu istirahat dan mengurangi aktivitas yang terlalu membebani.")
		} else {
			fmt.Println("  Tingkat stres anda rendah.")
			fmt.Println("  Bagus! Anda mampu mengelola tekanan dengan baik.")
		}
	}
}

func JadwalTugas() {
	var pilihan int
	fmt.Println("\n========== JADWAL TUGAS ==========")
	fmt.Println("1. Lihat Jadwal Tugas")
	fmt.Println("2. Tambah Tugas pada Milestone")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		JadwalTugasMilestone()
	} else if pilihan == 2 {
		TambahJadwalTugas()
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func TambahMoodTugas(indeks int) {
	fmt.Println("Bagaimana perasaan/mood anda dalam mengerjakan tugas ini(1-10): ")
	fmt.Scan(&TotalData[indeks].Mood.SkorMoodTugas)

	if TotalData[indeks].Mood.SkorMoodTugas < 1 || TotalData[indeks].Mood.SkorMoodTugas > 10 {
		fmt.Println("Masukan tidak valid")
	}
}

func JadwalTugasMilestone() {
	var pilihanID int
	if JumlahData == 0 {
		fmt.Println("Belum ada data yang ditambahkan")
		fmt.Println("Silahkan kembali ke halaman utama untuk menambahkan data terlebih dahulu ya!!")
		return
	}

	indeks := cariIdMilestone(pilihanID)
	m := TotalData[indeks]
	status := "Belum Selesai"

	fmt.Println("Masukan ID Milestone yang ingin dilihat jadwal tugasnya: ")
	fmt.Scan(&pilihanID)
	if indeks == -1 {
		fmt.Println("Milestone dengan ID tersebut tidak ditemukan.")
		return
	}
	if m.Selesai {
		status = "Selesai"
	}

	fmt.Printf("[%d] Nama Milestone: %s | Tanggal: %s | Tingkat Kepentingan: %d | Tingkat Stres: %d | Tingkat Mood: %d | %s | Deskripsi: %s | Perasaan: %s\n",
		m.id, m.NamaTugas, m.TanggalSelesai, m.Kepentingan, m.Mood.SkorStres, m.Mood.SkorMoodMilestone, status, m.Deskripsi, m.Mood.CatatanRasa)

	fmt.Printf("Tugas Milestone pada ID %d yang telah ditambahkan: %s\n", pilihanID, TotalData[pilihanID].tugas.jadwaltugas)

}

func TambahJadwalTugas() {
	var pilihanID int
	fmt.Print("Masukan ID Milestone yang ingin ditambahkan jadwal tugasnya: ")
	fmt.Scan(&pilihanID)
	cariIdMilestone(pilihanID)

	fmt.Print("Masukan Jadwal Tugas pada Milestone: ")
	fmt.Scan(&TotalData[pilihanID].tugas.jadwaltugas)
	indeks := cariIdMilestone(pilihanID)

	if indeks == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	fmt.Print("Masukan Jadwal Tugas pada Milestone: ")
	fmt.Scan(&TotalData[indeks].tugas.jadwaltugas)

	TambahMoodTugas(indeks)

}
