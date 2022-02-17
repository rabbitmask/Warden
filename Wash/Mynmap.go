package Wash

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap/v2"
	"strconv"
	"time"
)

func Mynmap(ip string) []map[string]string {
	fmt.Println("nmap start :" + ip)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		//nmap.WithPorts("80,443,843"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		fmt.Println("unable to create nmap scanner: %v", err)
	}
	result, warnings, err := scanner.Run()
	if err != nil {
		fmt.Println("unable to run nmap scan: %v", err)
	}
	if warnings != nil {
		fmt.Println("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	ris := [](map[string]string){}
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		//fmt.Printf("Host %q:\n", host.Addresses[0])

		//ris := [](map[string]string){}

		for _, port := range host.Ports {
			res := make(map[string]string)
			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			res["HOST"] = host.Addresses[0].Addr
			res["PORT"] = strconv.Itoa(int(port.ID))
			res["Protocol"] = port.Protocol
			res["State"] = port.State.State
			res["ServiceName"] = port.Service.Name
			//for k, v := range res {
			//	fmt.Printf("%s - %s\n", k, v)
			//}
			ris = append(ris, res)

			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
		//fmt.Println(ris)
		//return ris

	}
	//fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
	return ris
}
