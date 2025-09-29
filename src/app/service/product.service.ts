// product.service.ts
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, BehaviorSubject } from 'rxjs';
import { Productlist, CartItem, GetCartItem } from '../model/productlist.model';

@Injectable({ providedIn: 'root' })
export class ProductService {
  constructor(private http: HttpClient,) { }
  private cartSubject = new BehaviorSubject<CartItem[]>([]);

  cart$ = this.cartSubject.asObservable();

  getAllProducts(): Observable<Productlist[]> {
    return this.http.get<Productlist[]>('http://localhost:8180/api/products');
  }

  getCartProducts(cartIDs: number[]): Observable<GetCartItem[]> {
    return this.http.post<GetCartItem[]>('http://localhost:8180/api/cart/products', cartIDs);
  }

  createTransaction(body: {
    transaction: any,
    sub_transaction: any[],
    purchasing: any[],
    cartIds: number[]
  }): Observable<any> {
    return this.http.post<any>('http://localhost:8180/api/transaction', body);
  }

  redeemCode(body: {
    discount_code: string,
    seller_id: number,
    total: number
  }): Observable<any> {
    return this.http.post<any>('http://localhost:8180/api/redeem-code', body);
  }

  createProduct(body: {
    product_name: string,
    price: number,
    description: string,
    size: string,
    img: string
  }): Observable<any> {
    return this.http.post<any>('http://localhost:8180/api/add-product', body);
  }

  updateProduct(body: {
    product_name: string,
    price: number,
    description: string,
    size: string,
    img: string
  }, product_id: number): Observable<any> {
    const url = `http://localhost:8180/api/products/${product_id}`;
    return this.http.put<any>(url, body);
  }
}
