package image

type PreviewGenerator struct{}

func (g *PreviewGenerator) GenerateImageUrl(url string, height, width int) string {
	return url
}

func NewImgGenerator() *PreviewGenerator {
	return &PreviewGenerator{}
}

var DefaultGererator interface {
	GenerateImageUrl(url string, height, width int) string
} = NewImgGenerator()
