// Package ffmpeg provides a platform-specific embedded FFmpeg executable.
package ffmpeg

import (
	"errors"
	"runtime"
)

//go:generate sh scripts/compress-binaries.sh

// ErrUnsupported indicates that no FFmpeg binary is available for this platform.
var ErrUnsupported = errors.New("ffmpeg: no binary embedded for " + runtime.GOOS + "/" + runtime.GOARCH)
