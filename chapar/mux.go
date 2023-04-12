/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"bytes"
	"libgo/protocol"
)

// Multiplexer implement protocol.NetworkLink_Multiplexer interface
// Hardware implementation has some difference from Software(this) implementation:
// - multiplexers send frames on the connection to other mux not call other functionality
// - they must provide some congestion mechanism like cache to prevent sender frame.
// - mux must have some mechanism to drop frames on destination port unavailability (congestion, ...)
type Multiplexer struct {
	portNumber byte

	// Ports store all available link port to other physical or logical devices.
	ports [defaultPortNumber]port
	// UpperHandlers store all upper layer handlers
	upperHandlers [maxHeaderID]protocol.NetworkNetwork_Multiplexer

	connections Connections
}

func (mux *Multiplexer) HeaderID() (headerID protocol.NetworkPhysical_NextHeaderID) {
	return protocol.NetworkPhysical_Chapar
}

// Init initializes new Multiplexer object otherwise panic will occur on un-registered port or handler call.
func (mux *Multiplexer) Init(portNumber byte, pConnection protocol.NetworkPhysical_Connection, connections Connections) {
	// mux.physicalConnection = pConnection
	pConnection.RegisterLinkMultiplexer(mux)
	mux.portNumber = portNumber
	mux.connections = connections

	var i byte
	for i = 0; i <= 255; i++ {
		if i == portNumber {
			mux.ports[i].Init(portNumber, mux, pConnection)
		} else {
			// TODO::: get port info
			// mux.ports[i].Init(i, ?, ?)
		}

		var nonUH = UpperHandlerNonExist{headerID: protocol.NetworkLink_NextHeaderID(i)}
		mux.upperHandlers[i] = &nonUH
	}
}

// Deinit ready the connection pools to shutdown.
func (mux *Multiplexer) Deinit() {
	mux.connections.Deinit()
}

// RegisterNetworkMux registers new port on given ports pool.
func (mux *Multiplexer) RegisterNetworkMux(tm protocol.NetworkNetwork_Multiplexer) {
	// TODO::: check handler exist already and warn user??
	mux.upperHandlers[tm.HeaderID()] = tm
}

// UnRegisterNetworkMux delete the port by given port number on given ports pool.
func (mux *Multiplexer) UnRegisterNetworkMux(tm protocol.NetworkNetwork_Multiplexer) {
	var headerID = tm.HeaderID()
	var nonUH = UpperHandlerNonExist{headerID: headerID}
	mux.upperHandlers[headerID] = &nonUH
}

// removeTransportHandler delete the port by given port number on given ports pool.
func (mux *Multiplexer) removeTransportHandler(headerID protocol.NetworkLink_NextHeaderID) {
	var nonUH = UpperHandlerNonExist{headerID: headerID}
	mux.upperHandlers[headerID] = &nonUH
}

func (mux *Multiplexer) getTransportHandler(id protocol.NetworkTransport_HeaderID) protocol.NetworkNetwork_Multiplexer {
	return mux.upperHandlers[id]
}

// Send send the payload to all ports async.
func (mux *Multiplexer) Send(frame []byte) (err protocol.Error) {
	var f = Frame(frame)

	// err = f.CheckFrame()
	// if err != nil {
	// 	// TODO::: ???
	// 	return
	// }

	if f.IsBroadcastFrame() {
		err = mux.sendBroadcast(frame)
	} else {
		var portNum byte = f.NextPortNum()
		var port = mux.getPort(portNum)
		err = port.Send(frame)
	}
	return
}

// SendBroadcast send the payload to all ports async.
func (mux *Multiplexer) SendBroadcast(nexHeaderID NextHeaderID, payload protocol.Codec) (err protocol.Error) {
	var payloadLen int = payload.Len()
	if payloadLen > maxBroadcastPayloadLen {
		return &ErrMTU
	}

	var f Frame
	f.Init(nexHeaderID, nil, payloadLen)
	var framePayload = f.Payload()
	framePayload = framePayload[:0]
	_, err = payload.MarshalTo(framePayload)
	if err != nil {
		return
	}
	err = mux.sendBroadcast(f)
	return
}

// send the frame to all ports as BroadcastFrame!
func (mux *Multiplexer) sendBroadcast(frame []byte) (err protocol.Error) {
	// send the frame to all ports as BroadcastFrame!
	var portNum byte
	for portNum = 0; portNum <= 255; portNum++ {
		err = mux.getPort(portNum).Send(frame)
	}
	return
}

// Receive handles income frame to ports.
func (mux *Multiplexer) Receive(pConn protocol.NetworkPhysical_Connection, frame []byte) (err protocol.Error) {
	var f = Frame(frame)

	err = f.CheckFrame()
	if err != nil {
		// TODO::: ???
		return
	}

	var lastHop = f.IncrementNextHop(mux.portNumber)
	if lastHop {
		var path = f.Path()

		var conn *Connection
		conn, _ = mux.connections.GetConnectionByPath(path)
		if conn == nil {
			var newConn Connection
			newConn.Init(f, &mux.ports[mux.portNumber])
			conn = &newConn
			_ = mux.connections.RegisterConnection(conn)
		} else if !bytes.Equal(conn.pathFromPeer.Get(), path) {
			// TODO::: receive frame on alternative path, Any action needed??
		}

		var nextHeader = f.NextHeader()
		var payload = f.Payload()
		mux.getTransportHandler(nextHeader).Receive(conn, payload)
		return
	}

	if f.IsBroadcastFrame() {
		err = mux.sendBroadcast(frame)
	} else {
		var portNum byte = f.NextPortNum()
		err = mux.getPort(portNum).Receive(frame)
	}
	return
}

func (mux *Multiplexer) getPort(id byte) *port { return &mux.ports[id] }

func (mux *Multiplexer) registerPort(p port) {
	// TODO::: check port exist already and warn user??
	mux.ports[p.PortNumber()] = p
}

func (mux *Multiplexer) unRegisterPort(p port) {
	mux.removePort(p.PortNumber())
}

func (mux *Multiplexer) removePort(portNumber byte) {
	// mux.ports[portNumber].physicalConnection = TODO:::
}
