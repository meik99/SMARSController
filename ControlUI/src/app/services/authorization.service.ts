import { Injectable } from '@angular/core';
import {environment} from "../../environments/environment";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class AuthorizationService {

  constructor(
    private _http: HttpClient
  ) { }

  isHealthy(): Promise<boolean> {
    return new Promise<boolean>((resolve, reject) => {
      this._http.get(environment.authApi + '/health')
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
