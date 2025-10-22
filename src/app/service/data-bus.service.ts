import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";

@Injectable({ providedIn: 'root' })
export class DataBusService {
  private productID$ = new BehaviorSubject<number | null>(null);
  selectedProductID$ = this.productID$.asObservable();

  setProductID(p: number | null) { this.productID$.next(p); }
}
