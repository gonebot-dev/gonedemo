# gonedemo
A demo of gonebot.
### Simple Usage
There are two steps to build a gonebot.  
1. Load plugins.
```go
gonebot.LoadPlugin(builtinplugins.Echo)
```
2. Start gonebot backend.
```go
gonebot.StartBackend("onebot11")
```
If successfully connected to your frontend, then congratulations, your bot is alive.
### Default Onebot11 configuation
Take Lagrange.Onebot as an example.
```json
{
    "Type": "ReverseWebSocket",
    "Host": "127.0.0.1",
    "Port": 2048,
    "Suffix": "/onebot/v11/ws",
    "ReconnectInterval": 5000,
    "HeartBeatInterval": 5000,
    "AccessToken": ""
}
```