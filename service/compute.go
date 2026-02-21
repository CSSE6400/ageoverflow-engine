package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Compute struct {
}

type Report struct {
	ID      string        `json:"id"`
	Results ReportResults `json:"results"`
}

type ReportResults struct {
	Checksum          string            `json:"checksum"`
	Generations       ReportGenerations `json:"generations"`
	PrimaryGeneration string            `json:"primary_generation"`
	Age               int               `json:"age"`
}

type ReportGenerations struct {
	Silent      int `json:"silent"`
	BabyBoomers int `json:"baby_boomers"`
	X           int `json:"x"`
	Y           int `json:"y"`
	Z           int `json:"z"`
	Alpha       int `json:"alpha"`
}

type InternalPayload struct {
	Iterations int    `json:"iterations"`
	Age        int    `json:"age"`
	Silent     int    `json:"silent"`
	Boomers    int    `json:"boomers"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Z          int    `json:"z"`
	Alpha      int    `json:"alpha"`
	Filler     string `json:"filler"`
}

type Request struct {
	ID      string   `json:"id"`
	Content []string `json:"content"`
}

func NewCompute() Compute {
	return Compute{}
}

func (d Compute) Spin(cost int) error {
	_, err := bcrypt.GenerateFromPassword([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), cost)
	return err
}

func (d Compute) Process(request Request) (_ Report, err error) {
	// if the request contains no content then thats an err
	if len(request.Content) < 1 {
		err = errors.New("invalid request, no content")
		return
	}

	var payloads []InternalPayload
	for _, content := range request.Content {
		var payload InternalPayload
		payload, err = transformPayload(content)
		if err != nil {
			return
		}
		payloads = append(payloads, payload)
	}

	report := Report{
		ID: request.ID,
	}

	// Generate Hash
	var hash int
	for _, payload := range payloads {
		hash += payload.Age
	}
	report.Results = ReportResults{}
	report.Results.Checksum = fmt.Sprintf("0x%x", hash)
	report.Results.Generations = ReportGenerations{}

	generations := make(map[string]int)
	for _, payload := range payloads {
		report.Results.Age += payload.Age
		generations["silent"] += payload.Silent
		generations["baby_boomers"] += payload.Boomers
		generations["x"] += payload.X
		generations["y"] += payload.Y
		generations["z"] += payload.Z
		generations["alpha"] += payload.Alpha
	}
	// Normalise all the results
	for gen, value := range generations {
		generations[gen] = value / len(payloads)
	}
	report.Results.Age = report.Results.Age / len(payloads)
	report.Results.Generations.Silent = generations["silent"]
	report.Results.Generations.BabyBoomers = generations["baby_boomers"]
	report.Results.Generations.X = generations["x"]
	report.Results.Generations.Y = generations["y"]
	report.Results.Generations.Z = generations["z"]
	report.Results.Generations.Alpha = generations["alpha"]

	// Choose the winner
	highestGen, highest := "", 0
	for gen, value := range generations {
		if value > highest {
			highestGen = gen
			highest = value
		}
	}
	report.Results.PrimaryGeneration = highestGen

	// we are definitely doing a lot of import work right here
	for _, payload := range payloads {
		if payload.Iterations == 0 {
			// Skipping any compute
			continue
		}
		err = d.Spin(payload.Iterations)
		if err != nil {
			return
		}
	}

	return report, nil
}

func transformPayload(payload string) (_ InternalPayload, err error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return
	}
	parts := strings.Split(string(decodedBytes), "|")
	if len(parts) != 9 {
		err = fmt.Errorf("invalid payload, expected 9 parts, got %d, payload %s", len(parts), string(decodedBytes))
		return
	}
	result := InternalPayload{
		Filler: strings.Join(parts[8:], "|"),
	}
	// 0 -> interations
	result.Iterations, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	// 1 -> Age
	result.Age, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}
	// 2 -> Silent
	result.Silent, err = strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	// 3 -> Boomers
	result.Boomers, err = strconv.Atoi(parts[3])
	if err != nil {
		return
	}
	// 4 -> X
	result.X, err = strconv.Atoi(parts[4])
	if err != nil {
		return
	}
	// 5 -> Y
	result.Y, err = strconv.Atoi(parts[5])
	if err != nil {
		return
	}
	// 6 -> Z
	result.Z, err = strconv.Atoi(parts[6])
	if err != nil {
		return
	}
	// 7 -> Alpha
	result.Alpha, err = strconv.Atoi(parts[7])
	if err != nil {
		return
	}

	return result, nil
}
