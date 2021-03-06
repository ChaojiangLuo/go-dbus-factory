package hostname1

import "errors"
import "fmt"
import "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Hostname struct {
	hostname // interface org.freedesktop.hostname1
	proxy.Object
}

func NewHostname(conn *dbus.Conn) *Hostname {
	obj := new(Hostname)
	obj.Object.Init_(conn, "org.freedesktop.hostname1", "/org/freedesktop/hostname1")
	return obj
}

type hostname struct{}

func (v *hostname) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*hostname) GetInterfaceName_() string {
	return "org.freedesktop.hostname1"
}

// method SetHostname

func (v *hostname) GoSetHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHostname", flags, ch, arg0, arg1)
}

func (v *hostname) SetHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetStaticHostname

func (v *hostname) GoSetStaticHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetStaticHostname", flags, ch, arg0, arg1)
}

func (v *hostname) SetStaticHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetStaticHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPrettyHostname

func (v *hostname) GoSetPrettyHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrettyHostname", flags, ch, arg0, arg1)
}

func (v *hostname) SetPrettyHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetPrettyHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetIconName

func (v *hostname) GoSetIconName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIconName", flags, ch, arg0, arg1)
}

func (v *hostname) SetIconName(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetIconName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetChassis

func (v *hostname) GoSetChassis(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetChassis", flags, ch, arg0, arg1)
}

func (v *hostname) SetChassis(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetChassis(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetDeployment

func (v *hostname) GoSetDeployment(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeployment", flags, ch, arg0, arg1)
}

func (v *hostname) SetDeployment(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetDeployment(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetLocation

func (v *hostname) GoSetLocation(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocation", flags, ch, arg0, arg1)
}

func (v *hostname) SetLocation(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetLocation(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method GetProductUUID

func (v *hostname) GoGetProductUUID(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProductUUID", flags, ch, arg0)
}

func (*hostname) StoreGetProductUUID(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *hostname) GetProductUUID(flags dbus.Flags, arg0 bool) (arg1 []uint8, err error) {
	return v.StoreGetProductUUID(
		<-v.GoGetProductUUID(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// property Hostname s

func (v *hostname) Hostname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Hostname",
	}
}

// property StaticHostname s

func (v *hostname) StaticHostname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StaticHostname",
	}
}

// property PrettyHostname s

func (v *hostname) PrettyHostname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PrettyHostname",
	}
}

// property IconName s

func (v *hostname) IconName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IconName",
	}
}

// property Chassis s

func (v *hostname) Chassis() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Chassis",
	}
}

// property Deployment s

func (v *hostname) Deployment() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Deployment",
	}
}

// property Location s

func (v *hostname) Location() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Location",
	}
}

// property KernelName s

func (v *hostname) KernelName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KernelName",
	}
}

// property KernelRelease s

func (v *hostname) KernelRelease() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KernelRelease",
	}
}

// property KernelVersion s

func (v *hostname) KernelVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KernelVersion",
	}
}

// property OperatingSystemPrettyName s

func (v *hostname) OperatingSystemPrettyName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OperatingSystemPrettyName",
	}
}

// property OperatingSystemCPEName s

func (v *hostname) OperatingSystemCPEName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OperatingSystemCPEName",
	}
}

// property HomeURL s

func (v *hostname) HomeURL() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HomeURL",
	}
}
