package badge

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"io/ioutil"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/google/uuid"
	"github.com/volvofixthis/repository-badge/internal/domain/service"
	"golang.org/x/image/font"
)

//go:embed assets/*
var assetsFS embed.FS

type BadgeUsecaseI interface {
	GetBuildBadge(scopeID uuid.UUID) ([]byte, error)
	GetCoverageBadge(scopeID uuid.UUID) ([]byte, error)
}

type BadgeUsecase struct {
	ss service.ScopeServiceI
}

func loadImageFromFS(fs fs.FS, filename string) (image.Image, error) {
	f, err := assetsFS.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}

func loadFontFaceFromFS(fs fs.FS, filename string, points float64) (font.Face, error) {
	f, err := assetsFS.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fontBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	ft, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(ft, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}

func (bu *BadgeUsecase) GetBuildBadge(scopeID uuid.UUID) ([]byte, error) {
	bgImage := loadImageFromFS(assetsFS, "assets/backgrounds/build_passed.png")
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	face, err := loadFontFaceFromFS(assetsFS, "assets/fonts/TerminusTTF.ttf", 12.0)
	if err != nil {
		return nil, err
	}
	dc.SetFontFace(face)

	x := float64(imgWidth/2) + 10
	y := float64((imgHeight / 2) - 3)
	maxWidth := float64(imgWidth) - 10
	dc.SetColor(color.RGBA{255, 0, 0, 255})
	dc.DrawStringWrapped("passed", x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, dc.Image())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (bu *BadgeUsecase) GetCoverageBadge(scopeID uuid.UUID) ([]byte, error) {
	bgImage, err := loadImageFromFS(assetsFS, "assets/backgrounds/test_coverage.png")
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	face, err := loadFontFaceFromFS(assetsFS, "assets/fonts/TerminusTTF.ttf", 12.0)
	if err != nil {
		return nil, err
	}
	dc.SetFontFace(face)

	x := float64(imgWidth/2) + 30
	y := float64((imgHeight / 2) - 3)
	maxWidth := float64(imgWidth) - 10
	dc.SetColor(color.White)
	dc.DrawStringWrapped("1%", x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, dc.Image())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewBadgeUsecase() BadgeUsecaseI {
	return &BadgeUsecase{
		ss: service.NewScopeService(),
	}
}
