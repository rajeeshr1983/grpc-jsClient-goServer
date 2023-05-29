const {HelloRequest, RepeatHelloRequest} = require('./proto/helloworld_pb.js');
const {GreeterClient} = require('./proto/helloworld_grpc_web_pb.js');

var greeterService = new GreeterClient('http://localhost:8080');
var request = new HelloRequest();
request.setName('Rajeesh');
greeterService.sayHello(request, {}, function(err, response) {
  if(err != null)
  {
  console.log(response);
  console.log(response.getMessage())
  }
  else
  {
    console.log(err);
  }
});