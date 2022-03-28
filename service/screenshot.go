package service

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/alyakimenko/pageshot/models"
	"github.com/alyakimenko/pageshot/storage"
)

// defaultImageFormat is a default format for screenshot images.
const defaultImageFormat = models.PNGImageFormat

// Browser is a browser interface that is required for the ScreenshotService.
type Browser interface {
	Screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, error)
}

// FileStorage is a file storage interface that is required for the ScreenshotService.
type FileStorage interface {
	Upload(ctx context.Context, input storage.UploadInput) error
	Download(ctx context.Context, filename string) (io.Reader, error)
}

// ScreenshotService deals with screenshots related logic.
type ScreenshotService struct {
	browser     Browser
	fileStorage FileStorage
}

// ScreenshotServiceParams is an incoming params for the NewScreenshotService function.
type ScreenshotServiceParams struct {
	Browser     Browser
	FileStorage FileStorage
}

// NewScreenshotService creates new instance of the ScreenshotService.
func NewScreenshotService(params ScreenshotServiceParams) *ScreenshotService {
	return &ScreenshotService{
		browser:     params.Browser,
		fileStorage: params.FileStorage,
	}
}

// Screenshot takes a screenshot via the underlying Browser and uploads the resulting image to Storage.
func (s *ScreenshotService) Screenshot(ctx context.Context, opts models.ScreenshotOptions) (io.Reader, string, error) {
	if s.fileStorage == nil {
		return s.screenshotWithoutUpload(ctx, opts)
	}

	hash, err := opts.Hash()
	if err != nil {
		return nil, "", err
	}

	if !opts.Format.IsValid() {
		opts.Format = defaultImageFormat
	}

	// construct a filename
	filename := hash + opts.Format.Extension()

	// try to get cached image from a storage
	file, err := s.fileStorage.Download(ctx, filename)
	if err == nil {
		return file, opts.Format.ContentType(), nil
	}

	if !errors.Is(err, storage.ErrFileNotExist) {
		return nil, "", err
	}

	// if screenshot doesn't exist, create one
	screenshot, err := s.browser.Screenshot(ctx, opts)
	if err != nil {
		return nil, "", err
	}

	// and upload it to a storage
	var buf bytes.Buffer
	tee := io.TeeReader(bytes.NewReader(screenshot), &buf)

	err = s.fileStorage.Upload(ctx, storage.UploadInput{
		Filename: filename,
		File:     tee,
	})
	if err != nil {
		return nil, "", err
	}

	return &buf, opts.Format.ContentType(), nil
}

// screenshotWithoutUpload takes a screenshot without upload it to a storage.
func (s *ScreenshotService) screenshotWithoutUpload(
	ctx context.Context, opts models.ScreenshotOptions,
) (io.Reader, string, error) {
	screenshot, err := s.browser.Screenshot(ctx, opts)
	if err != nil {
		return nil, "", err
	}

	return bytes.NewBuffer(screenshot), opts.Format.ContentType(), nil
}
