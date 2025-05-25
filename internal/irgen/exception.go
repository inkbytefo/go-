package irgen

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
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

	// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
	// personalityFunc := g.getPersonalityFunction()
	// catchType := types.NewStruct(types.I8Ptr, types.I32)
	// landingPadInst := g.currentBB.NewLandingPad(catchType, personalityFunc)
	// landingPadInst.Cleanup = true

	for range catchBlocks {
		// Catch tipi için bir filtre ekle
		// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
		// typeInfo := g.getTypeInfo(catch.ExceptionType)
		// landingPadInst.Clauses = append(landingPadInst.Clauses, ir.NewCatch(typeInfo))
	}

	// İstisna tipine göre uygun catch bloğuna git
	// Not: Basitleştirilmiş yaklaşım, gerçek implementasyonda daha karmaşık olabilir
	// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
	// exceptionPtr := g.currentBB.NewExtractValue(landingPadInst, 0)

	// Catch bloklarını işle
	for i, catch := range catchBlocks {
		// Catch bloğunu işle
		g.currentBB = catch.Block

		// İstisna değişkenini tanımla
		exceptionVar := g.currentBB.NewAlloca(catch.ExceptionType)
		exceptionVar.SetName(catch.Variable)

		// İstisna değerini değişkene ata
		// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
		// typedExceptionPtr := g.currentBB.NewBitCast(exceptionPtr, types.NewPointer(catch.ExceptionType))
		// exceptionVal := g.currentBB.NewLoad(catch.ExceptionType, typedExceptionPtr)
		// g.currentBB.NewStore(exceptionVal, exceptionVar)

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
	// TODO: LLVM API değişikliği nedeniyle geçici olarak devre dışı
	// g.currentBB.NewResume(landingPadInst)

	// End bloğuna geç
	g.currentBB = endBlock

	// İstisna bilgisini yığından çıkar
	g.exceptionStack = g.exceptionStack[:len(g.exceptionStack)-1]
}
