import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { GreeterClientComponent } from './greeter-client/greeter-client.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, GreeterClientComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'client-angular';
}
