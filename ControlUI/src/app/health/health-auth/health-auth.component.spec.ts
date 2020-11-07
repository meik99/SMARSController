import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HealthAuthComponent } from './health-auth.component';

describe('HealthAuthComponent', () => {
  let component: HealthAuthComponent;
  let fixture: ComponentFixture<HealthAuthComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HealthAuthComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HealthAuthComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
