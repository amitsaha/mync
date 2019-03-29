# mync - My network checker

Small utility (especially for non-Unix environments) to check if a port on a remote host is reachable.

**Features**

- Check if a IP:port is reachable

# Usage

mync ip:port

# Tips

## Download on Windows

```
C:\> [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
C:\> Invoke-WebRequest https://github.com/amitsaha/mync/releases/download/0.1/mync.exe -OutFile mync.exe
C:\> mync.exe <ip>:port
```