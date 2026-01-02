#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <MQTTClient.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <linux/i2c-dev.h>


#define URL_1 "tcp://192.168.136.11:1883"
#define URL_2 "tcp://192.168.136.12:1883"
#define URL_3 "tcp://192.168.136.13:1883"
#define CLIENTID "BatterySensor"
#define TOPIC "sensors/battery"
#define USERNAME "cart_controller"
#define PASSWORD "cart_uniwa!@"
#define QOS 1
#define TIMEOUT 10000L



int main(int argc, char* argv){
    MQTTClient client;
    MQTTClient_connectOtions conn_opts = MQTTClient_connectOptions_initializer;
    MQTTClient_message pubmsg = MQTTClient_message_initializer;
    MQTTClient_deliveryToken token;
    int rc;

    MQTTClient_create(&client, URL_1, CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);
    conn_opts.keepAliveInterval = 20;
    conn_opts.cleansession = 1;
    conn_opts.username = USERNAME;
    conn_opts.password = PASSWORD;

    if((rc = MQTTClient_connect(client, &conn_opts)) != MQTTCLIENT_SUCCESS){
        printf("Failed to connect, return code %d\n", rc);
        MQTTClient_create(&client, URL_2, CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);
        if(rc = MQTTClient_connect(client, &conn_opts) != MQTTCLIENT_SUCCESS){
            printf("Failed to connect, return code %d\n", rc);
            MQTTClient_create(&client, URL_3, CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);
            if((rc = MQTTClient_connect(client, &conn_opts)) != MQTTCLIENT_SUCCESS){
                printf("Failed to connect, return code %d\n", rc);
                exit(EXIT_FAILURE);
            }
        }
    }

    while(1){
        int batt_level; //ina226 read battery level
        int isCharging = 0;

        char timeStamp[64];
        time_t now = time(NULL);
        strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));

        char payload[128];
        snprintf(payload, sizeof(payload),
                 "{\"battery_level\": %d, \"is_charging\": %s, \"timestamp_battery\": \"%s\"}",
                 batt_level,
                 isCharging ? "true" : "false",
                 timeStamp);

        pubmsg.payload = payload;
        pubmsg.payloadlen = strlen(payload);
        pubmsg.qos = QOS;
        pubmsg.retained = 0;

        MQTTClient_publish(client, TOPIC, pubmsg.payloadlen, pubmsg.payload, pubmsg.qos, pubmsg.retained, &token);
        MQTTClient_waitForCompletion(client, token, TIMEOUT);

        sleep(5);
    }

}