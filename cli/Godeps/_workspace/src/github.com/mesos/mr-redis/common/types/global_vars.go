package types

import (
	"container/list"

	"github.com/mesos/mr-redis/common/store"
)

var (
	Gdb   store.DB //Golabal variables related to db connection/instace
	MemDb *InMem   //In memory store

	OfferList *list.List       //list for having offer
	Cchan     chan TaskCreate  //Channel for Creator
	Mchan     chan *TaskUpdate //Channel for Maintainer
	Dchan     chan *Proc       //Channel for Destroyer
)

//Global db connection pointer, this will be initialized once abe be used everywhere

//global Constants releated to ETCD
const (
	ETC_BASE_DIR = "/MrRedis"
	ETC_INST_DIR = ETC_BASE_DIR + "/Instances"
	ETC_CONF_DIR = ETC_BASE_DIR + "/Config"
)

//Global constants for Instnace Status
//CREATING/ACTIVE/DELETED/DISABLED
const (
	INST_STATUS_CREATING = "CREATING"
	INST_STATUS_RUNNING  = "RUNNING"
	INST_STATUS_DISABLED = "DISABLED"
	INST_STATUS_DELETED  = "DELETED"
)

//Const for instance type
const (
	INST_TYPE_SINGLE       = "S"  //A Single instance redis-server
	INST_TYPE_MASTER_SLAVE = "MS" //A redis instnace with master-slave
)

//const for type of the redis-server
const (
	PROC_TYPE_MASTER = "M"
	PROC_TYPE_SLAVE  = "S"
)
