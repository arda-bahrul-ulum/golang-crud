package entities

type Pasien struct {
	Id int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	NIK string `validate:"required" label:"NIK"`
	JenisKelamin string `validate:"required" label:"Jenis Kelamin"`
	TempatLahir string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir string `validate:"required" label:"Tanggal Lahir"`
	Alamat string `validate:"required" label:"Alamat"`
	NoHp string `validate:"required" label:"Nomor HP"`
}