package semantic

// HashType, bir hash tipini temsil eder.
type HashType struct {
	KeyType   Type
	ValueType Type
}

// String, hash tipinin string temsilini döndürür.
func (ht *HashType) String() string {
	return "map[" + ht.KeyType.String() + "]" + ht.ValueType.String()
}

// Equals, iki hash tipinin eşit olup olmadığını kontrol eder.
func (ht *HashType) Equals(other Type) bool {
	if otherHash, ok := other.(*HashType); ok {
		return ht.KeyType.Equals(otherHash.KeyType) && ht.ValueType.Equals(otherHash.ValueType)
	}
	return false
}
