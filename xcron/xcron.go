package xcron

import "time"

type CronJob struct {
	ticker   *time.Ticker
	function func()
	running  bool
	name     string
}

/*
func SetInterval(call func(), sec int64) *time.Ticker {
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				call()
			}
		}
	}()
	return ticker
}
*/

func New() *CronJob {
	return &CronJob{
		running: false,
		name:    "",
	}
}

func (c *CronJob) Stop() {
	c.running = false
}

func (c *CronJob) SetInterval(sec int64) *CronJob {
	c.ticker = time.NewTicker(time.Duration(sec) * time.Second)
	return c
}

func (c *CronJob) SetCallback(call func()) *CronJob {
	c.function = call
	return c
}

func (c *CronJob) SetName(name string) *CronJob {
	c.name = name
	return c
}

func (c *CronJob) Run() *CronJob {
	if c.name != "" {
		if !c.running {
			c.running = true
			go c.run()
		}
	}
	return c
}

func (c *CronJob) run() {
	for {
		select {
		case <-c.ticker.C:
			if c.running {
				c.function()
			}
		}
	}
}
