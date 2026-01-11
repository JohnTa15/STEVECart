#include "lib.h"

#define TRIG_PIN 4
#define ECHO_PIN 5
#define CLIENTID "UltrasonicSensor"
#define TOPIC "sensors/ultrasonic"

int main(int argc, char* argv[]){
  time_t total_time = NULL;
  if(wiringPiSetup() == -1){
    printf("WiringPi setup failed!\n");
    return -1;
  }
  pinMode(TRIG_PIN, OUTPUT);
    pinMode(ECHO_PIN, INPUT);
    calculate_distance(&start, &end); 
    MQTT_init_and_connect(CLIENTID);
    char payload[128];
    char timeStamp[64];
    while(1){
      double distance = calculate_distance();
      time_t now = time(NULL);
      strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));
      snprintf(payload, sizeof(payload),
               "{\"distance_cm\": %.2f, \"timestamp_ultrasonic\": \"%s\"}",
               distance,
               timeStamp);
      MQTTClient_publisher(TOPIC, payload);
      printf("Published ultrasonic data: %s\n", payload);
      sleep(5);
      MQTT_disconnect();
    }
    return 0;
}


double calculate_distance(){
  digitalWrite(TRIG_PIN, LOW);
  delayMicroseconds(2);
  digitalWrite(TRIG_PIN,HIGH);
  digitalWrite(TRIG_PIN, LOW);
  delayMicroseconds(10);
  digitalWrite(TRIG_PIN, LOW);

  struct timespec start, end;
  while(digitalRead(ECHO_PIN) == 1) 
    clock_gettime(CLOCK_MONOTONIC, &start);
  while(digitalRead(ECHO_PIN) == 0) 
    clock_gettime(CLOCK_MONOTONIC, &end);

  double timeElapsed = (end.tv_sec - start.tv_sec) + (end.tv_nsec - start.tv_nsec) / 1e9;
  double distance = (timeElapsed * 34300) / 2;

  return distance;
}
