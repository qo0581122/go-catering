package model

type UserAddressTag struct {
	Model
	TagName         string `json:"tag_name"`
	TextColor       string `json:"text_color"`
	BorderColor     string `json:"border_color"`
	BackgroundColor string `json:"background_color"`
	Sort            uint32 `json:"sort"`
	Status          uint32 `json:"status"`
}

func (*UserAddressTag) TableName() string {
	return "user_address_tag"
}
