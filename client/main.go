package main

import (
	"context"
	"fmt"
	"grpc/api"
	"io"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()
	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println("cities:")
	for _, city := range resp.Items {
		fmt.Printf("\t%s:%s\n", city.GetCityCode(), city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "tm_mr",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Weather in Mary:")
	for {
		msg, err := stream.Recv()
		// fmt.Println(msg.String())
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temprature: %.2f\n", msg.GetTemperature())

	}
	fmt.Println("Server stopped sending")
}
