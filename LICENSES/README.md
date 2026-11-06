# Licenses and third-party notices

The Go wrapper in this repository is licensed under GNU GPL version 2 or, at
your option, any later version (SPDX: `GPL-2.0-or-later`). The full license is
in [../LICENSE](../LICENSE). The FFmpeg executables include x264 and are also
distributed under GPL version 2 or later. This does not replace the original
copyright notices or permissive licenses of their constituent code.

| Component | License / notices |
|---|---|
| FFmpeg 7.1.2 | [Upstream license overview](FFmpeg-LICENSE.md), [GPL v2](../LICENSE), [LGPL v2.1](LGPL-2.1.txt), [additional notices](FFmpeg-NOTICES.txt) |
| x264 (the supplied source snapshot) | [GPL v2](x264-GPL-2.0.txt), [additional notices](x264-NOTICES.txt); source headers permit later GPL versions |
| zlib 1.3.1 | [zlib license](zlib.txt) |
| musl from Zig 0.13.0 (Linux static runtime) | [Copyright and MIT license](musl-COPYRIGHT.txt), [additional notices](musl-NOTICES.txt) |
| MinGW-w64 from Zig 0.13.0 (Windows runtime) | [COPYING / ZPL v2.1](mingw-w64-COPYING.txt), [per-file notices](mingw-w64-NOTICES.txt) |
| Zig 0.13.0 compiler runtime and standard library | [MIT license](Zig-MIT.txt), [runtime notices](Zig-runtime-NOTICES.txt), [standard-library notices](Zig-std-NOTICES.txt) |
| LLVM-derived runtime code | [Legacy compiler-rt licenses](LLVM-compiler-rt-8-LICENSE.txt), [Apache v2.0 with LLVM exceptions](LLVM-Apache-2.0-with-exceptions.txt) |
| Go-derived math code referenced by Zig | [BSD 3-clause license](Go-BSD-3-Clause.txt) |

The bundled FFmpeg sources are copyright their respective authors, including
the FFmpeg developers (2000–2025). x264 source headers identify the x264 project
(2003–2024) and individual contributors. Original per-file copyright notices
and license terms are retained in the [source archives](../sources/README.md).

This software is based in part on the work of the Independent JPEG Group.
The FFmpeg snapshot retains the upstream IJG-derived files unchanged.

The `*-NOTICES.txt` files conservatively reproduce permissive notices from the
supplied source trees, including optional code not necessarily linked into
these executables. They are not a list of enabled codecs or runtime features.
The complete source files, including their original notices, are also provided.

The toolchain source archive includes library code for additional targets and
unused libraries. Those files retain their own license texts; their inclusion
does not imply that these executables link against all of them.

Applications that redistribute an extracted or embedded FFmpeg executable
must also supply the applicable license texts, copyright notices, and a
GPL-compliant way to obtain its corresponding sources. Importing this Go
package does not automatically include these documentation files in a built
application.
