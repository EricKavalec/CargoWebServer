package Server

import (
	"bufio"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
	"golang.org/x/net/websocket"
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

	// get the connection string...
	GetAddrStr() string

	// Return the port.
	GetPort() int
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

	// The number of connection try before give up...
	m_try int

	// The connection uuid
	m_uuid string
}

func NewTcpSocketConnection() *tcpSocketConnection {
	var conn = new(tcpSocketConnection)

	// The connection is close at start...
	conn.m_isOpen = false
	conn.m_try = 0

	// init members...
	conn.send = make(chan []byte /*, connection_channel_size*/)
	conn.m_uuid = Utility.RandomUUID()

	return conn
}

func (c *tcpSocketConnection) GetAddrStr() string {
	address := c.m_socket.RemoteAddr().String()
	address = address[:strings.Index(address, ":")] // Remove the port...
	return address
}

func (c *tcpSocketConnection) GetPort() int {
	address := c.m_socket.RemoteAddr().String()
	port, _ := strconv.Atoi(address[strings.LastIndex(address, ":")+1:])

	return port
}

func (c *tcpSocketConnection) GetUuid() string {
	return c.m_uuid
}

func (c *tcpSocketConnection) Open(host string, port int) (err error) {

	connectionId := host + ":" + strconv.Itoa(port)

	// Open the socket...
	c.m_socket, _ = net.Dial("tcp", connectionId)

	if err != nil {
		log.Println("Connection with host ", host, " on port ", strconv.Itoa(port), " fail!!!")
		return err
	}

	if c.m_socket == nil && c.m_try < 10 {
		c.m_try += 1
		time.Sleep(100 * time.Millisecond)
		c.Open(host, port)

	} else if c.m_try == 10 {
		return errors.New("fail to connect with " + connectionId)
	} else {
		c.m_isOpen = true
		GetServer().hub.register <- c

		// Start reading and writing loop's
		go c.Writer()
		go c.Reader()
	}

	return nil
}

func (c *tcpSocketConnection) Close() {
	if c.m_isOpen {
		c.m_isOpen = false
		c.m_socket.Close() // Close the socket..
		GetServer().hub.unregister <- c
		GetServer().removeAllOpenSubConnections(c.GetUuid())
	}
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
	// Set the buffer's
	var msgData []byte
	connbuf := bufio.NewReader(c.m_socket)
	for c.m_isOpen == true {
		b := make([]byte, 4096) // fix buffer size.
		size, err := connbuf.Read(b)
		if err == nil {
			msgData = append(msgData, b[0:size]...)
		}

		msg, err := NewMessageFromData(msgData, c)
		if err == nil {
			// The message is created so I will renew the buffer for the
			// next message to process.
			msgData = make([]byte, 0) // empty the buffer...
			GetServer().GetHub().receivedMsg <- msg
		}
	}

	// End the connection...
	c.Close()

}

func (c *tcpSocketConnection) Writer() {
	for c.m_isOpen == true {
		for msg := range c.send {
			// I will get the message here...
			c.m_socket.Write(msg)
		}
	}
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

	// number of try before giving up
	m_try int

	// The uuid of the connection
	m_uuid string
}

func NewWebSocketConnection() *webSocketConnection {
	var conn = new(webSocketConnection)
	conn.send = make(chan []byte /*, connection_channel_size*/)
	conn.m_uuid = Utility.RandomUUID()
	return conn
}

func (c *webSocketConnection) GetAddrStr() string {
	var address string
	if c.m_socket.IsServerConn() {
		address = c.m_socket.Request().RemoteAddr
	} else {
		address = c.m_socket.RemoteAddr().String()[5:] // ws:// (5 char to remove)
	}
	address = address[:strings.Index(address, ":")] // Remove the port...
	return address
}

func (c *webSocketConnection) GetPort() int {
	address := c.m_socket.RemoteAddr().String()
	port, _ := strconv.Atoi(address[strings.LastIndex(address, ":")+1:])
	return port
}

func (c *webSocketConnection) GetUuid() string {
	return c.m_uuid
}

func (c *webSocketConnection) Open(host string, port int) (err error) {

	// Open the socket...
	url := "http://" + host + ":" + strconv.Itoa(port)
	origin := "ws://" + host + ":" + strconv.Itoa(port)

	c.m_socket, err = websocket.Dial(origin, "", url)

	if err != nil && c.m_try < 10 {
		time.Sleep(100 * time.Millisecond)
		c.m_try += 1
		c.Open(host, port)
	} else if c.m_try == 10 {
		return errors.New("fail to connect with " + origin)
	} else if c.m_socket != nil {
		log.Println("--------> connection whit ", host, " at port ", port, " is now open!")
		c.m_isOpen = true
		GetServer().GetHub().register <- c

		// Start reading and writing loop's
		go c.Writer()
		go c.Reader()
	}

	return
}

func (c *webSocketConnection) Close() {
	if c.m_isOpen {
		c.m_isOpen = false
		c.m_socket.Close() // Close the socket..
		GetServer().GetHub().unregister <- c
		GetServer().removeAllOpenSubConnections(c.GetUuid())
	}
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
}

func (c *webSocketConnection) Writer() {
	for c.m_isOpen == true {
		for message := range c.send {
			// I will get the message here...
			websocket.Message.Send(c.m_socket, message)
		}
	}
}

// The web socket handler function...
func HttpHandler(ws *websocket.Conn) {

	// Here I will create the new connection...
	c := NewWebSocketConnection()
	c.m_socket = ws
	c.m_isOpen = true
	c.send = make(chan []byte)
	c.m_uuid = Utility.RandomUUID()
	GetServer().GetHub().register <- c

	defer func() {
		//  I will remove all sub-connection associated with the connection
		GetServer().removeAllOpenSubConnections(c.GetUuid())

		c.Close()
	}()

	// Start the writing loop...
	go c.Writer()

	// here the it stay in reader loop until the connection is close.
	c.Reader()
}
