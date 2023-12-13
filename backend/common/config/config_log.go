package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// _KEY_LOG             = "log"
	_KEY_LOG_PATH        = "log.path"
	_KEY_LOG_MAX_SIZE    = "log.max_size"
	_KEY_LOG_MAX_BACKUPS = "log.max_backups"
	_KEY_LOG_MAX_AGE     = "log.max_age"
	_KEY_LOG_LOCAL_TIEM  = "log.local_time"
	_KEY_LOG_COMPRESS    = "log.compress"
)

func LogGetPath() string {
	cfg := Instance()

	var path string
	if !cfg.v.IsSet(_KEY_LOG_PATH) {
		path = "./logs/log"

		cfg.v.Set(_KEY_LOG_PATH, path)
		cfg.v.WriteConfig()
	} else {
		path = cfg.v.GetString(_KEY_LOG_PATH)
	}

	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		panic(err)
	}

	if dirInfo, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else if !dirInfo.IsDir() {
		panic(fmt.Errorf("the target path \"%s\" already exists, but it's not a folder", dir))
	}

	return path
}

func LogGetMaxSize() int {
	cfg := Instance()

	if !cfg.v.IsSet(_KEY_LOG_MAX_SIZE) {
		size := 10

		cfg.v.Set(_KEY_LOG_MAX_SIZE, size)
		cfg.v.WriteConfig()

		return size
	}

	size := cfg.v.GetInt(_KEY_LOG_MAX_SIZE)
	return size
}

func LogGetMaxBackups() int {
	cfg := Instance()

	if !cfg.v.IsSet(_KEY_LOG_MAX_BACKUPS) {
		maxBackups := 7

		cfg.v.Set(_KEY_LOG_MAX_BACKUPS, maxBackups)
		cfg.v.WriteConfig()

		return maxBackups
	}

	maxBackups := cfg.v.GetInt(_KEY_LOG_MAX_BACKUPS)
	return maxBackups
}

func LogGetMaxAge() int {
	cfg := Instance()

	if !cfg.v.IsSet(_KEY_LOG_MAX_AGE) {
		maxAge := 30

		cfg.v.Set(_KEY_LOG_MAX_AGE, maxAge)
		cfg.v.WriteConfig()

		return maxAge
	}

	maxAge := cfg.v.GetInt(_KEY_LOG_MAX_AGE)
	return maxAge
}

func LogGetIsLocalTime() bool {
	cfg := Instance()

	if !cfg.v.IsSet(_KEY_LOG_LOCAL_TIEM) {
		isLocalTime := false

		cfg.v.Set(_KEY_LOG_LOCAL_TIEM, isLocalTime)
		cfg.v.WriteConfig()

		return isLocalTime
	}

	isLocalTime := cfg.v.GetBool(_KEY_LOG_LOCAL_TIEM)
	return isLocalTime
}

func LogGetIsComporess() bool {
	cfg := Instance()

	if !cfg.v.IsSet(_KEY_LOG_COMPRESS) {
		isCompress := false

		cfg.v.Set(_KEY_LOG_COMPRESS, isCompress)
		cfg.v.WriteConfig()

		return isCompress
	}

	isCompress := cfg.v.GetBool(_KEY_LOG_COMPRESS)
	return isCompress
}
