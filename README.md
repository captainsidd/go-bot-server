# go-bot-server

There's a lot of bots and crawlers out there. I'm curious to know where they come from.

This project is fairly straightforward: it accepts TCP connections on a variety of ports, and records the originating IP address of the request along with the port of the request. For example, if someone is trying to make a TCP connection on port 22, we can infer that they're looking for ssh servers to potentially exploit.

## Day 1 log

Spent about 3-4 hours today working on this
Spent a good amount of time trying to figure out a clean way to spoof the calling IP address - resigned myself to faking it while running locally
Spent some time trying to write a more complex script that pinged the service to generate test data while spoofing the calling IP address, decided to just spoof it in the code itself and make a simple bash script tomorrow
Learned a decent amount about the way port management works on the OS level when calling TCP level resources
Was concerned about listening on the HTTP vs the TCP/IP level, but given the complexity around handling http requests gracefully opted to listen on HTTP instead
Annoyingly spent a lot of time trying to figure out why the port value I had wasn't the port I was listening on - turns out I was logging the client port, not the server port.
Haven't used Go in ages so took a little time to shake off the rust


## Things To Do
* Identify the ports and import them cleanly
* Move the DB creds to env vars
* Have an endpoint to return last 25 requests
* Create bash script that pings service every 10 seconds
* Create easy build command that starts service and scripts (Makefile, Dockerize)?
* Write up README with intentions, challenges, learnings, future work