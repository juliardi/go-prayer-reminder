#!/bin/sh
echo "Start Building Prayer Reminder..."
echo "Step 1: Installing dependency"
dep ensure
echo "Step 2: Build Prayer Reminder"
go build -ldflags="-s -w" PrayerReminder.go
echo "Step 3: Compressing executables"
upx --brute PrayerReminder
echo "Finish"