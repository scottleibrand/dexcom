/*
Package dexcom provides functions to access a Dexcom CGM receiver
over a USB or BLE connection.

Based on the Python version at github.com/bewest/decoding-dexcom
*/
package dexcom

import (
	"log"

	"github.com/ecc1/usbserial"
)

type Connection interface {
	Send([]byte) error
	Receive([]byte) error
	Close()
}

type Cgm struct {
	conn Connection
	err  error
}

// Open first attempts to open a USB connection;
// if that fails it tries a BLE connection.
func Open() *Cgm {
	cgm := &Cgm{}
	cgm.conn, cgm.err = OpenUSB()
	if cgm.err == nil {
		return cgm
	}
	_, ok := cgm.err.(usbserial.DeviceNotFoundError)
	if !ok {
		log.Print(cgm.err)
	}
	cgm.conn, cgm.err = OpenBLE()
	return cgm
}

func (cgm *Cgm) Error() error {
	return cgm.err
}

func (cgm *Cgm) SetError(err error) {
	cgm.err = err
}
