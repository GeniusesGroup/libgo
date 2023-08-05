/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
	"memar/time/monotonic"
)

// https://www.rfc-editor.org/rfc/rfc0793
// https://www.rfc-editor.org/rfc/rfc1122
type timingUserTimeout struct {
	enable         bool
	retransmission int
	idle           protocol.Duration
	synIdle        protocol.Duration
	lastUse        monotonic.Time
	nextCheck      monotonic.Time
	retryCount     int
}

//memar:impl memar/protocol.ObjectLifeCycle
func (ut *timingUserTimeout) Init(now monotonic.Time) (next protocol.Duration, err protocol.Error) {
	ut.enable = CNF_UserTimeout_PerStream
	ut.idle = CNF_KeepAlive_Idle
	now.Add(CNF_KeepAlive_Idle)
	ut.nextCheck = now
	return CNF_KeepAlive_Idle, nil
}
func (ut *timingUserTimeout) Reinit() (err protocol.Error) {
	ut.enable = CNF_UserTimeout_PerStream
	ut.nextCheck = 0
	ut.retryCount = 0
	return
}
func (ut *timingUserTimeout) Deinit() (err protocol.Error) {
	return
}

// Don't block the caller
func (ut *timingUserTimeout) CheckInterval(st *Stream, now monotonic.Time) (next protocol.Duration) {
	if !ut.enable {
		return -1
	}

	next = ut.nextCheck.Until(now)
	if next > 0 {
		return
	}

	var streamStatus = st.status.Load()
	// TODO::: check other stream status??
	if streamStatus != StreamStatus_Established {
		return -1
	}

	// check last use of stream and compare with our state
	if ut.lastUse != st.lastUse {
		ut.lastUse = st.lastUse
		now.Add(ut.idle)
		ut.nextCheck = now
		ut.retryCount = 0
		next = ut.idle
		return
	}

	ut.nextCheck.Add(next)

	// if (tp->packets_out || !tcp_write_queue_empty(sk))

	if ut.retryCount <= CNF_KeepAlive_Probes {
		// TODO::: send ack segment (keepalive message)
		var err = st.sendQuickACK()
		if err != nil {
			// TODO:::
		}
		ut.retryCount++
	} else {
		var err = st.close()
		if err != nil {
			// TODO:::
		}
		return
	}
	return
}
