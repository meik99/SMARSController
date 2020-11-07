import { Component, OnInit } from '@angular/core';
import {AlarmService} from "../../services/alarm.service";

@Component({
  selector: 'app-health-alarm',
  templateUrl: './health-alarm.component.html',
  styleUrls: ['./health-alarm.component.scss']
})
export class HealthAlarmComponent implements OnInit {
  isHealthy = false;

  private interval = null;

  constructor(
    private _alarmService: AlarmService
  ) { }

  ngOnInit(): void {
    this.checkHealthStatus();
    this.interval = setInterval(() => {
       this.checkHealthStatus();
      }, 5000);
  }

  checkHealthStatus() {
    this.
    _alarmService.
    isHealthy().
    then(result => this.isHealthy = result).
    catch(err => this.isHealthy = err);
  }
}
