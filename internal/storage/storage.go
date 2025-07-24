package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

type Storage interface {
	Create(name, user, project string) error
	Get(name string) (*Profile, error)
	List() map[string]Profile
	Delete(name string) error
}

type FileStorage struct {
	dir string
}

func NewFileStorage(dir string) *FileStorage {
	return &FileStorage{dir: dir}
}

func (fs *FileStorage) Create(name, user, project string) error {
	fp := fs.getProfilePath(name)
	if _, err := os.Stat(fp); !os.IsNotExist(err) {
		return errors.New("profile already exists")
	}

	profile := Profile{User: user, Project: project}
	data, err := yaml.Marshal(profile)
	if err != nil {
		return fmt.Errorf("yaml marshal error: %w", err)
	}

	return os.WriteFile(fp, data, 0o600)
}

func (fs *FileStorage) Get(name string) (*Profile, error) {
	data, err := os.ReadFile(fs.getProfilePath(name))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("profile not found")
		}
		return nil, fmt.Errorf("read error: %w", err)
	}

	var profile Profile
	if err := yaml.Unmarshal(data, &profile); err != nil {
		return nil, fmt.Errorf("yaml parse error: %w", err)
	}

	return &profile, nil
}

func (fs *FileStorage) List() map[string]Profile {
	profiles := make(map[string]Profile)
	files, _ := filepath.Glob(filepath.Join(fs.dir, "*.yaml"))

	for _, f := range files {
		name := strings.TrimSuffix(filepath.Base(f), ".yaml")
		if p, err := fs.Get(name); err == nil {
			profiles[name] = *p
		}
	}
	return profiles
}

func (fs *FileStorage) Delete(name string) error {
	fp := fs.getProfilePath(name)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return errors.New("profile not found")
	}
	return os.Remove(fp)
}

func (fs *FileStorage) getProfilePath(name string) string {
	return filepath.Join(fs.dir, name+".yaml")
}
