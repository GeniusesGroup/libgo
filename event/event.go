/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

// Event implement memar/event/protocol.Event
//
//memar:impl memar/event/protocol.Event
type Event struct {
	dt   datatype_p.DataType
	time unix.Time
}

//memar:impl memar/computer/language/object/protocol.LifeCycle
func (self *Event) Init(dt datatype_p.DataType, time unix.Time) {
	self.dt = dt
	self.time = time
}

//memar:impl memar/event/protocol.Event
func (self *Event) DataType() datatype_p.DataType { return self.dt }
func (self *Event) Time() time_p.Time             { return &self.time }
func (self *Event) Cancelable() bool              { return false }
func (self *Event) DefaultPrevented() bool        { return false }
func (self *Event) Bubbles() bool                 { return false }

//memar:impl memar/event/protocol.Event_Methods
func (self *Event) PreventDefault() (err error_p.Error) { return nil }
