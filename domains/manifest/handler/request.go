package manifesthandler

type Request struct {
	CustomURL string `json:"custom_url" form:"custom_url" validate:"required,min=3"`
	RealURL   string `json:"real_url" form:"real_url" validate:"required"`
}
