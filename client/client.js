const {
  HelloRequest,
  HelloReply,
  RepeatHelloRequest,
} = require("./greeter-service_pb.js");
const { GreeterClient } = require("./greeter-service_grpc_web_pb.js");

var client = new GreeterClient("http://localhost:8080");

// simple unary call
var request = new HelloRequest();
request.setName("UNARY Request Visitor");

client.sayHello(request, {}, (err, response) => {
  if (err) {
    console.log(
      `Unexpected error for sayHello: code = ${err.code}` +
        `, message = "${err.message}"`
    );
  } else {
    console.log("Say hello called with name: " + response.getMessage());
  }
});

// server streaming call - Stream 1
var streamRequest1 = new RepeatHelloRequest();
streamRequest1.setName("Stream Request Visitor - 1");
streamRequest1.setCount(5);

console.log("Calling stream server expecting 5 returns");
subscribeToStream(client.sayRepeatedHello(streamRequest1, {}));

// server streaming call - Stream 2
var streamRequest2 = new RepeatHelloRequest();
streamRequest2.setName("Stream Request Visitor - 2");
streamRequest2.setCount(7);

console.log("Calling stream server expecting 7 returns");
subscribeToStream(client.sayRepeatedHello(streamRequest2, {}));

function subscribeToStream(stream) {
  stream.on("data", (response) => {
    console.log("Stream data received: " + response.getMessage());
  });
  stream.on("error", (err) => {
    console.log(
      `Unexpected stream error: code = ${err.code}` +
        `, message = "${err.message}"`
    );
  });
  stream.on("end", (end) => {
    console.log("stream ended");
  });
}
