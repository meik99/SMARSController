import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HealthAlarmComponent } from './health-alarm.component';

describe('HealthAlarmComponent', () => {
  let component: HealthAlarmComponent;
  let fixture: ComponentFixture<HealthAlarmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HealthAlarmComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HealthAlarmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
