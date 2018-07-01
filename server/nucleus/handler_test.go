package nucleus_test

func TestServer_nucleusCheckUser(t *testing.T) {
	ts := servertest.StartTestServer()
	defer ts.Close()

	req, _ := http.NewRequest(
		http.MethodGet,
		ts.URL+`/nucleus/check/user/`+nucleusID,
		nil,
	)
	servertest.SetCommonTestHeaders(req)
	servertest.AddTestGameClientHeaders(req)

	res, _ := ts.Client().Do(req)

	t.Log(res)
}
