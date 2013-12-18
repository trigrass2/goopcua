package builtin

type Boolean bool
type SByte int8
type Byte uint8
type Int16 int16
type UInt16 uint16
type Int32 int32
type UInt32 uint32
type Int64 int64
type UInt64 uint64
type Float float32
type Double float64
type String string

type IdType int

const (
	Numeric IdType = iota
	String
	Guid
	Opaque
)

type NodeId struct {
	namespaceIndex UInt16
	identifierType IdType
	identifier     interface{}
}

type QualifiedName struct {
	namespaceIndex UInt16
	name           String
}

type LocaleId String

type LocalText struct {
	text   String
	locale LocaleId
}
