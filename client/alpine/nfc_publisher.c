#include "lib.h"

#include <nfc/nfc.h>
#include "utils/nfc-utils.h"
#include "libnfc/chips/pn53.h"

#define CLIENTID "NFCSensor"
#define TOPIC "sensors/nfc"

#define SPI_CLK 23
#define SPI_MISO 21
#define SPI_MOSI 19
#define SPI_CEO 24
#define RESET_PIN 20

static size_t simCards = 1; //total cards that they'll be scanned simultaneously
static uint8_t uiPollNr = 15; //total tries
static uint8_t uiPeriod = 2; //total seconds 

int main(int argc, char *argv[]){
  //https://github.com/nfc-tools/libnfc/blob/master/examples/
  const nfc_device *pnd = NULL; //NFC device(PN532)
  const nfc_context *context; //env libraries
  const nfc_target nt; //NFC tag
  init_NFC(&pnd, &context, nt);
  MQTT_init_and_connect(CLIENTID, TOPIC);
  char payload[128]; 
  char topic[64];
  char scanned_tag[64] = polling_data(&pnd, &context, nt);

  
  while (1) {
    time_t now = time(NULL);
    strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));

    snprintf(payload, sizeof(payload),
             "{\"NFC_data\": %s, \"timestamp_NFC\": \"%s\", \"PN532\"}",
             str_uid,
             timeStamp);

    MQTTClient_publisher(TOPIC, payload);
    printf("Waiting to remove the card from the reader..");
    fflush(stdout);
  }
  nfc_close(pnd);
  nfc_exit(context);
  MQTT_disconnect();
  return 0;
}

void init_NFC(nfc_device &pnd,nfc_context &context,nfc_target nt){
  nfc_init(&context);
  if(context == NULL){
    perror("Unable to initialize libnfc..");
    exit(0);
  }
  pnd = nfc_open(context, NULL);
  if(pnd == NULL){
    perror("Can't recognise the NFC device.. exiting");
    nfc_exit(context);
    exit(0);
  } 
  if(nfc_initiator_init(pnd) < 0){
    nfc_perror(pnd, "nfc_initiator_init");
    nfc_close(pnd);
    nfc_exit(context);
    exit(0);
  }

  printf("NFC device: %s", nfc_device_get_name(pnd));
}

char polling_data(&pnd, &context, nt){
  int res = 0;
  char str_uid[64] = "";
  char tmp[8];
  const nfc_modulation Modulations = { //Modulations for NTAG215 for NFC tags
    .nmt = NFC_ISO14443A, .nbr = NBR_106
  };
  if(res = nfc_initiator_poll_target(pnd, &Modulations,simCards, uiPollNr, uiPeriod, &nt) < 0){
    nfc_perror(pnd, "nfc_initiator_poll_target"); 
    nfc_close(pnd);
    nfc_exit(context);
    exit(0);
  }
  if(res > 0){
    for(int i = 0; i < nt.nti.nai.szUidLen; i++){ 
      sprintf(temp, "%02X", nt.nti.nai.abtUid[i]); //conversion byte to HEX
      strcat(str_uid, temp);
    }

    printf("UID is: ", str_uid);
  }
  return str_uid;
}


