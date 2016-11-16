package gifutil

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
)

var palette = []color.Color{ color.White, color.Black, color.RGBA{ 0x3b, 0xaa, 0x5b, 0xff} }

const (
    whiteIndex = 0
    balckIndex = 1
    chateauGreenIndex = 2
)

func Lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )

    freq := rand.Float64() * 3.0
    anim := gif.GIF{ LoopCount: nframes }
    phase := 0.0
    for i := 1; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)
        for j := 0.0; j < cycles * 2 * math.Pi; j += res {
            x := math.Sin(j)
            y := math.Sin(j * freq + phase)
            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5),
                chateauGreenIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
