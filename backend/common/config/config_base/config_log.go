package config_base

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
	if !cfg.V.IsSet(_KEY_LOG_PATH) {
		path = "./logs/log"

		cfg.V.Set(_KEY_LOG_PATH, path)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default log path failed with error: %v\n", err)
		}
	} else {
		path = cfg.V.GetString(_KEY_LOG_PATH)
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

	if !cfg.V.IsSet(_KEY_LOG_MAX_SIZE) {
		size := 10

		cfg.V.Set(_KEY_LOG_MAX_SIZE, size)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default max log size failed with error: %v\n", err)
		}

		return size
	}

	size := cfg.V.GetInt(_KEY_LOG_MAX_SIZE)
	return size
}

func LogGetMaxBackups() int {
	cfg := Instance()

	if !cfg.V.IsSet(_KEY_LOG_MAX_BACKUPS) {
		maxBackups := 7

		cfg.V.Set(_KEY_LOG_MAX_BACKUPS, maxBackups)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default max log backup failed with error: %v\n", err)
		}

		return maxBackups
	}

	maxBackups := cfg.V.GetInt(_KEY_LOG_MAX_BACKUPS)
	return maxBackups
}

func LogGetMaxAge() int {
	cfg := Instance()

	if !cfg.V.IsSet(_KEY_LOG_MAX_AGE) {
		maxAge := 30

		cfg.V.Set(_KEY_LOG_MAX_AGE, maxAge)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default max log age failed with error: %v\n", err)
		}

		return maxAge
	}

	maxAge := cfg.V.GetInt(_KEY_LOG_MAX_AGE)
	return maxAge
}

func LogGetIsLocalTime() bool {
	cfg := Instance()

	if !cfg.V.IsSet(_KEY_LOG_LOCAL_TIEM) {
		isLocalTime := false

		cfg.V.Set(_KEY_LOG_LOCAL_TIEM, isLocalTime)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default log local time failed with error: %v\n", err)
		}

		return isLocalTime
	}

	isLocalTime := cfg.V.GetBool(_KEY_LOG_LOCAL_TIEM)
	return isLocalTime
}

func LogGetIsComporess() bool {
	cfg := Instance()

	if !cfg.V.IsSet(_KEY_LOG_COMPRESS) {
		isCompress := false

		cfg.V.Set(_KEY_LOG_COMPRESS, isCompress)
		err := cfg.V.WriteConfig()
		if err != nil {
			fmt.Printf("set default log comporess failed with error: %v\n", err)
		}

		return isCompress
	}

	isCompress := cfg.V.GetBool(_KEY_LOG_COMPRESS)
	return isCompress
}
