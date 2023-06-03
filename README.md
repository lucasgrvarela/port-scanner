# Port Scanner

The Port Scanner is a command-line tool that allows you to scan ports on a target host. It can determine whether a specific port is open or closed by establishing a connection to it.

## Usage

To use the Port Scanner tool, follow the steps below:

1. Clone the repository:
```bash
git clone https://github.com/lucasgrvarela/port-scanner.git
```

2. Change into the project directory:
```bash
cd port-scanner
```

3. Build the executable:
```bash
go build -o port-scanner
```

4. Run the scanner:
```bash
./port-scanner [protocol] [hostname] [port]
```
- [protocol]: The protocol to use (e.g., tcp or udp).
- [hostname]: The hostname or IP address of the target.
- [port]: The port number to scan. Specify 0 to scan all ports from 1 to 1024.

## Example
1. Scan one specific port
```bash
./port-scanner tcp localhost 80
{tcp/80 Open}
```

2. Scan all ports from 1 to 1024
```bash
./port-scanner tcp localhost 0
Port result: {tcp/1 Closed}
Port result: {tcp/2 Closed}
Port result: {tcp/3 Closed}
...
Port result: {tcp/80 Open}
Port result: {tcp/81 Closed}
...
```

# Contributing
Contributions to this project are welcome! If you have any suggestions, bug reports, or want to add new features, feel free to open an issue or submit a pull request.
