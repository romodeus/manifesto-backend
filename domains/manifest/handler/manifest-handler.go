package manifesthandler

import (
	"manifesto/config"
	manifestentity "manifesto/domains/manifest/entities"
	"manifesto/exceptions"
	"manifesto/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type manifestHandler struct {
	Service manifestentity.IServiceManifest
}

func New(service manifestentity.IServiceManifest) *manifestHandler {
	return &manifestHandler{
		Service: service,
	}
}

func (s *manifestHandler) Post(c echo.Context) error {
	cfg := config.GetConfig()

	request := Request{}
	err := c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&request)
	if err != nil {
		return err
	}

	result, err := s.Service.CustomURLCreate(manifestentity.Manifest{
		CustomURL: request.CustomURL,
		RealURL:   request.RealURL,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(map[string]interface{}{
		"custom_url": cfg.BASE_URL + "/" + result.CustomURL,
		"real_url":   result.RealURL,
	}))
}

func (s *manifestHandler) Get(c echo.Context) error {
	cfg := config.GetConfig()
	manifest := manifestentity.Manifest{}
	manifest.CustomURL = c.Param("key")

	result, err := s.Service.Get(manifest)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(map[string]interface{}{
		"custom_url": cfg.BASE_URL + "/" + result.CustomURL,
		"real_url":   result.RealURL,
	}))
}

func (s *manifestHandler) Redirect(c echo.Context) error {
	manifest := manifestentity.Manifest{}
	manifest.CustomURL = c.Param("key")

	result, err := s.Service.Get(manifest)

	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, config.NOT_FOUND_URL_REDIRECT)
	}

	return c.Redirect(http.StatusPermanentRedirect, result.RealURL)
}
