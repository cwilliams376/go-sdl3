//go:build darwin && arm64

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img_arm64.dylib
	imgBlob    []byte
	imgLibName = "libSDL3_image.dylib"
)
