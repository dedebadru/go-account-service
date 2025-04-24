package dto

type RegisterRequest struct {
	Name						string `json:"nama" validate:"required"`
	IdentityNumber	string `json:"nik" validate:"required"`
	PhoneNumber 		string `json:"no_hp" validate:"required"`
}
