package filesystem

import "github.com/minio/minio-go/v7"

type Config struct {
	Driver string
	Name   string
	Dir    string
	Cloud  *minio.Client
}
