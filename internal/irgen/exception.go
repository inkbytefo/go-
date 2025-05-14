package irgen

import (
	"fmt"
	"goplus/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// ExceptionInfo, bir istisna bloğu hakkında bilgi tutar.
type ExceptionInfo struct {
	TryBlock     *ir.Block
	CatchBlocks  []*CatchBlockInfo
	FinallyBlock *ir.Block
	LandingPad   *ir.Block
	ResumeBlock  *ir.Block
}

// CatchBlockInfo, bir catch bloğu hakkında bilgi tutar.
type CatchBlockInfo struct {
	ExceptionType types.Type
	Block         *ir.Block
	Variable      string
}

// generateTryCatchStatement, bir try-catch deyimi için IR üretir.
func (g *IRGenerator) generateTryCatchStatement(stmt *ast.TryCatchStatement) {
	if g.currentFunc == nil {
		g.ReportError("Geçerli bir fonksiyon yok, try-catch deyimi değerlendirilemiyor")
		return
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, try-catch deyimi değerlendirilemiyor")
		return
	}

	// Blokları oluştur
	tryBlock := g.currentFunc.NewBlock("try")
	landingPad := g.currentFunc.NewBlock("landingpad")
	resumeBlock := g.currentFunc.NewBlock("resume")
	endBlock := g.currentFunc.NewBlock("try.end")

	// Catch blokları oluştur
	catchBlocks := make([]*CatchBlockInfo, len(stmt.Catches))
	for i, catch := range stmt.Catches {
		catchBlock := g.currentFunc.NewBlock(fmt.Sprintf("catch.%d", i))
		
		// İstisna tipini belirle
		var exceptionType types.Type = types.I8Ptr // Varsayılan olarak void*
		if catch.Type != nil {
			if typeIdent, ok := catch.Type.(*ast.Identifier); ok {
				if t, exists := g.typeTable[typeIdent.Value]; exists {
					exceptionType = t
				}
			}
		}
		
		catchBlocks[i] = &CatchBlockInfo{
			ExceptionType: exceptionType,
			Block:         catchBlock,
			Variable:      catch.Parameter.Value,
		}
	}

	// Finally bloğu oluştur
	var finallyBlock *ir.Block
	if stmt.Finally != nil {
		finallyBlock = g.currentFunc.NewBlock("finally")
	}

	// İstisna bilgisini oluştur
	exceptionInfo := &ExceptionInfo{
		TryBlock:     tryBlock,
		CatchBlocks:  catchBlocks,
		FinallyBlock: finallyBlock,
		LandingPad:   landingPad,
		ResumeBlock:  resumeBlock,
	}

	// İstisna bilgisini yığına ekle
	g.exceptionStack = append(g.exceptionStack, exceptionInfo)

	// Try bloğuna git
	g.currentBB.NewBr(tryBlock)

	// Try bloğunu işle
	g.currentBB = tryBlock
	g.generateBlockStatement(stmt.Try)

	// Eğer try bloğu bir dönüş ifadesi ile bitmiyorsa, finally bloğuna veya end bloğuna git
	if g.currentBB.Term == nil {
		if finallyBlock != nil {
			g.currentBB.NewBr(finallyBlock)
		} else {
			g.currentBB.NewBr(endBlock)
		}
	}

	// Landing pad bloğunu işle
	g.currentBB = landingPad
	
	// Personality fonksiyonunu al
	personalityFunc := g.getPersonalityFunction()
	
	// Landing pad oluştur
	// Not: LLVM IR'da landing pad, istisna yakalama mekanizmasıdır
	// Burada basitleştirilmiş bir yaklaşım kullanıyoruz
	catchType := types.NewStruct(types.I8Ptr, types.I32) // {i8*, i32} tipinde
	landingPadInst := g.currentBB.NewLandingPad(catchType, personalityFunc, 0)
	
	// Catch tipleri için cleanup ve catch filtreleri ekle
	landingPadInst.Cleanup = true
	
	for _, catch := range catchBlocks {
		// Catch tipi için bir filtre ekle
		typeInfo := g.getTypeInfo(catch.ExceptionType)
		landingPadInst.Clauses = append(landingPadInst.Clauses, ir.NewCatch(typeInfo))
	}
	
	// İstisna tipine göre uygun catch bloğuna git
	// Not: Basitleştirilmiş yaklaşım, gerçek implementasyonda daha karmaşık olabilir
	exceptionPtr := g.currentBB.NewExtractValue(landingPadInst, 0) // {i8*, i32} tipinden i8* al
	
	// Catch bloklarını işle
	for i, catch := range catchBlocks {
		// Catch bloğunu işle
		g.currentBB = catch.Block
		
		// İstisna değişkenini tanımla
		exceptionVar := g.currentBB.NewAlloca(catch.ExceptionType)
		exceptionVar.SetName(catch.Variable)
		
		// İstisna değerini değişkene ata
		typedExceptionPtr := g.currentBB.NewBitCast(exceptionPtr, types.NewPointer(catch.ExceptionType))
		exceptionVal := g.currentBB.NewLoad(catch.ExceptionType, typedExceptionPtr)
		g.currentBB.NewStore(exceptionVal, exceptionVar)
		
		// Değişkeni sembol tablosuna ekle
		g.symbolTable[catch.Variable] = exceptionVar
		
		// Catch bloğunu işle
		g.generateBlockStatement(stmt.Catches[i].Body)
		
		// Eğer catch bloğu bir dönüş ifadesi ile bitmiyorsa, finally bloğuna veya end bloğuna git
		if g.currentBB.Term == nil {
			if finallyBlock != nil {
				g.currentBB.NewBr(finallyBlock)
			} else {
				g.currentBB.NewBr(endBlock)
			}
		}
	}
	
	// Finally bloğunu işle
	if finallyBlock != nil {
		g.currentBB = finallyBlock
		g.generateBlockStatement(stmt.Finally)
		
		// Eğer finally bloğu bir dönüş ifadesi ile bitmiyorsa, end bloğuna git
		if g.currentBB.Term == nil {
			g.currentBB.NewBr(endBlock)
		}
	}
	
	// Resume bloğunu işle
	g.currentBB = resumeBlock
	g.currentBB.NewResume(landingPadInst)
	
	// End bloğuna geç
	g.currentBB = endBlock
	
	// İstisna bilgisini yığından çıkar
	g.exceptionStack = g.exceptionStack[:len(g.exceptionStack)-1]
}

// generateThrowStatement, bir throw deyimi için IR üretir.
func (g *IRGenerator) generateThrowStatement(stmt *ast.ThrowStatement) {
	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, throw deyimi değerlendirilemiyor")
		return
	}
	
	// Fırlatılacak değeri değerlendir
	exceptionVal := g.generateExpression(stmt.Exception)
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
	
	// Personality fonksiyonunu al
	personalityFunc := g.getPersonalityFunction()
	
	// Landing pad oluştur
	catchType := types.NewStruct(types.I8Ptr, types.I32) // {i8*, i32} tipinde
	landingPadInst := g.currentBB.NewLandingPad(catchType, personalityFunc, 0)
	landingPadInst.Cleanup = true
	
	// Resume bloğuna git
	g.currentBB.NewBr(resumeBlock)
	
	// Resume bloğunu işle
	g.currentBB = resumeBlock
	g.currentBB.NewResume(landingPadInst)
	
	// End bloğuna geç
	g.currentBB = endBlock
	
	// İstisna bilgisini yığından çıkar
	g.exceptionStack = g.exceptionStack[:len(g.exceptionStack)-1]
	
	// Sonuç değerini yükle ve döndür
	return g.currentBB.NewLoad(resultType, resultVar)
}
