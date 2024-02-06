package bluesdp_proto

import (
	"bluesdp/database"
	"fmt"

	"golang.org/x/net/context"
)

type BlueSDPServer struct {
	BlueSDPHeartBeatServiceServer
}

func (server *BlueSDPServer) SayAlive(ctx context.Context, message *BlueHeartBeat) (*BlueHeartBeatResponse, error) {
	response_string := fmt.Sprintf("Hello From the  %v", message.IpAddress)
	database.CreateOrUpdate(message.ServiceName, message.IpAddress)
	return &BlueHeartBeatResponse{Reply: response_string}, nil
}

func (server *BlueSDPServer) GetList(ctx context.Context, message *GetRegistredServiceList) (*RespondRegistredServiceList, error) {
	var response RespondRegistredServiceList
	var records []database.ServiceRecord
	// fmt.Println(message.ServiceName)
	dbcon := database.ReturnSession()
	if err := dbcon.Model(database.ServiceRecord{}).Where("service_name =?", message.ServiceName).Find(&records).Error; err != nil {
		fmt.Println("Some Error Occoured")
	}

	for _, value := range records {
		response.Ips = append(response.Ips, &ServiceDetail{IpAddress: value.IpAddress})
	}

	return &response, nil
}
