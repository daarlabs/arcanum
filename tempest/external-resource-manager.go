package tempest

import (
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type externalResourceManager struct {
	styles  [][]byte
	scripts [][]byte
}

func createExternalResourceManager() *externalResourceManager {
	m := &externalResourceManager{
		styles:  make([][]byte, 0),
		scripts: make([][]byte, 0),
	}
	return m
}

func (m *externalResourceManager) run() error {
	readStyles := make([]string, 0)
	readScripts := make([]string, 0)
	fetchStyles := make([]string, 0)
	fetchScripts := make([]string, 0)
	for _, styleFile := range GlobalConfig.Styles {
		isHttp := strings.HasPrefix(styleFile, "http")
		if !isHttp {
			readStyles = append(readStyles, styleFile)
		}
		if isHttp {
			fetchStyles = append(fetchStyles, styleFile)
		}
	}
	for _, scriptFile := range GlobalConfig.Scripts {
		isHttp := strings.HasPrefix(scriptFile, "http")
		if !isHttp {
			readScripts = append(readScripts, scriptFile)
		}
		if isHttp {
			fetchScripts = append(fetchScripts, scriptFile)
		}
	}
	if len(fetchStyles) > 0 {
		r, err := m.fetch(fetchStyles)
		if err != nil {
			return err
		}
		m.styles = append(m.styles, r...)
	}
	if len(readStyles) > 0 {
		r, err := m.read(readStyles)
		if err != nil {
			return err
		}
		m.styles = append(m.styles, r...)
	}
	if len(fetchScripts) > 0 {
		r, err := m.fetch(fetchScripts)
		if err != nil {
			return err
		}
		m.scripts = append(m.scripts, r...)
	}
	if len(readScripts) > 0 {
		r, err := m.read(readScripts)
		if err != nil {
			return err
		}
		m.scripts = append(m.scripts, r...)
	}
	return nil
}

func (m *externalResourceManager) mustRun() {
	err := m.run()
	if err != nil {
		panic(err)
	}
}

func (m *externalResourceManager) read(paths []string) ([][]byte, error) {
	n := len(paths)
	errs := make(chan error, n)
	result := make([][]byte, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i, filePath := range paths {
		go func(i int, filePath string, errs chan error) {
			defer wg.Done()
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				errs <- err
				return
			}
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				errs <- err
				return
			}
			result[i] = fileBytes
		}(i, filePath, errs)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return [][]byte{}, err
		}
	}
	return result, nil
}

func (m *externalResourceManager) fetch(paths []string) ([][]byte, error) {
	n := len(paths)
	errs := make(chan error, n)
	result := make([][]byte, n)
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i, script := range paths {
		go func(errs chan error, i int, script string) {
			defer wg.Done()
			r, err := http.Get(script)
			if err != nil {
				errs <- err
				return
			}
			scriptBytes, err := io.ReadAll(r.Body)
			if err != nil {
				errs <- err
				return
			}
			result[i] = scriptBytes
		}(errs, i, script)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return [][]byte{}, err
		}
	}
	return result, nil
}
