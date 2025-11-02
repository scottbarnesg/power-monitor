#!/bin/bash

# Copy base file to the actual file
cp scripts/power-monitor.base.service scripts/power-monitor.service
# Set current user as user to run service as
echo "Setting user $USER as power monitor user..."
sed -i "/^User=/ s/$/$USER\n/" scripts/power-monitor.service
# Use current path to set execution path for service file
echo "Setting execution path for power monitor service..."
path=$(pwd)
echo "ExecStart=$path/run.sh" >> scripts/power-monitor.service
echo "WorkingDirectory=$path" >> scripts/power-monitor.service
# Copy service file to systemd directory
echo "Copying service file to systemd service directory..."
sudo cp scripts/power-monitor.service /etc/systemd/system/
# Enable and start service
echo "Enabling and starting service..."
sudo systemctl enable power-monitor
sudo systemctl start power-monitor
echo "Done!"