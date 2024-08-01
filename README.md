# go-bot-server

There's a lot of bots and crawlers out there. I'm curious to know where they come from.

This project is fairly straightforward: it accepts HTTP connections on a variety of ports, and records the originating IP address of the request along with the port of the request. For example, if someone is trying to make a HTTP connection on port 22, we can infer that they're looking for ssh servers to potentially exploit.

## Usage

This is a web server that listens for HTTP requests on the ports listed below. Not all of ever

# Notes

`lookup/ipLookup.go` is an example of the type of documentation I would have for a normal method. I'm a firm believer that comments don't exist to explain HOW some code works - that's only needed if its a complicated section of code. They should rather explain WHY a function exists, and WHAT it does.

I've commited the env vars file to the repo, even though that's terrible practice. This is so that y'all at Rockbot could run this yourselves. Don't worry, I'll rotate the keys and scrub the commit history when I'm done - there's a reminder in my calendar for it. The env vars only contains keys for [Supabase](https://supabase.com/), which is my database of choice.

I chose Supabase because its a free cloud DB ala Firebase. They allow for querying the DB through API calls, which makes it simple to use, even if there isn't an officially supported client in Go. Thankfully, there's [supabase-go](https://github.com/nedpals/supabase-go), which provides the minimal level of functionality I needed to make this MVP work.

To look up geolocation data, I'm using [IP-API](https://ip-api.com/docs/api:json). It's simple and lightweight - I don't need an API key, I just pass the IP address in a GET call. It returns json data containing many things, but I'm only using Country Code and City for now.

## Learnings & Challenges



## Future Work

I haven't tested this in the real world, which means my IP address splitting functionality may not actually work with real data. Spoofing the IP address is trickier as it can't be done from the application layer, which is why I couldn't test that locally.

The next thing to do would be to put this up in the cloud. Dockerizing this and having it sit behind a publicly exposed, passthrough network load balancer would allow for it to gather real data. Unit tests and integration tests would also be good to have.

## Day 1 log

* Spent about 3-4 hours today working on this
* Spent a good amount of time trying to figure out a clean way to spoof the calling IP address - resigned myself to faking it while running locally
* Spent some time trying to write a more complex script that pinged the service to generate test data while spoofing the calling IP address, decided to just spoof it in the code itself and make a simple bash script tomorrow
* Learned a decent amount about the way port management works on the OS level when calling TCP level resources
* Was concerned about listening on the HTTP vs the TCP/IP level, but given the complexity around handling http requests gracefully opted to listen on HTTP instead
* Annoyingly spent a lot of time trying to figure out why the port value I had wasn't the port I was listening on - turns out I was logging the client port, not the server port.
* Haven't used Go in ages so took a little time to shake off the rust

## Day 2 log
* Move the DB creds to env vars
* Have an endpoint to return snapshot of db data
* Cleaned up logging, error handling in DB layer, file structure and comments
* Created bash script that pings service every 10 seconds
* Created a Makefile that starts the service and runs the test script
* Write up README with intentions, challenges, learnings, future work
