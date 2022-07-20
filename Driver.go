package main

/*
#cgo  LDFLAGS: -lwsock32
#cgo  LDFLAGS: -lWs2_32
#cgo  CFLAGS:  -I  ./include
#cgo  LDFLAGS:  -L ./lib  -lnfapi

#include <stdio.h>
#include <stdlib.h>
#include "nfapi.h"
#include "Driver.h"
*/

/*
#cgo darwin LDFLAGS:  -Lc -L/usr/lib -lnfproxy  -lboost_system -lboost_thread -lstdc++ -v
#cgo darwin CFLAGS: -Iinclude

#define _C_API
#include <stddef.h>
#include <stdlib.h>

#include "nfapi_macos.h"
#include "Driver.h"



*/
import "C"

import (
	"fmt"
	"sync"
)

var uMap sync.Map

//void (NFAPI_CC *threadStart)();
//export go_threadStart
func go_threadStart() {
	fmt.Println("func go_threadStart() ")
}

//void (NFAPI_CC *threadEnd)();
//export go_threadEnd
func go_threadEnd() {
	fmt.Println("func go_threadEnd() ")
}

//void (NFAPI_CC *tcpConnectRequest)(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
//export go_tcpConnectRequest
func go_tcpConnectRequest(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) {
	fmt.Println("func go_tcpConnectRequest(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) ")
}

//void (NFAPI_CC *tcpConnected)(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
//export go_tcpConnected
func go_tcpConnected(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) {
	fmt.Println("func go_tcpConnected(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) ")
}

//void (NFAPI_CC *tcpClosed)(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
//export go_tcpClosed
func go_tcpClosed(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) {
	fmt.Println("func go_tcpClosed(id C.ENDPOINT_ID, pConnInfo C.PNF_TCP_CONN_INFO) ")
}

//void (NFAPI_CC *tcpReceive)(ENDPOINT_ID id, const char * buf, int len);
//export go_tcpReceive
func go_tcpReceive(id C.ENDPOINT_ID, buf *C.char, len C.int) {
	fmt.Println("func go_tcpReceive(id C.ENDPOINT_ID, buf *C.char, len C.int) ")
}

//void (NFAPI_CC *tcpSend)(ENDPOINT_ID id, const char * buf, int len);
//export go_tcpSend
func go_tcpSend(id C.ENDPOINT_ID, buf *C.char, len C.int) {
	fmt.Println("func go_tcpSend(id C.ENDPOINT_ID, buf *C.char, len C.int) ")
}

//void (NFAPI_CC *tcpCanReceive)(ENDPOINT_ID id);
//export go_tcpCanReceive
func go_tcpCanReceive(id C.ENDPOINT_ID) {
	fmt.Println("func go_tcpCanReceive(id C.ENDPOINT_ID) ")
}

//void (NFAPI_CC *tcpCanSend)(ENDPOINT_ID id);
//export go_tcpCanSend
func go_tcpCanSend(id C.ENDPOINT_ID) {
	fmt.Println("func go_tcpCanSend(id C.ENDPOINT_ID) ")
}

//void (NFAPI_CC *udpCreated)(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo);
//export go_udpCreated
func go_udpCreated(id C.ENDPOINT_ID, pConnInfo C.PNF_UDP_CONN_INFO) {
	fmt.Println("func go_udpCreated(id C.ENDPOINT_ID, pConnInfo C.PNF_UDP_CONN_INFO) ")
}

//void (NFAPI_CC *udpConnectRequest)(ENDPOINT_ID id, PNF_UDP_CONN_REQUEST pConnReq);
//export go_udpConnectRequest
func go_udpConnectRequest(id C.ENDPOINT_ID, pConnReq C.PNF_UDP_CONN_REQUEST) {
	fmt.Println("func go_udpConnectRequest(id C.ENDPOINT_ID, pConnReq C.PNF_UDP_CONN_REQUEST) ")
}

//void (NFAPI_CC *udpClosed)(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo);
//export go_udpClosed
func go_udpClosed(id C.ENDPOINT_ID, pConnInfo C.PNF_UDP_CONN_INFO) {
	fmt.Println("func go_udpClosed(id C.ENDPOINT_ID, pConnInfo C.PNF_UDP_CONN_INFO) ")
}

//void (NFAPI_CC *udpReceive)(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options);
//export go_udpReceive
func go_udpReceive(id C.ENDPOINT_ID, remoteAddress *C.uchar, buf *C.char, len C.int, options C.PNF_UDP_OPTIONS) {
	fmt.Println("func go_udpReceive(id C.ENDPOINT_ID, remoteAddress *C.uchar, buf *C.char, len C.int, options C.PNF_UDP_OPTIONS) ")
}

//void (NFAPI_CC *udpSend)(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options);
//export go_udpSend
func go_udpSend(id C.ENDPOINT_ID, remoteAddress *C.uchar, buf *C.char, len C.int, options C.PNF_UDP_OPTIONS) {
	fmt.Println("func go_udpSend(id C.ENDPOINT_ID, remoteAddress *C.uchar, buf *C.char, len C.int, options C.PNF_UDP_OPTIONS) ")
}

//void (NFAPI_CC *udpCanReceive)(ENDPOINT_ID id);
//export go_udpCanReceive
func go_udpCanReceive(id C.ENDPOINT_ID) {
	fmt.Println("func go_udpCanReceive(id C.ENDPOINT_ID) ")
}

//void (NFAPI_CC *udpCanSend)(ENDPOINT_ID id);
//export go_udpCanSend
func go_udpCanSend(id C.ENDPOINT_ID) {
	fmt.Println("func go_udpCanSend(id C.ENDPOINT_ID) ")
}

func DriverInit() C.int {

	return C.DriverMain()
}
