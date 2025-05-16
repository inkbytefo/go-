package token

// Position, kaynak koddaki bir konumu temsil eder.
type Position struct {
	Line   int // Satır numarası (1-tabanlı)
	Column int // Sütun numarası (1-tabanlı)
	Offset int // Dosyadaki byte offset (0-tabanlı)
}

// IsValid, konumun geçerli olup olmadığını kontrol eder.
func (p Position) IsValid() bool {
	return p.Line > 0
}
