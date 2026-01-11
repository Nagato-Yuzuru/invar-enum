package renderer

import (
	"context"

	"github.com/Nagato-Yuzuru/invar-enum/pkg/strutil"
)

type GeneratedFile struct {
	FileName string
	Content  []byte
}
type FileModel struct {
	Path string // e.g. "pkg/payment/payment_invar.go"
	// PackageName Java / go need. e.g. "payment" (Go package) or namespace
	PackageName string
	Imports     []string // e.g. ["database/sql", "fmt"]

	Enums []EnumModel

	Options RenderOptions
}

type (
	EnumModel struct {
		// ProtoName original proto info
		ProtoName    string
		ProtoRawName string
		Name         strutil.CaseSet
		Comments     []string
		DefaultItem  *EnumItem
		Items        []EnumItem
		// Options if nil inherit from FileModel
		Options *RenderOptions
	}

	EnumItem struct {
		Key          strutil.CaseSet
		StrValue     string
		Numeric      int32 // Protobuf always positive
		Label        string
		Comments     []string
		IsDeprecated bool
		ProtoKey     string
	}
)

type RenderOptions struct {
	UseNumericValue  bool
	UseStrValue      bool
	GenerateLabel    bool
	GenerateProtoMap bool
	FileHeader       string
	Extras           ExtraOptions
}

// ExtraOptions Unuse in MVP
type ExtraOptions map[string]any

type Renderer interface {
	Render(ctx context.Context, files []FileModel) ([]GeneratedFile, error)
}
