package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const timeout = 10
const tictoc = 12

// Connection struct
type Connection struct {
	LocalIP  string `json:"localip"`
	RemoteIP string `json:"remoteip"`
}

// Report struct
type Report struct {
	Service    Service    `json:"service"`
	Connection Connection `json:"connection"`
	Errors     string     `json:"errors"`
	Elapsed    string     `json:"elapsed"`
	Status     string     `json:"status"`
	ConnTime   string     `json:"connection time"`
}

// Service Struct
type Service struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Network string `json:"network"`
}

func main() {

	csvFilename := flag.String("csv", "uat_services.csv", "A file with a list of services and ports")
	flag.Parse()

	go backgroundTask(csvFilename)
	fmt.Println("The rest of my application can continue")
	select {}

}

// func backgroundTask(csvFile *string) {
// 	ticker := time.NewTicker(tictoc * time.Second)

// 	for range ticker.C {
// 		file, err := os.Open(*csvFile)
// 		defer file.Close()
// 		if err != nil {
// 			log.Printf("Failed to open csv file: %s \n", *csvFile)
// 		}
// 		services := getServices(file)
// 		rawReports, err := makeRequest(services)
// 		if err != nil {
// 			log.Println("error converting making requests")
// 		}
// 		// fmt.Println(rawReports)
// 		reports, err := json.Marshal(rawReports)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		log.Println(string(reports))
// 	}
// }

// func makeRequest(services []Service) (reports []Report, err error) {
// 	ch := make(chan Report)
// 	var wg sync.WaitGroup
// 	timeOut := time.Duration(timeout) * time.Second
// 	for _, s := range services {
// 		wg.Add(1)
// 		// worker
// 		go func(s Service) {
// 			defer wg.Done()
// 			var r Report
// 			start := time.Now()
// 			conn, err := net.DialTimeout(s.Network, s.Address+":"+s.Port, timeOut)
// 			elapsed := time.Since(start)
// 			r.Elapsed = fmt.Sprintf("%dms", elapsed.Nanoseconds()/1000000)
// 			r.ConnTime = (time.Now().String())
// 			if err != nil {
// 				r.Errors = error.Error(err)
// 				r.Service = s
// 				r.Status = "failure"
// 				ch <- r
// 				return
// 			}
// 			r.Errors = "none"
// 			r.Service = s
// 			r.Connection.LocalIP = fmt.Sprintf(conn.LocalAddr().String())
// 			r.Connection.RemoteIP = fmt.Sprintf(conn.RemoteAddr().String())
// 			r.Status = "success"
// 			conn.Close()
// 			ch <- r
// 		}(s)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for range ch {
// 		rep := <-ch
// 		reports = append(reports, rep)
// 	}
// 	return reports, nil
// }

// func parseLines(lines [][]string) []Service {
// 	ret := make([]Service, len(lines))
// 	for i, line := range lines {
// 		ret[i] = Service{
// 			Name:    line[0],
// 			Profile: line[1],
// 			Address: line[2],
// 			Port:    line[3],
// 			Network: line[4],
// 		}
// 	}
// 	return ret
// }

// func getServices(file *os.File) []Service {
// 	csvReader := csv.NewReader(file)
// 	// eliminates first row
// 	if _, err := csvReader.Read(); err != nil {
// 		panic(err)
// 	}

// 	// reads the rest of the body
// 	log.Println("Reading csv file.")
// 	lines, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Printf(err.Error())
// 	}
// 	return parseLines(lines)
}
