#define _C_API

#include <stdio.h>
#include <string.h>
#include "Driver.h"
#include "nfapi_macos.h"






//netfilter2

#define NFDRIVER_NAME "nfext"

//C调Golang函数
extern void go_threadStart();
extern void go_threadEnd();
extern void go_tcpConnectRequest(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
extern void go_tcpConnected(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
extern void go_tcpClosed(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo);
extern void go_tcpReceive(ENDPOINT_ID id, const char * buf, int len);
extern void go_tcpSend(ENDPOINT_ID id, const char * buf, int len);
extern void go_tcpCanReceive(ENDPOINT_ID id);
extern void go_tcpCanSend(ENDPOINT_ID id);
extern void go_udpCreated(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo);
extern void go_udpConnectRequest(ENDPOINT_ID id, PNF_UDP_CONN_REQUEST pConnReq);
extern void go_udpClosed(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo);
extern void go_udpReceive(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options);
extern void go_udpSend(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options);
extern void go_udpCanReceive(ENDPOINT_ID id);
extern void go_udpCanSend(ENDPOINT_ID id);




void threadStart(){
	printf("void threadStart(){");
	go_threadStart();
}
void threadEnd(){
	printf("void threadEnd(){");
	go_threadEnd();
}
void tcpConnectRequest(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){
	printf("void tcpConnectRequest(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){");
	go_tcpConnectRequest(id,pConnInfo);
}
void tcpConnected(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){
	printf("void tcpConnected(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){");
	go_tcpConnected(id,pConnInfo);
}
void tcpClosed(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){
	printf("void tcpClosed(ENDPOINT_ID id, PNF_TCP_CONN_INFO pConnInfo){");
	go_tcpClosed(id,pConnInfo);
}
void tcpReceive(ENDPOINT_ID id, const char * buf, int len){
	printf("void tcpReceive(ENDPOINT_ID id, const char * buf, int len){");
	go_tcpReceive(id,buf,len);
}
void tcpSend(ENDPOINT_ID id, const char * buf, int len){
	printf("void tcpSend(ENDPOINT_ID id, const char * buf, int len){");
	go_tcpSend(id,buf,len);
}
void tcpCanReceive(ENDPOINT_ID id){
	printf("void tcpCanReceive(ENDPOINT_ID id){");
	go_tcpCanReceive(id);
}
void tcpCanSend(ENDPOINT_ID id){
	printf("void tcpCanSend(ENDPOINT_ID id){");
	go_tcpCanSend(id);
}
void udpCreated(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo){
	printf("void udpCreated(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo){");
	go_udpCreated(id,pConnInfo);
}
void udpConnectRequest(ENDPOINT_ID id, PNF_UDP_CONN_REQUEST pConnReq){
	printf("void udpConnectRequest(ENDPOINT_ID id, PNF_UDP_CONN_REQUEST pConnReq){");
	go_udpConnectRequest(id,pConnReq);
}
void udpClosed(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo){
	printf("void udpClosed(ENDPOINT_ID id, PNF_UDP_CONN_INFO pConnInfo){");
	go_udpClosed(id,pConnInfo);
}
void udpReceive(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options){
	printf("void udpReceive(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options){");
	go_udpReceive(id,remoteAddress,buf,len,options);
}
void udpSend(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options){
    printf("void udpSend(ENDPOINT_ID id, const unsigned char * remoteAddress, const char * buf, int len, PNF_UDP_OPTIONS options){");
    go_udpSend(id,remoteAddress,buf,len,options);
}
void udpCanReceive(ENDPOINT_ID id){
	printf("void udpCanReceive(ENDPOINT_ID id){");
	go_udpCanReceive(id);
}
void udpCanSend(ENDPOINT_ID id){
	printf("void udpCanSend(ENDPOINT_ID id){");
	go_udpCanSend(id);
}


//////////////////////////////////控制程序////////////////////////////////

NF_EventHandler eh = {
		threadStart,
		threadEnd,
		tcpConnectRequest,
		tcpConnected,
		tcpClosed,
		tcpReceive,
		tcpSend,
		tcpCanReceive,
		tcpCanSend,
		udpCreated,
		udpConnectRequest,
		udpClosed,
		udpReceive,
		udpSend,
		udpCanReceive,
		udpCanSend
	};


int DriverMain(){
    printf("\naaaaaaaaaaaaa\n");
    printf("NFDRIVER_NAME = %s\n",NFDRIVER_NAME);


    NF_STATUS a = nf_init("nfext", &eh);

	NFEXT_RULE rule;

	memset(&rule, 0, sizeof(rule));
	rule.fieldsMask = NFEXT_USE_REMOTE_PORTS;
	rule.ip_family = AF_INET;
	rule.remotePorts.from = 1;
	rule.remotePorts.to = 65535;
	rule.filteringFlag = NFEXT_REDIRECT;

	nf_addRule(&rule);

	memset(&rule, 0, sizeof(rule));
	rule.fieldsMask = NFEXT_USE_REMOTE_PORTS;
	rule.ip_family = AF_INET6;
	rule.remotePorts.from = 1;
	rule.remotePorts.to = 65535;
	rule.filteringFlag = NFEXT_REDIRECT;

	nf_addRule(&rule);

	return a;
}
