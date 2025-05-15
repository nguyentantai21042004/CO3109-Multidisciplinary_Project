import { Inject, Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { HttpClient } from "@angular/common/http";
import { inject } from "@angular/core"; 
import { DOCUMENT } from "@angular/common"; 
import { HttpUtilService } from "./http.util.service";
import { Observable } from "rxjs";
import { RegisterParam } from "../params/register.param"; 

@Injectable({
  providedIn: 'root',
})
export class UserService {

  private http = inject(HttpClient);
  private httpUtilService = inject(HttpUtilService);

  localStorage?: Storage;

  private apiConfig = {
    headers: this.httpUtilService.createHeaders(),
  };

  private apiBaseUrl = environment.apiBaseUrl;

  constructor(@Inject(DOCUMENT) private document: Document) {
    this.localStorage = document.defaultView?.localStorage;
  }

  private apiRegister = `${this.apiBaseUrl}/auth/register`;
  register(registerParam: RegisterParam): Observable<any> {
    return this.http.post(this.apiRegister, registerParam, this.apiConfig);
  }

  private apiSendOtp = `${this.apiBaseUrl}/auth/send-otp`;
  sendOtp(email: string, password: string): Observable<any> {
    const payload = { email, password };
    return this.http.post(this.apiSendOtp, payload, this.apiConfig);
  }

  private apiVerifyOtp = `${this.apiBaseUrl}/auth/verify-otp`;
  verifyOtp(email: string, otp: string): Observable<any> {
    const payload = { email, otp };
    return this.http.post(this.apiVerifyOtp, payload, this.apiConfig);
  }

  private apiLogin = `${this.apiBaseUrl}/auth/login`;
  login(email: string, password: string): Observable<any> {
    const payload = { email, password };
    return this.http.post(this.apiLogin, payload, this.apiConfig);
  }
}