package xnetwork

import (
	"fmt"
	"fx/pkg/e"
	"fx/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/gopacket/pcap"
)

func Find(c *gin.Context) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		logging.Error(err.Error())
	}
	results := []map[string]string{}
	for _, device := range devices {
		ip := ""
		netmask := ""
		for _, address := range device.Addresses {
			ip = address.IP.String()
			netmask = address.Netmask.String()
		}
		tmp := map[string]string{"name": device.Name, "description": device.Description, "address": device.Description, "ip": ip, "netmask": netmask}
		results = append(results, tmp)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": e.SUCCESS,
		"msg":    e.GetMsg(e.SUCCESS),
		"result": results,
	})

	c.Abort()
}
