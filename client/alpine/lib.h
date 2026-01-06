#ifndef LIB_H
#define LIB_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <time.h>
#include <MQTTClient.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <linux/i2c-dev.h>
#include <wiringPi.h>


#define URL_1 "tcp://192.168.136.11:1883"
#define URL_2 "tcp://192.168.136.12:1883"
#define URL_3 "tcp://192.168.136.13:1883"
#define USERNAME "cart_controller"
#define PASSWORD "cart_uniwa!@"
#define QOS 1
#define TIMEOUT 10000L

extern MQTT_Client client;
extern MQTTClient_connectOptions conn_opts;

void MQTT_init_and_connect(const char *CLIENTID);
void MQTT_publisher(const char *TOPIC, char *payload);
void MQTT_disconnect();
#endif