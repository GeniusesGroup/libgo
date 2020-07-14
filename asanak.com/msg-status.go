/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../http"
	"../json"
	"../log"
)

// MsgStatusReq is request structure of MsgStatus()
type MsgStatusReq struct {
	MsgIDs []string // max 100 number
}

// MsgStatusRes is response structure of MsgStatus()
type MsgStatusRes []struct {
	MsgID       string
	Status      string
	SendTime    string
	DeliverTime string
}

// Msg status
const (
	StatusNotFound         = "-1"
	StatusInQueue          = "1"
	StatusSent             = "2"
	StatusPending          = "4"
	StatusFailed           = "5"
	StatusSuccess          = "6"
	StatusNoResponse       = "7"
	StatusReject           = "8"
	StatusSentToDest       = "9"
	StatusPartiallySent    = "10"
	StatusNoMsgID          = "11"
	StatusPartiallySuccess = "12"
	StatusNoDelivery       = "13"
	// Status  = ""
)

// MsgStatus send req to asanak and return res to caller! It block caller until finish proccess.
func (a *Asanak) MsgStatus(req *MsgStatusReq) (res MsgStatusRes, err error) {
	// Due to legal paper not allow to send more than 100 destination at a request!
	var ln = len(req.MsgIDs)
	if ln == 0 || ln > 99 {
		log.Fatal("Message IDs number can't be 0 or more than 100")
	}

	return a.msgStatusByHTTP(req)
}

// https://panel.asanak.com/webservice/v1rest/msgstatus?Username=$USERNAME&Password=$PASSWORD&msgid=$MSGID
func (a *Asanak) msgStatusByHTTP(req *MsgStatusReq) (res MsgStatusRes, err error) {
	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Authority = domain
	httpReq.URI.Path = sendSMSPath
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.SetValue(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.SetValue(http.HeaderKeyContentType, "application/x-www-form-urlencoded")
	httpReq.Header.SetValue(http.HeaderKeyCacheControl, "no-cache")

	var msgIDLen = len(req.MsgIDs)
	// make a slice by needed size to avoid copy data on append! '+msgIDLen' use for destination seprators(',')!
	httpReq.Body = make([]byte, 0, sendSMSFixedSize+a.fixSizeDatalen+msgIDLen+(msgIDLen*len(req.MsgIDs[0])))
	// set username
	httpReq.Body = append(httpReq.Body, username...)
	httpReq.Body = append(httpReq.Body, a.UserName...)
	// set password
	httpReq.Body = append(httpReq.Body, password...)
	httpReq.Body = append(httpReq.Body, a.Password...)
	// set destination numbers
	httpReq.Body = append(httpReq.Body, msgID...)
	for i := 0; i < msgIDLen; i++ {
		httpReq.Body = append(httpReq.Body, req.MsgIDs[i]...)
		httpReq.Body = append(httpReq.Body, ',') // "," || "-"
	}

	var serverRes []byte
	serverRes, err = a.sendHTTPRequest(httpReq.Marshal())
	if err != nil {
		return nil, err
	}

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(serverRes)
	if err != nil {
		return nil, err
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return nil, ErrBadRequest
	case http.StatusUnauthorizedCode:
		log.Fatal("Authorization failed with asanak services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return nil, ErrAsanakServerInternalError
	}

	res = MsgStatusRes{}
	err = res.jsonDecoder(httpRes.Body)

	return
}

func (res *MsgStatusRes) jsonDecoder(buf []byte) (err error) {
	// TODO::: Help to complete json generator package to have better performance!
	err = json.UnMarshal(buf, res)
	return
}

func (a *Asanak) msgStatusBySRPC(req *MsgStatusReq) (res MsgStatusRes, err error) {
	// TODO::: when asanak.com impelement SRPC, complete SDK here!
	return
}
