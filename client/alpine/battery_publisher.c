#include "lib.h"

#define CLIENTID "BatterySensor"
#define TOPIC "sensors/battery"

#define INA226_ADDRESS 0x40
#define INA226_REG_BUS_VOLTAGE 0x02
#define INA226_REG_SHUNT_VOLTAGE 0x01
#define INA226_REG_POWER 0x03
#define INA226_REG_CURRENT 0x04
#define INA226_REG_CALIBRATION 0x05

int main(int argc, char* argv){
    int i2c_fd;
    char *fn = "/dev/i2c-1";

    if((i2c_fd = open(fn, O_RDWR)) < 0){ //reading i2c bus from the i2c-1 file
        perror("Failed to open the i2c bus");
        exit(1);
    }

    if(ioctl(i2c_fd, I2C_SLAVE, INA226_ADDRESS) < 0){
        perror("Failed to have bus access.."); 
        exit(1);
    }
    MQTT_init_and_connect(CLIENTID);
    char payload[128];
    char timeStamp[64];
    while(1){
        int voltage_level = read_voltage(i2c_fd); //ina226 read battery level
        if(voltage_level != -1) int batt_level = calculate_percentage(voltage_level);
        int isCharging = 0;
        time_t now = time(NULL);
        strftime(timeStamp, sizeof(timeStamp), "%Y-%m-%dT%H:%M:%S%z", localtime(&now));

        snprintf(payload, sizeof(payload),
                 "{\"battery_level\": %d, \"is_charging\": %s, \"timestamp_battery\": \"%s\"}",
                 batt_level,
                 isCharging ? "true" : "false",
                 timeStamp);

        MQTTClient_publisher(TOPIC, payload);
        printf("Published battery data: %s\n", payload);
        sleep(5);

        close(i2c_fd);
        MQTT_disconnect();
        return 0;
    }

}

int read_voltage(int i2c_fd){
    uint8_t reg = INA226_REG_BUS_VOLTAGE;
    uint8_t buf[2];

    //writing to the driver what register we need to read
    if(write(i2c_fd, &reg, 1) != 1){ 
        perror("Failed to write to I2C Bus of RPI");
        return -1;
    }

    //reading 2 bytes(MSB, LSB) of the register
    if(read(i2c_fd, buf, 2) != 2){ 
        perror("Failed to read from I2C Bus of RPI");
        return -1;
    }

    //Merging MSB,lSB με αποτέλεσμα 16-bit
    uint16_t raw_value = (buf[0] << 8) | buf[1];
    //Converting to mV(LSB = 1.25mV)
    int voltage_mV = (int) (raw_value * 1.25);

    return voltage_mV;
}

int calculate_percentage(int voltage_mV){
    const int min_voltage = 10500;
    const int max_voltage = 12600;
    if(voltage_mV >= max_voltage) return 100; 
    if(voltage_mV <= min_voltage) return 0;

    int percentage = (voltage_mV - min_voltage) * 100 / (max_voltage - min_voltage);
    return percentage;
}