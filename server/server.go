package server

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"goymdial/conf"
	"goymdial/gen-go/rpc/dial/yamutech/com"
	"log"
	"os"
)

type Handle struct {
	com.Dial
}

func Dial_server_start() {
	serverTransport, err := thrift.NewTServerSocket(conf.Conf.DialServer)
	if err != nil {
		log.Println("Error!", err)
		os.Exit(1)
	}
	processor := com.NewDialProcessor(Handle{})
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	log.Println("Dial thrift server in", conf.Conf.DialServer)

	server.Serve()
}

func (h Handle) AddHealthPolicy(policy *com.HealthPolicyInfo) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) ModHealthPolicy(policy *com.HealthPolicyInfo) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelHealthPolicy(policy *com.HealthPolicyInfo) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddHealthGroup(groupName string, policyName string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelHealthGroup(groupName string, policyName string) (r com.RetCode, err error) {
	return com.RetCode_FAIL, nil
}

func (h Handle) AddHealthRecord(groupName string, records []*com.DialRecord) (r com.RetCode, err error) {
	return com.RetCode_FAIL, nil
}

func (h Handle) DelHealthRecord(groupName string, records []*com.DialRecord) (r com.RetCode, err error) {
	return com.RetCode_FAIL, nil
}

func (h Handle) AddDialServer(rid com.ObjectId, ip *com.IpAddr, typ com.DialServerType) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelDialServer(rid com.ObjectId) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddNginxGroup(groupName string, policyName string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelNginxGroup(groupName string, policyName string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddNginxServer(groupName string, servers []*com.DialNginxServer) (r com.RetCode, err error) {
	return com.RetCode_FAIL, nil
}

func (h Handle) DelNginxServer(groupName string, servers []*com.DialNginxServer) (r com.RetCode, err error) {
	return com.RetCode_FAIL, nil
}

func (h Handle) HeartBeat() (r *com.HeartBeatState, err error) {
	state := com.NewHeartBeatState()
	return state, nil
}

func (h Handle) SetServerState(enable bool) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddSnmpGroupInfo(snmp *com.SnmpGroupInfo) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelSnmpGroupInfo(snmp string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddSnmpProcessInfo(snmp string, processname string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelSnmpProcessInfo(snmp string, processname string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddIpSec(ipsec *com.SysIpSec, interval int32) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelIpSec(ipsecid string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddDcInfo(dc *com.DcInfo) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) DelDcInfo(id string) (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}
