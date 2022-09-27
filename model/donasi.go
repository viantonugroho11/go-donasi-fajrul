package model

type ResponseKampanyeDonasi struct {
	DonasiID       string                 `json:"donasi_id"`
	NamaKampanye   string                 `json:"nama_kampanye"`
	Slug           string                 `json:"slug"`
	Kategori       ResponseKategoriDonasi `json:"kategori"`
	Deskripsi      string                 `json:"deskripsi"`
	TargetDonasi   int                    `json:"target_donasi"`
	DonasiSaatIni  int                    `json:"donasi_saat_ini"`
	DonaturSaatIni int                    `json:"donatur_saat_ini"`
	Gambar         string                 `json:"gambar"`
	BatasWaktu     string                 `json:"batas_waktu"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
}

type ResponseKategoriDonasi struct {
	KategoriID string `json:"kategori_id"`
	Nama       string `json:"nama"`
	Slug       string `json:"slug"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
