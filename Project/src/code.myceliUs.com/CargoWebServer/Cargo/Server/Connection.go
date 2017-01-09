package Server

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go.net/websocket"
)

/**
 * The connection interface is an abstraction of both
 * TCP connection and WebSocket connection.
 */
type connection interface {
	// Open a connection from a client...
	Open(host string, port int) (err error)

	// Close the connection
	Close()

	// The writting and reading loop...
	Reader()
	Writer()

	// id is the message id...
	Send(data []byte)

	// Tell if the connection is open...
	IsOpen() bool

	// Return the uuid for that connection.
	GetUuid() string

	// Generate a unique id...
	GenerateUuid()

	// get the connection string...
	GetAddrStr() string
}

////////////////////////////////////////////////////////////////////////////////
//									TCP
////////////////////////////////////////////////////////////////////////////////
/**
 * The tcp socket connection...
 */
type tcpSocketConnection struct {
	// The tcp socket connection.
	m_socket net.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The socket state.
	m_isOpen bool

	// The channel uuid...
	m_uuid string
}

func NewTcpSocketConnection() *tcpSocketConnection {
	var conn = new(tcpSocketConnection)

	// The connection is close at start...
	conn.m_isOpen = false

	// init members...
	conn.send = make(chan []byte /*, connection_channel_size*/)
	conn.GenerateUuid()

	return conn
}

func (c *tcpSocketConnection) GetAddrStr() string {
	address := c.m_socket.RemoteAddr().String()
	address = address[:strings.Index(address, ":")] // Remove the port...
	return address
}

func (c *tcpSocketConnection) GetUuid() string {
	return c.m_uuid
}

func (c *tcpSocketConnection) GenerateUuid() {
	c.m_uuid = uuid.NewRandom().String()
}

func (c *tcpSocketConnection) Open(host string, port int) (err error) {
	// Open the socket...
	c.m_socket, _ = net.Dial("tcp", host+":"+strconv.Itoa(port))

	if err != nil {
		log.Println("Connection with host ", host, " on port ", strconv.Itoa(port), " fail!!!")
		return err
	}
	log.Println("Connection with host ", host, " on port ", strconv.Itoa(port), " is open")
	c.m_isOpen = true
	return nil
}

func (c *tcpSocketConnection) Close() {
	c.m_socket.Close() // Close the socket..
	c.m_isOpen = false
}

/**
 * The connection state...
 */
func (c *tcpSocketConnection) IsOpen() bool {
	return c.m_isOpen
}

func (c *tcpSocketConnection) Send(data []byte) {
	msgSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgSize, uint32(len(data)))
	var data_ []byte
	data_ = append(data_, msgSize...)
	data_ = append(data_, data...)
	c.send <- data_

}

func (c *tcpSocketConnection) Reader() {
	//log.Println("Open new tcp connection whit id ", c.GetUuid())
	for c.m_isOpen == true {
		var in []byte

		// The input read the maximum message input...
		in = make([]byte, getMaxMessageSize()+200)

		if _, err := c.m_socket.Read(in); err != nil {
			log.Println("error!!! ", err)
			break
		}

		msgSize := int32(uint32(in[0]) | uint32(in[1])<<8 | uint32(in[2])<<16 | uint32(in[3])<<24)
		msgData := in[4 : msgSize+4] // The message start at 4 so it end four byte after...

		msg, err := NewMessageFromData(msgData, c)
		if err == nil {
			GetServer().GetHub().receivedMsg <- msg
		} else {
			log.Println("error: ", err)
		}
	}

	// End the connection...
	c.Close()
}

func (c *tcpSocketConnection) Writer() {
	for c.m_isOpen == true {
		for message := range c.send {
			// I will get the message here...
			c.m_socket.Write(message)
		}
		time.Sleep(time.Duration(time.Millisecond))
	}
	//log.Println("Close the tcp connection writer ", c.GetUuid())
	c.Close()
}

////////////////////////////////////////////////////////////////////////////////
//									WebSocket
////////////////////////////////////////////////////////////////////////////////
/**
 * The web socket connection...
 */
type webSocketConnection struct {
	// The websocket connection.
	m_socket *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The socket state.
	m_isOpen bool

	// The channel uuid...
	m_uuid string
}

func NewWebSocketConnection() *webSocketConnection {
	var conn = new(webSocketConnection)
	conn.send = make(chan []byte /*, connection_channel_size*/)
	conn.GenerateUuid()
	return conn
}

func (c *webSocketConnection) GetAddrStr() string {
	address := c.m_socket.Request().RemoteAddr
	address = address[:strings.Index(address, ":")] // Remove the port...
	return address
}

func (c *webSocketConnection) GetUuid() string {
	return c.m_uuid
}

func (c *webSocketConnection) GenerateUuid() {
	c.m_isOpen = true
	c.m_uuid = uuid.NewRandom().String()
}

func (c *webSocketConnection) Open(host string, port int) (err error) {
	c.m_isOpen = true
	// Open the socket...
	url := "http://" + host + ":" + strconv.Itoa(port)
	origin := "ws://" + host + ":" + strconv.Itoa(port)
	c.m_socket, err = websocket.Dial(origin, "", url)
	if err != nil {
		return err
	}
	return nil
}

func (c *webSocketConnection) Close() {
	c.m_socket.Close() // Close the socket..
	c.m_isOpen = false
}

/**
 * The connection state...
 */
func (c *webSocketConnection) IsOpen() bool {
	return c.m_isOpen
}

func (c *webSocketConnection) Send(data []byte) {
	c.send <- data
}

func (c *webSocketConnection) Reader() {
	for c.m_isOpen == true {
		var in []byte
		if err := websocket.Message.Receive(c.m_socket, &in); err != nil {
			break
		}
		msg, err := NewMessageFromData(in, c)
		if err == nil {
			GetServer().GetHub().receivedMsg <- msg
		}

	}
	// End the connection...
	c.Close()
	c.m_isOpen = false
}

func (c *webSocketConnection) Writer() {
	for c.m_isOpen == true {
		for message := range c.send {
			// I will get the message here...
			websocket.Message.Send(c.m_socket, message)
		}
		time.Sleep(time.Duration(time.Millisecond))
	}
	c.Close()
}

// The web socket handler function...
func HttpHandler(ws *websocket.Conn) {
	// Here I will create the new connection...
	c := NewWebSocketConnection()
	c.m_socket = ws

	GetServer().GetHub().register <- c

	defer func() {
		GetServer().GetHub().unregister <- c
	}()

	// Start the writing loop...
	go c.Writer()

	// continue to the reading loop...
	c.Reader()
}