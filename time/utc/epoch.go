/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	string_p "memar/string/protocol"
)

var Epoch = epoch{}

type epoch struct{}

//memar:impl memar/datatype/protocol.Field_ID
func (e *epoch) DataTypeID() datatype_p.ID { return 0 }

//memar:impl memar/datatype/protocol.DataType_ID
func (e *epoch) DataTypeID_Base64() string { return "" }

//memar:impl memar/datatype/protocol.Detail
func (e *epoch) DevNote() string  { return "" }
func (e *epoch) Domain() string   { return "" }
func (e *epoch) Overview() string { return "" }
func (e *epoch) Summary() string  { return "" }
func (e *epoch) TAGS() []string   { return nil }
func (e *epoch) UserNote() string { return "" }

//memar:impl memar/datatype/protocol.Quiddity
func (e *epoch) Abbreviation() string { return "" }
func (e *epoch) Aliases() []string    { return nil }
func (e *epoch) Name() string         { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (e *epoch) ExpireInFavorOf() datatype_p.DataType { return nil }
func (e *epoch) ExpiryDate() string                   { return "" }
func (e *epoch) IssueDate() string                    { return "" }
func (e *epoch) ReferenceURI() string                 { return "" }

//memar:impl memar/datatype/protocol.Field_LifeCycle
func (e *epoch) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }

//memar:impl memar/string/protocol.Stringer_To
func (e *epoch) ToString() (str string_p.String, err error_p.Error) {
	return
}
