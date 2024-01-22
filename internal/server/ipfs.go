package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

type Client interface {
	Post(path string) (*resty.Response, error)
}

type Web3Storage struct {
	BaseURL string
}

func (w *Web3Storage) Post(path string) (*resty.Response, error) {
	panic("implement me")
}

type Localhost struct {
	BaseURL string
}

func (l *Localhost) Post(path string) (*resty.Response, error) {
	client := resty.New()

	r, err := client.R().
		SetFile("path", path).
		Post(fmt.Sprintf("%s/api/v0/add?cid-version=1", l.BaseURL))

	if err != nil {
		return nil, err
	}

	return r, nil
}

type AddFileResponse struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
	Size string `json:"size"`
}

func (s *Server) IPFSAddFiles() error {
	form, _ := s.ctx.MultipartForm()
	f := form.File["path"]

	local := &Localhost{
		BaseURL: IPFSAddr(),
	}

	var fs []AddFileResponse
	for _, file := range f {
		dst := "/tmp/" + file.Filename
		if err := s.ctx.SaveUploadedFile(file, dst); err != nil {
			return fmt.Errorf("failed to save file: %w", err)
		}

		r, err := local.Post(dst)
		if err != nil {
			return fmt.Errorf("failed to post ipfs: %w", err)
		}

		if r.IsError() {
			return fmt.Errorf("failed to post ipfs: %w", err)
		}

		var fr AddFileResponse
		if err := json.Unmarshal(r.Body(), &fr); err != nil {
			return fmt.Errorf("failed to parse response body: %w", err)
		}

		fs = append(fs, fr)
		if err := os.Remove(dst); err != nil {
			return fmt.Errorf("failed to remove file: %w", err)
		}
	}

	Success(s.ctx, fs)
	return nil
}
