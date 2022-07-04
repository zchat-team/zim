package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/panjf2000/gnet"
	"github.com/zmicro-team/zmicro/core/log"

	"github.com/zchat-team/zim/app/conn/protocol"
)

type TcpServer struct {
	gnet.EventHandler
	addr  string
	codec gnet.ICodec
	srv   *Server
}

func NewTcpServer(srv *Server, addr string) *TcpServer {
	ts := new(TcpServer)
	ts.addr = addr
	ts.codec = &TcpCodec{}
	ts.srv = srv
	return ts
}

func (s *TcpServer) Start() error {
	return gnet.Serve(s, s.addr, gnet.WithMulticore(true), gnet.WithTCPKeepAlive(time.Minute*5), gnet.WithCodec(s.codec))
}

func (s *TcpServer) Stop() error {
	//return gnet.Stop(context.Background(), s.addr)
	return nil
}

func (s *TcpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Infof("tcp server is listening on %s (multi-cores: %t, loops: %d)",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (s *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Info("TCP OnOpened ...")
	conn := &Connection{
		Status: AuthPending,
		Conn:   c,
	}
	c.SetContext(conn)

	s.srv.OnOpen(conn)

	return
}

func (s *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Info("TCP OnClose ...")

	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}

	s.srv.OnClose(conn)
	return
}

func (s *TcpServer) React(data []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}

	s.srv.OnMessage(data, conn)

	return
}

// ==================================== Codec ==============================================

type TcpCodec struct {
}

func (_ *TcpCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

func (_ *TcpCodec) Decode(c gnet.Conn) ([]byte, error) {
	if size, header := c.ReadN(protocol.HeaderLen); size == protocol.HeaderLen {
		byteBuffer := bytes.NewBuffer(header)
		var p protocol.Packet
		if err := binary.Read(byteBuffer, binary.BigEndian, &p.HeaderLen); err != nil {
			return nil, err
		}
		if err := binary.Read(byteBuffer, binary.BigEndian, &p.Version); err != nil {
			return nil, err
		}
		if err := binary.Read(byteBuffer, binary.BigEndian, &p.Cmd); err != nil {
			return nil, err
		}
		if err := binary.Read(byteBuffer, binary.BigEndian, &p.Seq); err != nil {
			return nil, err
		}
		if err := binary.Read(byteBuffer, binary.BigEndian, &p.BodyLen); err != nil {
			return nil, err
		}

		protocolLen := int(protocol.HeaderLen + p.BodyLen)
		if size, data := c.ReadN(protocolLen); size == protocolLen {
			c.ShiftN(protocolLen)
			return data, nil
		}
		return nil, errors.New("not enough payload data")
	}

	return nil, errors.New("not enough header data")
}
