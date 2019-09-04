import http from "k6/http";
import { check } from "k6";

export default function() {
	// Send a JSON encoded POST request

	let json_obj = {query_ty:{}, result_ty:{}, fnc:{}, params:{ id:{}, hashed_value:{}}};
	json_obj.query_ty="invoke";
	json_obj.result_ty="object";
	json_obj.fnc="insert";
	json_obj.params.id="test_id"
	json_obj.params.hashed_value="test_hashed_value";

	let body = JSON.stringify(json_obj);
	let res = http.post("http://211.253.31.18:18800/fabric1.3.0/epc090385610214/sc/arom_test2", body, { headers: { "apiKey": "RVBDR1JJRDE5MDMyODIxMDNmNHJwaTE3dHw0ZDJhNzVjNjAwNDc2ODZjY2IzNTc4MzM1MzhhYTExNWJmZTA4MmFlODA5ZjEyNGU0MjMzNDEyYzVmMjE3MWIx", "Content-Type":"application/json", "cache-control":"no-cacheRequest" }});

	// Use JSON.parse to deserialize the JSON (instead of using the r.json() method)
	let j = JSON.parse(res.body);

	// Verify response
	check(res, {
			"status is 200": (r) => r.status === 200,
			"is key correct": (r) => j.result === "true"
			});



	console.log(res);

}
