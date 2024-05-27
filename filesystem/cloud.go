package filesystem

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/minio/minio-go/v7"
)

type cloud struct {
	ctx         context.Context
	dir         string
	storageName string
	client      *minio.Client
}

func createCloud(ctx context.Context, dir string, storageName string, client *minio.Client) Client {
	if strings.HasPrefix(dir, "/") {
		dir = strings.TrimPrefix(dir, "/")
	}
	s := &cloud{
		ctx:         ctx,
		dir:         dir,
		client:      client,
		storageName: storageName,
	}
	return s
}

func (s *cloud) GetList() ([]string, error) {
	result := make([]string, 0)
	for item := range s.client.ListObjects(s.ctx, s.storageName, minio.ListObjectsOptions{Recursive: true}) {
		result = append(result, item.Key)
	}
	return make([]string, 0), nil
}

func (s *cloud) MustGetList() []string {
	list, err := s.GetList()
	if err != nil {
		panic(err)
	}
	return list
}

func (s *cloud) Read(path string) ([]byte, error) {
	object, err := s.client.GetObject(s.ctx, s.storageName, s.createPath(path), minio.GetObjectOptions{})
	if err != nil {
		return []byte{}, err
	}
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(object); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func (s *cloud) MustRead(path string) []byte {
	data, err := s.Read(path)
	if err != nil {
		panic(err)
	}
	return data
}

func (s *cloud) Create(path string, data []byte) error {
	_, err := s.client.PutObject(
		s.ctx, s.storageName, s.createPath(path), bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
			ContentType: http.DetectContentType(data),
		},
	)
	return err
}

func (s *cloud) MustCreate(path string, data []byte) {
	err := s.Create(path, data)
	if err != nil {
		panic(err)
	}
}

func (s *cloud) Remove(path string) error {
	return s.client.RemoveObject(s.ctx, s.storageName, s.createPath(path), minio.RemoveObjectOptions{})
}

func (s *cloud) MustRemove(path string) {
	err := s.Remove(path)
	if err != nil {
		panic(err)
	}
}

func (s *cloud) createPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	return fmt.Sprintf("%s/%s", s.dir, path)
}
