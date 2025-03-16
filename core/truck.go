package core

import (
	"app/structure"
	"fmt"
	"sync"
)

type Set = structure.Set

type Truck struct {
	tenants *Set
}

func Process(group *sync.WaitGroup) {
	truck := Truck{structure.NewSet()}
	for i := 0; i < 5; i++ {
		tenantId := fmt.Sprintf("%d", 10001001+i)
		truck.tenants.Add(tenantId)
	}

	fmt.Println(truck.tenants.Len())
	arr := truck.tenants.Values()
	for _, v := range arr {
		fmt.Println(v)
	}

	tinyTest()
	group.Done()
}

func tinyTest() {
	initDebug()
	c1, c2 := make(chan string), make(chan string)
	tenm := TenantManager{c1, c2, true}
	busm := newBusinessWorkers()

	go busm.listenerRegister(c1)
	go busm.listenerUnregister(c2)
	go tenm.scan()

	for i := 0; i < 5; i++ {
		tenantId := fmt.Sprintf("%d", 10001001+i)
		newChannel(tenantId)
		c1 <- tenantId
	}

	c := getChannel("10001001")
	c <- "tt"
	for i := 0; i < 5; i++ {
		tenantId := fmt.Sprintf("%d", 10001001+i)
		c2 <- tenantId
		freeChannel(tenantId)
	}
	busm.wg.Wait()
}
