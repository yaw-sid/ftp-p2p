# ftp-p2p

## Overview
This is a stripped down peer to peer version of FTP.

### Run
1. First build the project
```
go build .
```
2. Run first peer
```
./ftp-p2p -l 5000
```

3. Run second peer 
```
./ftp-p2p -l 5001 -d <first peer multiaddress>
```

4. Enter FTP commands
```
> PWD
```