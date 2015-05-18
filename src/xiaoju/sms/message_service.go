// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sms

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"xiaoju"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var _ = xiaoju.GoUnusedProtection__

type MessageService interface {
	xiaoju.BaseService

	// Parameters:
	//  - Message
	SendMessage(message *Message) (r SerialNumber, se *ServiceError, ie *InputError, err error)
	// Parameters:
	//  - BusinessId
	//  - JobId
	//  - RuleId
	//  - StartTime
	//  - EndTime
	CreateBreaker(businessId int32, jobId int32, ruleId int32, startTime int32, endTime int32) (r int64, se *ServiceError, err error)
	// Parameters:
	//  - Breaker
	DestoryBreaker(breaker int64) (se *ServiceError, err error)
}

type MessageServiceClient struct {
	*xiaoju.BaseServiceClient
}

func NewMessageServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MessageServiceClient {
	return &MessageServiceClient{BaseServiceClient: xiaoju.NewBaseServiceClientFactory(t, f)}
}

func NewMessageServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MessageServiceClient {
	return &MessageServiceClient{BaseServiceClient: xiaoju.NewBaseServiceClientProtocol(t, iprot, oprot)}
}

// Parameters:
//  - Message
func (p *MessageServiceClient) SendMessage(message *Message) (r SerialNumber, se *ServiceError, ie *InputError, err error) {
	if err = p.sendSendMessage(message); err != nil {
		return
	}
	return p.recvSendMessage()
}

func (p *MessageServiceClient) sendSendMessage(message *Message) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("sendMessage", thrift.CALL, p.SeqId)
	args1 := NewSendMessageArgs()
	args1.Message = message
	err = args1.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *MessageServiceClient) recvSendMessage() (value SerialNumber, se *ServiceError, ie *InputError, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error3 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error4 error
		error4, err = error3.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error4
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result2 := NewSendMessageResult()
	err = result2.Read(iprot)
	iprot.ReadMessageEnd()
	value = result2.Success
	if result2.Se != nil {
		se = result2.Se
	}
	if result2.Ie != nil {
		ie = result2.Ie
	}
	return
}

// Parameters:
//  - BusinessId
//  - JobId
//  - RuleId
//  - StartTime
//  - EndTime
func (p *MessageServiceClient) CreateBreaker(businessId int32, jobId int32, ruleId int32, startTime int32, endTime int32) (r int64, se *ServiceError, err error) {
	if err = p.sendCreateBreaker(businessId, jobId, ruleId, startTime, endTime); err != nil {
		return
	}
	return p.recvCreateBreaker()
}

func (p *MessageServiceClient) sendCreateBreaker(businessId int32, jobId int32, ruleId int32, startTime int32, endTime int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("createBreaker", thrift.CALL, p.SeqId)
	args5 := NewCreateBreakerArgs()
	args5.BusinessId = businessId
	args5.JobId = jobId
	args5.RuleId = ruleId
	args5.StartTime = startTime
	args5.EndTime = endTime
	err = args5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *MessageServiceClient) recvCreateBreaker() (value int64, se *ServiceError, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error7 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error8 error
		error8, err = error7.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error8
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result6 := NewCreateBreakerResult()
	err = result6.Read(iprot)
	iprot.ReadMessageEnd()
	value = result6.Success
	if result6.Se != nil {
		se = result6.Se
	}
	return
}

// Parameters:
//  - Breaker
func (p *MessageServiceClient) DestoryBreaker(breaker int64) (se *ServiceError, err error) {
	if err = p.sendDestoryBreaker(breaker); err != nil {
		return
	}
	return p.recvDestoryBreaker()
}

func (p *MessageServiceClient) sendDestoryBreaker(breaker int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("destoryBreaker", thrift.CALL, p.SeqId)
	args9 := NewDestoryBreakerArgs()
	args9.Breaker = breaker
	err = args9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *MessageServiceClient) recvDestoryBreaker() (se *ServiceError, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error11 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error12 error
		error12, err = error11.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error12
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result10 := NewDestoryBreakerResult()
	err = result10.Read(iprot)
	iprot.ReadMessageEnd()
	if result10.Se != nil {
		se = result10.Se
	}
	return
}

type MessageServiceProcessor struct {
	*xiaoju.BaseServiceProcessor
}

func NewMessageServiceProcessor(handler MessageService) *MessageServiceProcessor {
	self13 := &MessageServiceProcessor{xiaoju.NewBaseServiceProcessor(handler)}
	self13.AddToProcessorMap("sendMessage", &messageServiceProcessorSendMessage{handler: handler})
	self13.AddToProcessorMap("createBreaker", &messageServiceProcessorCreateBreaker{handler: handler})
	self13.AddToProcessorMap("destoryBreaker", &messageServiceProcessorDestoryBreaker{handler: handler})
	return self13
}

type messageServiceProcessorSendMessage struct {
	handler MessageService
}

func (p *messageServiceProcessorSendMessage) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSendMessageArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("sendMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSendMessageResult()
	if result.Success, result.Se, result.Ie, err = p.handler.SendMessage(args.Message); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing sendMessage: "+err.Error())
		oprot.WriteMessageBegin("sendMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("sendMessage", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type messageServiceProcessorCreateBreaker struct {
	handler MessageService
}

func (p *messageServiceProcessorCreateBreaker) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewCreateBreakerArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("createBreaker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewCreateBreakerResult()
	if result.Success, result.Se, err = p.handler.CreateBreaker(args.BusinessId, args.JobId, args.RuleId, args.StartTime, args.EndTime); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing createBreaker: "+err.Error())
		oprot.WriteMessageBegin("createBreaker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("createBreaker", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type messageServiceProcessorDestoryBreaker struct {
	handler MessageService
}

func (p *messageServiceProcessorDestoryBreaker) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDestoryBreakerArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("destoryBreaker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDestoryBreakerResult()
	if result.Se, err = p.handler.DestoryBreaker(args.Breaker); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing destoryBreaker: "+err.Error())
		oprot.WriteMessageBegin("destoryBreaker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("destoryBreaker", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type SendMessageArgs struct {
	Message *Message `thrift:"message,1"`
}

func NewSendMessageArgs() *SendMessageArgs {
	return &SendMessageArgs{}
}

func (p *SendMessageArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SendMessageArgs) readField1(iprot thrift.TProtocol) error {
	p.Message = NewMessage()
	if err := p.Message.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Message)
	}
	return nil
}

func (p *SendMessageArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("sendMessage_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *SendMessageArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Message != nil {
		if err := oprot.WriteFieldBegin("message", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:message: %s", p, err)
		}
		if err := p.Message.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Message)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:message: %s", p, err)
		}
	}
	return err
}

func (p *SendMessageArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendMessageArgs(%+v)", *p)
}

type SendMessageResult struct {
	Success SerialNumber  `thrift:"success,0"`
	Se      *ServiceError `thrift:"se,1"`
	Ie      *InputError   `thrift:"ie,2"`
}

func NewSendMessageResult() *SendMessageResult {
	return &SendMessageResult{}
}

func (p *SendMessageResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SendMessageResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = SerialNumber(v)
	}
	return nil
}

func (p *SendMessageResult) readField1(iprot thrift.TProtocol) error {
	p.Se = NewServiceError()
	if err := p.Se.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Se)
	}
	return nil
}

func (p *SendMessageResult) readField2(iprot thrift.TProtocol) error {
	p.Ie = NewInputError()
	if err := p.Ie.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Ie)
	}
	return nil
}

func (p *SendMessageResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("sendMessage_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Ie != nil:
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	case p.Se != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *SendMessageResult) writeField0(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
		return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Success)); err != nil {
		return fmt.Errorf("%T.success (0) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 0:success: %s", p, err)
	}
	return err
}

func (p *SendMessageResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Se != nil {
		if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:se: %s", p, err)
		}
		if err := p.Se.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Se)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:se: %s", p, err)
		}
	}
	return err
}

func (p *SendMessageResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Ie != nil {
		if err := oprot.WriteFieldBegin("ie", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:ie: %s", p, err)
		}
		if err := p.Ie.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Ie)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:ie: %s", p, err)
		}
	}
	return err
}

func (p *SendMessageResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendMessageResult(%+v)", *p)
}

type CreateBreakerArgs struct {
	BusinessId int32 `thrift:"businessId,1"`
	JobId      int32 `thrift:"jobId,2"`
	RuleId     int32 `thrift:"ruleId,3"`
	StartTime  int32 `thrift:"startTime,4"`
	EndTime    int32 `thrift:"endTime,5"`
}

func NewCreateBreakerArgs() *CreateBreakerArgs {
	return &CreateBreakerArgs{}
}

func (p *CreateBreakerArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CreateBreakerArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.BusinessId = v
	}
	return nil
}

func (p *CreateBreakerArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.JobId = v
	}
	return nil
}

func (p *CreateBreakerArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.RuleId = v
	}
	return nil
}

func (p *CreateBreakerArgs) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.StartTime = v
	}
	return nil
}

func (p *CreateBreakerArgs) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 5: %s")
	} else {
		p.EndTime = v
	}
	return nil
}

func (p *CreateBreakerArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createBreaker_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
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
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *CreateBreakerArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("businessId", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:businessId: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.BusinessId)); err != nil {
		return fmt.Errorf("%T.businessId (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:businessId: %s", p, err)
	}
	return err
}

func (p *CreateBreakerArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("jobId", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:jobId: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.JobId)); err != nil {
		return fmt.Errorf("%T.jobId (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:jobId: %s", p, err)
	}
	return err
}

func (p *CreateBreakerArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ruleId", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:ruleId: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.RuleId)); err != nil {
		return fmt.Errorf("%T.ruleId (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:ruleId: %s", p, err)
	}
	return err
}

func (p *CreateBreakerArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("startTime", thrift.I32, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:startTime: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.StartTime)); err != nil {
		return fmt.Errorf("%T.startTime (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:startTime: %s", p, err)
	}
	return err
}

func (p *CreateBreakerArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("endTime", thrift.I32, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:endTime: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.EndTime)); err != nil {
		return fmt.Errorf("%T.endTime (5) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:endTime: %s", p, err)
	}
	return err
}

func (p *CreateBreakerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateBreakerArgs(%+v)", *p)
}

type CreateBreakerResult struct {
	Success int64         `thrift:"success,0"`
	Se      *ServiceError `thrift:"se,1"`
}

func NewCreateBreakerResult() *CreateBreakerResult {
	return &CreateBreakerResult{}
}

func (p *CreateBreakerResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CreateBreakerResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = v
	}
	return nil
}

func (p *CreateBreakerResult) readField1(iprot thrift.TProtocol) error {
	p.Se = NewServiceError()
	if err := p.Se.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Se)
	}
	return nil
}

func (p *CreateBreakerResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createBreaker_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Se != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *CreateBreakerResult) writeField0(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
		return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Success)); err != nil {
		return fmt.Errorf("%T.success (0) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 0:success: %s", p, err)
	}
	return err
}

func (p *CreateBreakerResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Se != nil {
		if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:se: %s", p, err)
		}
		if err := p.Se.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Se)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:se: %s", p, err)
		}
	}
	return err
}

func (p *CreateBreakerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateBreakerResult(%+v)", *p)
}

type DestoryBreakerArgs struct {
	Breaker int64 `thrift:"breaker,1"`
}

func NewDestoryBreakerArgs() *DestoryBreakerArgs {
	return &DestoryBreakerArgs{}
}

func (p *DestoryBreakerArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DestoryBreakerArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Breaker = v
	}
	return nil
}

func (p *DestoryBreakerArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("destoryBreaker_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *DestoryBreakerArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("breaker", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:breaker: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Breaker)); err != nil {
		return fmt.Errorf("%T.breaker (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:breaker: %s", p, err)
	}
	return err
}

func (p *DestoryBreakerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DestoryBreakerArgs(%+v)", *p)
}

type DestoryBreakerResult struct {
	Se *ServiceError `thrift:"se,1"`
}

func NewDestoryBreakerResult() *DestoryBreakerResult {
	return &DestoryBreakerResult{}
}

func (p *DestoryBreakerResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DestoryBreakerResult) readField1(iprot thrift.TProtocol) error {
	p.Se = NewServiceError()
	if err := p.Se.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Se)
	}
	return nil
}

func (p *DestoryBreakerResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("destoryBreaker_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Se != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *DestoryBreakerResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Se != nil {
		if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:se: %s", p, err)
		}
		if err := p.Se.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Se)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:se: %s", p, err)
		}
	}
	return err
}

func (p *DestoryBreakerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DestoryBreakerResult(%+v)", *p)
}
