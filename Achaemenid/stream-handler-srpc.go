/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	errorr "../error"
	"../srpc"
)

// Indicate standard listen and send port number register for sRPC protocol.
const (
	ProtocolPortSRPCReceive uint16 = 4
	ProtocolPortSRPCSend    uint16 = 5
)

// SRPCHandler use to standard services handlers in any layer!
type SRPCHandler func(*Stream)

// SrpcIncomeRequestHandler handle incoming sRPC request streams!
func SrpcIncomeRequestHandler(s *Server, st *Stream) {
	st.Err = srpc.CheckPacket(st.IncomePayload, 4)
	if st.Err != nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.OutcomePayload = make([]byte, 4)
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st)
		return
	}

	st.Service = s.Services.GetServiceHandlerByID(srpc.GetID(st.IncomePayload))
	if st.Service == nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.OutcomePayload = make([]byte, 4)
		st.Err = srpc.ErrSRPCServiceNotFound
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st)
		return
	}

	// Handle request stream
	st.Service.SRPCHandler(st)
	if st.Err != nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.OutcomePayload = make([]byte, 4)
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st)
		return
	}

	// Handle response stream
	SrpcOutcomeResponseHandler(s, st)
}

// SrpcIncomeResponseHandler use to handle incoming sRPC response streams!
func SrpcIncomeResponseHandler(s *Server, st *Stream) {
	// Get error code from
	st.Err = errorr.GetErrByCode(srpc.GetID(st.IncomePayload))

	st.SetState(StateReady)
}

// SrpcOutcomeRequestHandler use to handle outcoming sRPC request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func SrpcOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	srpc.SetID(st.OutcomePayload, st.Service.ID)

	err = st.SendAndWait()
	// TODO::: handle send error almost due to no network available or connection closed!

	return
}

// SrpcOutcomeResponseHandler use to handle outcoming sRPC response stream!
func SrpcOutcomeResponseHandler(s *Server, st *Stream) (err error) {
	// Close request stream!
	st.Connection.StreamPool.CloseStream(st)

	// write error code to stream payload if exist.
	srpc.SetID(st.OutcomePayload, errorr.GetCode(st.Err))

	// send stream to outcome pool by weight
	err = st.Send()
	// TODO::: handle send error almost due to no network available or connection closed!

	return
}
