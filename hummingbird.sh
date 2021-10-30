#!/usr/bin/env bash

setup(){
	apt update 2> /dev/null > /dev/null
	apt install -y sysstat >/dev/null 2>/dev/null
}

_date(){
	echo "@#@#@#@#@ begin logs"
	date
}

memory(){
	echo "@#@#@#@#@ free memory"
	free -h
}
cpu(){
	echo "@#@#@#@#@ processor core count"
	grep processor /proc/cpuinfo | wc -l
}

disk(){
	echo "@#@#@#@#@ free disk space"
	df
	echo  "@#@#@#@#@ disk utilization"
	iostat -p -d 1 2 | tail -n +4 | awk '/Device/{y=1;p}y'
}


network(){
	echo "@#@#@#@#@ networking information"
	ifconfig
	# vnstat is required for usage
}

packages(){
	echo "@#@#@#@#@ installed packages"
	apt list --installed 2> /dev/null | wc -l
	echo "@#@#@#@#@ packages with updates"
	apt list –upgradable 2> /dev/null | wc -l
}

services(){
	echo "@#@#@#@#@ systemctl status"
	systemctl list-units --failed -q
	echo "@#@#@#@#@ docker stats"
	docker stats --no-stream --no-trunc
}

_uptime(){
	echo "@#@#@#@#@ uptime"
	uptime

	today=`date`
	strday=${1:-${today:0:10}}
	month=$(echo $strday | awk '{ print $2 }')
	N_Crash=`last -F  |grep crash  |  grep "$month"  | sort -k 7 -u | wc -l`
	N_Reboot=`last -F | grep reboot | grep "$month"  | wc -l `
	echo "@#@#@#@#@ crashes this month"
	echo "$N_Crash"
	echo "@#@#@#@#@ reboots this month"
	echo "$N_Reboot"
}


{ setup ;
	_date ;
	memory ;
	cpu ;
	disk ;
	network ;
	packages ;
	services ;
	_uptime ;  } >> ~/hummingbird.log
