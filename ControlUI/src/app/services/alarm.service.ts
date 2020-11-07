import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class AlarmService {

  constructor(
    private _http: HttpClient
  ) { }

  isHealthy(): Promise<boolean> {
    return new Promise<boolean>((resolve, reject) => {
      this._http.get(environment.alarmApi + '/health')
        .subscribe(
          (value: any) => {
            console.log(value);
            resolve(value.status === 200);
          },
          (err) => {
            console.log(err);
            resolve(false);
          }
        );
    });
  }
}
