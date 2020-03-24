package main

import(
	 "os/exec"
	 "strconv"
)

var cmd string = "cat /proc/meminfo | awk '{if(NR==2)print $2}'"

func GetDevicesCnt()(int,error){
	str := GetRDMA()
	rdma_size,_ := strconv.Atoi(str)
	return int(rdma_size/1000000),nil
}

func GetRDMA()(string){
	c := exec.Command("/bin/bash","-c",cmd)
	out,_ := c.CombinedOutput()
	return string(out)
}

