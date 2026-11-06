package ffmpeg

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetFFmpeg(t *testing.T) {
	binary, err := GetFFmpeg()
	supportedOS := runtime.GOOS == "linux" || runtime.GOOS == "darwin" || runtime.GOOS == "windows"
	supportedArch := runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	if !supportedOS || !supportedArch {
		if !errors.Is(err, ErrUnsupported) || len(binary) != 0 {
			t.Fatalf("want empty binary and ErrUnsupported, got %d bytes and %v", len(binary), err)
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	sourceOS := runtime.GOOS
	if sourceOS == "darwin" {
		sourceOS = "macos"
	}
	sourceArch := runtime.GOARCH
	if sourceArch == "amd64" {
		sourceArch = "x86_64"
	}
	name := "ffmpeg"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	want, err := os.ReadFile(filepath.Join(sourceOS+"-"+sourceArch, name))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(binary, want) {
		t.Fatal("unpacked FFmpeg does not match the original executable")
	}
	// Each call must return independent bytes that the caller can modify.
	binary[0] ^= 0xff
	again, err := GetFFmpeg()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(again, want) {
		t.Fatal("modifying one result changed a subsequent result")
	}
}

func TestUnpackedFFmpegRuns(t *testing.T) {
	binary, err := GetFFmpeg()
	if errors.Is(err, ErrUnsupported) {
		t.Skip(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	name := "ffmpeg"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, binary, 0o700); err != nil {
		t.Fatal(err)
	}
	out, err := exec.Command(path, "-version").CombinedOutput()
	if err != nil {
		t.Fatalf("%v\n%s", err, out)
	}
	if !bytes.Contains(out, []byte("ffmpeg version 7.1.2")) {
		t.Fatalf("unexpected version output: %s", out)
	}
}
