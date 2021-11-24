package config

import (
	"fmt"
	"runtime"
)

var (
	version        = "XXXX"
	buildDate      = "1970-01-01T00:00:00Z"
	gitCommit      = ""
	gitTag         = ""
	kubectlVersion = ""
)

type BuildInfo struct {
	Version        string `json:"version"`
	BuildDate      string `json:"build_data"`
	GitCommit      string `json:"git_commit"`
	GitTag         string `json:"git_tag"`
	Go             string `json:"go-version"`
	KubectlVersion string `json:"kubectl_version"`
	Platform       string `json:"platform"`
}

func (v BuildInfo) String() string {
	return fmt.Sprintf(
		"BuildInfo(Version='%v', GitCommit='%v', BuildDate='%v', Go='%v', KubectlVersion='%v', Platform='%v')",
		v.Version,
		v.GitCommit,
		v.BuildDate,
		v.Go,
		v.KubectlVersion,
		v.Platform,
	)
}

func Get() BuildInfo {
	return BuildInfo{
		Version:        version,
		BuildDate:      buildDate,
		GitCommit:      gitCommit,
		GitTag:         gitTag,
		Go:             runtime.Version(),
		Platform:       fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		KubectlVersion: kubectlVersion,
	}
}
