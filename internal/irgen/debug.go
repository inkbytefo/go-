package irgen

import (
	"fmt"
	"goplus/internal/ast"
	"goplus/internal/token"
	"path/filepath"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// DebugInfo, hata ayıklama bilgisi üretimi için kullanılan yapıdır.
type DebugInfo struct {
	module           *ir.Module
	compileUnit      *metadata.DICompileUnit
	fileMetadata     map[string]*metadata.DIFile
	typeMetadata     map[types.Type]metadata.Metadata
	subprogramStack  []*metadata.DISubprogram
	lexicalBlockStack []*metadata.DILexicalBlock
	currentLine      int
	currentCol       int
	currentFile      string
}

// NewDebugInfo, yeni bir DebugInfo oluşturur.
func NewDebugInfo(module *ir.Module) *DebugInfo {
	return &DebugInfo{
		module:           module,
		fileMetadata:     make(map[string]*metadata.DIFile),
		typeMetadata:     make(map[types.Type]metadata.Metadata),
		subprogramStack:  make([]*metadata.DISubprogram, 0),
		lexicalBlockStack: make([]*metadata.DILexicalBlock, 0),
	}
}

// InitCompileUnit, derleme birimi meta verisini başlatır.
func (d *DebugInfo) InitCompileUnit(filename, directory, producer string, isOptimized bool, flags string, runtimeVersion int) {
	// Dosya meta verisini oluştur
	file := d.getOrCreateFileMetadata(filename, directory)

	// Derleme birimi meta verisini oluştur
	d.compileUnit = &metadata.DICompileUnit{
		Language:       metadata.DwarfLangC, // Şimdilik C olarak işaretliyoruz
		File:           file,
		Producer:       producer,
		IsOptimized:    isOptimized,
		Flags:          flags,
		RuntimeVersion: runtimeVersion,
	}

	// Modüle derleme birimi meta verisini ekle
	d.module.NamedMetadata["llvm.dbg.cu"] = &metadata.NamedMetadata{
		Name: "llvm.dbg.cu",
		Nodes: []metadata.Metadata{
			d.compileUnit,
		},
	}

	// Geçerli dosyayı ayarla
	d.currentFile = filename
}

// getOrCreateFileMetadata, dosya meta verisini döndürür veya oluşturur.
func (d *DebugInfo) getOrCreateFileMetadata(filename, directory string) *metadata.DIFile {
	key := filepath.Join(directory, filename)
	if file, exists := d.fileMetadata[key]; exists {
		return file
	}

	file := &metadata.DIFile{
		Filename:  filename,
		Directory: directory,
	}
	d.fileMetadata[key] = file
	return file
}

// CreateFunction, fonksiyon meta verisini oluşturur.
func (d *DebugInfo) CreateFunction(fn *ir.Func, name string, linkageName string, file *metadata.DIFile, line int, isLocal bool, isDefinition bool, scopeLine int, flags metadata.DIFlag, isOptimized bool) *metadata.DISubprogram {
	// Fonksiyon tipi meta verisini oluştur
	fnType := d.getOrCreateFunctionTypeMetadata(fn.Sig)

	// Fonksiyon meta verisini oluştur
	subprogram := &metadata.DISubprogram{
		Name:         name,
		LinkageName:  linkageName,
		File:         file,
		Line:         line,
		Type:         fnType,
		IsLocal:      isLocal,
		IsDefinition: isDefinition,
		ScopeLine:    scopeLine,
		Flags:        flags,
		IsOptimized:  isOptimized,
		Unit:         d.compileUnit,
	}

	// Fonksiyona meta veri ekle
	fn.Metadata = append(fn.Metadata, &metadata.Attachment{
		Name: "dbg",
		Node: subprogram,
	})

	// Fonksiyon yığınına ekle
	d.subprogramStack = append(d.subprogramStack, subprogram)

	return subprogram
}

// FinishFunction, fonksiyon meta verisini tamamlar.
func (d *DebugInfo) FinishFunction() {
	if len(d.subprogramStack) > 0 {
		d.subprogramStack = d.subprogramStack[:len(d.subprogramStack)-1]
	}
}

// CreateLexicalBlock, sözcüksel blok meta verisini oluşturur.
func (d *DebugInfo) CreateLexicalBlock(file *metadata.DIFile, line, column int) *metadata.DILexicalBlock {
	var scope metadata.Metadata
	if len(d.lexicalBlockStack) > 0 {
		scope = d.lexicalBlockStack[len(d.lexicalBlockStack)-1]
	} else if len(d.subprogramStack) > 0 {
		scope = d.subprogramStack[len(d.subprogramStack)-1]
	} else {
		scope = d.compileUnit
	}

	block := &metadata.DILexicalBlock{
		File:   file,
		Line:   line,
		Column: column,
		Scope:  scope,
	}

	d.lexicalBlockStack = append(d.lexicalBlockStack, block)
	return block
}

// FinishLexicalBlock, sözcüksel blok meta verisini tamamlar.
func (d *DebugInfo) FinishLexicalBlock() {
	if len(d.lexicalBlockStack) > 0 {
		d.lexicalBlockStack = d.lexicalBlockStack[:len(d.lexicalBlockStack)-1]
	}
}

// CreateLocalVariable, yerel değişken meta verisini oluşturur.
func (d *DebugInfo) CreateLocalVariable(name string, file *metadata.DIFile, line, column int, typ types.Type, argIndex int, align int) *metadata.DILocalVariable {
	var scope metadata.Metadata
	if len(d.lexicalBlockStack) > 0 {
		scope = d.lexicalBlockStack[len(d.lexicalBlockStack)-1]
	} else if len(d.subprogramStack) > 0 {
		scope = d.subprogramStack[len(d.subprogramStack)-1]
	} else {
		scope = d.compileUnit
	}

	typeMetadata := d.getOrCreateTypeMetadata(typ)

	localVar := &metadata.DILocalVariable{
		Name:     name,
		File:     file,
		Line:     line,
		Type:     typeMetadata,
		ArgIndex: argIndex,
		Align:    align,
		Scope:    scope,
	}

	return localVar
}

// InsertDeclare, değişken bildirimi meta verisini ekler.
func (d *DebugInfo) InsertDeclare(block *ir.Block, value value.Value, localVar *metadata.DILocalVariable, expr metadata.Metadata, line, column int) {
	// Geçerli dosya meta verisini al
	file := d.getOrCreateFileMetadata(d.currentFile, "")

	// Konum meta verisini oluştur
	location := &metadata.DILocation{
		Line:   line,
		Column: column,
		Scope:  localVar.Scope,
		File:   file,
	}

	// Declare intrinsic çağrısı oluştur
	declareFunc := d.getOrCreateIntrinsicFunction("llvm.dbg.declare", types.Void, types.Metadata, types.Metadata)
	call := block.NewCall(declareFunc, metadata.Value{Value: value}, metadata.Value{Value: localVar}, metadata.Value{Value: expr})

	// Çağrıya konum meta verisi ekle
	call.Metadata = append(call.Metadata, &metadata.Attachment{
		Name: "dbg",
		Node: location,
	})
}

// InsertValue, değer meta verisini ekler.
func (d *DebugInfo) InsertValue(block *ir.Block, value value.Value, localVar *metadata.DILocalVariable, expr metadata.Metadata, line, column int) {
	// Geçerli dosya meta verisini al
	file := d.getOrCreateFileMetadata(d.currentFile, "")

	// Konum meta verisini oluştur
	location := &metadata.DILocation{
		Line:   line,
		Column: column,
		Scope:  localVar.Scope,
		File:   file,
	}

	// Value intrinsic çağrısı oluştur
	valueFunc := d.getOrCreateIntrinsicFunction("llvm.dbg.value", types.Void, types.Metadata, types.Metadata, types.Metadata)
	call := block.NewCall(valueFunc, metadata.Value{Value: value}, metadata.Value{Value: localVar}, metadata.Value{Value: expr})

	// Çağrıya konum meta verisi ekle
	call.Metadata = append(call.Metadata, &metadata.Attachment{
		Name: "dbg",
		Node: location,
	})
}

// SetLocation, geçerli konum bilgisini ayarlar.
func (d *DebugInfo) SetLocation(line, column int, filename string) {
	d.currentLine = line
	d.currentCol = column
	if filename != "" {
		d.currentFile = filename
	}
}

// AttachLocation, bir talimata konum meta verisi ekler.
func (d *DebugInfo) AttachLocation(inst ir.Instruction) {
	// Geçerli dosya meta verisini al
	file := d.getOrCreateFileMetadata(d.currentFile, "")

	// Geçerli kapsamı belirle
	var scope metadata.Metadata
	if len(d.lexicalBlockStack) > 0 {
		scope = d.lexicalBlockStack[len(d.lexicalBlockStack)-1]
	} else if len(d.subprogramStack) > 0 {
		scope = d.subprogramStack[len(d.subprogramStack)-1]
	} else {
		scope = d.compileUnit
	}

	// Konum meta verisini oluştur
	location := &metadata.DILocation{
		Line:   d.currentLine,
		Column: d.currentCol,
		Scope:  scope,
		File:   file,
	}

	// Talimata konum meta verisi ekle
	inst.Metadata = append(inst.Metadata, &metadata.Attachment{
		Name: "dbg",
		Node: location,
	})
}

// getOrCreateIntrinsicFunction, intrinsic fonksiyonu döndürür veya oluşturur.
func (d *DebugInfo) getOrCreateIntrinsicFunction(name string, returnType types.Type, paramTypes ...types.Type) *ir.Func {
	// Fonksiyonu modülde ara
	for _, fn := range d.module.Funcs {
		if fn.Name() == name {
			return fn
		}
	}

	// Fonksiyonu oluştur
	fn := d.module.NewFunc(name, returnType, ir.NewParam("", paramTypes[0]))
	for i := 1; i < len(paramTypes); i++ {
		fn.Params = append(fn.Params, ir.NewParam("", paramTypes[i]))
	}
	return fn
}

// getOrCreateTypeMetadata, tip meta verisini döndürür veya oluşturur.
func (d *DebugInfo) getOrCreateTypeMetadata(typ types.Type) metadata.Metadata {
	if meta, exists := d.typeMetadata[typ]; exists {
		return meta
	}

	var typeMeta metadata.Metadata

	switch t := typ.(type) {
	case *types.VoidType:
		typeMeta = &metadata.DIBasicType{
			Name:    "void",
			Size:    0,
			Align:   0,
			Encoding: metadata.DwarfAteAddress,
		}
	case *types.IntType:
		typeMeta = &metadata.DIBasicType{
			Name:    fmt.Sprintf("i%d", t.BitSize),
			Size:    t.BitSize,
			Align:   t.BitSize / 8,
			Encoding: metadata.DwarfAteUnsigned,
		}
	case *types.FloatType:
		typeMeta = &metadata.DIBasicType{
			Name:    "float",
			Size:    32,
			Align:   4,
			Encoding: metadata.DwarfAteFloat,
		}
	case *types.DoubleType:
		typeMeta = &metadata.DIBasicType{
			Name:    "double",
			Size:    64,
			Align:   8,
			Encoding: metadata.DwarfAteFloat,
		}
	case *types.PointerType:
		elemTypeMeta := d.getOrCreateTypeMetadata(t.ElemType)
		typeMeta = &metadata.DIDerivedType{
			Tag:      metadata.DwarfTagPointerType,
			BaseType: elemTypeMeta,
			Size:     64, // 64-bit işaretçi
			Align:    8,
		}
	case *types.ArrayType:
		elemTypeMeta := d.getOrCreateTypeMetadata(t.ElemType)
		typeMeta = &metadata.DICompositeType{
			Tag:      metadata.DwarfTagArrayType,
			Elements: []metadata.Metadata{elemTypeMeta},
			Size:     0, // Boyut bilinmiyor
			Align:    0,
		}
	case *types.StructType:
		elements := make([]metadata.Metadata, len(t.Fields))
		for i, field := range t.Fields {
			elements[i] = d.getOrCreateTypeMetadata(field)
		}
		typeMeta = &metadata.DICompositeType{
			Tag:      metadata.DwarfTagStructType,
			Name:     "struct",
			Elements: elements,
			Size:     0, // Boyut bilinmiyor
			Align:    0,
		}
	case *types.FuncType:
		typeMeta = d.getOrCreateFunctionTypeMetadata(t)
	default:
		// Bilinmeyen tip için void döndür
		typeMeta = &metadata.DIBasicType{
			Name:    "unknown",
			Size:    0,
			Align:   0,
			Encoding: metadata.DwarfAteAddress,
		}
	}

	d.typeMetadata[typ] = typeMeta
	return typeMeta
}

// getOrCreateFunctionTypeMetadata, fonksiyon tipi meta verisini döndürür veya oluşturur.
func (d *DebugInfo) getOrCreateFunctionTypeMetadata(typ *types.FuncType) *metadata.DISubroutineType {
	// Dönüş tipi meta verisini al
	returnTypeMeta := d.getOrCreateTypeMetadata(typ.RetType)

	// Parametre tipleri meta verilerini al
	paramTypeMetas := make([]metadata.Metadata, len(typ.Params)+1)
	paramTypeMetas[0] = returnTypeMeta
	for i, paramType := range typ.Params {
		paramTypeMetas[i+1] = d.getOrCreateTypeMetadata(paramType)
	}

	// Fonksiyon tipi meta verisini oluştur
	return &metadata.DISubroutineType{
		Types: paramTypeMetas,
	}
}
