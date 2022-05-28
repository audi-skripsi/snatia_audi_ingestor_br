import http from "k6/http";

export const options = {
  vus: 10,
  iterations: 10000,
};

export default function () {
  const mode = __ENV.MODE;
  let reqPath = "";
  if (mode == "microbatch") {
    reqPath = "/v1/microbatch";
  } else if (mode == "non-microbatch") {
    reqPath = "/v1/non-microbatch";
  } else {
    console.error("ENV should be microbatch or non-microbatch");
    return;
  }

  http.post("http://localhost:8080" + reqPath, {
    level: "info",
    app_name: "test-appname",
    message: "test-message",
    timestamp: 1653307474,
  });
}
