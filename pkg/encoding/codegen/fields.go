package codegen

import (
	"errors"
	"strings"
)

var ErrInvalidField = errors.New("invalid TlvField. Please check the annotation (including type and arguments)")
var ErrWrongTypeNumber = errors.New("invalid type number")

type TlvField interface {
	Name() string
	TypeNum() uint64

	// codegen encoding length of the field
	//   - in(value): struct being encoded
	//   - in(encoder): encoder struct
	//   - out(l): length variable to update
	GenEncodingLength() (string, error)
	// codegen encoding length for nocopy
	GenEncodingWirePlan() (string, error)
	// codegen encoding the field
	//   - in(value): struct being encoded
	//   - out(buf): buffer to write to
	GenEncodeInto() (string, error)
	// codegen fields in encoder struct
	GenEncoderStruct() (string, error)
	// codegen encoder initialization
	//   - in(value): struct being encoded
	//   - in(encoder): encoder struct
	GenInitEncoder() (string, error)
	// codegen parsing context struct fields
	GenParsingContextStruct() (string, error)
	// codegen parsing context initialization
	GenInitContext() (string, error)
	// codegen reading from buffer
	//   - in(value): struct being decoded
	//   - in(context): parsing context (if with context)
	//   - in(reader): byte reader
	//   - in(ignoreCritical): ignore critical flag
	GenReadFrom() (string, error)
	// codegen skipping reading a field
	//   - in(value): struct being decoded
	//   - out(err): error variable
	GenSkipProcess() (string, error)
	// codegen converting to dict
	GenToDict() (string, error)
	// codegen converting from dict
	GenFromDict() (string, error)
}

// BaseTlvField is a base class for all TLV fields.
// Golang's inheritance is not supported, so we use this class to disable
// optional functions.
type BaseTlvField struct {
	name    string
	typeNum uint64
}

func (f *BaseTlvField) Name() string {
	return f.name
}

func (f *BaseTlvField) TypeNum() uint64 {
	return f.typeNum
}

func (*BaseTlvField) GenEncodingLength() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenEncodingWirePlan() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenEncodeInto() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenEncoderStruct() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenInitEncoder() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenParsingContextStruct() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenInitContext() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenReadFrom() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenSkipProcess() (string, error) {
	return "// base - skip", nil
}

func (*BaseTlvField) GenToDict() (string, error) {
	return "", nil
}

func (*BaseTlvField) GenFromDict() (string, error) {
	return "", nil
}

func CreateField(className string, name string, typeNum uint64, annotation string, model *TlvModel) (TlvField, error) {
	fieldList := map[string]func(string, uint64, string, *TlvModel) (TlvField, error){
		"natural":           NewNaturalField,
		"fixedUint":         NewFixedUintField,
		"time":              NewTimeField,
		"binary":            NewBinaryField,
		"string":            NewStringField,
		"wire":              NewWireField,
		"name":              NewNameField,
		"bool":              NewBoolField,
		"procedureArgument": NewProcedureArgument,
		"offsetMarker":      NewOffsetMarker,
		"rangeMarker":       NewRangeMarker,
		"sequence":          NewSequenceField,
		"struct":            NewStructField,
		"signature":         NewSignatureField,
		"interestName":      NewInterestNameField,
		"map":               NewMapField,
	}

	for k, f := range fieldList {
		if strings.HasPrefix(className, k) {
			return f(name, typeNum, annotation, model)
		}
	}
	return nil, ErrInvalidField
}
