# GO-Minus Encoding Package

This package provides encoding and decoding functionality for various data formats in the GO-Minus programming language. It includes support for common formats such as JSON, XML, CSV, Base64, and more.

## Features

- JSON encoding and decoding
- XML encoding and decoding
- CSV encoding and decoding
- Base64 encoding and decoding
- Hex encoding and decoding
- URL encoding and decoding
- YAML encoding and decoding
- Protocol Buffers encoding and decoding
- MessagePack encoding and decoding
- BSON encoding and decoding
- TOML encoding and decoding

## Usage

### JSON

```go
import (
    "encoding/json"
    "fmt"
)

// Define a struct
class Person {
    public:
        var Name string `json:"name"`
        var Age int `json:"age"`
        var Email string `json:"email,omitempty"`
}

func main() {
    // Create a Person
    person := Person{
        Name: "John Doe",
        Age: 30,
        Email: "john@example.com",
    }
    
    // Marshal (encode) to JSON
    data, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }
    
    fmt.Println("JSON:", string(data))
    
    // Unmarshal (decode) from JSON
    var decodedPerson Person
    err = json.Unmarshal(data, &decodedPerson)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    
    fmt.Println("Decoded person:", decodedPerson.Name, decodedPerson.Age, decodedPerson.Email)
    
    // Pretty print JSON
    prettyData, err := json.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Println("Error marshaling JSON with indent:", err)
        return
    }
    
    fmt.Println("Pretty JSON:")
    fmt.Println(string(prettyData))
}
```

### XML

```go
import (
    "encoding/xml"
    "fmt"
)

// Define a struct
class Person {
    public:
        var Name string `xml:"name"`
        var Age int `xml:"age"`
        var Email string `xml:"email,omitempty"`
}

func main() {
    // Create a Person
    person := Person{
        Name: "John Doe",
        Age: 30,
        Email: "john@example.com",
    }
    
    // Marshal (encode) to XML
    data, err := xml.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling XML:", err)
        return
    }
    
    fmt.Println("XML:", string(data))
    
    // Unmarshal (decode) from XML
    var decodedPerson Person
    err = xml.Unmarshal(data, &decodedPerson)
    if err != nil {
        fmt.Println("Error unmarshaling XML:", err)
        return
    }
    
    fmt.Println("Decoded person:", decodedPerson.Name, decodedPerson.Age, decodedPerson.Email)
    
    // Pretty print XML
    prettyData, err := xml.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Println("Error marshaling XML with indent:", err)
        return
    }
    
    fmt.Println("Pretty XML:")
    fmt.Println(string(prettyData))
}
```

### CSV

```go
import (
    "encoding/csv"
    "fmt"
    "strings"
)

func main() {
    // Create CSV data
    csvData := `Name,Age,Email
John Doe,30,john@example.com
Jane Smith,25,jane@example.com
Bob Johnson,40,bob@example.com`
    
    // Create a CSV reader
    reader := csv.NewReader(strings.NewReader(csvData))
    
    // Read all records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV:", err)
        return
    }
    
    // Print records
    for i, record := range records {
        if i == 0 {
            fmt.Println("Header:", record)
        } else {
            fmt.Printf("Record %d: %v\n", i, record)
        }
    }
    
    // Create new CSV data
    newRecords := [][]string{
        {"Name", "Age", "Email"},
        {"Alice", "28", "alice@example.com"},
        {"Bob", "35", "bob@example.com"},
    }
    
    // Create a string builder
    var sb strings.Builder
    
    // Create a CSV writer
    writer := csv.NewWriter(&sb)
    
    // Write all records
    err = writer.WriteAll(newRecords)
    if err != nil {
        fmt.Println("Error writing CSV:", err)
        return
    }
    
    // Print the CSV data
    fmt.Println("Generated CSV:")
    fmt.Println(sb.String())
}
```

### Base64

```go
import (
    "encoding/base64"
    "fmt"
)

func main() {
    // Original data
    data := []byte("Hello, World!")
    
    // Encode to Base64
    encoded := base64.StdEncoding.EncodeToString(data)
    fmt.Println("Base64 encoded:", encoded)
    
    // Decode from Base64
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Error decoding Base64:", err)
        return
    }
    
    fmt.Println("Base64 decoded:", string(decoded))
    
    // URL-safe Base64 encoding
    urlEncoded := base64.URLEncoding.EncodeToString(data)
    fmt.Println("URL-safe Base64 encoded:", urlEncoded)
    
    // URL-safe Base64 decoding
    urlDecoded, err := base64.URLEncoding.DecodeString(urlEncoded)
    if err != nil {
        fmt.Println("Error decoding URL-safe Base64:", err)
        return
    }
    
    fmt.Println("URL-safe Base64 decoded:", string(urlDecoded))
}
```

### Hex

```go
import (
    "encoding/hex"
    "fmt"
)

func main() {
    // Original data
    data := []byte("Hello, World!")
    
    // Encode to Hex
    encoded := hex.EncodeToString(data)
    fmt.Println("Hex encoded:", encoded)
    
    // Decode from Hex
    decoded, err := hex.DecodeString(encoded)
    if err != nil {
        fmt.Println("Error decoding Hex:", err)
        return
    }
    
    fmt.Println("Hex decoded:", string(decoded))
}
```

## Classes and Interfaces

### JSON Package

```go
// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error)

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error

// MarshalIndent is like Marshal but applies Indent to format the output.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

// Encoder writes JSON values to an output stream.
class Encoder {
    // NewEncoder returns a new encoder that writes to w.
    static func NewEncoder(w io.Writer) *Encoder
    
    // Encode writes the JSON encoding of v to the stream.
    func Encode(v interface{}) error
}

// Decoder reads and decodes JSON values from an input stream.
class Decoder {
    // NewDecoder returns a new decoder that reads from r.
    static func NewDecoder(r io.Reader) *Decoder
    
    // Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
    func Decode(v interface{}) error
}
```

### XML Package

```go
// Marshal returns the XML encoding of v.
func Marshal(v interface{}) ([]byte, error)

// Unmarshal parses the XML-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error

// MarshalIndent is like Marshal but applies Indent to format the output.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

// Encoder writes XML values to an output stream.
class Encoder {
    // NewEncoder returns a new encoder that writes to w.
    static func NewEncoder(w io.Writer) *Encoder
    
    // Encode writes the XML encoding of v to the stream.
    func Encode(v interface{}) error
}

// Decoder reads and decodes XML values from an input stream.
class Decoder {
    // NewDecoder returns a new decoder that reads from r.
    static func NewDecoder(r io.Reader) *Decoder
    
    // Decode reads the next XML-encoded value from its input and stores it in the value pointed to by v.
    func Decode(v interface{}) error
}
```

### CSV Package

```go
// Reader reads records from a CSV-encoded file.
class Reader {
    // NewReader returns a new Reader that reads from r.
    static func NewReader(r io.Reader) *Reader
    
    // Read reads one record (a slice of fields) from r.
    func Read() ([]string, error)
    
    // ReadAll reads all the remaining records from r.
    func ReadAll() ([][]string, error)
}

// Writer writes records to a CSV-encoded file.
class Writer {
    // NewWriter returns a new Writer that writes to w.
    static func NewWriter(w io.Writer) *Writer
    
    // Write writes a single CSV record to w along with any necessary quoting.
    func Write(record []string) error
    
    // WriteAll writes multiple CSV records to w using Write.
    func WriteAll(records [][]string) error
}
```

### Base64 Package

```go
// StdEncoding is the standard base64 encoding.
var StdEncoding *Encoding

// URLEncoding is the alternate base64 encoding defined in RFC 4648.
var URLEncoding *Encoding

// Encoding is a base64 encoding.
class Encoding {
    // EncodeToString returns the base64 encoding of src.
    func EncodeToString(src []byte) string
    
    // DecodeString returns the bytes represented by the base64 string s.
    func DecodeString(s string) ([]byte, error)
    
    // Encode encodes src using the encoding enc, writing EncodedLen(len(src)) bytes to dst.
    func Encode(dst, src []byte)
    
    // Decode decodes src using the encoding enc.
    func Decode(dst, src []byte) (n int, err error)
}
```

### Hex Package

```go
// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string

// DecodeString returns the bytes represented by the hexadecimal string s.
func DecodeString(s string) ([]byte, error)

// Encode encodes src into dst, returning the number of bytes written to dst.
func Encode(dst, src []byte) int

// Decode decodes src into dst, returning the number of bytes written to dst.
func Decode(dst, src []byte) (int, error)
```

## Error Handling

The encoding package uses GO-Minus's exception handling mechanism for error handling.

```go
import (
    "encoding/json"
    "fmt"
)

func main() {
    try {
        // Invalid JSON data
        data := []byte(`{"name": "John", "age": 30,}`)
        
        var person Person
        err := json.Unmarshal(data, &person)
        if err != nil {
            throw err
        }
        
        fmt.Println("Decoded person:", person.Name, person.Age)
    } catch (err) {
        fmt.Println("Error:", err)
    }
}
```
