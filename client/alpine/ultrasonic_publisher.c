#include "lib.h"

#define TRIG_PIN 4
#define ECHO_PIN 5
#define CLIENTID "UltrasonicSensor"
#define TOPIC "sensors/ultrasonic"

int main(){
    MQTT_init_and_connect(CLIENTID); 
    uint8_t trigPin;
    uint8_t echoPin;
    char payload[128];
    char timeStamp[64];
    if (trigPin == 0 && echoPin == 0){
        while (1){
            trigPin = 1;
            sleep_ms(10);
            trigPin = 0;
        }
        double distance = calculate_pulse_distance(trigPin, echoPin);
        strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));
        snprintf(payload, sizeof(payload),
                 "{\"Distance\": \"%fcm\", \"Timestamp\": \"%s\"}",
                 distance,
                 timeStamp);
        MQTT_publisher(TOPIC, payload);
        printf("Published ultrasonic data: %s\n", payload);
        sleep(5);
        MQTT_disconnect();
    } else{
        perror("Something went wrong..");
        exit(1);
    }
    return 0;
}

double get_distance(){
    //sending pulse in order to activate the sensor
    digitalWrite(TRIG_PIN, LOW);
    delayMicroseconds(2);
    digitalWrite(TRIG_PIN, HIGH);
    delayMicroseconds(10);
    digitalWrite(TRIG_PIN, LOW);
    double distance;
    while (echoPin != 0){
        if (echoPin == 1){
            startTime = clock_gettime;
            break;
        }
    }
    while (echoPin != 1){
        if (echoPin == 0){
            endTime = clock_gettime;
            break;
        }
    }
    time = endTime - startTime;
    distance = (double)(time * 34300) / 2.0; //speed of sound 34300 divided by 2 cause of the sound went and came..
    return distance;
}
