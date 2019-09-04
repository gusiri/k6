import http from "k6/http";
import { check } from "k6";

export default function() {
    // Send a JSON encoded POST request
    let body = JSON.stringify({

	"userId":"645344",
	"userName":"유병현",
	"userRegnum":"590505-1910522",
	"verifyId":"123452",
	"version":"1.0",
	"type":"2",
	"fromDate":"20190101",
	"toDate":"20190904"
}
);
    let res = http.post("http://211.251.250.166/v1/api/mylist", body, { headers: { "Content-Type": "application/json", "Authorization":"Bearer dummy" }});

    // Use JSON.parse to deserialize the JSON (instead of using the r.json() method)
    let j = JSON.parse(res.body);

    // print response body for debugging
    console.log(res.body);

    // Verify response
    check(res, {
        //"status is 200": (r) => r.status === 200
	"retCode is ok": (res) => j.retCode === "ok"
    });

}
