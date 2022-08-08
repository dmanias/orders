package main

import (
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/pampatzoglou/orders/orders/pb"

	//	pb "bookshop/server/pb/inventory"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// kafka
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//kafka
// produce
func produce(messages string, topicName string) {
	fmt.Printf("Starting producer...")
	/*	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"})
	*/

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		// "client.id":         socket.gethostname(),
		"client.id": "test",
		"acks":      "all"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := topicName
	data := []byte(messages)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

//Send data to Kafka topic
type Data struct {
	Id          int
	Message     string
	Source      string
	Destination string
}

// end kafka

//grpc
type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetOrderList(ctx context.Context, in *pb.GetOrderListRequest) (*pb.GetOrderListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetOrderListResponse{
		Orders: getSampleOrders(),
	}, nil
}

// end grpc

func main() {

	//kafka
	topicName := "myTopic"

	data := []Data{
		{Id: 2, Message: "World", Source: "1", Destination: "A"},
		{Id: 2, Message: "Earth", Source: "1", Destination: "B"},
		{Id: 2, Message: "Planets", Source: "2", Destination: "C"},
	}
	stringJson, _ := json.Marshal(data)

	fmt.Println(string(stringJson))

	produce(string(stringJson), topicName)
	//end kafka

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleOrders() []*pb.Order {
	sampleOrders := []*pb.Order{
		{
			Name:     "Banana",
			Category: "Food",
			Price:    12.1,
			Quantity: 3,
		},
		{
			Name:     "Banana",
			Category: "Food",
			Price:    12.1,
			Quantity: 3,
		},
	}
	return sampleOrders
}
