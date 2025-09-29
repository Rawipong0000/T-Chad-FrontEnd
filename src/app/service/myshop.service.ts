import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class MyShopService {
    private baseUrl = 'http://localhost:8180/api/users';

    constructor(private http: HttpClient) { }

    getShopName(): Observable<any> {
        return this.http.get<any>(`${this.baseUrl}/MyShop`);
    }

    editShopName(Shopname: string): Observable<any> {
        return this.http.put<any>(`${this.baseUrl}/MyShop/Edit/Shopname`, { shopname: Shopname });
    }

    getMyShopAllProducts(): Observable<any> {
        return this.http.get<any>(`${this.baseUrl}/MyShop/products`);
    }

    getMyShopTransaction(): Observable<any> {
        return this.http.get<any>(`${this.baseUrl}/MyShop/order-manage`);
    }

    editTracking(body: {
        Tran_id: number,
        Tracking: string
    }): Observable<any> {
        return this.http.put<any>(`${this.baseUrl}/MyShop/order-manage/edit-tracking`, body);
    }

    approveRefund(SubTranID: number): Observable<any> {
        const body = { sub_tran_id: SubTranID }
        return this.http.put<any>(`${this.baseUrl}/History/update-refund-approve`, body);
    }

    rejectRefund(SubTranID: number): Observable<any> {
        const body = { sub_tran_id: SubTranID }
        return this.http.put<any>(`${this.baseUrl}/History/update-refund-reject`, body);
    }

    cancelTransaction(SubTranID: number): Observable<any> {
        const body = { sub_tran_id: SubTranID }
        return this.http.put<any>(`${this.baseUrl}/History/update-cancel`, body);
    }

    GetPromoCode(): Observable<any> {
        return this.http.get<any>(`${this.baseUrl}/MyShop/promocodes`);
    }

    GetPromoCodeByID(DiscountID: number): Observable<any> {
        const body = { discount_id: DiscountID }
        return this.http.put<any>(`${this.baseUrl}/MyShop/promocodes/get-by-id`,body);
    }

    CreatePromoCode(body: {
        discount_code: string,
        limit: number,
        discount_by_percent: number,
        discount_by_number: number,
        minimum_total: number,
        maximum_discount: number
    }): Observable<any> {
        return this.http.post<any>(`${this.baseUrl}/MyShop/promocodes/create`,body);
    }

    UpdatePromoCode(body: {
        discount_id: number,
        discount_code: string,
        discount_by_percent: number,
        discount_by_number: number,
        minimum_total: number,
        maximum_discount: number,
        limit: number
    }): Observable<any> {
        return this.http.put<any>(`${this.baseUrl}/MyShop/promocodes/update`,body);
    }

    DeactivatePromoCodeByID(DiscountID: number): Observable<any> {
        const body = { discount_id: DiscountID }
        return this.http.put<any>(`${this.baseUrl}/MyShop/promocodes/deactivate`,body);
    }
}