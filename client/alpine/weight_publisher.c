#include "lib.h"

#define CLIENTID "ScaleSensor"
#define TOPIC "sensors/ScaleSensor"

#define DOUT_PIN 2;
#define SCK_PIN 3;

int main(){
  float scale_factor = 1;
  long offset = ;
  digitalWrite(DOUT_PIN, LOW);
  digitalWrite(DOUT_PIN, HIGH);
  offset = digitalRead(DOUT_PIN);
  while(1){
    if(offset == 0){
      printf("Scale is empty..");
    }

  }
  return 1;
}
