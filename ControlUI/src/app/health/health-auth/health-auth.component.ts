import { Component, OnInit } from '@angular/core';
import {AuthorizationService} from "../../services/authorization.service";

@Component({
  selector: 'app-health-auth',
  templateUrl: './health-auth.component.html',
  styleUrls: ['./health-auth.component.scss']
})
export class HealthAuthComponent implements OnInit {
  isHealthy = false;

  private interval = null;

  constructor(
    private _authService: AuthorizationService
  ) { }

  ngOnInit(): void {
    this.checkHealthStatus();
    this.interval = setInterval(() => {
      this.checkHealthStatus();
    }, 5000);
  }

  checkHealthStatus() {
    this.
      _authService.
      isHealthy().
      then(result => this.isHealthy = result).
      catch(err => this.isHealthy = err);
  }
}
