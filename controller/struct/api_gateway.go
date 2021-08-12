package _struct

type ApiGateway struct {
	Token       string `json:"token" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Endpoind    string `json:"endpoind" binding:"required"`
	Method      string `json:"method" binding:"required"`
}

type ResultData struct {
	ID          string `json:"ID,omitempty"`
	IdentityID  string `json:"identity_id,omitempty"`
	DUK         string `json:"duk,omitempty"`
	Name        string `json:"name,omitempty"`
	KTP         string `json:"ktp,omitempty"`
	NRP         string `json:"nrp,omitempty"`
	NIP         string `json:"nip,omitempty"`
	BirthPlace  string `json:"birth_place,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Address     string `json:"address,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Religion    string `json:"religion,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Link        string `json:"link,omitempty"`
	Formula     string `json:"formula,omitempty"`
	Position    string `json:"position,omitempty"`
	Sector      string `json:"sector,omitempty"`
	Workunit    string `json:"workunit,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

type Result struct {
	Data    ResultData `json:"data,omitempty"`
	Message string     `json:"message,omitempty"`
	Status  int        `json:"status,omitempty"`
}
