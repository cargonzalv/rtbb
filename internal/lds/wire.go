package lds

import (
	dsplds "github.com/adgear/dsp-lds/pkg/lds"
	"github.com/adgear/rtb-bidder/config"
)

// ProvideMetrics loading metrics from go_commons library.
func ProvideLocalDataStore(cfg *config.Config) dsplds.LocalDataStore {
	l, err := dsplds.NewLocalDataStore(&cfg.LocalDataStoreConfig)

	if err != nil {
		panic(err)
	}
	return dsplds.LocalDataStore(l)
}
