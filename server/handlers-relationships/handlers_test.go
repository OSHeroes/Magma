package relationships

// func TestServer_relationshipRosterServer(t *testing.T) {
// 	ts := servertest.StartTestServer()
// 	defer ts.Close()

// 	req, _ := http.NewRequest(
// 		http.MethodGet,
// 		ts.URL+`/relationships/roster/server:`+serverName+`/bvip/1,3`,
// 		nil,
// 	)
// 	servertest.SetCommonTestHeaders(req)
// 	servertest.AddTestGameServerHeaders(req)

// 	res, _ := ts.Client().Do(req)

// 	t.Log(res)
// }

// func TestServer_relationshipRosterNucleus(t *testing.T) {
// 	ts := servertest.StartTestServer()
// 	defer ts.Close()

// 	req, _ := http.NewRequest(
// 		http.MethodGet,
// 		ts.URL+`/relationships/roster/nucleus:`+nucleusID,
// 		nil,
// 	)
// 	servertest.SetCommonTestHeaders(req)
// 	servertest.AddTestGameClientHeaders(req)

// 	res, _ := ts.Client().Do(req)

// 	t.Log(res)
// }
