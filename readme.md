# CLI aplication for handling yaml profiles

Author: Gosha Ivanov 

You can download the binary file from the bin folder to use this application (it was created for Windows 11 operating system)
Also you can build binary file using this commands (If you have Go installed on your computer): 

For MacOS/Linux: 
```
git clone https://github.com/prikokes/mws.git
go build -o bin/mws cmd/mws/main.go
```
For Windows: 
```
git clone https://github.com/prikokes/mws.git
go build -o bin\mws.exe cmd/mws/main.go
```

Application can handle such commands: 

You can create a profile:
```
.\mws.exe profile create --name=NAME --user=USER --project=PROJECT
```

You can get information about existing profile: 
```
.\mws.exe profile get --name=NAME
```

You can delete existing profile:
```
.\mws.exe profile delete --name=NAME
```

You can list existing profiles: 
```
.\mws.exe profile list
```

On MacOS/Linux you should use commands like this:
```
./mws profile create --name=NAME --user=USER --project=PROJECT
./mws profile get --name=NAME
./mws profile delete --name=NAME
./mws profile list
```
