package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	groupSize   = "GROUP_SIZE"
	storageFlag = "STORAGE"
)

// EnvConfig entity for config
type EnvConfig interface {
	GroupSize() int32
	StorageFlag() bool
}

type envConfig struct {
	groupSize   int32
	storageFlag bool // true - mem; false - db
}

// NewEnvConfig constructor
func NewEnvConfig() (EnvConfig, error) {
	groupSize, err := strconv.Atoi(os.Getenv(groupSize))
	if err != nil {
		return nil, errors.New("group size not found")
	}

	storageFlag, err := strconv.ParseBool(os.Getenv(storageFlag))
	if err != nil {
		return nil, errors.New("storage flag not found")
	}

	return &envConfig{
		groupSize:   int32(groupSize),
		storageFlag: storageFlag,
	}, nil
}

func (cfg *envConfig) GroupSize() int32 {
	return cfg.groupSize
}

func (cfg *envConfig) StorageFlag() bool {
	return cfg.storageFlag
}
