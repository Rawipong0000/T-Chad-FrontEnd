import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private baseUrl = 'http://localhost:8180/api/users';

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<any> {
    const body = { email, password };
    return this.http.post<any>(`${this.baseUrl}/email`, body);
  }

  checkEmail(email: string): Observable<any> {
    return this.http.post(`${this.baseUrl}/check-email`, { email });
  }

  register(data: {
    email: string;
    Name: string;
    Surname: string;
    Password: string;
  }): Observable<{ message: string }> {
    return this.http.post<{ message: string }>(this.baseUrl, data);
  }
}