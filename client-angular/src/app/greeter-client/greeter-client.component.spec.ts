import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GreeterClientComponent } from './greeter-client.component';

describe('GreeterClientComponent', () => {
  let component: GreeterClientComponent;
  let fixture: ComponentFixture<GreeterClientComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GreeterClientComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GreeterClientComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
