package encoding

import "time"

// 编码接口
type IEncoder interface {
	// Pushes a namespace into the namespace stack.
	PushNamespace(namespaceUri string)
	// Pops a namespace from the namespace stack.
	PopNamespace() string

	//Writes a boolean to the stream.
	WriteBoolean(fieldName string, value bool)

	//Writes a sbyte to the stream
	WriteSByte(fieldName string, value int8)

	//Writes a byte to the stream.
	WriteByte(fieldName string, value uint8)

	//Writes a short to the stream.
	WriteInt16(fieldName string, value int16)

	//Writes a ushort to the stream.
	WriteUInt16(fieldName string, value uint16)

	//Writes an int to the stream.
	WriteInt32(fieldName string, value int32)

	//Write an uint to the stream.
	WriteUInt32(fieldName string, value uint32)

	//Writes an long to stream.
	WriteInt64(fieldName string, value int64)

	//Writes an ulong to stream.
	WriteUInt64(fieldName string, value uint64)

	//Writes a float to the stream.
	WriteFloat(fieldName string, value float32)

	//Writes a double to the stream.
	WriteDouble(fieldName string, value float64)

	//Writes a string to the stream.
	WriteString(fieldName string, value string)

	//Writes a UTC date/time to the stream.
	WriteDateTime(fieldName string, value time.Time)
}
