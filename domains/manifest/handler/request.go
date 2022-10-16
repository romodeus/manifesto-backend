package manifesthandler

type Request struct {
	CustomURL string `json:"custom_url" form:"custom_url" validate:"required,max=10,min=3"`
	RealURL   string `json:"real_url" form:"real_url" validate:"required"`
}
