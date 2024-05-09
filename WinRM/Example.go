package WinRM

import (
	"fmt"
	"github.com/masterzen/winrm"
	"golang.org/x/net/context"
	"time"
)

func WinRM() {
	// Define the WinRM client configuration

	username := "dhvani"
	password := "Mind@123"

	// Create a new WinRM client
	endpointConfig := winrm.NewEndpoint("172.16.8.113", 5985, false, true, nil, nil, nil, time.Second*30)
	client, err := winrm.NewClient(endpointConfig, username, password)
	if err != nil {
		fmt.Printf("Failed to create WinRM client: %v\n", err)
		return
	}

	// Define the command to execute
	command :=
		"(Get-Counter -Counter \"\\System\\System Up Time\").CounterSamples.CookedValue;" +
			"(Get-WmiObject Win32_ComputerSystem).Model;" +
			"(Get-WmiObject -Class Win32_OperatingSystem).Version;" +
			"(Get-WmiObject -Class Win32_ComputerSystem).NumberOfProcessors;" +
			"(Get-WmiObject -Class Win32_Processor| Select-Object -Property Description)\n"

	command = "(Get-Counter -Counter \"\\Processor(_Total)\\% Interrupt Time\").CounterSamples.CookedValue;" +
		"(Get-Counter -Counter \"\\Processor(_total)\\% Idle Time\").CounterSamples.CookedValue;" +
		"(Get-Counter -Counter \"\\Processor(_Total)\\% User Time\").CounterSamples.CookedValue;" +
		"(Get-WmiObject -Class Win32_Processor| Select-Object -Property Description);"
	//"(Get-WmiObject -Class Win32_Processor | Measure-Object -Property NumberOfCores -Sum).Sum;"

	//command := "Get-Counter -Counter \"\\Network Interface(*)\\Packets Sent/sec\" | Select-Object -ExpandProperty CounterSamples | Select-Object InstanceName, CookedValue,RawValue"

	//cmd, err := shell.ExecuteWithContext(context.Background(), command)

	output, errorOutput, exitCode, err := client.RunPSWithContext(context.Background(), command)

	if err != nil {
		fmt.Printf("Failed to execute command: %v\n", err)
		return
	}

	if exitCode != 0 {
		fmt.Printf("Failed to execute command: %v\n", errorOutput)
	} else {
		fmt.Printf("Successfully executed command:\n%v\n", output)
	}

	// Read the output of the command
	//output := cmd.Stdout
	//
	//buf := make([]byte, 4096)
	//for {
	//
	//	n, err := output.Read(buf)
	//	//fmt.Println(n, err)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Print(string(buf[:n]))
	//}

	// Close the command
	//err = cmd.Close()
	//if err != nil {
	//	fmt.Printf("Failed to close command: %v\n", err)
	//	return
	//}
}
