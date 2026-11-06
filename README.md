# ffmpeg-stripped

A Go library that embeds stripped-down FFmpeg 7.1.2 executables for local video
processing. Only the target platform’s gzip-compressed binary is embedded and
unpacked on demand. No CGO, separate FFmpeg installation, or runtime downloads.

## Install

```sh
go get github.com/Warp-net/ffmpeg-stripped
```

## Usage

```go
import ffmpeg "github.com/Warp-net/ffmpeg-stripped"

binary, err := ffmpeg.GetFFmpeg()
if err != nil {
    return err
}
// Use ffmpeg.exe on Windows. Choose a writable location for the executable.
if err := os.WriteFile("ffmpeg", binary, 0o700); err != nil {
    return err
}
cmd := exec.Command("./ffmpeg", "-version")
cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
return cmd.Run()
```

The example also uses `os` and `os/exec`. Each `GetFFmpeg()` call returns a fresh
`[]byte`. Linux, macOS, and Windows support both `amd64` and `arm64`; Android,
iOS, and other targets return `ffmpeg.ErrUnsupported`.

## Included features

- Local MP4/MOV input with H.264 video and PNG images.
- Frame-rate conversion, image overlays, and MP4/H.264 output via x264.
- Copying MP4-compatible audio without decoding or re-encoding.

Other video decoders, networking, capture devices, `ffprobe`, and `ffplay` are
excluded. x264 supports 8-bit YUV 4:2:0 only. Assembly optimizations are disabled,
which can reduce performance.

Linux binaries are fully static (musl). Windows requires system DLLs, including
the Universal CRT. macOS requires version 11 or later; binaries are not notarized.
No separate FFmpeg, x264, or zlib shared libraries are needed.

## License

The Go wrapper and bundled FFmpeg/x264 executables are distributed under
GPL-2.0-or-later; see [LICENSE](LICENSE). Third-party license texts and notices
are in [LICENSES](LICENSES/README.md). Corresponding sources, build instructions,
and checksums are in [sources](sources/README.md).

When redistributing the embedded executables, preserve the applicable licenses
and notices and provide corresponding source code as required by the GPL.
