package builder

type GsJsonPatchBuilder struct {
}

func NewGsJsonPatchBuilder() *GsJsonPatchBuilder {
	return &GsJsonPatchBuilder{}
}

func (g *GsJsonPatchBuilder) Build() []byte {
	return []byte{}
}
