// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thriftliontesting

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - One
//  - Two
//  - StringField
//  - Int32Field
//  - Bar
type Foo struct {
  One *string `thrift:"one,1" db:"one" json:"one,omitempty"`
  Two *int32 `thrift:"two,2" db:"two" json:"two,omitempty"`
  StringField *string `thrift:"string_field,3" db:"string_field" json:"string_field,omitempty"`
  Int32Field *int32 `thrift:"int32_field,4" db:"int32_field" json:"int32_field,omitempty"`
  Bar *Bar `thrift:"bar,5" db:"bar" json:"bar,omitempty"`
}

func NewFoo() *Foo {
  return &Foo{}
}

var Foo_One_DEFAULT string
func (p *Foo) GetOne() string {
  if !p.IsSetOne() {
    return Foo_One_DEFAULT
  }
return *p.One
}
var Foo_Two_DEFAULT int32
func (p *Foo) GetTwo() int32 {
  if !p.IsSetTwo() {
    return Foo_Two_DEFAULT
  }
return *p.Two
}
var Foo_StringField_DEFAULT string
func (p *Foo) GetStringField() string {
  if !p.IsSetStringField() {
    return Foo_StringField_DEFAULT
  }
return *p.StringField
}
var Foo_Int32Field_DEFAULT int32
func (p *Foo) GetInt32Field() int32 {
  if !p.IsSetInt32Field() {
    return Foo_Int32Field_DEFAULT
  }
return *p.Int32Field
}
var Foo_Bar_DEFAULT *Bar
func (p *Foo) GetBar() *Bar {
  if !p.IsSetBar() {
    return Foo_Bar_DEFAULT
  }
return p.Bar
}
func (p *Foo) IsSetOne() bool {
  return p.One != nil
}

func (p *Foo) IsSetTwo() bool {
  return p.Two != nil
}

func (p *Foo) IsSetStringField() bool {
  return p.StringField != nil
}

func (p *Foo) IsSetInt32Field() bool {
  return p.Int32Field != nil
}

func (p *Foo) IsSetBar() bool {
  return p.Bar != nil
}

func (p *Foo) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Foo)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.One = &v
}
  return nil
}

func (p *Foo)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Two = &v
}
  return nil
}

func (p *Foo)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.StringField = &v
}
  return nil
}

func (p *Foo)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Int32Field = &v
}
  return nil
}

func (p *Foo)  ReadField5(iprot thrift.TProtocol) error {
  p.Bar = &Bar{}
  if err := p.Bar.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Bar), err)
  }
  return nil
}

func (p *Foo) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Foo"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Foo) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetOne() {
    if err := oprot.WriteFieldBegin("one", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:one: ", p), err) }
    if err := oprot.WriteString(string(*p.One)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.one (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:one: ", p), err) }
  }
  return err
}

func (p *Foo) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetTwo() {
    if err := oprot.WriteFieldBegin("two", thrift.I32, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:two: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Two)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.two (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:two: ", p), err) }
  }
  return err
}

func (p *Foo) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetStringField() {
    if err := oprot.WriteFieldBegin("string_field", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:string_field: ", p), err) }
    if err := oprot.WriteString(string(*p.StringField)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.string_field (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:string_field: ", p), err) }
  }
  return err
}

func (p *Foo) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetInt32Field() {
    if err := oprot.WriteFieldBegin("int32_field", thrift.I32, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:int32_field: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Int32Field)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.int32_field (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:int32_field: ", p), err) }
  }
  return err
}

func (p *Foo) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetBar() {
    if err := oprot.WriteFieldBegin("bar", thrift.STRUCT, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:bar: ", p), err) }
    if err := p.Bar.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Bar), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:bar: ", p), err) }
  }
  return err
}

func (p *Foo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Foo(%+v)", *p)
}

// Attributes:
//  - One
//  - Two
//  - StringField
//  - Int32Field
type Bar struct {
  One *string `thrift:"one,1" db:"one" json:"one,omitempty"`
  Two *string `thrift:"two,2" db:"two" json:"two,omitempty"`
  StringField *string `thrift:"string_field,3" db:"string_field" json:"string_field,omitempty"`
  Int32Field *int32 `thrift:"int32_field,4" db:"int32_field" json:"int32_field,omitempty"`
}

func NewBar() *Bar {
  return &Bar{}
}

var Bar_One_DEFAULT string
func (p *Bar) GetOne() string {
  if !p.IsSetOne() {
    return Bar_One_DEFAULT
  }
return *p.One
}
var Bar_Two_DEFAULT string
func (p *Bar) GetTwo() string {
  if !p.IsSetTwo() {
    return Bar_Two_DEFAULT
  }
return *p.Two
}
var Bar_StringField_DEFAULT string
func (p *Bar) GetStringField() string {
  if !p.IsSetStringField() {
    return Bar_StringField_DEFAULT
  }
return *p.StringField
}
var Bar_Int32Field_DEFAULT int32
func (p *Bar) GetInt32Field() int32 {
  if !p.IsSetInt32Field() {
    return Bar_Int32Field_DEFAULT
  }
return *p.Int32Field
}
func (p *Bar) IsSetOne() bool {
  return p.One != nil
}

func (p *Bar) IsSetTwo() bool {
  return p.Two != nil
}

func (p *Bar) IsSetStringField() bool {
  return p.StringField != nil
}

func (p *Bar) IsSetInt32Field() bool {
  return p.Int32Field != nil
}

func (p *Bar) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Bar)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.One = &v
}
  return nil
}

func (p *Bar)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Two = &v
}
  return nil
}

func (p *Bar)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.StringField = &v
}
  return nil
}

func (p *Bar)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Int32Field = &v
}
  return nil
}

func (p *Bar) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Bar"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Bar) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetOne() {
    if err := oprot.WriteFieldBegin("one", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:one: ", p), err) }
    if err := oprot.WriteString(string(*p.One)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.one (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:one: ", p), err) }
  }
  return err
}

func (p *Bar) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetTwo() {
    if err := oprot.WriteFieldBegin("two", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:two: ", p), err) }
    if err := oprot.WriteString(string(*p.Two)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.two (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:two: ", p), err) }
  }
  return err
}

func (p *Bar) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetStringField() {
    if err := oprot.WriteFieldBegin("string_field", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:string_field: ", p), err) }
    if err := oprot.WriteString(string(*p.StringField)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.string_field (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:string_field: ", p), err) }
  }
  return err
}

func (p *Bar) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetInt32Field() {
    if err := oprot.WriteFieldBegin("int32_field", thrift.I32, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:int32_field: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Int32Field)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.int32_field (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:int32_field: ", p), err) }
  }
  return err
}

func (p *Bar) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Bar(%+v)", *p)
}

// Attributes:
//  - StringField
//  - Int32Field
type Ban struct {
  StringField *string `thrift:"string_field,1" db:"string_field" json:"string_field,omitempty"`
  Int32Field *int32 `thrift:"int32_field,2" db:"int32_field" json:"int32_field,omitempty"`
}

func NewBan() *Ban {
  return &Ban{}
}

var Ban_StringField_DEFAULT string
func (p *Ban) GetStringField() string {
  if !p.IsSetStringField() {
    return Ban_StringField_DEFAULT
  }
return *p.StringField
}
var Ban_Int32Field_DEFAULT int32
func (p *Ban) GetInt32Field() int32 {
  if !p.IsSetInt32Field() {
    return Ban_Int32Field_DEFAULT
  }
return *p.Int32Field
}
func (p *Ban) IsSetStringField() bool {
  return p.StringField != nil
}

func (p *Ban) IsSetInt32Field() bool {
  return p.Int32Field != nil
}

func (p *Ban) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Ban)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.StringField = &v
}
  return nil
}

func (p *Ban)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Int32Field = &v
}
  return nil
}

func (p *Ban) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Ban"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Ban) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetStringField() {
    if err := oprot.WriteFieldBegin("string_field", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:string_field: ", p), err) }
    if err := oprot.WriteString(string(*p.StringField)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.string_field (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:string_field: ", p), err) }
  }
  return err
}

func (p *Ban) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetInt32Field() {
    if err := oprot.WriteFieldBegin("int32_field", thrift.I32, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:int32_field: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Int32Field)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.int32_field (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:int32_field: ", p), err) }
  }
  return err
}

func (p *Ban) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Ban(%+v)", *p)
}

// Attributes:
//  - Ban
type Bat struct {
  Ban *Ban `thrift:"ban,1" db:"ban" json:"ban,omitempty"`
}

func NewBat() *Bat {
  return &Bat{}
}

var Bat_Ban_DEFAULT *Ban
func (p *Bat) GetBan() *Ban {
  if !p.IsSetBan() {
    return Bat_Ban_DEFAULT
  }
return p.Ban
}
func (p *Bat) IsSetBan() bool {
  return p.Ban != nil
}

func (p *Bat) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Bat)  ReadField1(iprot thrift.TProtocol) error {
  p.Ban = &Ban{}
  if err := p.Ban.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ban), err)
  }
  return nil
}

func (p *Bat) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Bat"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Bat) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetBan() {
    if err := oprot.WriteFieldBegin("ban", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ban: ", p), err) }
    if err := p.Ban.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ban), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ban: ", p), err) }
  }
  return err
}

func (p *Bat) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Bat(%+v)", *p)
}

// Attributes:
//  - Bat
type Baz struct {
  Bat *Bat `thrift:"bat,1" db:"bat" json:"bat,omitempty"`
}

func NewBaz() *Baz {
  return &Baz{}
}

var Baz_Bat_DEFAULT *Bat
func (p *Baz) GetBat() *Bat {
  if !p.IsSetBat() {
    return Baz_Bat_DEFAULT
  }
return p.Bat
}
func (p *Baz) IsSetBat() bool {
  return p.Bat != nil
}

func (p *Baz) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Baz)  ReadField1(iprot thrift.TProtocol) error {
  p.Bat = &Bat{}
  if err := p.Bat.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Bat), err)
  }
  return nil
}

func (p *Baz) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Baz"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Baz) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetBat() {
    if err := oprot.WriteFieldBegin("bat", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:bat: ", p), err) }
    if err := p.Bat.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Bat), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:bat: ", p), err) }
  }
  return err
}

func (p *Baz) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Baz(%+v)", *p)
}

type Empty struct {
}

func NewEmpty() *Empty {
  return &Empty{}
}

func (p *Empty) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Empty) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Empty"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Empty) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Empty(%+v)", *p)
}

// Attributes:
//  - One
type NoStdJson struct {
  One map[int64]string `thrift:"one,1" db:"one" json:"one,omitempty"`
}

func NewNoStdJson() *NoStdJson {
  return &NoStdJson{}
}

var NoStdJson_One_DEFAULT map[int64]string

func (p *NoStdJson) GetOne() map[int64]string {
  return p.One
}
func (p *NoStdJson) IsSetOne() bool {
  return p.One != nil
}

func (p *NoStdJson) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *NoStdJson)  ReadField1(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[int64]string, size)
  p.One =  tMap
  for i := 0; i < size; i ++ {
var _key0 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.One[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *NoStdJson) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("NoStdJson"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NoStdJson) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetOne() {
    if err := oprot.WriteFieldBegin("one", thrift.MAP, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:one: ", p), err) }
    if err := oprot.WriteMapBegin(thrift.I64, thrift.STRING, len(p.One)); err != nil {
      return thrift.PrependError("error writing map begin: ", err)
    }
    for k, v := range p.One {
      if err := oprot.WriteI64(int64(k)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteMapEnd(); err != nil {
      return thrift.PrependError("error writing map end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:one: ", p), err) }
  }
  return err
}

func (p *NoStdJson) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NoStdJson(%+v)", *p)
}
