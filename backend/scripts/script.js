import http from 'k6/http';
import { check } from 'k6';
export default function () {
  const params = {
    headers: {
      'Content-Type': 'application/json;charset=UTF-8',
      },
      cookies: {
        _cookie: 'faa0de164780d6079f58758657adcd4be851911c70bd1a65cad6b08dd4df361f',
      },
    };
  let res = http.get('http://localhost:8080/api/task?id=673&page=0', params);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  // console.log(res.json())
}
// docker exec -it kartenspielen-table go test ./interface/controllers/tests/... -count 50