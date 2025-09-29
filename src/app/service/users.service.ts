import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  private baseUrl = 'http://localhost:8180/api/users';

  constructor(private http: HttpClient) {}

  getUserByID(): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/id`);
  }

  updateUser(body: {
        Name: string,
        Last_name: string,
        phone: string,
        address: string,
        subdistrict: string,
        district: string,
        province: string,
        postal_code: string
    }): Observable<any> {
      return this.http.put<any>(`${this.baseUrl}/profile/update`, body);
  }

  getProvince(): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/profile/province`);
  }

  getDistrict(ProvinceID: number): Observable<any>{
    const body = { province_id: ProvinceID }
        return this.http.put<any>(`${this.baseUrl}/profile/district`, body);
  }

  getSubdistrict(DistrictID: number): Observable<any>{
    const body = { district_id: DistrictID }
        return this.http.put<any>(`${this.baseUrl}/profile/subdistrict`, body);
  }
}