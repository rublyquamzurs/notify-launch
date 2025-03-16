package core

import (
	"time"
)

type TenantManager struct {
	tenantRegisterChannel   chan<- string
	tenantUnregisterChannel chan<- string
	running                 bool
}

func (t *TenantManager) scan() {
	for {
		n, f := getRegisterInfo()
		if registerRequest, ok := n.(string); ok {
			t.tenantRegisterChannel <- registerRequest
		}
		if unregisterRequest, ok := f.(string); ok {
			t.tenantUnregisterChannel <- unregisterRequest
		}
		time.Sleep(5 * time.Second)
		if !t.running {
			break
		}
	}
}

func getRegisterInfo() (interface{}, interface{}) {
	var o1, o2 interface{}
	select {
	case ctl := <-registers:
		o1 = ctl
	case <-time.After(1 * time.Second):
		o1 = nil
	}
	select {
	case ctl := <-unregister:
		o2 = ctl
	case <-time.After(1 * time.Second):
		o2 = nil
	}
	return o1, o2
}
