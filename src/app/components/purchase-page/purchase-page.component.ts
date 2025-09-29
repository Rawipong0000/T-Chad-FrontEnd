import { Component, OnInit } from '@angular/core';
import { ProductService } from 'src/app/service/product.service';
import { Router } from '@angular/router';
import { GetCartItem } from 'src/app/model/productlist.model';
import { DiscountCode, Purchasing, Transaction } from 'src/app/model/transaction.model';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-purchase-page',
  templateUrl: './purchase-page.component.html',
  styleUrls: ['./purchase-page.component.css']
})
export class PurchasePageComponent implements OnInit {
  discount_code: string = "";
  getCart: GetCartItem[] = [];
  groupedCart: Record<string, { items: GetCartItem[]; name: string, shopname: string, discount_id: number, dis_percent: number, dis_number: number, maxDis: number, total: number }> = {};
  discountCodes: { [sellerID: string]: string } = {};
  discountDetail: DiscountCode[] = [];
  cartID: number[] = [];
  transaction: Transaction[] = [];
  purchasing: Purchasing[] = [];
  private cartSubject = new BehaviorSubject<GetCartItem[]>([]);

  constructor(
    private productService: ProductService,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.cartID = (JSON.parse(localStorage.getItem('cart') || '[]') as unknown[]).map(id => Number(id));

    if (this.cartID.length > 0) {
      this.productService.getCartProducts(this.cartID).subscribe({
        next: (data) => {
          console.log('Fetched Cart data:', data);
          this.getCart = data;
          this.groupCartBySeller();
          console.log('GrupCart: ', this.groupedCart)
        },
        error: (err) => {
          console.error('Failed to fetch products', err);
        }
      });
    }
  }

  groupCartBySeller() {
    this.groupedCart = this.getCart.reduce((acc: Record<string, {
      items: GetCartItem[],
      name: string,
      shopname: string,
      discount_id: number,
      dis_percent: number,
      dis_number: number,
      maxDis: number,
      total: number
    }>, item) => {
      if (!acc[item.product_user_id]) {
        acc[item.product_user_id] = {
          items: [],
          name: "",
          shopname: "",
          discount_id: 0,
          dis_percent: 0,
          dis_number: 0,
          maxDis: 999999,
          total: 0
        };
      }
      acc[item.product_user_id].items.push(item);
      acc[item.product_user_id].name = item.name;
      acc[item.product_user_id].shopname = item.shopname;
      acc[item.product_user_id].total += item.price;

      const discount = this.discountDetail.find(
        d => d.seller_id === item.product_user_id
      );
      console.log('Discount:', discount);

      if (discount) {
        acc[item.product_user_id].discount_id = discount.discount_id || 0;
        acc[item.product_user_id].dis_percent = discount.discount_by_percent || 0;
        acc[item.product_user_id].dis_number = discount.discount_by_number || 0;
        acc[item.product_user_id].maxDis = discount.maximum_discount || 999999;

        const total = acc[item.product_user_id].total;
        const percent = acc[item.product_user_id].dis_percent;
        const fix = acc[item.product_user_id].dis_number;
        const maxDis = acc[item.product_user_id].maxDis;
        const calDis = this.calculateDiscount(total, percent, fix, maxDis);
        console.log('Discount:', calDis);

        acc[item.product_user_id].total = total - calDis;
      }

      return acc;
    }, {} as typeof this.groupedCart);
  }

  removeItem(itemToRemove: GetCartItem): void {
    this.removeFromCart(itemToRemove.product_id); // ใช้ id ใน CartItem
    this.getCart = this.getCartItems(); // รีเฟรชรายการใหม่
  }

  purchase(): void {
    const transaction = {
      Total: this.getTotalPrice(),
    };

    const sub_transaction = Object.entries(this.groupedCart).map(([sellerId, group]) => {
      return {
        seller_id: parseInt(sellerId, 10),
        discount_code: this.discountCodes[sellerId] || '',
        sub_total: group.total
      };
    });
    console.log(sub_transaction);

    const purchasing = this.getCart.map(item => {
      const discountCode = this.discountCodes[item.product_user_id] || '';

      return {
        Tran_id: 0,
        Product_ID: item.product_id,
        Tracking: "",
        Discount_Code: discountCode
      };
    });

    const code: string[] = Object.values(this.discountCodes);

    const body = {
      transaction,
      sub_transaction,
      purchasing,
      cartIds: this.cartID,
      discountcode: code
    };

    this.productService.createTransaction(body).subscribe({
      next: (response) => {
        console.log('Purchase successful:', response);
        alert('Purchase successful!');
        this.clearCart();
        this.getCart = [];
        localStorage.removeItem("cart");
        this.goToDashboard();
      },
      error: (err) => {
        const errorMsg = (err.error?.error || '').toLowerCase();
        switch (errorMsg) {
            case "address blank":
              alert("Address blank");
              break;
            case "sold":
              alert("Some products have alredy sold");
              break;
            case "delete":
              alert("Some products have alredy deleted");
              break;
            default:
              alert("Unknown error: " + errorMsg);
          }
        console.error("purchase failed:", err.error);
        alert("purchase failed: " + err.error);
      }
    })
  }

  getCartProductIds(): number[] {
    return this.getCart.map(item => item.product_id);
  }

  removeFromCart(itemId: number) {
    this.getCart = this.getCart.filter(item => item.product_id !== itemId);
    this.cartID = this.cartID.filter(item => item !== itemId)
    this.cartSubject.next(this.getCart);
    localStorage.setItem("cart", JSON.stringify(this.cartID));
    this.groupCartBySeller();
  }

  getCartItems(): GetCartItem[] {
    return [...this.getCart];
  }

  Redeem(sellerID: string) {
    const code = this.discountCodes[sellerID] || '';
    const sellerTotal = this.groupedCart[sellerID]?.total || 0;

    const body = {
      discount_code: code,
      seller_id: +sellerID,
      total: sellerTotal
    };
    if (code == "") {
      alert('Discount code should not be blank');
    } else {
      this.productService.redeemCode(body).subscribe({
        next: (response) => {
          console.log('Redeem successful:', response);
          const discount: DiscountCode = Array.isArray(response) ? response[0] : response;
          const index = this.discountDetail.findIndex(d => d.seller_id === +sellerID);
          if (index >= 0) {
            this.discountDetail[index] = discount;
          } else {
            this.discountDetail.push(discount);
          }
          this.groupCartBySeller();
          alert('Redeem successful!');
        },
        error: (err) => {
          const errorMsg = (err.error?.error || '').toLowerCase();
          console.error("redeem failed:", err.error);
          switch (errorMsg) {
            case "invalid code":
              alert("Invalid Code");
              break;
            case "invalid seller":
              alert("Invalid Seller");
              break;
            case "less than minimum":
              alert("Less than Minimum");
              break;
            case "exceed limit used":
              alert("Exceed Limit Used");
              break;
            default:
              alert("Unknown error: " + errorMsg);
          }
        }
      })
    }
  }

  calculateDiscount(Total: number, Percent: number, Fix: number, MaxDis: number) {
    const discount = (Total * (Percent / 100)) + Fix;
    if (discount > MaxDis) {
      return MaxDis;
    } else {
      return discount;
    }
  }

  clearCart() {
    this.getCart = [];
    this.cartSubject.next(this.getCart);
  }

  getTotalPrice(): number {
    return Object.values(this.groupedCart).reduce(
      (sum, seller) => sum + seller.total,
      0
    );
  }

  goToDashboard() {
    this.router.navigate(['/dashboard']);
  }
}
