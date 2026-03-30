#include "lib.h"

#define CLIENTID "ScaleSensor"
#define TOPIC "sensors/ScaleSensor"

#define DOUT_PIN 2;
#define SCK_PIN 3;

int main(){
  float scale_factor = 1;
  long offset = 0; // Fix syntax error
  digitalWrite(DOUT_PIN, LOW);
  digitalWrite(DOUT_PIN, HIGH);
  offset = digitalRead(DOUT_PIN);

  char timeStamp[64];
  while(1){
    if(offset == 0){
      // Scale is empty or error
      sleep(1);
      continue;
    }
    
    // Fake weight data logic
    float weight = 1.5; 
    time_t now = time(NULL);
    strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));

    printf("{\"weight_kg\": %.2f, \"timestamp_weight\": \"%s\"}\n",
             weight,
             timeStamp);
    fflush(stdout);
    sleep(1);
  }
  return 0;
}
