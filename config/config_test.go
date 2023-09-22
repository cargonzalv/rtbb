package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Test Suite")
}

var _ = Describe("rtb-bidder configuration setup", func() {
	BeforeEach(func() {
		parentDir, _ := os.Getwd()
		parentDir = strings.TrimRight(parentDir, "/config")
		os.Setenv(configParentDirectEnvVarName, parentDir)
	})

	It("Returns the correct yaml config file path", func() {
		path := configPath()
		Expect(path).ShouldNot(BeNil())
		Expect(path).Should(HaveSuffix("rtb-bidder/config"))
	})

	Context("Configuration is loaded", func() {
		It("Default configuration is loaded when folder path is provided", func() {
			path := configPath()
			cfg, err := loadConfig(path)
			Expect(cfg).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
			Expect(cfg.App.Name).Should(BeIdenticalTo("rtb-bidder"))
		})

		It("Path is resolved and Configuration is loaded", func() {
			cfg := ProvideConfig()
			Expect(cfg).ShouldNot(BeNil())
			Expect(cfg.App.Name).Should(BeIdenticalTo("rtb-bidder"))
		})
	})

	Context("Returns the environment variable if present else the fallback value", func() {
		f := "fallback"
		It("Returns env variable value", func() {
			expectedVal := getEnv(configParentDirectEnvVarName, f)
			Expect(expectedVal).Should(HaveSuffix("rtb-bidder"))
		})

		It("Returns fallback variable value", func() {
			expectedVal := getEnv("MISSING_ENV_VARIABLE", f)
			Expect(expectedVal).Should(BeIdenticalTo(f))
		})
	})
})
