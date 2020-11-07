import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ControlPageComponent } from './control-page/control-page.component';
import { HealthComponent } from './health/health.component';
import { HealthAlarmComponent } from './health/health-alarm/health-alarm.component';
import {HttpClientModule} from "@angular/common/http";
import { HealthStatusComponent } from './health/health-status/health-status.component';
import { HealthAuthComponent } from './health/health-auth/health-auth.component';

@NgModule({
  declarations: [
    AppComponent,
    ControlPageComponent,
    HealthComponent,
    HealthAlarmComponent,
    HealthStatusComponent,
    HealthAuthComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
