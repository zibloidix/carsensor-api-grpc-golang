package main

import (
	"io"
	"log"
	"net"

	"github.com/zibloidix/carsensor-api-grpc-golang/carsensorpb"
	"google.golang.org/grpc"
)

type CarRoute struct {
	car    int32
	route  int32
	points []float32
}

type server struct{}

func (*server) SendPoint(stream carsensorpb.CarSensorService_SendPointServer) error {
	var car int32
	var route int32
	var lat float32
	var lon float32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&carsensorpb.SendPointResponse{
				Route:  route,
				Status: 200,
			})
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		car = req.GetCar()
		route = req.GetRoute()
		lat = req.GetLatitude()
		lon = req.GetLongitude()

		log.Printf("Car: %d, Route: %d, Lat: %d, Lon: %d \n", car, route, lat, lon)
	}
	return nil
}

func main() {
	log.Printf("CarSensor service start")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := grpc.NewServer()
	carsensorpb.RegisterCarSensorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
