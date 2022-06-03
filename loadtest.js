import http from "k6/http";

export const options = {
  vus: 10,
  iterations: 10000,
};

const rawData = open("data-sample.json");
const testData = JSON.parse(rawData);

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

  const randInt = Math.floor(Math.random() * testData.length);

  http.post(
    "http://localhost:8080" + reqPath,
    JSON.stringify(testData[randInt]),
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
}
