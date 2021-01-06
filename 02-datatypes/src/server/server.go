package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "../../proto/cognologix.com/datatypespb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDataTypeServiceServer
}

func (s *server) TestScalarDataType(ctx context.Context, req *pb.ScalarDataTypeRequest) (*pb.ScalarDataTypeResponse, error) {
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
	res := &pb.ScalarDataTypeResponse{
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

func (s *server) TestEnumDataType(ctx context.Context, req *pb.EnumerationRequest) (*pb.EnumerationResponse, error) {
	planets := req.GetPlanet()
	fmt.Printf("d: %v, Type d: %T\n", planets, planets)
	var planetsStrArray []string
	for _, planet := range planets {
		planetsStrArray = append(planetsStrArray, pb.EnumerationRequest_Planet_name[int32(planet)])
	}
	res := &pb.EnumerationResponse{
		Planet: planetsStrArray,
	}
	return res, nil
}

func main() {
	fmt.Println("DataTypes gRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listen to localhost: %v", err)
	}
	gServer := grpc.NewServer()
	pb.RegisterDataTypeServiceServer(gServer, &server{})
	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
