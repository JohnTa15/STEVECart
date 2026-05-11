#!/bin/bash 
interface="wlan0"

# making cart_id from the last 2 bytes of the mac address
cart_id=$(cat /sys/class/net/$interface/address | awk -F: '{print $5$6}')

#exporting cart_id to telegraf config file
echo "CART_ID=\"$cart_id\"" > /etc/default/telegraf 

#hostname is made of "cart-" + cart_id
hostname="cart-$cart_id"
ip_address=$(ip -o -f inet addr show $interface | awk '{print $4}' | head -n 1)

current_ip=${ip_address%/*}
mac_address=$(cat /sys/class/net/$interface/address)

current_gateway=$(ip route | grep default | awk '{print $3}')

ID=$(echo $current_ip | awk '{print $1}' | cut -d '.' -f 4)

ACTIVE_NODE=0

if [ -n "$current_ip" && -n "$current_gateway" ]; then
    echo "I don't have any ip available returning to the backup ip"
    current_ip="10.10.10.10"

echo "Seting static IP configuration for $interface"
echo "auto $interface
    iface $interface inet static
    address $current_ip
    gateway $current_gateway
    hostname $hostname" >> /etc/network/interfaces 
echo "Static IP configuration set:"
ip addr show $interface

echo "Setting hostname to $hostname"
echo $hostname > /etc/hostname
hostname -f /etc/hostname 

#Registering cart_ID to server's backend and pings if the sites are alive
if ping -c 2 -W 2 "$MASTER_IP" &> /dev/null; then
    echo "Master node ($MASTER_IP) is reachable."
	ACTIVE_NODE=$MASTER_IP
elif ping -c 2 -W 2 "$SECONDARY_IP" &> /dev/null; then
    echo "Master down. Secondary ($SECONDARY_IP) reachable."
    # sed -i "s/$MASTER_IP/$SECONDARY_IP/g" "$TELEGRAF_CONF"
	ACTIVE_NODE=$SECONDARY_IP
elif ping -c 2 -W 2 "$THIRD_IP" &> /dev/null; then
    echo "Master/Secondary down. Third ($THIRD_IP) reachable."
    # sed -i "s/$MASTER_IP/$THIRD_IP/g" "$TELEGRAF_CONF"
	ACTIVE_NODE=$THIRD_IP
else
    echo "All nodes unreachable."
fi
curl -X POST http://$ACTIVE_NODE:8089/registerCartID -H "Content-Type: application/json" -d "{\"cart_id\": \"$cart_id\"}"    
curl -X POST http://$ACTIVE_NODE:8089/registerCartID -H "Content-Type: application/json" -d "{\"mac_address\": \"$mac_address\"}"    

while true; do
	if ping -c 1 -W 5 "$ACTIVE_NODE" &> /dev/null; then
		curl -X POST http://$ACTIVE_NODE:8089/registerCartID -H "Content-Type: application/json" -d "{\"is_active\": \"true\"}"    
    fi
	
	sleep 30
done

# 5. Update /etc/hosts (Required for 'sudo' to not complain)
# echo "Updating /etc/hosts..."
# cat > /etc/hosts <<EOF
# 127.0.0.1   localhost localhost.localdomain
# 127.0.1.1   $NEW_HOSTNAME
# ::1         localhost localhost.localdomain
# EOF