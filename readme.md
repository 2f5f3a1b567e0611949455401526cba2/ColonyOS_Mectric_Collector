# Example readme with more to come.

Pages : https://github.com/users/2f5f3a1b567e0611949455401526cba2/projects/2


GOOD THINGS TO KNOW!: 
Man kan easily curla en get!
curl -X POST http://localhost:8080/metrics \
  -H "Content-Type: application/json" \
  -d '{"host":"test", "cpu":80, "memory":70, "timestamp":"2025-01-01T00:00:00Z"}'




What @firrenbirren02 wants!


[ Metric Agent ] -> [ Backend ] -> [ DB ] -> [ Grafana ]


MVP!
METRIC AGENT: 
  SKICKA JSONS PERIODVIS BESTÅENDE AV HOST,CPU%, MEMORY, TIMESTAMP !DONE!
  LOGS OUTGOING METRICS                                            !DONE!
  HTTP BACKEND                                                     !DONE!
BACKEND:
  FÅ IN POSTS TILL /metrics                                        !DONE!
  STORE METRICS I MONGODB                                          !DONE!
  HISTORISK METRICS                                                ?DONE?
  JSON API INTEGRATION (GIN)                                       !DONE!
  LOGGING                                                            X
GRAFANA: 
  BARA VISUALIZE HONESTLY                                          !DONE!



BACKEND: 
LOCALHOST:27017 KÖRS
