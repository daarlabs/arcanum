package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type local struct {
	dir string
	mu  *sync.Mutex
}

func createLocal(dir string) Client {
	if !strings.HasPrefix(dir, "/") && !strings.HasPrefix(dir, "./") {
		dir = "/" + dir
	}
	s := &local{
		dir: dir,
		mu:  new(sync.Mutex),
	}
	if err := s.validateDir(dir); err != nil {
		panic(err)
	}
	return s
}

func (s *local) GetList() ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]string, 0)
	err := filepath.Walk(
		s.dir, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			result = append(result, path)
			return nil
		},
	)
	return result, err
}

func (s *local) MustGetList() []string {
	list, err := s.GetList()
	if err != nil {
		panic(err)
	}
	return list
}

func (s *local) Read(path string) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return os.ReadFile(s.createPath(path))
}

func (s *local) MustRead(path string) []byte {
	data, err := s.Read(path)
	if err != nil {
		panic(err)
	}
	return data
}

func (s *local) Create(path string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path = s.createPath(path)
	if err := s.validateDir(path[:strings.LastIndex(path, "/")]); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err = f.Write(data); err != nil {
		return err
	}
	if err = f.Chmod(0600); err != nil {
		return err
	}
	return f.Close()
}

func (s *local) MustCreate(path string, data []byte) {
	err := s.Create(path, data)
	if err != nil {
		panic(err)
	}
}

func (s *local) Remove(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return os.Remove(s.createPath(path))
}

func (s *local) MustRemove(path string) {
	err := s.Remove(path)
	if err != nil {
		panic(err)
	}
}

func (s *local) createPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if strings.HasPrefix(path, "./") {
		path = strings.TrimPrefix(path, "./")
	}
	return fmt.Sprintf("%s/%s", s.dir, path)
}

func (s *local) validateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}
	return nil
}
