using Greeter.API;
using Grpc.Core;
using grpcWeb;
using NATS.Client.Core;

namespace grpcWeb.Services;

public class GreeterService : Greeter.API.Greeter.GreeterBase
{
    private readonly ILogger<GreeterService> _logger;
    public GreeterService(ILogger<GreeterService> logger)
    {
        _logger = logger;
    }

    public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
    {
        _logger.LogInformation("Say Hello receive");
        return Task.FromResult(new HelloReply
        {
            Message = "Hello " + request.Name
        });
    }
    
    public override async Task SayRepeatedHello(RepeatHelloRequest request,
        IServerStreamWriter<HelloReply> responseStream, ServerCallContext context)
    {
        _logger.LogInformation("Say Repeated Hello received");
        for (var i = 0; i < request.Count; i++)
        {
            await responseStream.WriteAsync(new HelloReply
            {
                Message = $"Name: {request.Name}, Count: {i+1}"
            });
            await Task.Delay(TimeSpan.FromSeconds(1));
        }
    }

    public override async Task SubscribeRepeatedHello(SubscribeHelloRequest request,
        IServerStreamWriter<HelloReply> responseStream, ServerCallContext context)
    {
        _logger.LogInformation("Subscribe Repeated Hello received");

        var ctx = context.CancellationToken;

        // NATS core M:N messaging example
        await using var nats = new NatsConnection();

        // Subscribe on one terminal
        await Task.Run(async () =>
        {
            await foreach (var msg in nats.SubscribeAsync<string>(subject: "foo.bar.>").WithCancellation(ctx))
            {
                await responseStream.WriteAsync(new HelloReply
                {
                    Message = $"NATS -{msg.Subject}- Stream Received: {msg.Data}"
                }, ctx);
            }
        }, ctx);
    }
}
