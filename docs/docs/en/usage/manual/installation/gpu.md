---
title: Perform Inference (3/3)
---

With CPU-Only

```
  execution: local
     script: post.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 30s (gracefulStop: 30s)


running (0m33.2s), 000/100 VUs, 1009 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  30s

     data_received..................: 291 kB 8.7 kB/s
     data_sent......................: 172 kB 5.2 kB/s
     http_req_blocked...............: avg=128.71µs min=713ns    med=1.08µs  max=18.93ms  p(90)=8.73µs  p(95)=583.83µs
     http_req_connecting............: avg=125.64µs min=0s       med=0s      max=18.91ms  p(90)=0s      p(95)=565.82µs
     http_req_duration..............: avg=3.18s    min=388.86ms med=3.27s   max=3.95s    p(90)=3.4s    p(95)=3.45s   
       { expected_response:true }...: avg=3.18s    min=388.86ms med=3.27s   max=3.95s    p(90)=3.4s    p(95)=3.45s   
     http_req_failed................: 0.00%  ✓ 0     ✗ 1009 
     http_req_receiving.............: avg=12.53ms  min=11.5µs   med=71.01µs max=47.05ms  p(90)=42.47ms p(95)=43.06ms 
     http_req_sending...............: avg=21.11µs  min=4.39µs   med=6.98µs  max=290.55µs p(90)=24.55µs p(95)=147.92µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=3.17s    min=388.64ms med=3.26s   max=3.94s    p(90)=3.4s    p(95)=3.41s   
     http_reqs......................: 1009   30.35459/s
     iteration_duration.............: avg=3.18s    min=389.88ms med=3.27s   max=3.95s    p(90)=3.4s    p(95)=3.45s   
     iterations.....................: 1009   30.35459/s
     vus............................: 48     min=48  max=100
     vus_max........................: 100    min=100 max=100
```

With GPU

```
  execution: local
     script: post.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 30s (gracefulStop: 30s)


running (0m31.3s), 000/100 VUs, 1013 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  30s

     data_received..................: 288 kB 9.2 kB/s
     data_sent......................: 172 kB 5.5 kB/s
     http_req_blocked...............: avg=115.27µs min=697ns    med=880ns   max=23.32ms  p(90)=6.32µs  p(95)=121.38µs
     http_req_connecting............: avg=101µs    min=0s       med=0s      max=23.3ms   p(90)=0s      p(95)=47.09µs 
     http_req_duration..............: avg=3.06s    min=461.11ms med=3.02s   max=3.9s     p(90)=3.07s   p(95)=3.9s    
       { expected_response:true }...: avg=3.06s    min=461.11ms med=3.02s   max=3.9s     p(90)=3.07s   p(95)=3.9s    
     http_req_failed................: 0.00%  ✓ 0     ✗ 1013 
     http_req_receiving.............: avg=34.1µs   min=10.55µs  med=32.88µs max=439.36µs p(90)=41.02µs p(95)=44.6µs  
     http_req_sending...............: avg=27.08µs  min=4.22µs   med=5.4µs   max=1.93ms   p(90)=17.84µs p(95)=81.18µs 
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=3.06s    min=461.06ms med=3.02s   max=3.9s     p(90)=3.07s   p(95)=3.9s    
     http_reqs......................: 1013   32.395143/s
     iteration_duration.............: avg=3.06s    min=461.17ms med=3.02s   max=3.9s     p(90)=3.07s   p(95)=3.9s    
     iterations.....................: 1013   32.395143/s
     vus............................: 59     min=59  max=100
     vus_max........................: 100    min=100 max=100
```