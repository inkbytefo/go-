# GO-Minus Crypto Package

This package provides cryptographic functionality for the GO-Minus programming language. It includes support for various cryptographic algorithms, secure random number generation, and cryptographic utilities.

## Features

- Symmetric encryption (AES, ChaCha20, etc.)
- Asymmetric encryption (RSA, ECDSA, etc.)
- Hash functions (SHA-256, SHA-3, BLAKE2, etc.)
- Message authentication codes (HMAC, Poly1305, etc.)
- Digital signatures
- Key derivation functions (PBKDF2, Argon2, etc.)
- Secure random number generation
- TLS/SSL support
- X.509 certificate handling
- Password hashing (bcrypt, scrypt, etc.)

## Usage

### Hash Functions

```go
import (
    "crypto/sha256"
    "fmt"
)

func main() {
    // Create a new hash
    hash := sha256.New()
    
    // Write data to the hash
    data := []byte("Hello, World!")
    hash.Write(data)
    
    // Get the hash sum
    sum := hash.Sum(nil)
    
    // Print the hash as a hexadecimal string
    fmt.Printf("SHA-256: %x\n", sum)
    
    // One-line hashing
    oneLineSum := sha256.Sum256(data)
    fmt.Printf("SHA-256 (one-line): %x\n", oneLineSum)
}
```

### HMAC

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "fmt"
)

func main() {
    // Create a key
    key := []byte("secret-key")
    
    // Create a new HMAC
    h := hmac.New(sha256.New, key)
    
    // Write data to the HMAC
    data := []byte("Hello, World!")
    h.Write(data)
    
    // Get the HMAC sum
    sum := h.Sum(nil)
    
    // Print the HMAC as a hexadecimal string
    fmt.Printf("HMAC-SHA256: %x\n", sum)
}
```

### AES Encryption

```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
    // Create a key (AES-256 requires a 32-byte key)
    key := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
        fmt.Println("Error generating key:", err)
        return
    }
    
    // Create a new cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println("Error creating cipher block:", err)
        return
    }
    
    // Create a plaintext
    plaintext := []byte("Hello, World!")
    
    // Create a ciphertext with room for the IV
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    
    // Create an IV
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println("Error generating IV:", err)
        return
    }
    
    // Encrypt the plaintext
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
    
    // Print the ciphertext
    fmt.Printf("Ciphertext: %x\n", ciphertext)
    
    // Decrypt the ciphertext
    decrypted := make([]byte, len(plaintext))
    decryptStream := cipher.NewCFBDecrypter(block, iv)
    decryptStream.XORKeyStream(decrypted, ciphertext[aes.BlockSize:])
    
    // Print the decrypted plaintext
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### RSA Encryption

```go
import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
)

func main() {
    // Generate a new RSA key pair
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println("Error generating key pair:", err)
        return
    }
    
    // Get the public key
    publicKey := &privateKey.PublicKey
    
    // Create a plaintext
    plaintext := []byte("Hello, World!")
    
    // Encrypt the plaintext
    ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
    if err != nil {
        fmt.Println("Error encrypting plaintext:", err)
        return
    }
    
    // Print the ciphertext
    fmt.Printf("Ciphertext: %x\n", ciphertext)
    
    // Decrypt the ciphertext
    decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
    if err != nil {
        fmt.Println("Error decrypting ciphertext:", err)
        return
    }
    
    // Print the decrypted plaintext
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### Digital Signatures

```go
import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
)

func main() {
    // Generate a new RSA key pair
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println("Error generating key pair:", err)
        return
    }
    
    // Get the public key
    publicKey := &privateKey.PublicKey
    
    // Create a message
    message := []byte("Hello, World!")
    
    // Hash the message
    hash := sha256.Sum256(message)
    
    // Sign the hash
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
    if err != nil {
        fmt.Println("Error signing message:", err)
        return
    }
    
    // Print the signature
    fmt.Printf("Signature: %x\n", signature)
    
    // Verify the signature
    err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
    if err != nil {
        fmt.Println("Error verifying signature:", err)
        return
    }
    
    fmt.Println("Signature verified successfully")
}
```

### Password Hashing

```go
import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "fmt"
    "golang.org/x/crypto/argon2"
)

func main() {
    // Password to hash
    password := []byte("password123")
    
    // Generate a random salt
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        fmt.Println("Error generating salt:", err)
        return
    }
    
    // Argon2 parameters
    time := uint32(3)
    memory := uint32(64 * 1024)
    threads := uint8(4)
    keyLen := uint32(32)
    
    // Hash the password
    hash := argon2.IDKey(password, salt, time, memory, threads, keyLen)
    
    // Encode the salt and hash
    saltBase64 := base64.StdEncoding.EncodeToString(salt)
    hashBase64 := base64.StdEncoding.EncodeToString(hash)
    
    // Print the encoded salt and hash
    fmt.Printf("Salt: %s\n", saltBase64)
    fmt.Printf("Hash: %s\n", hashBase64)
    
    // Verify the password
    // In a real application, you would retrieve the salt and hash from storage
    decodedSalt, _ := base64.StdEncoding.DecodeString(saltBase64)
    decodedHash, _ := base64.StdEncoding.DecodeString(hashBase64)
    
    // Hash the password with the same salt
    verifyHash := argon2.IDKey(password, decodedSalt, time, memory, threads, keyLen)
    
    // Compare the hashes
    if subtle.ConstantTimeCompare(decodedHash, verifyHash) == 1 {
        fmt.Println("Password verified successfully")
    } else {
        fmt.Println("Password verification failed")
    }
}
```

### Secure Random Numbers

```go
import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
)

func main() {
    // Generate a random byte slice
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        fmt.Println("Error generating random bytes:", err)
        return
    }
    
    // Print the random bytes
    fmt.Printf("Random bytes: %x\n", bytes)
    
    // Generate a random uint64
    var randomUint64 uint64
    binary.Read(rand.Reader, binary.BigEndian, &randomUint64)
    fmt.Printf("Random uint64: %d\n", randomUint64)
    
    // Generate a random number in a range (0-99)
    randomInRange := int(randomUint64 % 100)
    fmt.Printf("Random number in range (0-99): %d\n", randomInRange)
}
```

## Packages

### crypto/aes

The `aes` package implements the Advanced Encryption Standard (AES) encryption algorithm.

```go
// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (cipher.Block, error)

// BlockSize is the AES block size in bytes.
const BlockSize = 16
```

### crypto/cipher

The `cipher` package implements standard block cipher modes and stream ciphers.

```go
// Block represents an implementation of block cipher using a given key.
interface Block {
    // BlockSize returns the cipher's block size.
    func BlockSize() int
    
    // Encrypt encrypts the first block in src into dst.
    func Encrypt(dst, src []byte)
    
    // Decrypt decrypts the first block in src into dst.
    func Decrypt(dst, src []byte)
}

// NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining mode.
func NewCBCEncrypter(b Block, iv []byte) BlockMode

// NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining mode.
func NewCBCDecrypter(b Block, iv []byte) BlockMode

// NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode.
func NewCFBEncrypter(block Block, iv []byte) Stream

// NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode.
func NewCFBDecrypter(block Block, iv []byte) Stream

// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode.
func NewGCM(cipher Block) (AEAD, error)
```

### crypto/rsa

The `rsa` package implements RSA encryption as specified in PKCS #1.

```go
// GenerateKey generates an RSA keypair of the given bit size.
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)

// EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS #1 v1.5.
func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS #1 v1.5.
func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)

// SignPKCS1v15 calculates the signature of hashed using RSASSA-PKCS1-V1_5-SIGN.
func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)

// VerifyPKCS1v15 verifies an RSA PKCS #1 v1.5 signature.
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
```

### crypto/sha256

The `sha256` package implements the SHA-224 and SHA-256 hash algorithms.

```go
// New returns a new hash.Hash computing the SHA-256 checksum.
func New() hash.Hash

// Sum256 returns the SHA-256 checksum of the data.
func Sum256(data []byte) [32]byte
```

### crypto/hmac

The `hmac` package implements the Keyed-Hash Message Authentication Code (HMAC).

```go
// New returns a new HMAC hash using the given hash function and key.
func New(h func() hash.Hash, key []byte) hash.Hash
```

### crypto/rand

The `rand` package implements a cryptographically secure random number generator.

```go
// Read fills the given byte slice with random bytes.
func Read(b []byte) (n int, err error)

// Reader is a global, shared instance of a cryptographically secure random number generator.
var Reader io.Reader
```

## Error Handling

The crypto package uses GO-Minus's exception handling mechanism for error handling.

```go
import (
    "crypto/aes"
    "fmt"
)

func main() {
    try {
        // Invalid key size (AES-256 requires a 32-byte key)
        key := []byte("too-short-key")
        
        // This will fail because the key is too short
        _, err := aes.NewCipher(key)
        if err != nil {
            throw err
        }
    } catch (err) {
        fmt.Println("Crypto error:", err)
    }
}
```
