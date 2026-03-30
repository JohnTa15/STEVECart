#include "lib.h"

MQTTClient client;
MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;

void MQTT_init_and_connect(const char *CLIENTID)
{
    int rc;
    char *URL[] = {URL_1, URL_2, URL_3};
    int connected = 0;
    for (int i = 0; i < 3; i++)
    {
        MQTTClient_create(&client, URL[i], CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);
        conn_opts.keepAliveInterval = 20;
        conn_opts.cleansession = 1;
        conn_opts.username = USERNAME;
        conn_opts.password = PASSWORD;
        for (int i = 0; i < 3; i++)
        {
            printf("Trying to connect to broker: %s\n", URL[i]);

            MQTTClient_create(&client, URL[i], CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);

            if ((rc = MQTTClient_connect(client, &conn_opts)) == MQTTCLIENT_SUCCESS)
            {
                printf("Connected successfully to %s\n", URL[i]);
                connected = 1;
                break;
            }
            else
            {
                printf("Failed to connect to %s, return code %d\n", URL[i], rc);
                MQTTClient_destroy(&client);
            }
        }
    }

    if (!connected)
    {
        printf("Could not connect to any broker. Exiting.\n");
        exit(EXIT_FAILURE);
    }
}

void MQTT_publisher(const char *TOPIC, char *payload)
{
    MQTTClient_message pubmsg = MQTTClient_message_initializer;
    MQTTClient_deliveryToken token;
    pubmsg.payload = payload;
    pubmsg.payloadlen = (int)strlen(payload);
    pubmsg.qos = QOS;
    pubmsg.retained = 0;
    if (MQTTClient_isConnected(client))
    {
        int rc = MQTT_publisherClient_publishMessage(client, TOPIC, &pubmsg, &token);
        if (rc == MQTTCLIENT_SUCCESS)
        {
            rc = MQTTClient_waitForCompletion(client, token, TIMEOUT);
            printf("Message with token %d delivered\n", token);
        }
        else
        {
            printf("Failed to publish, return code %d\n", rc);
        }
    }
    else
    {
        printf("Client is not connected. Cannot publish message.\n");
    }
}

void MQTT_disconnect()
{
    MQTTClient_disconnect(client, 10000);
    MQTTClient_destroy(&client);
}