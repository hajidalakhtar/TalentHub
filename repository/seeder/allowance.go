// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

func CreateAllowance(db *gorm.DB,
	tunjangan_lain_lain int,
	tunjangan_jabatan int,
	tunjangan_bulanan int,

) error {
	return db.Create(&entity.Allowance{
		TunjanganLainlain: tunjangan_lain_lain,
		TunjanganJabatan:  tunjangan_jabatan,
		TunjanganBulanan:  tunjangan_bulanan,
	}).Error
}
