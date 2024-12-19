// @generated by the gondn_tlv_gen, DO NOT modify manually
package crdt

import (
	"encoding/binary"
	"io"
	"strings"

	enc "github.com/zjkmxy/go-ndn/pkg/encoding"
)

type IDTypeEncoder struct {
	length uint
}

type IDTypeParsingContext struct {
}

func (encoder *IDTypeEncoder) Init(value *IDType) {

	l := uint(0)
	l += 1
	switch x := value.Producer; {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += 1
	switch x := value.Clock; {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	encoder.length = l

}

func (context *IDTypeParsingContext) Init() {

}

func (encoder *IDTypeEncoder) EncodeInto(value *IDType, buf []byte) {

	pos := uint(0)

	buf[pos] = byte(161)
	pos += 1
	switch x := value.Producer; {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	buf[pos] = byte(163)
	pos += 1
	switch x := value.Clock; {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
}

func (encoder *IDTypeEncoder) Encode(value *IDType) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *IDTypeParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*IDType, error) {
	if reader == nil {
		return nil, enc.ErrBufferOverflow
	}

	var handled_Producer bool = false
	var handled_Clock bool = false

	progress := -1
	_ = progress

	value := &IDType{}
	var err error
	var startPos int
	for {
		startPos = reader.Pos()
		if startPos >= reader.Length() {
			break
		}
		typ := enc.TLNum(0)
		l := enc.TLNum(0)
		typ, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}
		l, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}

		err = nil
		if handled := false; true {
			switch typ {
			case 161:
				if true {
					handled = true
					handled_Producer = true
					value.Producer = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x := byte(0)
							x, err = reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.Producer = uint64(value.Producer<<8) | uint64(x)
						}
					}
				}
			case 163:
				if true {
					handled = true
					handled_Clock = true
					value.Clock = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x := byte(0)
							x, err = reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.Clock = uint64(value.Clock<<8) | uint64(x)
						}
					}
				}
			default:
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				handled = true
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}

	startPos = reader.Pos()
	err = nil

	if !handled_Producer && err == nil {
		err = enc.ErrSkipRequired{Name: "Producer", TypeNum: 161}
	}
	if !handled_Clock && err == nil {
		err = enc.ErrSkipRequired{Name: "Clock", TypeNum: 163}
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (value *IDType) Encode() enc.Wire {
	encoder := IDTypeEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *IDType) Bytes() []byte {
	return value.Encode().Join()
}

func ParseIDType(reader enc.ParseReader, ignoreCritical bool) (*IDType, error) {
	context := IDTypeParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type RecordEncoder struct {
	length uint

	ID_encoder          IDTypeEncoder
	Origin_encoder      IDTypeEncoder
	RightOrigin_encoder IDTypeEncoder
}

type RecordParsingContext struct {
	ID_context          IDTypeParsingContext
	Origin_context      IDTypeParsingContext
	RightOrigin_context IDTypeParsingContext
}

func (encoder *RecordEncoder) Init(value *Record) {

	if value.ID != nil {
		encoder.ID_encoder.Init(value.ID)
	}
	if value.Origin != nil {
		encoder.Origin_encoder.Init(value.Origin)
	}
	if value.RightOrigin != nil {
		encoder.RightOrigin_encoder.Init(value.RightOrigin)
	}

	l := uint(0)
	l += 1
	switch x := value.RecordType; {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	if value.ID != nil {
		l += 1
		switch x := encoder.ID_encoder.length; {
		case x <= 0xfc:
			l += 1
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
		l += encoder.ID_encoder.length
	}
	if value.Origin != nil {
		l += 1
		switch x := encoder.Origin_encoder.length; {
		case x <= 0xfc:
			l += 1
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
		l += encoder.Origin_encoder.length
	}
	if value.RightOrigin != nil {
		l += 1
		switch x := encoder.RightOrigin_encoder.length; {
		case x <= 0xfc:
			l += 1
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
		l += encoder.RightOrigin_encoder.length
	}
	l += 1
	switch x := len(value.Content); {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += uint(len(value.Content))
	encoder.length = l

}

func (context *RecordParsingContext) Init() {

	context.ID_context.Init()
	context.Origin_context.Init()
	context.RightOrigin_context.Init()

}

func (encoder *RecordEncoder) EncodeInto(value *Record, buf []byte) {

	pos := uint(0)

	buf[pos] = byte(165)
	pos += 1
	switch x := value.RecordType; {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	if value.ID != nil {
		buf[pos] = byte(167)
		pos += 1
		switch x := encoder.ID_encoder.length; {
		case x <= 0xfc:
			buf[pos] = byte(x)
			pos += 1
		case x <= 0xffff:
			buf[pos] = 0xfd
			binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
			pos += 3
		case x <= 0xffffffff:
			buf[pos] = 0xfe
			binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
			pos += 5
		default:
			buf[pos] = 0xff
			binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
			pos += 9
		}
		if encoder.ID_encoder.length > 0 {
			encoder.ID_encoder.EncodeInto(value.ID, buf[pos:])
			pos += encoder.ID_encoder.length
		}
	}
	if value.Origin != nil {
		buf[pos] = byte(169)
		pos += 1
		switch x := encoder.Origin_encoder.length; {
		case x <= 0xfc:
			buf[pos] = byte(x)
			pos += 1
		case x <= 0xffff:
			buf[pos] = 0xfd
			binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
			pos += 3
		case x <= 0xffffffff:
			buf[pos] = 0xfe
			binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
			pos += 5
		default:
			buf[pos] = 0xff
			binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
			pos += 9
		}
		if encoder.Origin_encoder.length > 0 {
			encoder.Origin_encoder.EncodeInto(value.Origin, buf[pos:])
			pos += encoder.Origin_encoder.length
		}
	}
	if value.RightOrigin != nil {
		buf[pos] = byte(170)
		pos += 1
		switch x := encoder.RightOrigin_encoder.length; {
		case x <= 0xfc:
			buf[pos] = byte(x)
			pos += 1
		case x <= 0xffff:
			buf[pos] = 0xfd
			binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
			pos += 3
		case x <= 0xffffffff:
			buf[pos] = 0xfe
			binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
			pos += 5
		default:
			buf[pos] = 0xff
			binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
			pos += 9
		}
		if encoder.RightOrigin_encoder.length > 0 {
			encoder.RightOrigin_encoder.EncodeInto(value.RightOrigin, buf[pos:])
			pos += encoder.RightOrigin_encoder.length
		}
	}
	buf[pos] = byte(172)
	pos += 1
	switch x := len(value.Content); {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	copy(buf[pos:], value.Content)
	pos += uint(len(value.Content))
}

func (encoder *RecordEncoder) Encode(value *Record) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *RecordParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*Record, error) {
	if reader == nil {
		return nil, enc.ErrBufferOverflow
	}

	var handled_RecordType bool = false
	var handled_ID bool = false
	var handled_Origin bool = false
	var handled_RightOrigin bool = false
	var handled_Content bool = false

	progress := -1
	_ = progress

	value := &Record{}
	var err error
	var startPos int
	for {
		startPos = reader.Pos()
		if startPos >= reader.Length() {
			break
		}
		typ := enc.TLNum(0)
		l := enc.TLNum(0)
		typ, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}
		l, err = enc.ReadTLNum(reader)
		if err != nil {
			return nil, enc.ErrFailToParse{TypeNum: 0, Err: err}
		}

		err = nil
		if handled := false; true {
			switch typ {
			case 165:
				if true {
					handled = true
					handled_RecordType = true
					value.RecordType = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x := byte(0)
							x, err = reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.RecordType = uint64(value.RecordType<<8) | uint64(x)
						}
					}
				}
			case 167:
				if true {
					handled = true
					handled_ID = true
					value.ID, err = context.ID_context.Parse(reader.Delegate(int(l)), ignoreCritical)
				}
			case 169:
				if true {
					handled = true
					handled_Origin = true
					value.Origin, err = context.Origin_context.Parse(reader.Delegate(int(l)), ignoreCritical)
				}
			case 170:
				if true {
					handled = true
					handled_RightOrigin = true
					value.RightOrigin, err = context.RightOrigin_context.Parse(reader.Delegate(int(l)), ignoreCritical)
				}
			case 172:
				if true {
					handled = true
					handled_Content = true
					{
						var builder strings.Builder
						_, err = io.CopyN(&builder, reader, int64(l))
						if err == nil {
							value.Content = builder.String()
						}
					}
				}
			default:
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				handled = true
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}

	startPos = reader.Pos()
	err = nil

	if !handled_RecordType && err == nil {
		err = enc.ErrSkipRequired{Name: "RecordType", TypeNum: 165}
	}
	if !handled_ID && err == nil {
		value.ID = nil
	}
	if !handled_Origin && err == nil {
		value.Origin = nil
	}
	if !handled_RightOrigin && err == nil {
		value.RightOrigin = nil
	}
	if !handled_Content && err == nil {
		err = enc.ErrSkipRequired{Name: "Content", TypeNum: 172}
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (value *Record) Encode() enc.Wire {
	encoder := RecordEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *Record) Bytes() []byte {
	return value.Encode().Join()
}

func ParseRecord(reader enc.ParseReader, ignoreCritical bool) (*Record, error) {
	context := RecordParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}
