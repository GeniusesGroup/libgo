/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Stream use to send or receive data on specific connection.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
// If OutcomePayload not present stream is UnidirectionalStream otherwise it is BidirectionalStream!
type Stream struct {
	ID             uint32 // Even number for Peer(who start connection). Odd number for server(who accept connection).
	Connection     *Connection
	Service        *Service
	IncomePayload  []byte // Income||Request data buffer.
	OutcomePayload []byte // Outcome||Response data buffer. Will divide to n packet to respect network MTU!

	/* State */
	Err          error      // Decode||Encode by ErrorID
	State        state      // States locate in const of this file.
	StateChannel chan state // States locate in const of this file.
	Weight       weight     // 16 queue for priority weight of the streams exist.

	/* Metrics */
	TotalPacket     uint32 // Expected packets count that must received!
	PacketReceived  uint32 // Count of packets received!
	LastPacketID    uint32 // Last send or received Packet use to know order of packets!
	PacketDropCount uint8  // Count drop packets to prevent some attacks type!
}

// SetState change state of stream and send notification on stream StateChannel.
func (st *Stream) SetState(state state) {
	st.State = state
	// notify stream listener that stream ready to use!
	st.StateChannel <- StateReady
}

// Send register stream in send pool. Usually use to send response stream.
func (st *Stream) Send() (err error) {
	// First Check st.Connection.Status to ability send stream over it

	// TODO::: remove this check when remove IP support
	if st.Connection.IPAddr.IP != nil {
		// Send by IP
	} else {
		// Send stream by GP
	}

	// send stream by weight

	st.Connection.BytesSent += uint64(len(st.OutcomePayload))

	// Last Close stream!
	st.Connection.StreamPool.CloseStream(st)

	return
}

// SendAndWait register stream in send pool and block caller until response ready to read.
func (st *Stream) SendAndWait() (err error) {
	st.Send()

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus state = <-st.StateChannel
	if responseStatus == StateReady {

	} else {

	}

	// Last Close response stream!
	st.Connection.StreamPool.CloseStream(st)

	return
}

// Authorize authorize request by data in related stream and connection.
func (st *Stream) Authorize() (err error) {
	err = st.Connection.AccessControl.AuthorizeWhen(0, 1)
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.AuthorizeWhich(st.Service.ID, st.Service.CRUD)
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.AuthorizeWhere(st.Connection.SocietyID, st.Connection.RouterID)
	if err != nil {
		return
	}
	return
}

// MakeNewStream use to make new stream without any connection exist yet!
// TODO::: due to IPv4&&IPv6 support we need this func! Remove it when remove those support!
func MakeNewStream() (st *Stream, err error) {
	// TODO::: check server allow make new connection and has enough resource

	st = &Stream{
		// ID:           0,
		State:        StateOpening,
		StateChannel: make(chan state),
	}
	return
}
