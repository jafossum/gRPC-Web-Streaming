import { Component, OnInit } from '@angular/core';

import { grpc } from '@improbable-eng/grpc-web';
import { GrpcWebImpl, GreeterClientImpl } from './proto/greeter-service';
import { MatListModule } from '@angular/material/list';

@Component({
  selector: 'app-greeter-client',
  standalone: true,
  imports: [MatListModule],
  templateUrl: './greeter-client.component.html',
  styleUrl: './greeter-client.component.css',
})
export class GreeterClientComponent implements OnInit {
  repeatedResults: string[] = [];
  streamResults: string[] = [];

  ngOnInit(): void {
    // Create the gRPC Client
    const rpc = new GrpcWebImpl('http://127.0.0.1:8080', {
      debug: false,
      metadata: new grpc.Metadata(),
    });

    // Create the GreeterClient
    const client = new GreeterClientImpl(rpc);

    // Call Unary
    client
      .SayHello({ name: 'FINALLY!' })
      .then((value) => {
        console.log('Say hello response:' + value.message);
      })
      .catch((error) => {
        console.log(error);
      });

    // Call a stream endpoint
    client.SayRepeatedHello({ name: 'Repeat count 5', count: 5 }).subscribe({
      next: (value) => {
        console.log(
          'SayRepeatedHello emitted the next value: ' + value.message
        );
        this.repeatedResults.push(value.message);
      },
      error: (err) =>
        console.error('SayRepeatedHello emitted an error: ' + err),
      complete: () =>
        console.log('SayRepeatedHello emitted the complete notification'),
    });

    // Call a stream subscribe endpoint
    client.SubscribeRepeatedHello({ name: 'Subscribe stream' }).subscribe({
      next: (value) => {
        console.log(
          'SubscribeRepeatedHello emitted the next value: ' + value.message
        );
        this.streamResults.push(value.message);
      },
      error: (err) =>
        console.error('SubscribeRepeatedHello emitted an error: ' + err),
      complete: () =>
        console.log('SubscribeRepeatedHello emitted the complete notification'),
    });
  }
}
