package actions

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/robfig/cron"
)

// Job is a struct used for handling
// scheduled actions using CRON
type Job struct {
	Cron *cron.Cron
}

// NewJob is a function used to initialized
// a Job struct with a Cron that can be reused
func NewJob() *Job {
	return &Job{
		Cron: cron.New(),
	}
}

// TurnOnRelay uses the duration string passed by parameter to know
// when to turn on the water pump.
func (j *Job) TurnOnRelay(duration string) {

	j.Cron.AddFunc(duration, func() {
		started := time.Now()
		fmt.Println("*** [*] CRON job 'TurnOnRelay' started ***")
		fmt.Printf("*** [*] CRON job 'TurnOnRelay' start time: %v ***\n", started)
		client := &http.Client{}

		value := `1`

		req, err := http.NewRequest("PUT", "https://watercontrol-9eaa1.firebaseio.com/motor.json", strings.NewReader(value))
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'TurnOnRelay' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' time elapsed: %v ***\n", ended.Sub(started))
		}

		res, err := client.Do(req)
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'TurnOnRelay' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'TurnOnRelay' time elapsed: %v ***\n", ended.Sub(started))
		}

		defer res.Body.Close()

		ended := time.Now()
		fmt.Println("*** [*] CRON job 'TurnOnRelay' finished succesfully ***")
		fmt.Printf("*** [*] CRON job 'TurnOnRelay' end time: %v ***\n", ended)
		fmt.Printf("*** [*] CRON job 'TurnOnRelay' time elapsed: %v ***\n", ended.Sub(started))
	})

	j.Cron.Start()
}

// TurnOffRelay uses the duration string passed by parameter to know
// when to turn off the water pump.
func (j *Job) TurnOffRelay(duration string) {

	j.Cron.AddFunc(duration, func() {
		started := time.Now()
		fmt.Println("*** [*] CRON job 'TurnOffRelay' started ***")
		fmt.Printf("*** [*] CRON job 'TurnOffRelay' start time: %v ***\n", started)

		client := &http.Client{}

		value := `2`

		req, err := http.NewRequest("PUT", "https://watercontrol-9eaa1.firebaseio.com/motor.json", strings.NewReader(value))
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'TurnOffRelay' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' time elapsed: %v ***\n", ended.Sub(started))
		}

		res, err := client.Do(req)
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'TurnOffRelay' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'TurnOffRelay' time elapsed: %v ***\n", ended.Sub(started))
		}

		defer res.Body.Close()

		ended := time.Now()
		fmt.Println("*** [*] CRON job 'TurnOnRelay' finished succesfully ***")
		fmt.Printf("*** [*] CRON job 'TurnOnRelay' end time: %v ***\n", ended)
		fmt.Printf("*** [*] CRON job 'TurnOnRelay' time elapsed: %v ***\n", ended.Sub(started))
	})

	j.Cron.Start()
}
