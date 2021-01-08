package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	datatypespb "../../proto/cognologix.com/datatypespb"
	personpb "../../proto/cognologix.com/human/personpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// var (
// 	d float64 = 1.0
// 	f float32 = 1.0
// )

// type cognoDate struct {
// 	day, month, year int
// }

// func (x *cognoDate) ProtoReflect() protoreflect.Message {
// 	mi := &file_person_proto_msgTypes[3]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

func main() {
	fmt.Println("Person gRPC client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	// gClient := datatypespb.NewDataTypeServiceClient(conn)
	gClient := personpb.NewPersonServiceClient(conn)
	// doScalarCall(gClient)
	// doEnumCall(gClient)
	// a, err := anypb.New(&cognoDate{
	// 	day:   17,
	// 	month: 12,
	// 	year:  1994,
	// })
	// addPerson(gClient, &personpb.PutPersonRequest{Person: getPeople()})
	// getPerson(gClient, &personpb.GetPersonRequest{
	// 	Id: 3,
	// })
	// putPeople(gClient)
	// getAllPeople(gClient, &personpb.GetPersonRequest{Id: 2})
	chat(gClient)
}

func putPeople(client personpb.PersonServiceClient) {
	stream, err := client.PutPeople(context.Background())
	peopleSlice := getPeople()
	requests := []*personpb.PutPersonRequest{
		&personpb.PutPersonRequest{
			Person: []*personpb.Person{
				peopleSlice[0],
			},
		},
		&personpb.PutPersonRequest{
			Person: []*personpb.Person{
				peopleSlice[1],
			},
		},
		&personpb.PutPersonRequest{
			Person: []*personpb.Person{
				peopleSlice[2],
			},
		},
	}
	if err != nil {
		log.Fatalf("Error calling PutPeople stream API: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending -> %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from PutPeople stream: %v", err)
	}
	fmt.Printf("Reponse: %v\n", res)
}

func getAllPeople(client personpb.PersonServiceClient, req *personpb.GetPersonRequest) {
	resultStream, err := client.GetAllPeople(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling TestRPC: %v", err)
	}
	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error getting data from stream: %v", err)
		}
		fmt.Printf("Got Person: %v\n", msg.GetPerson())
	}
	fmt.Println(resultStream)
}

func addPerson(client personpb.PersonServiceClient, req *personpb.PutPersonRequest) {
	res, err := client.PutPerson(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling TestRPC: %v", err)
	}
	fmt.Println(res)
}

func getPerson(client personpb.PersonServiceClient, req *personpb.GetPersonRequest) {
	res, err := client.GetPerson(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling TestRPC: %v", err)
	}
	fmt.Println(res)
}

func chat(client personpb.PersonServiceClient) {
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error getting client stream: %v", err)
	}
	waitc := make(chan struct{})
	// send a bunch of messages to the server using go routine
	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Print("Send: ")
			if scanner.Text() != "\n" {
				replyMsg := scanner.Text()
				err := stream.Send(&personpb.ChatRequest{Message: replyMsg})
				if err != nil {
					log.Fatalf("Error sending message to chat server: %v", err)
					return
				}
			}

			if strings.ToLower(strings.Trim(scanner.Text(), " ")) == "end" {
				stream.CloseSend()
			}
		}
	}()

	// received a bunch of message from the server using go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error sending message to chat server: %v", err)
				return
			}
			receivedMsg := res.GetMessage()
			fmt.Printf("Received: %v\n", receivedMsg)
			if (err == io.EOF) || (strings.ToLower(strings.Trim(receivedMsg, " ")) == "bye") {
				close(waitc)
			}
			if err != nil {
				log.Fatalf("Error receiving message from chat client: %v", err)
				close(waitc)
			}
		}
	}()

	<-waitc
}

func doScalarCall(client datatypespb.DataTypeServiceClient) {
	req := &datatypespb.ScalarDataTypeRequest{
		D:      1.61803398875, //Golden Ratio
		F:      9.10938356,    // Mass of an electron
		I32:    5,
		I64:    299792458, //Speed of light
		Si32:   -1241422,
		Si64:   -42212414221241422,
		Fxd32:  9091941,  //Dennis Ritchie
		Fxd64:  10071856, //Nikola Tesla
		Sfxd32: 28121969, //Linus Torvalds
		Sfxd64: 16031953, // Richard Stallman
		B:      true,
		S:      "I am not one of the X-Men :(",
		Bts:    []byte{5, 2, 5, 2, 8, 43},
	}
	res, err := client.TestScalarDataType(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling TestRPC: %v", err)
	}
	d := res.GetD()
	f := res.GetF()
	i32 := res.GetI32()
	i64 := res.GetI64()
	ui32 := res.GetUi32()
	ui64 := res.GetUi64()
	si32 := res.GetSi32()
	si64 := res.GetSi64()
	fxd32 := res.GetFxd32()
	fxd64 := res.GetFxd64()
	sfxd32 := res.GetSfxd32()
	sfxd64 := res.GetSfxd64()
	b := res.GetB()
	s := res.GetS()
	bts := res.GetBts()
	fmt.Printf("d: %v, Type d: %T\n", d, d)
	fmt.Printf("f: %v, Type f: %T\n", f, f)
	fmt.Printf("i32: %v, Type i32: %T\n", i32, i32)
	fmt.Printf("i64: %v, Type i64: %T\n", i64, i64)
	fmt.Printf("ui32: %v, Type ui32: %T\n", ui32, ui32)
	fmt.Printf("ui64: %v, Type ui64: %T\n", ui64, ui64)
	fmt.Printf("si32: %v, Type si32: %T\n", si32, si32)
	fmt.Printf("si64: %v, Type si64: %T\n", si64, si64)
	fmt.Printf("f32: %v, Type f32: %T\n", fxd32, fxd32)
	fmt.Printf("f64: %v, Type f64: %T\n", fxd64, fxd64)
	fmt.Printf("sf32: %v, Type sf32: %T\n", sfxd32, sfxd32)
	fmt.Printf("sf64: %v, Type sf64: %T\n", sfxd64, sfxd64)
	fmt.Printf("b: %v, Type b: %T\n", b, b)
	fmt.Printf("s: %v, Type s: %T\n", s, s)
	fmt.Printf("bts: %v, Type bts: %T\n", bts, bts)
}

func doEnumCall(client datatypespb.DataTypeServiceClient) {
	req := &datatypespb.EnumerationRequest{
		Planet: []datatypespb.EnumerationRequest_Planet{datatypespb.EnumerationRequest_EARTH, datatypespb.EnumerationRequest_MARS},
	}
	res, err := client.TestEnumDataType(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling TestRPC: %v", err)
	}
	fmt.Println(res)
}

func getPeople() []*personpb.Person {
	return []*personpb.Person{
		{
			FirstName: "Linus",
			LastName:  "Torvalds",
			Phones: []*personpb.PhoneNumber{
				{
					Number: "1111111111",
					PhType: personpb.PhoneNumber_HOME,
				},
			},
			Age:         51,
			LastUpdated: timestamppb.Now(),
			// Extra:       []*anypb.Any{a},
			Family: map[string]*personpb.Person{
				"father": {
					FirstName: "Nils",
					LastName:  "Torvalds",
				},
				"mother": {
					FirstName: "Anna",
					LastName:  "Torvalds",
				},
			},
		},
		{
			FirstName: "Dennis",
			LastName:  "Ritchie",
			Phones: []*personpb.PhoneNumber{
				{
					Number: "2222222222",
					PhType: personpb.PhoneNumber_WORK,
				},
			},
			Age:         70,
			LastUpdated: timestamppb.Now(),
			// Extra:       []*anypb.Any{a},
			Family: map[string]*personpb.Person{
				"father": {
					FirstName: "Alistair",
					LastName:  "Ritchie",
				},
				"mother": {
					FirstName: "Jean McGee",
					LastName:  "Ritchie",
				},
			},
		},
		{
			FirstName: "Vishal",
			LastName:  "Zambare",
			Phones: []*personpb.PhoneNumber{
				{
					Number: "3333333333",
					PhType: personpb.PhoneNumber_MOBILE,
				},
			},
			Age:         27,
			LastUpdated: timestamppb.Now(),
			// Extra:       []*anypb.Any{a},
			Family: map[string]*personpb.Person{
				"father": {
					FirstName: "Suhdhakar",
					LastName:  "Zambare",
				},
				"mother": {
					FirstName: "Prarthana",
					LastName:  "Zambare",
				},
			},
		},
	}
}
