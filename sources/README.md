# Corresponding sources and build information

These source snapshots accompany the six FFmpeg executables in `../bin/`.
They are provided alongside the binaries rather than through a future written
offer. Preserve access to these sources and the applicable license notices
when redistributing the executables.

## Contents

- [ffmpeg-minimal-sources.tar.xz](ffmpeg-minimal-sources.tar.xz): the FFmpeg
  7.1.2, x264, and zlib 1.3.1 source trees used for the build, plus the original
  build scripts and verification programs.
- [zig-0.13.0-library-sources.tar.xz](zig-0.13.0-library-sources.tar.xz): the
  library source tree from the actual Zig 0.13.0 installation, including musl,
  MinGW-w64, compiler runtime, standard library, platform headers, and notices.
  The compiler executable is not included.
- [build-configs.tar.xz](build-configs.tar.xz): the original `config.mak` for
  each target, retained as a record of the actual build configuration.
- [ffmpeg-configure.patch](ffmpeg-configure.patch): changes relative to the
  pristine FFmpeg 7.1.2 release. These are already applied in the source archive.
- [SHA256SUMS](SHA256SUMS): checksums of the supplied source/configuration archives.
- [UPSTREAM-SHA256SUMS](UPSTREAM-SHA256SUMS): checksums of the original downloaded
  archives; these filenames refer to the original build workspace.
- [BINARY-SHA256SUMS](BINARY-SHA256SUMS): checksums of the **uncompressed**
  executables, named after the archives in `../bin/` without the `.gz` suffix.

The six binaries were matched byte-for-byte to the original build output
before the duplicate uncompressed files were removed from this repository.
The x264 snapshot is identified by its supplied sources and original archive
checksum, not by a moving upstream branch.

## Changes to upstream sources

On November 6, 2026, FFmpeg's `configure` was changed for Zig's linker:

1. Remove Darwin's `-Wl,-dynamic,-search_paths_first` options.
2. Remove Windows x86-64's `-Wl,--pic-executable,-e,mainCRTStartup` workaround.

No other existing FFmpeg source file differed from the downloaded release.
x264 and zlib source files matched their downloaded archives without changes.
The distributed `configure` also contains a dated comment documenting these
changes; that comment was added when packaging the sources and changes no
build behavior.

## Build

Use a Linux x86-64 host with Bash, make, pkg-config, Python 3, and tar/xz.
Extract both source archives into the same **new** directory:

```sh
sha256sum -c SHA256SUMS
mkdir rebuild
tar -xJf ffmpeg-minimal-sources.tar.xz -C rebuild
tar -xJf zig-0.13.0-library-sources.tar.xz -C rebuild
cd rebuild
```

Obtain the original Linux x86-64 [Zig 0.13.0 distribution](https://ziglang.org/download/0.13.0/zig-linux-x86_64-0.13.0.tar.xz),
verify it, and extract it into `work/` to supply the compiler executable:

```sh
sha256sum zig-linux-x86_64-0.13.0.tar.xz
# Expected: d45312e61ebcc48032b77bc4cf7fd6915c11fa16e4aad116b66c9468211230ea
tar -xJf zig-linux-x86_64-0.13.0.tar.xz -C work
```

Run the original build commands sequentially (each uses `make -j4`):

```sh
bash work/build-minimal.sh linux-x86_64 x86_64-linux-musl x86_64-linux-musl linux x86_64
bash work/build-minimal.sh linux-arm64 aarch64-linux-musl aarch64-linux-musl linux aarch64
bash work/build-minimal.sh windows-x86_64 x86_64-windows-gnu x86_64-w64-mingw32 mingw32 x86_64
bash work/build-minimal.sh windows-arm64 aarch64-windows-gnu aarch64-w64-mingw32 mingw32 aarch64
bash work/build-minimal.sh macos-x86_64 x86_64-macos.11.0 x86_64-apple-darwin darwin x86_64
bash work/build-minimal.sh macos-arm64 aarch64-macos.11.0 aarch64-apple-darwin darwin aarch64
```

Executables are written to `outputs/ffmpeg-minimal/<target>/`. The supplied
scripts reproduce the build procedure; identical output bytes across different
hosts or working-directory paths have not been verified.

To replace an embedded executable, gzip it with `gzip -n -9 -c` and use the
existing `bin/ffmpeg-<GOOS>-<GOARCH>[.exe].gz` filename. Update the corresponding
sources, checksums, and notices at the same time.
