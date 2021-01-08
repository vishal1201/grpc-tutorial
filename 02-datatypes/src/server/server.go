package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	datatypespb "../../proto/cognologix.com/datatypespb"
	personpb "../../proto/cognologix.com/human/personpb"
	"google.golang.org/grpc"
)

var people = make(map[int64]personpb.Person)

type dataTypeServer struct {
	datatypespb.UnimplementedDataTypeServiceServer
}

func (s *dataTypeServer) TestScalarDataType(ctx context.Context, req *datatypespb.ScalarDataTypeRequest) (*datatypespb.ScalarDataTypeResponse, error) {
	d := req.GetD()
	f := req.GetF()
	i32 := req.GetI32()
	i64 := req.GetI64()
	ui32 := req.GetUi32()
	ui64 := req.GetUi64()
	si32 := req.GetSi32()
	si64 := req.GetSi64()
	fxd32 := req.GetFxd32()
	fxd64 := req.GetFxd64()
	sfxd32 := req.GetSfxd32()
	sfxd64 := req.GetSfxd64()
	b := req.GetB()
	str := req.GetS()
	bts := req.GetBts()
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
	fmt.Printf("s: %v, Type s: %T\n", str, str)
	fmt.Printf("bts: %v, Type bts: %T\n", bts, bts)
	// fmt.Printf("req *.pb.DataTypeRequest -\n %v", req)
	res := &datatypespb.ScalarDataTypeResponse{
		D:      d,
		F:      f,
		I32:    i32,
		I64:    i64,
		Si32:   si32,
		Si64:   si64,
		Fxd32:  fxd32,
		Fxd64:  fxd64,
		Sfxd32: sfxd32,
		Sfxd64: sfxd64,
		B:      b,
		S:      str,
		Bts:    bts,
	}
	return res, nil
}

func (s *dataTypeServer) TestEnumDataType(ctx context.Context, req *datatypespb.EnumerationRequest) (*datatypespb.EnumerationResponse, error) {
	planets := req.GetPlanet()
	fmt.Printf("d: %v, Type d: %T\n", planets, planets)
	var planetsStrArray []string
	for _, planet := range planets {
		planetsStrArray = append(planetsStrArray, datatypespb.EnumerationRequest_Planet_name[int32(planet)])
	}
	res := &datatypespb.EnumerationResponse{
		Planet: planetsStrArray,
	}
	return res, nil
}

type personServer struct {
	personpb.UnimplementedPersonServiceServer
}

func (p *personServer) GetPerson(ctx context.Context, req *personpb.GetPersonRequest) (*personpb.GetPersonResponse, error) {
	id := req.GetId()
	fmt.Printf("d: %v, Type d: %T\n", req, req)
	person := people[id]
	fmt.Printf("Family: %v | Type:  %T", person.Family, person.Family)
	res := &personpb.GetPersonResponse{
		Person: &person,
	}
	return res, nil
}

func (p *personServer) PutPerson(ctx context.Context, req *personpb.PutPersonRequest) (*personpb.PutPersonResponse, error) {
	// person := req.GetPerson()
	fmt.Printf("d: %v, Type d: %T\n", req, req)
	var ids []int64
	for _, person := range req.GetPerson() {
		id := int64(len(people) + 1)
		people[id] = *person
		ids = append(ids, id)
	}
	res := &personpb.PutPersonResponse{
		Result: true,
		Id:     ids,
	}
	return res, nil
}

func (p *personServer) GetAllPeople(req *personpb.GetPersonRequest, stream personpb.PersonService_GetAllPeopleServer) error {
	log.Println("GetAllPeople() invoked")
	for _, person := range people {
		res := &personpb.GetPersonResponse{
			Person: &person,
		}
		stream.Send(res)
		time.Sleep(2000 * time.Millisecond)
	}
	return nil
}

func (p *personServer) PutPeople(stream personpb.PersonService_PutPeopleServer) error {
	var ids []int64
	for {
		req, err := stream.Recv()
		// fmt.Printf("Received!: %v\n", req)
		if err == io.EOF {
			return stream.SendAndClose(&personpb.PutPersonResponse{
				Result: true,
				Id:     ids,
			})
		}
		if err != nil {
			log.Fatalf("Error receving data from stream: %v", err)
		}
		for _, person := range req.GetPerson() {
			id := int64(len(people) + 1)
			people[id] = *person
			ids = append(ids, id)
		}
		fmt.Printf("People: %v \n ids: %v\n", people, ids)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listen to localhost: %v", err)
	}
	gServer := grpc.NewServer()
	// registerDataTypeServer(gServer, lis)
	registerPersonServer(gServer, lis)
}

func registerDataTypeServer(gServer *grpc.Server, listener net.Listener) {
	fmt.Println("DataTypes gRPC server")
	datatypespb.RegisterDataTypeServiceServer(gServer, &dataTypeServer{})
	if err := gServer.Serve(listener); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
func registerPersonServer(gServer *grpc.Server, listener net.Listener) {
	fmt.Println("Person gRPC server")
	personpb.RegisterPersonServiceServer(gServer, &personServer{})
	if err := gServer.Serve(listener); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
