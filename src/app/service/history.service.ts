import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class HistoryService {
    private baseUrl = 'http://localhost:8180/api/users';

    constructor(private http: HttpClient) { }

    getHistoryTransaction(): Observable<any> {
        return this.http.get<any>(`${this.baseUrl}/History`);
    }

    completeTransaction(SubTranID: number): Observable<any> {
        const body = { sub_tran_id: SubTranID }
        return this.http.put<any>(`${this.baseUrl}/History/update-complete`, body);
    }

    refundTransaction(SubTranID: number): Observable<any> {
        const body = { sub_tran_id: SubTranID }
        return this.http.put<any>(`${this.baseUrl}/History/update-refund`, body);
    }
}