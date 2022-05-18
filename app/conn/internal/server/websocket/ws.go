package websocket

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"

	"github.com/panjf2000/gnet"
)

func Upgrade(conn gnet.Conn, r *Request) (out []byte, err error) {
	if r.Method != http.MethodGet {
		err = ErrHandshakeBadMethod
		return
	}

	if !tokenListContainsValue(r.Header, headerConnectionCanonical, specHeaderValueConnection) {
		err = ErrHandshakeBadConnection
		return
	}

	if !tokenListContainsValue(r.Header, headerUpgradeCanonical, specHeaderValueUpgrade) {
		err = ErrHandshakeBadUpgrade
		return
	}

	if r.Method != http.MethodGet {
		err = ErrHandshakeBadMethod
		return
	}

	if !tokenListContainsValue(r.Header, headerSecVersionCanonical, specHeaderValueSecVersion) {
		err = ErrHandshakeBadSecVersion
		//return
	}

	challengeKey := r.Header.Get(headerSecKeyCanonical)
	if challengeKey == "" {
		err = ErrHandshakeBadSecKey
		return
	}

	if r.Header.Get(headerSecProtocolCanonical) != "" {

	}

	if r.Header.Get(headerSecExtensionsCanonical) != "" {

	}

	if err != nil {
		var code int
		if rej, ok := err.(*rejectConnectionError); ok {
			code = rej.code
			//header[1] = rej.header
		}
		if code == 0 {
			code = http.StatusInternalServerError
		}
		var buf bytes.Buffer
		bw := bufio.NewWriter(&buf)
		httpWriteResponseError(bw, err, code, nil)
		bw.Flush()
		out = buf.Bytes()
	} else {
		var buf bytes.Buffer
		bw := bufio.NewWriter(&buf)
		bw.WriteString(textHeadUpgrade)
		bw.WriteString(headerSecAccept)
		bw.WriteString(colonAndSpace)
		bw.WriteString(computeAcceptKey(challengeKey))
		bw.WriteString(crlf)

		bw.WriteString(crlf)

		bw.Flush()
		out = buf.Bytes()
	}

	return
}

// Errors used by both client and server when preparing WebSocket handshake.
var (
	ErrHandshakeBadProtocol = RejectConnectionError(
		RejectionStatus(http.StatusHTTPVersionNotSupported),
		RejectionReason(fmt.Sprintf("handshake error: bad HTTP protocol version")),
	)
	ErrHandshakeBadMethod = RejectConnectionError(
		RejectionStatus(http.StatusMethodNotAllowed),
		RejectionReason(fmt.Sprintf("handshake error: bad HTTP request method")),
	)
	ErrHandshakeBadHost = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerHost)),
	)
	ErrHandshakeBadUpgrade = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerUpgrade)),
	)
	ErrHandshakeBadConnection = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerConnection)),
	)
	ErrHandshakeBadSecAccept = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecAccept)),
	)
	ErrHandshakeBadSecKey = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecKey)),
	)
	ErrHandshakeBadSecVersion = RejectConnectionError(
		RejectionStatus(http.StatusBadRequest),
		RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecVersion)),
	)
)

var ErrShortPackaet = fmt.Errorf("short packet")

// ErrMalformedResponse is returned by Dialer to indicate that server response
// can not be parsed.
var ErrMalformedResponse = fmt.Errorf("malformed HTTP response")

// ErrMalformedRequest is returned when HTTP request can not be parsed.
var ErrMalformedRequest = RejectConnectionError(
	RejectionStatus(http.StatusBadRequest),
	RejectionReason("malformed HTTP request"),
)

// ErrHandshakeUpgradeRequired is returned by Upgrader to indicate that
// connection is rejected because given WebSocket version is malformed.
//
// According to RFC6455:
// If this version does not match a version understood by the server, the
// server MUST abort the WebSocket handshake described in this section and
// instead send an appropriate HTTP error code (such as 426 Upgrade Required)
// and a |Sec-WebSocket-ClientVersion| header field indicating the version(s) the
// server is capable of understanding.
var ErrHandshakeUpgradeRequired = RejectConnectionError(
	RejectionStatus(http.StatusUpgradeRequired),
	RejectionHeader(HandshakeHeaderString(headerSecVersion+": 13\r\n")),
	RejectionReason(fmt.Sprintf("handshake error: bad %q header", headerSecVersion)),
)

// ErrNotHijacker is an error returned when http.ResponseWriter does not
// implement http.Hijacker interface.
var ErrNotHijacker = RejectConnectionError(
	RejectionStatus(http.StatusInternalServerError),
	RejectionReason("given http.ResponseWriter is not a http.Hijacker"),
)

func FrameToBytes(f *Frame) (ret []byte, err error) {
	ret, err = WriteHeader(f.Header)
	ret = append(ret, f.Payload...)
	return
}
