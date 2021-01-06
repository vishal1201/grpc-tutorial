package main

import (
	"context"
	"fmt"
	"log"

	datatypespb "../../proto/cognologix.com/datatypespb"
	personpb "../../proto/cognologix.com/human/personpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// var (
// 	d float64 = 1.0
// 	f float32 = 1.0
// )

func main() {
	fmt.Println("DataTypes gRPC client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	// gClient := datatypespb.NewDataTypeServiceClient(conn)
	gClient := personpb.NewPersonServiceClient(conn)
	// doScalarCall(gClient)
	// doEnumCall(gClient)
	addPerson(gClient, &personpb.PutPersonRequest{
		Person: &personpb.Person{
			FirstName: "Linus",
			LastName:  "Torvalds",
			Phones: []*personpb.PhoneNumber{
				&personpb.PhoneNumber{
					Number: "1111111111",
					PhType: personpb.PhoneNumber_HOME,
				},
			},
			Age:         51,
			LastUpdated: timestamppb.Now(),
		},
	})
	getPerson(gClient, &personpb.GetPersonRequest{
		Id: 1,
	})
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
