package irgen

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// generateThrowStatement, bir throw deyimi için IR üretir.
func (g *IRGenerator) generateThrowStatement(stmt *ast.ThrowStatement) {
	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, throw deyimi değerlendirilemiyor")
		return
	}

	// Fırlatılacak değeri değerlendir
	exceptionVal := g.generateExpression(stmt.Value)
	if exceptionVal == nil {
		return
	}

	// İstisna fırlatma fonksiyonunu al
	throwFunc := g.getThrowFunction()

	// İstisna değerini void* tipine dönüştür
	exceptionPtr := g.currentBB.NewAlloca(exceptionVal.Type())
	g.currentBB.NewStore(exceptionVal, exceptionPtr)
	voidPtr := g.currentBB.NewBitCast(exceptionPtr, types.I8Ptr)

	// İstisna fırlatma fonksiyonunu çağır
	g.currentBB.NewCall(throwFunc, voidPtr)

	// Unreachable ekle
	g.currentBB.NewUnreachable()
}

// getPersonalityFunction, istisna işleme için personality fonksiyonunu döndürür.
func (g *IRGenerator) getPersonalityFunction() *ir.Func {
	// Personality fonksiyonunu bul veya oluştur
	personalityFunc := g.getFunction("__gxx_personality_v0")
	if personalityFunc == nil {
		// Personality fonksiyonunu tanımla
		personalityFunc = g.module.NewFunc("__gxx_personality_v0", types.I32,
			ir.NewParam("version", types.I32),
			ir.NewParam("actions", types.I32),
			ir.NewParam("exceptionClass", types.I64),
			ir.NewParam("exceptionObject", types.I8Ptr),
			ir.NewParam("context", types.I8Ptr),
		)
		g.symbolTable["__gxx_personality_v0"] = personalityFunc
	}
	return personalityFunc
}

// getThrowFunction, istisna fırlatma fonksiyonunu döndürür.
func (g *IRGenerator) getThrowFunction() *ir.Func {
	// Throw fonksiyonunu bul veya oluştur
	throwFunc := g.getFunction("__cxa_throw")
	if throwFunc == nil {
		// Throw fonksiyonunu tanımla
		throwFunc = g.module.NewFunc("__cxa_throw", types.Void,
			ir.NewParam("exceptionObject", types.I8Ptr),
			ir.NewParam("exceptionType", types.I8Ptr),
			ir.NewParam("destructor", types.I8Ptr),
		)
		g.symbolTable["__cxa_throw"] = throwFunc
	}
	return throwFunc
}

// getTypeInfo, bir tip için tip bilgisini döndürür.
func (g *IRGenerator) getTypeInfo(typ types.Type) value.Value {
	// Tip bilgisi için global değişken oluştur
	typeName := fmt.Sprintf("_ZTI%s", typ.String())
	typeInfo := g.module.NewGlobalDef(typeName, constant.NewNull(types.I8Ptr))
	return typeInfo
}

// generateTryExpression, bir try ifadesi için IR üretir.
func (g *IRGenerator) generateTryExpression(expr *ast.TryExpression) value.Value {
	if g.currentFunc == nil {
		g.ReportError("Geçerli bir fonksiyon yok, try ifadesi değerlendirilemiyor")
		return nil
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, try ifadesi değerlendirilemiyor")
		return nil
	}

	// Blokları oluştur
	tryBlock := g.currentFunc.NewBlock("try.expr")
	landingPad := g.currentFunc.NewBlock("landingpad.expr")
	resumeBlock := g.currentFunc.NewBlock("resume.expr")
	endBlock := g.currentFunc.NewBlock("try.expr.end")

	// Sonuç değişkeni oluştur
	resultType := g.getExpressionType(expr.Expression)
	if resultType == nil {
		resultType = types.I32 // Varsayılan olarak int32
	}
	resultVar := g.currentBB.NewAlloca(resultType)
	resultVar.SetName("try.result")

	// İstisna bilgisini oluştur
	exceptionInfo := &ExceptionInfo{
		TryBlock:     tryBlock,
		CatchBlocks:  []*CatchBlockInfo{},
		FinallyBlock: nil,
		LandingPad:   landingPad,
		ResumeBlock:  resumeBlock,
	}

	// İstisna bilgisini yığına ekle
	g.exceptionStack = append(g.exceptionStack, exceptionInfo)

	// Try bloğuna git
	g.currentBB.NewBr(tryBlock)

	// Try bloğunu işle
	g.currentBB = tryBlock
	result := g.generateExpression(expr.Expression)
	if result != nil {
		g.currentBB.NewStore(result, resultVar)
	}

	// End bloğuna git
	g.currentBB.NewBr(endBlock)

	// Landing pad bloğunu işle
	g.currentBB = landingPad

	// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
	// personalityFunc := g.getPersonalityFunction()
	// catchType := types.NewStruct(types.I8Ptr, types.I32)
	// landingPadInst := g.currentBB.NewLandingPad(catchType, personalityFunc)
	// landingPadInst.Cleanup = true

	// Resume bloğuna git
	g.currentBB.NewBr(resumeBlock)

	// Resume bloğunu işle
	g.currentBB = resumeBlock
	// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
	// g.currentBB.NewResume(landingPadInst)

	// End bloğuna geç
	g.currentBB = endBlock

	// İstisna bilgisini yığından çıkar
	g.exceptionStack = g.exceptionStack[:len(g.exceptionStack)-1]

	// Sonuç değerini yükle ve döndür
	return g.currentBB.NewLoad(resultType, resultVar)
}
