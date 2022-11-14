import http from 'k6/http';

export default function () {
    http.post('http://god-jay-micro-api-service:8080/user', {"id": 1});
}