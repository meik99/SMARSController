import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-health-status',
  templateUrl: './health-status.component.html',
  styleUrls: ['./health-status.component.scss']
})
export class HealthStatusComponent implements OnInit {
  @Input()
  isHealthy = false;

  constructor() { }

  ngOnInit(): void {
  }

}
