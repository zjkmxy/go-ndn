// Generated by the generator, DO NOT modify manually
package gen_basic

import (
	"encoding/binary"
	"io"
	"time"

	enc "github.com/zjkmxy/go-ndn/pkg/encoding"
)

type FakeMetaInfoEncoder struct {
	length uint
}

type FakeMetaInfoParsingContext struct {
}

func (encoder *FakeMetaInfoEncoder) Init(value *FakeMetaInfo) {

	l := uint(0)
	l += 1
	switch x := value.Number; {
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
	switch x := uint64(value.Time / time.Millisecond); {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}

	if value.Binary != nil {
		l += 1
		switch x := len(value.Binary); {
		case x <= 0xfc:
			l += 1
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
		l += uint(len(value.Binary))
	}

	encoder.length = l

}

func (context *FakeMetaInfoParsingContext) Init() {

}

func (encoder *FakeMetaInfoEncoder) EncodeInto(value *FakeMetaInfo, buf []byte) {

	pos := uint(0)
	buf[pos] = byte(24)
	pos += 1
	switch x := value.Number; {
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

	buf[pos] = byte(25)
	pos += 1
	switch x := uint64(value.Time / time.Millisecond); {
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

	if value.Binary != nil {
		buf[pos] = byte(26)
		pos += 1
		switch x := len(value.Binary); {
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
		copy(buf[pos:], value.Binary)
		pos += uint(len(value.Binary))
	}

}

func (encoder *FakeMetaInfoEncoder) Encode(value *FakeMetaInfo) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *FakeMetaInfoParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*FakeMetaInfo, error) {
	progress := -1
	value := &FakeMetaInfo{}
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
		for handled := false; !handled; progress++ {
			switch typ {
			case 24:
				if progress+1 == 0 {
					handled = true
					value.Number = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x, err := reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.Number = uint64(value.Number<<8) | uint64(x)
						}
					}
				}
			case 25:
				if progress+1 == 1 {
					handled = true
					{
						timeInt := uint64(0)
						timeInt = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x, err := reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								timeInt = uint64(timeInt<<8) | uint64(x)
							}
						}
						value.Time = time.Duration(timeInt) * time.Millisecond
					}

				}
			case 26:
				if progress+1 == 2 {
					handled = true
					value.Binary = make([]byte, l)
					_, err = io.ReadFull(reader, value.Binary)

				}
			default:
				handled = true
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
				switch progress {
				case 0 - 1:
					err = enc.ErrSkipRequired{TypeNum: 24}
				case 1 - 1:
					err = enc.ErrSkipRequired{TypeNum: 25}
				case 2 - 1:
					value.Binary = nil
				}
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}
	startPos = reader.Pos()
	for ; progress < 3; progress++ {
		switch progress {
		case 0 - 1:
			err = enc.ErrSkipRequired{TypeNum: 24}
		case 1 - 1:
			err = enc.ErrSkipRequired{TypeNum: 25}
		case 2 - 1:
			value.Binary = nil
		}
	}
	return value, nil
}

func (value *FakeMetaInfo) Encode() enc.Wire {
	encoder := FakeMetaInfoEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *FakeMetaInfo) Bytes() []byte {
	return value.Encode().Join()
}

func ParseFakeMetaInfo(reader enc.ParseReader, ignoreCritical bool) (*FakeMetaInfo, error) {
	context := FakeMetaInfoParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type OptFieldEncoder struct {
	length uint
}

type OptFieldParsingContext struct {
}

func (encoder *OptFieldEncoder) Init(value *OptField) {

	l := uint(0)
	if value.Number != nil {
		l += 1
		switch x := *value.Number; {
		case x <= 0xff:
			l += 2
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
	}

	if value.Time != nil {
		l += 1
		switch x := uint64(*value.Time / time.Millisecond); {
		case x <= 0xff:
			l += 2
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
	}

	if value.Binary != nil {
		l += 1
		switch x := len(value.Binary); {
		case x <= 0xfc:
			l += 1
		case x <= 0xffff:
			l += 3
		case x <= 0xffffffff:
			l += 5
		default:
			l += 9
		}
		l += uint(len(value.Binary))
	}

	if value.Bool {
		l += 1
		l += 1
	}

	encoder.length = l

}

func (context *OptFieldParsingContext) Init() {

}

func (encoder *OptFieldEncoder) EncodeInto(value *OptField, buf []byte) {

	pos := uint(0)
	if value.Number != nil {
		buf[pos] = byte(24)
		pos += 1
		switch x := *value.Number; {
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

	if value.Time != nil {
		buf[pos] = byte(25)
		pos += 1
		switch x := uint64(*value.Time / time.Millisecond); {
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

	if value.Binary != nil {
		buf[pos] = byte(26)
		pos += 1
		switch x := len(value.Binary); {
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
		copy(buf[pos:], value.Binary)
		pos += uint(len(value.Binary))
	}

	if value.Bool {
		buf[pos] = byte(48)
		pos += 1
		buf[pos] = byte(0)
		pos += 1
	}

}

func (encoder *OptFieldEncoder) Encode(value *OptField) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *OptFieldParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*OptField, error) {
	progress := -1
	value := &OptField{}
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
		for handled := false; !handled; progress++ {
			switch typ {
			case 24:
				if progress+1 == 0 {
					handled = true
					{
						tempVal := uint64(0)
						tempVal = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x, err := reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								tempVal = uint64(tempVal<<8) | uint64(x)
							}
						}
						value.Number = &tempVal
					}

				}
			case 25:
				if progress+1 == 1 {
					handled = true
					{
						timeInt := uint64(0)
						timeInt = uint64(0)
						{
							for i := 0; i < int(l); i++ {
								x, err := reader.ReadByte()
								if err != nil {
									if err == io.EOF {
										err = io.ErrUnexpectedEOF
									}
									break
								}
								timeInt = uint64(timeInt<<8) | uint64(x)
							}
						}
						tempVal := time.Duration(timeInt) * time.Millisecond
						value.Time = &tempVal
					}

				}
			case 26:
				if progress+1 == 2 {
					handled = true
					value.Binary = make([]byte, l)
					_, err = io.ReadFull(reader, value.Binary)

				}
			case 48:
				if progress+1 == 3 {
					handled = true
					value.Bool = true
				}
			default:
				handled = true
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
				switch progress {
				case 0 - 1:
					value.Number = nil
				case 1 - 1:
					value.Time = nil
				case 2 - 1:
					value.Binary = nil
				case 3 - 1:
					value.Bool = false
				}
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}
	startPos = reader.Pos()
	for ; progress < 4; progress++ {
		switch progress {
		case 0 - 1:
			value.Number = nil
		case 1 - 1:
			value.Time = nil
		case 2 - 1:
			value.Binary = nil
		case 3 - 1:
			value.Bool = false
		}
	}
	return value, nil
}

func (value *OptField) Encode() enc.Wire {
	encoder := OptFieldEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *OptField) Bytes() []byte {
	return value.Encode().Join()
}

func ParseOptField(reader enc.ParseReader, ignoreCritical bool) (*OptField, error) {
	context := OptFieldParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type WireNameFieldEncoder struct {
	length uint

	Wire_length uint
	Name_length uint
}

type WireNameFieldParsingContext struct {
}

func (encoder *WireNameFieldEncoder) Init(value *WireNameField) {
	encoder.Wire_length = 0
	for _, c := range value.Wire {
		encoder.Wire_length += uint(len(c))
	}

	encoder.Name_length = 0
	for _, c := range value.Name {
		encoder.Name_length += uint(c.EncodingLength())
	}

	l := uint(0)
	l += 1
	switch x := encoder.Wire_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Wire_length

	l += 1
	switch x := encoder.Name_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Name_length

	encoder.length = l

}

func (context *WireNameFieldParsingContext) Init() {

}

func (encoder *WireNameFieldEncoder) EncodeInto(value *WireNameField, buf []byte) {

	pos := uint(0)
	buf[pos] = byte(1)
	pos += 1
	switch x := encoder.Wire_length; {
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
	for _, w := range value.Wire {
		copy(buf[pos:], w)
		pos += uint(len(w))
	}

	buf[pos] = byte(2)
	pos += 1
	switch x := encoder.Name_length; {
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
	for _, c := range value.Name {
		pos += uint(c.EncodeInto(buf[pos:]))
	}

}

func (encoder *WireNameFieldEncoder) Encode(value *WireNameField) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *WireNameFieldParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*WireNameField, error) {
	progress := -1
	value := &WireNameField{}
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
		for handled := false; !handled; progress++ {
			switch typ {
			case 1:
				if progress+1 == 0 {
					handled = true
					value.Wire, err = reader.ReadWire(int(l))

				}
			case 2:
				if progress+1 == 1 {
					handled = true
					value.Name = make(enc.Name, 0)
					startName := reader.Pos()
					endName := startName + int(l)
					for reader.Pos() < endName {
						c, err := enc.ReadComponent(reader)
						if err != nil {
							break
						}
						value.Name = append(value.Name, *c)
					}
					if err != nil && reader.Pos() != endName {
						err = enc.ErrBufferOverflow
					}

				}
			default:
				handled = true
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
				switch progress {
				case 0 - 1:
					value.Wire = nil
				case 1 - 1:
					value.Name = nil
				}
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}
	startPos = reader.Pos()
	for ; progress < 2; progress++ {
		switch progress {
		case 0 - 1:
			value.Wire = nil
		case 1 - 1:
			value.Name = nil
		}
	}
	return value, nil
}

func (value *WireNameField) Encode() enc.Wire {
	encoder := WireNameFieldEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *WireNameField) Bytes() []byte {
	return value.Encode().Join()
}

func ParseWireNameField(reader enc.ParseReader, ignoreCritical bool) (*WireNameField, error) {
	context := WireNameFieldParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}

type MarkersEncoder struct {
	length uint

	startMarker int
	Wire_length uint
	argument    int
	Name_length uint
	endMarker   int
}

type MarkersParsingContext struct {
	startMarker int

	argument int

	endMarker int
}

func (encoder *MarkersEncoder) Init(value *Markers) {

	encoder.Wire_length = 0
	for _, c := range value.Wire {
		encoder.Wire_length += uint(len(c))
	}

	encoder.Name_length = 0
	for _, c := range value.Name {
		encoder.Name_length += uint(c.EncodingLength())
	}

	l := uint(0)

	l += 1
	switch x := encoder.Wire_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Wire_length

	l += 1
	switch x := encoder.Name_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Name_length

	encoder.length = l

}

func (context *MarkersParsingContext) Init() {

}

func (encoder *MarkersEncoder) EncodeInto(value *Markers, buf []byte) {

	pos := uint(0)
	encoder.startMarker = int(pos)
	buf[pos] = byte(1)
	pos += 1
	switch x := encoder.Wire_length; {
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
	for _, w := range value.Wire {
		copy(buf[pos:], w)
		pos += uint(len(w))
	}

	buf[pos] = byte(2)
	pos += 1
	switch x := encoder.Name_length; {
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
	for _, c := range value.Name {
		pos += uint(c.EncodeInto(buf[pos:]))
	}

	encoder.endMarker = int(pos)
}

func (encoder *MarkersEncoder) Encode(value *Markers) enc.Wire {

	wire := make(enc.Wire, 1)
	wire[0] = make([]byte, encoder.length)
	buf := wire[0]
	encoder.EncodeInto(value, buf)

	return wire
}

func (context *MarkersParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*Markers, error) {
	progress := -1
	value := &Markers{}
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
		for handled := false; !handled; progress++ {
			switch typ {
			case 1:
				if progress+1 == 1 {
					handled = true
					value.Wire, err = reader.ReadWire(int(l))

				}
			case 2:
				if progress+1 == 3 {
					handled = true
					value.Name = make(enc.Name, 0)
					startName := reader.Pos()
					endName := startName + int(l)
					for reader.Pos() < endName {
						c, err := enc.ReadComponent(reader)
						if err != nil {
							break
						}
						value.Name = append(value.Name, *c)
					}
					if err != nil && reader.Pos() != endName {
						err = enc.ErrBufferOverflow
					}

				}
			default:
				handled = true
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
				switch progress {
				case 0 - 1:
					context.startMarker = int(startPos)
				case 1 - 1:
					value.Wire = nil
				case 2 - 1:

				case 3 - 1:
					value.Name = nil
				case 4 - 1:
					context.endMarker = int(startPos)
				}
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}
	startPos = reader.Pos()
	for ; progress < 5; progress++ {
		switch progress {
		case 0 - 1:
			context.startMarker = int(startPos)
		case 1 - 1:
			value.Wire = nil
		case 2 - 1:

		case 3 - 1:
			value.Name = nil
		case 4 - 1:
			context.endMarker = int(startPos)
		}
	}
	return value, nil
}

type NoCopyStructEncoder struct {
	length uint

	wirePlan []uint

	Wire1_length uint

	Wire2_length uint
}

type NoCopyStructParsingContext struct {
}

func (encoder *NoCopyStructEncoder) Init(value *NoCopyStruct) {
	encoder.Wire1_length = 0
	for _, c := range value.Wire1 {
		encoder.Wire1_length += uint(len(c))
	}

	encoder.Wire2_length = 0
	for _, c := range value.Wire2 {
		encoder.Wire2_length += uint(len(c))
	}

	l := uint(0)
	l += 1
	switch x := encoder.Wire1_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Wire1_length

	l += 1
	switch x := value.Number; {
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
	switch x := encoder.Wire2_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += encoder.Wire2_length

	encoder.length = l

	wirePlan := make([]uint, 0)
	l = uint(0)
	l += 1
	switch x := encoder.Wire1_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	wirePlan = append(wirePlan, l)
	l = 0
	for range value.Wire1 {
		wirePlan = append(wirePlan, l)
		l = 0
	}

	l += 1
	switch x := value.Number; {
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
	switch x := encoder.Wire2_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	wirePlan = append(wirePlan, l)
	l = 0
	for range value.Wire2 {
		wirePlan = append(wirePlan, l)
		l = 0
	}

	if l > 0 {
		wirePlan = append(wirePlan, l)
	}
	encoder.wirePlan = wirePlan
}

func (context *NoCopyStructParsingContext) Init() {

}

func (encoder *NoCopyStructEncoder) EncodeInto(value *NoCopyStruct, wire enc.Wire) {

	wireIdx := 0
	buf := wire[wireIdx]

	pos := uint(0)
	buf[pos] = byte(1)
	pos += 1
	switch x := encoder.Wire1_length; {
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
	wireIdx++
	pos = 0
	if wireIdx < len(wire) {
		buf = wire[wireIdx]
	} else {
		buf = nil
	}
	for _, w := range value.Wire1 {
		wire[wireIdx] = w
		wireIdx++
		pos = 0
		if wireIdx < len(wire) {
			buf = wire[wireIdx]
		} else {
			buf = nil
		}
	}

	buf[pos] = byte(2)
	pos += 1
	switch x := value.Number; {
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

	buf[pos] = byte(3)
	pos += 1
	switch x := encoder.Wire2_length; {
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
	wireIdx++
	pos = 0
	if wireIdx < len(wire) {
		buf = wire[wireIdx]
	} else {
		buf = nil
	}
	for _, w := range value.Wire2 {
		wire[wireIdx] = w
		wireIdx++
		pos = 0
		if wireIdx < len(wire) {
			buf = wire[wireIdx]
		} else {
			buf = nil
		}
	}

}

func (encoder *NoCopyStructEncoder) Encode(value *NoCopyStruct) enc.Wire {

	wire := make(enc.Wire, len(encoder.wirePlan))
	for i, l := range encoder.wirePlan {
		if l > 0 {
			wire[i] = make([]byte, l)
		}
	}
	encoder.EncodeInto(value, wire)

	return wire
}

func (context *NoCopyStructParsingContext) Parse(reader enc.ParseReader, ignoreCritical bool) (*NoCopyStruct, error) {
	progress := -1
	value := &NoCopyStruct{}
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
		for handled := false; !handled; progress++ {
			switch typ {
			case 1:
				if progress+1 == 0 {
					handled = true
					value.Wire1, err = reader.ReadWire(int(l))

				}
			case 2:
				if progress+1 == 1 {
					handled = true
					value.Number = uint64(0)
					{
						for i := 0; i < int(l); i++ {
							x, err := reader.ReadByte()
							if err != nil {
								if err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								break
							}
							value.Number = uint64(value.Number<<8) | uint64(x)
						}
					}
				}
			case 3:
				if progress+1 == 2 {
					handled = true
					value.Wire2, err = reader.ReadWire(int(l))

				}
			default:
				handled = true
				if !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
					return nil, enc.ErrUnrecognizedField{TypeNum: typ}
				}
				err = reader.Skip(int(l))
			}
			if err == nil && !handled {
				switch progress {
				case 0 - 1:
					value.Wire1 = nil
				case 1 - 1:
					err = enc.ErrSkipRequired{TypeNum: 2}
				case 2 - 1:
					value.Wire2 = nil
				}
			}
			if err != nil {
				return nil, enc.ErrFailToParse{TypeNum: typ, Err: err}
			}
		}
	}
	startPos = reader.Pos()
	for ; progress < 3; progress++ {
		switch progress {
		case 0 - 1:
			value.Wire1 = nil
		case 1 - 1:
			err = enc.ErrSkipRequired{TypeNum: 2}
		case 2 - 1:
			value.Wire2 = nil
		}
	}
	return value, nil
}

func (value *NoCopyStruct) Encode() enc.Wire {
	encoder := NoCopyStructEncoder{}
	encoder.Init(value)
	return encoder.Encode(value)
}

func (value *NoCopyStruct) Bytes() []byte {
	return value.Encode().Join()
}

func ParseNoCopyStruct(reader enc.ParseReader, ignoreCritical bool) (*NoCopyStruct, error) {
	context := NoCopyStructParsingContext{}
	context.Init()
	return context.Parse(reader, ignoreCritical)
}
