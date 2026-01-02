#!/bin/bash 
interface="wlan0"

cart_id=$(cat /sys/class/net/$interface/address | awk -F: '{print $5$6}')
hostname="cart-$cart_id"
ip_address=$(ip -o -f inet addr show $interface | awk '{print $4}' | head -n 1)

current_ip=${ip_address%/*}

current_gateway=$(ip route | grep default | awk '{print $3}')

ID=$(echo $current_ip | awk '{print $1}' | cut -d '.' -f 4)

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


# 5. Update /etc/hosts (Required for 'sudo' to not complain)
# echo "Updating /etc/hosts..."
# cat > /etc/hosts <<EOF
# 127.0.0.1   localhost localhost.localdomain
# 127.0.1.1   $NEW_HOSTNAME
# ::1         localhost localhost.localdomain
# EOF