Output to file:
```
go tool pprof cpu.prof.pb.gz
```
```
(pprof) tags
```

Output to HTTP endpoint:
```
go tool pprof http://localhost:6060/debug/pprof/profile
```
(default sample size is 30 sec, can be changed by adding `?seconds=<custom_sample_size>` to endpoint URL)

Use custom CPU sampling rate (dafault 100 Hz may happen not to be enough to fetch any samples):
https://github.com/valteem/go-my-playground/blob/f450a867fcac03b04db24b0160e8eb662cd06296/my-snippets/ticker/main.go#L19