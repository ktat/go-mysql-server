// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dolthub/go-mysql-server/sql/types"
	"gopkg.in/src-d/go-errors.v1"

	"github.com/dolthub/go-mysql-server/sql"
)

// ErrConvertExpression is returned when a conversion is not possible.
var ErrConvertExpression = errors.NewKind("expression '%v': couldn't convert to %v")

const (
	// ConvertToBinary is a conversion to binary.
	ConvertToBinary = "binary"
	// ConvertToChar is a conversion to char.
	ConvertToChar = "char"
	// ConvertToNChar is a conversion to nchar.
	ConvertToNChar = "nchar"
	// ConvertToDate is a conversion to date.
	ConvertToDate = "date"
	// ConvertToDatetime is a conversion to datetime.
	ConvertToDatetime = "datetime"
	// ConvertToDecimal is a conversion to decimal.
	ConvertToDecimal = "decimal"
	// ConvertToDouble is a conversion to double.
	ConvertToDouble = "double"
	// ConvertToJSON is a conversion to json.
	ConvertToJSON = "json"
	// ConvertToReal is a conversion to double.
	ConvertToReal = "real"
	// ConvertToSigned is a conversion to signed.
	ConvertToSigned = "signed"
	// ConvertToTime is a conversion to time.
	ConvertToTime = "time"
	// ConvertToUnsigned is a conversion to unsigned.
	ConvertToUnsigned = "unsigned"
)

// Convert represent a CAST(x AS T) or CONVERT(x, T) operation that casts x expression to type T.
type Convert struct {
	UnaryExpression
	// Type to cast
	castToType string
}

// NewConvert creates a new Convert expression.
func NewConvert(expr sql.Expression, castToType string) *Convert {
	return &Convert{
		UnaryExpression: UnaryExpression{Child: expr},
		castToType:      strings.ToLower(castToType),
	}
}

// IsNullable implements the Expression interface.
func (c *Convert) IsNullable() bool {
	switch c.castToType {
	case ConvertToDate, ConvertToDatetime:
		return true
	default:
		return c.Child.IsNullable()
	}
}

// Type implements the Expression interface.
func (c *Convert) Type() sql.Type {
	switch c.castToType {
	case ConvertToBinary:
		return sql.LongBlob
	case ConvertToChar, ConvertToNChar:
		return sql.LongText
	case ConvertToDate:
		return types.Date
	case ConvertToDatetime:
		return types.Datetime
	case ConvertToDecimal:
		//TODO: these values are completely arbitrary, we need to get the given precision/scale and store it
		return sql.MustCreateDecimalType(65, 10)
	case ConvertToDouble, ConvertToReal:
		return sql.Float64
	case ConvertToJSON:
		return sql.JSON
	case ConvertToSigned:
		return sql.Int64
	case ConvertToTime:
		return sql.Time
	case ConvertToUnsigned:
		return sql.Uint64
	default:
		return sql.Null
	}
}

// String implements the Stringer interface.
func (c *Convert) String() string {
	return fmt.Sprintf("convert(%v, %v)", c.Child, c.castToType)
}

// DebugString implements the Expression interface.
func (c *Convert) DebugString() string {
	pr := sql.NewTreePrinter()
	_ = pr.WriteNode("convert")
	children := []string{
		fmt.Sprintf("type: %v", c.castToType),
		fmt.Sprintf(sql.DebugString(c.Child)),
	}
	_ = pr.WriteChildren(children...)
	return pr.String()
}

// WithChildren implements the Expression interface.
func (c *Convert) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	if len(children) != 1 {
		return nil, sql.ErrInvalidChildrenNumber.New(c, len(children), 1)
	}
	return NewConvert(children[0], c.castToType), nil
}

// Eval implements the Expression interface.
func (c *Convert) Eval(ctx *sql.Context, row sql.Row) (interface{}, error) {
	val, err := c.Child.Eval(ctx, row)
	if err != nil {
		return nil, err
	}

	if val == nil {
		return nil, nil
	}

	// Should always return nil, and a warning instead
	casted, err := convertValue(val, c.castToType, c.Child.Type())
	if err != nil {
		if c.castToType == ConvertToJSON {
			return nil, ErrConvertExpression.Wrap(err, c.String(), c.castToType)
		}
		ctx.Warn(1292, "Incorrect %s value: %v", c.castToType, val)
		return nil, nil
	}

	return casted, nil
}

// convertValue only returns an error if converting to JSON, Date, and Datetime;
// the zero value is returned for float types.
// Nil is returned in all other cases.
func convertValue(val interface{}, castTo string, originType sql.Type) (interface{}, error) {
	switch strings.ToLower(castTo) {
	case ConvertToBinary:
		b, err := sql.LongBlob.Convert(val)
		if err != nil {
			return nil, nil
		}
		if types.IsTextOnly(originType) {
			// For string types we need to re-encode the string as we want the binary representation of the character set
			encoder := originType.(sql.StringType).Collation().CharacterSet().Encoder()
			encodedBytes, ok := encoder.Encode(b.([]byte))
			if !ok {
				return nil, fmt.Errorf("unable to re-encode string to convert to binary")
			}
			b = encodedBytes
		}
		return b, nil
	case ConvertToChar, ConvertToNChar:
		s, err := sql.LongText.Convert(val)
		if err != nil {
			return nil, nil
		}
		return s, nil
	case ConvertToDate:
		_, isTime := val.(time.Time)
		_, isString := val.(string)
		_, isBinary := val.([]byte)
		if !(isTime || isString || isBinary) {
			return nil, nil
		}
		d, err := types.Date.Convert(val)
		if err != nil {
			return nil, err
		}
		return d, nil
	case ConvertToDatetime:
		_, isTime := val.(time.Time)
		_, isString := val.(string)
		_, isBinary := val.([]byte)
		if !(isTime || isString || isBinary) {
			return nil, nil
		}
		d, err := types.Datetime.Convert(val)
		if err != nil {
			return nil, err
		}
		return d, nil
	case ConvertToDecimal:
		value, err := convertHexBlobToDecimalForNumericContext(val, originType)
		if err != nil {
			return nil, err
		}
		d, err := sql.InternalDecimalType.Convert(value)
		if err != nil {
			return "0", nil
		}
		return d, nil
	case ConvertToDouble, ConvertToReal:
		value, err := convertHexBlobToDecimalForNumericContext(val, originType)
		if err != nil {
			return nil, err
		}
		d, err := sql.Float64.Convert(value)
		if err != nil {
			return sql.Float64.Zero(), nil
		}
		return d, nil
	case ConvertToJSON:
		js, err := sql.JSON.Convert(val)
		if err != nil {
			return nil, err
		}
		return js, nil
	case ConvertToSigned:
		value, err := convertHexBlobToDecimalForNumericContext(val, originType)
		if err != nil {
			return nil, err
		}
		num, err := sql.Int64.Convert(value)
		if err != nil {
			return sql.Int64.Zero(), nil
		}

		return num, nil
	case ConvertToTime:
		t, err := sql.Time.Convert(val)
		if err != nil {
			return nil, nil
		}
		return t, nil
	case ConvertToUnsigned:
		value, err := convertHexBlobToDecimalForNumericContext(val, originType)
		if err != nil {
			return nil, err
		}
		num, err := sql.Uint64.Convert(value)
		if err != nil {
			num, err = sql.Int64.Convert(value)
			if err != nil {
				return sql.Uint64.Zero(), nil
			}
			return uint64(num.(int64)), nil
		}
		return num, nil
	default:
		return nil, nil
	}
}

// convertHexBlobToDecimalForNumericContext converts byte array value to unsigned int value if originType is BLOB type.
// This function is called when convertTo type is number type only. The hex literal values are parsed into blobs as
// binary string as default, but for numeric context, the value should be a number.
// Byte arrays of other SQL types are not handled here.
func convertHexBlobToDecimalForNumericContext(val interface{}, originType sql.Type) (interface{}, error) {
	if bin, isBinary := val.([]byte); isBinary && types.IsBlobType(originType) {
		stringVal := hex.EncodeToString(bin)
		decimalNum, err := strconv.ParseUint(stringVal, 16, 64)
		if err != nil {
			return nil, errors.NewKind("failed to convert hex blob value to unsigned int").New()
		}
		val = decimalNum
	}
	return val, nil
}
