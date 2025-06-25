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