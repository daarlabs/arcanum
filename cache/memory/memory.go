package memory

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Client struct {
	sync.RWMutex
	data map[string]data
	dir  string
}

type data struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
}

const (
	jsonSuffix    = ".json"
	watchInterval = time.Second
)

func New(dir string) *Client {
	m := &Client{
		data: make(map[string]data),
		dir:  getDir(dir),
	}
	go m.load()
	go m.watch()
	return m
}

func (m *Client) Get(key string) string {
	m.Lock()
	d, ok := m.data[key]
	m.Unlock()
	if !ok {
		return ""
	}
	return d.Value
}

func (m *Client) Set(key string, value string, expiration time.Duration) error {
	d := data{
		Value:      value,
		Expiration: time.Now().Add(expiration),
	}
	m.Lock()
	m.data[key] = d
	m.Unlock()
	if err := m.setTempFile(key, d); err != nil {
		return err
	}
	return nil
}

func (m *Client) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *Client) Destroy(key string) error {
	delete(m.data, key)
	if err := m.deleteTempFile(key); err != nil {
		return err
	}
	return nil
}

func (m *Client) load() {
	if _, err := os.Stat(m.dir); os.IsNotExist(err) {
		return
	}
	if err := filepath.Walk(
		m.dir, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), jsonSuffix) {
				return nil
			}
			key := strings.TrimSuffix(info.Name(), jsonSuffix)
			fbts, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if len(fbts) == 0 {
				return nil
			}
			var d data
			if err := json.Unmarshal(fbts, &d); err != nil {
				return err
			}
			m.data[key] = d
			return nil
		},
	); err != nil {
		log.Fatalln(err)
	}
}

func (m *Client) watch() {
	ticker := time.NewTicker(watchInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			expired := make([]string, 0)
			m.RLock()
			for key, d := range m.data {
				if t.Before(d.Expiration) {
					continue
				}
				expired = append(expired, key)
			}
			m.RUnlock()
			if len(expired) > 0 {
				m.Lock()
				for _, key := range expired {
					if err := m.Destroy(key); err != nil {
						log.Fatalln(err)
					}
				}
				expired = nil
				m.Unlock()
			}
		}
	}
}

func (m *Client) setTempFile(key string, d data) error {
	if _, err := os.Stat(m.dir); os.IsNotExist(err) {
		if err := os.MkdirAll(m.dir, os.ModePerm); err != nil {
			return err
		}
	}
	path := fmt.Sprintf("%s/%s%s", m.dir, key, jsonSuffix)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	dbts, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, dbts, 0444); err != nil {
		return err
	}
	return nil
}

func (m *Client) deleteTempFile(key string) error {
	if _, err := os.Stat(m.dir); os.IsNotExist(err) {
		return nil
	}
	path := fmt.Sprintf("%s/%s%s", m.dir, key, jsonSuffix)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

func getDir(tmpDir string) string {
	if strings.HasSuffix(tmpDir, "/") {
		tmpDir = strings.TrimSuffix(tmpDir, "/")
	}
	return fmt.Sprintf("%s/.arcanum/memory", tmpDir)
}
