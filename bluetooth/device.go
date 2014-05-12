package main

import (
	"dbus/org/bluez"
	"dlib/dbus"
)

const deviceRssiNotInRange = -1000 // -1000db means device not in range

type device struct {
	bluezDevice *bluez.Device1

	Adapter dbus.ObjectPath
	Path    dbus.ObjectPath

	Alias     string
	Trusted   bool
	Paired    bool
	Connected bool

	// optinal
	Icon string
	RSSI int16
}

func (b *Bluetooth) newDevice(dpath dbus.ObjectPath, data map[string]dbus.Variant) (d *device) {
	d = &device{Path: dpath}
	d.bluezDevice, _ = bluezNewDevice(dpath)
	d.Adapter = d.bluezDevice.Adapter.Get()
	d.Alias = d.bluezDevice.Alias.Get()
	d.Trusted = d.bluezDevice.Trusted.Get()
	d.Paired = d.bluezDevice.Paired.Get()
	d.Connected = d.bluezDevice.Connected.Get()

	// optional properties, read from dbus object data in order to check if property is exists
	d.Icon = getDBusObjectValueString(data, "Icon")
	if isDBusObjectKeyExists(data, "RSSI") {
		d.RSSI = getDBusObjectValueInt16(data, "RSSI")
	} else {
		d.RSSI = deviceRssiNotInRange
	}

	// connect properties
	d.bluezDevice.Alias.ConnectChanged(func() {
		d.Alias = d.bluezDevice.Alias.Get()
		b.updatePropDevices()
	})
	d.bluezDevice.Trusted.ConnectChanged(func() {
		d.Trusted = d.bluezDevice.Trusted.Get()
		b.updatePropDevices()
	})
	d.bluezDevice.Paired.ConnectChanged(func() {
		d.Paired = d.bluezDevice.Paired.Get()
		b.updatePropDevices()
	})
	d.bluezDevice.Connected.ConnectChanged(func() {
		d.Connected = d.bluezDevice.Connected.Get()
		b.updatePropDevices()
	})
	d.bluezDevice.Icon.ConnectChanged(func() {
		d.Icon = d.bluezDevice.Icon.Get()
		b.updatePropDevices()
	})
	d.bluezDevice.RSSI.ConnectChanged(func() {
		d.RSSI = d.bluezDevice.RSSI.Get()
		b.updatePropDevices()
	})

	return
}

func (b *Bluetooth) addDevice(dpath dbus.ObjectPath, data map[string]dbus.Variant) {
	d := b.newDevice(dpath, data)
	if b.isDeviceExists(b.devices[d.Adapter], dpath) {
		logger.Warning("repeat add device:", dpath)
		return
	}
	b.devices[d.Adapter] = append(b.devices[d.Adapter], d)
	b.updatePropDevices()

	// send signal, DeviceAdded()
	if dbus.ObjectPath(b.PrimaryAdapter) == d.Adapter {
		if b.DeviceAdded != nil {
			b.DeviceAdded(marshalJSON(d))
		}
	}
}

func (b *Bluetooth) removeDevice(dpath dbus.ObjectPath) {
	// find adapter of the device
	for apath, devices := range b.devices {
		if b.isDeviceExists(devices, dpath) {
			b.devices[apath] = b.doRemoveDevice(devices, dpath)
			b.updatePropDevices()

			// send signal, DeviceRemoved()
			if dbus.ObjectPath(b.PrimaryAdapter) == apath {
				if b.DeviceRemoved != nil {
					d := device{Adapter: apath, Path: dpath}
					b.DeviceRemoved(marshalJSON(d))
				}
			}
			return
		}
	}
}
func (b *Bluetooth) doRemoveDevice(devices []*device, dpath dbus.ObjectPath) []*device {
	i := b.getDeviceIndex(devices, dpath)
	if i < 0 {
		logger.Warning("repeat remove device:", dpath)
		return devices
	}

	copy(devices[i:], devices[i+1:])
	devices[len(devices)-1] = nil
	devices = devices[:len(devices)-1]
	return devices
}

func (b *Bluetooth) isDeviceExists(devices []*device, dpath dbus.ObjectPath) bool {
	if b.getDeviceIndex(devices, dpath) >= 0 {
		return true
	}
	return false
}

func (b *Bluetooth) getDeviceIndex(devices []*device, dpath dbus.ObjectPath) int {
	for i, d := range devices {
		if d.Path == dpath {
			return i
		}
	}
	return -1
}

// GetDevices return all devices object that marshaled by json.
func (b *Bluetooth) GetDevices() (devicesJSON string) {
	return b.Devices
}

func (b *Bluetooth) ConnectDeivce(dpath dbus.ObjectPath) (err error) {
	err = bluezConnectDevice(dpath)
	return
}

func (b *Bluetooth) DisconnectDeivce(dpath dbus.ObjectPath) (err error) {
	err = bluezDisconnectDevice(dpath)
	return
}

func (b *Bluetooth) RemoveDevice(dpath dbus.ObjectPath) (err error) {
	err = bluezRemoveDevice(dbus.ObjectPath(b.PrimaryAdapter), dpath)
	return
}

func (b *Bluetooth) SetDeivceAlias(dpath dbus.ObjectPath, alias string) (err error) {
	err = bluezSetDeviceAlias(dpath, alias)
	return
}

func (b *Bluetooth) SetDeivceTrusted(dpath dbus.ObjectPath, trusted bool) (err error) {
	err = bluezSetDeviceTrusted(dpath, trusted)
	return
}