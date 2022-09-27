package model

type PayloadTransaksiRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	Nominal  int    `json:"nominal"`
	DonasiID string `json:"donasi_id"`
	Doa      string `json:"doa"`
	Status   bool   `json:"status"`
}


