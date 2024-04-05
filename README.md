# Cli interface for DMP Virtual Keypad alarm systems
This is a tool I threw together for interfacing with DMP alarm systems that
can be managed through the "Virtual Keypad" app. These are resold white label,
to mostly commercial buildings and some larger residential installations. Mostly
by ADT.

## Usage
At time of writing, two factor authentication signins aren't supported. However,
you still need to provide either an authtoken or a username/password combination.

You can provide the auth details through environment variables or command line
flags. You must provide one. The command args override the env. See `.env.example`
for a list of supported variables

## Buildling
To build for your current machine/architecture simply run the go build command
1. `git clone https://github.com/NotM32/keypadgo.git`
2. `cd ./keypadgo`
3. `go build *.go`
4. Execute the generated binary (`./main`)

### Examples
All examples assume you have configured the `.env` file copied from `.env.example`
and `source`d it with `source .env` (or `fish` equivelant)

#### Open / Lock / Access doors

1. run`./keypad doors` to list doors and their ids
   ```
   Authenticated via email & password
   [1] | FRONT DOOR READER
   [2] | MAN TRAP DOOR
   [3] | NOC DOOR
   [4] | DATA CENTER DOOR
   [6] | DATA CNTR RAMP DOOR
   [7] | ELECTRICAL ROOM DOOR
   [8] | SHIP/ REC INSIDE DOOR
   [9] | SHIP/ REC OUTSIDE DOOR
	```
2. run `./keypad doors access 2 3 5` to trigger the electric trim (like you scanned in)
3. run `./keypad doors unlock 2 4` to unlock permanently the selected doors
3. run `./keypad doors lock 2 4` to lock the doors again

## Contributers
This was a hastly thrown together hobby project; if anyone ends up using this in a
larger capacity and wants to contribute, you are welcome to submit a pull request.

## Support
You're welcome to open issues, however due to time constraints I may not be able to
respond right away.
