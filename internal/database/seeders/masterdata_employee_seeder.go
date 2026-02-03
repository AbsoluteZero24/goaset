package seeders

import (
	"github.com/AbsoluteZero24/goaset/internal/models"
	"gorm.io/gorm"
)

func SeedMasterDataEmployee(db *gorm.DB) error {
	// 1. KSO PUSAT
	pusat := models.MasterBranch{Name: "KSO PUSAT"}
	db.Where(models.MasterBranch{Name: pusat.Name}).FirstOrCreate(&pusat)

	// 1.1. Operasi Luar Negeri
	operasiLN := models.MasterDepartment{Name: "Operasi Luar Negeri", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: operasiLN.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&operasiLN)

	// Sub-departments for 1.1
	subLN := []string{"Cabang Luar Negeri", "Surveyor Luar Negeri"}
	for _, name := range subLN {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: operasiLN.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: operasiLN.ID,
		})
	}

	// 1.2. Keuangan & Administrasi
	keuanganAd := models.MasterDepartment{Name: "Keuangan & Administrasi", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: keuanganAd.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&keuanganAd)

	// Sub-departments for 1.2
	subKA := []string{"SDM, Umum, dan Hukum", "Keuangan & Akutansi"}
	for _, name := range subKA {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: keuanganAd.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: keuanganAd.ID,
		})
	}

	// 1.3. Sistem, Kepatuhan & Pelayanan Pelanggan
	sistemKepatuhan := models.MasterDepartment{Name: "Sistem, Kepatuhan & Pelayanan Pelanggan", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: sistemKepatuhan.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&sistemKepatuhan)

	// Sub-departments for 1.3
	subSK := []string{"Pelayanan Pelanggan", "Standarisasi & Pelaporan", "Sistem, Jaminan Mutu & Kepatuhan"}
	for _, name := range subSK {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: sistemKepatuhan.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: sistemKepatuhan.ID,
		})
	}

	// 1.4. Operasi Dalam Negeri
	operasiDN := models.MasterDepartment{Name: "Operasi Dalam Negeri", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: operasiDN.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&operasiDN)

	// Sub-departments for 1.4
	subDN := []string{"Operasi 1", "Operasi 2"}
	for _, name := range subDN {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: operasiDN.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: operasiDN.ID,
		})
	}

	// 1.5. Pengembangan Usaha, Hubungan Pemangku Kepentingan & Penjualan
	pengembangan := models.MasterDepartment{Name: "Pengembangan Usaha, Hubungan Pemangku Kepentingan & Penjualan", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: pengembangan.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&pengembangan)

	// Sub-departments for 1.5
	subP := []string{"Penjualan", "Pengembangan Usaha", "Hubungan Pemangku Kepentingan"}
	for _, name := range subP {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: pengembangan.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: pengembangan.ID,
		})
	}

	// 1.6. Sistem Informasi
	it := models.MasterDepartment{Name: "Sistem Informasi", MasterBranchID: pusat.ID}
	db.Where(models.MasterDepartment{Name: it.Name, MasterBranchID: pusat.ID}).FirstOrCreate(&it)

	// Sub-departments for 1.6
	subIT := []string{"Pengembangan Aplikasi", "Otomasi Sistem dan Pengelolaan Infrastruktur"}
	for _, name := range subIT {
		db.Where(models.MasterSubDepartment{Name: name, MasterDepartmentID: it.ID}).FirstOrCreate(&models.MasterSubDepartment{
			Name:               name,
			MasterDepartmentID: it.ID,
		})
	}

	// Other branches
	otherBranches := []string{
		"KSO Cabang Hongkong",
		"KSO Cabang Korea",
		"KSO Cabang Malaysia",
		"KSO Cabang Shenzhen",
		"KSO Cabang Singapura",
		"KSO Cabang Thailand",
		"KSO Cabang Vietnam",
	}
	for _, name := range otherBranches {
		db.Where(models.MasterBranch{Name: name}).FirstOrCreate(&models.MasterBranch{Name: name})
	}

	return nil
}
