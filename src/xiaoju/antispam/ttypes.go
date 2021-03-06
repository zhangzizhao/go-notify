// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package antispam

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"xiaoju"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = xiaoju.GoUnusedProtection__
var GoUnusedProtection__ int

type SerialNumber int64

func SerialNumberPtr(v SerialNumber) *SerialNumber { return &v }

type Request struct {
	PhoneNumber string            `thrift:"phoneNumber,1,required" json:"phoneNumber"`
	Imei        string            `thrift:"imei,2,required" json:"imei"`
	Ip          string            `thrift:"ip,3,required" json:"ip"`
	Info        map[string]string `thrift:"info,4" json:"info"`
}

func NewRequest() *Request {
	return &Request{}
}

func (p *Request) GetPhoneNumber() string {
	return p.PhoneNumber
}

func (p *Request) GetImei() string {
	return p.Imei
}

func (p *Request) GetIp() string {
	return p.Ip
}

var Request_Info_DEFAULT map[string]string

func (p *Request) GetInfo() map[string]string {
	return p.Info
}
func (p *Request) IsSetInfo() bool {
	return p.Info != nil
}

func (p *Request) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *Request) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.PhoneNumber = v
	}
	return nil
}

func (p *Request) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Imei = v
	}
	return nil
}

func (p *Request) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Ip = v
	}
	return nil
}

func (p *Request) ReadField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Info = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
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
		p.Info[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Request"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("phoneNumber", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:phoneNumber: ", p), err)
	}
	if err := oprot.WriteString(string(p.PhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.phoneNumber (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:phoneNumber: ", p), err)
	}
	return err
}

func (p *Request) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("imei", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:imei: ", p), err)
	}
	if err := oprot.WriteString(string(p.Imei)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.imei (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:imei: ", p), err)
	}
	return err
}

func (p *Request) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ip: ", p), err)
	}
	if err := oprot.WriteString(string(p.Ip)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ip: ", p), err)
	}
	return err
}

func (p *Request) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetInfo() {
		if err := oprot.WriteFieldBegin("info", thrift.MAP, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:info: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Info)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Info {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:info: ", p), err)
		}
	}
	return err
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}

type ServiceError struct {
	ErrorA1 string `thrift:"error,1,required" json:"error"`
}

func NewServiceError() *ServiceError {
	return &ServiceError{}
}

func (p *ServiceError) GetErrorA1() string {
	return p.ErrorA1
}
func (p *ServiceError) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *ServiceError) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ErrorA1 = v
	}
	return nil
}

func (p *ServiceError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ServiceError"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ServiceError) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("error", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err)
	}
	if err := oprot.WriteString(string(p.ErrorA1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.error (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err)
	}
	return err
}

func (p *ServiceError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceError(%+v)", *p)
}

func (p *ServiceError) Error() string {
	return p.String()
}
