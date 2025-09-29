import { Component, Output, Input, EventEmitter, OnInit} from '@angular/core';
import { Router } from '@angular/router';
import { MyShopService } from 'src/app/service/myshop.service';
import { DiscountCode } from 'src/app/model/transaction.model';
import { PromoCodePageComponent } from '../promo-code-page.component';

@Component({
  selector: 'app-add-promo-code-page',
  templateUrl: './add-promo-code-page.component.html',
  styleUrls: ['./add-promo-code-page.component.css']
})
export class AddPromoCodePageComponent implements OnInit{
  @Output() closed = new EventEmitter<void>();
  @Input() DiscountID: number = 0;

  Discount_ID: number = 0;

  promocode: DiscountCode = {};

  Discount_Code: string = "";
  Discount_Percent: number = 0;
  Discount_Fix: number = 0;
  Discount_Minimum: number = 0;
  Discount_Maximum: number = 999999;
  Discount_Limit: number = 999999;
  Delflag: Boolean = true;

  constructor(
    private myShopService: MyShopService,
    private promocodepage: PromoCodePageComponent,
    private router: Router,
  ) { }

  closePopup() {
    this.closed.emit();
  }

  ngOnInit(): void {
    this.Discount_ID = this.DiscountID;
    console.log('Discount_ID:', this.Discount_ID);
    if (this.Discount_ID != 0) {
      this.myShopService.GetPromoCodeByID(this.Discount_ID).subscribe({
        next: (data) => {
          console.log('Fetched promocode data:', data);
          this.promocode = data;
          this.LoadData();
          console.log('Code', this.Discount_Code);
        },
        error: (err) => {
          console.error('Failed to fetch promocode', err);
        }
      });
    }
  }

  LoadData() {
  this.Discount_Code = this.promocode.discount_code ?? "";
  this.Discount_Percent = this.promocode.discount_by_percent ?? 0;
  this.Discount_Fix = this.promocode.discount_by_number ?? 0;
  this.Discount_Minimum = this.promocode.minimum_total ?? 0;
  this.Discount_Maximum = this.promocode.maximum_discount ?? 999999;
  this.Discount_Limit = this.promocode.limit ?? 999999;
  this.Delflag = this.promocode.delflag ?? true;
  }

  SaveEditing() {
    if (this.Discount_Code == "") {
      alert("promo code cannot be blank")
    } else {
      const body = {
        discount_id: this.Discount_ID,
        discount_code: this.Discount_Code,
        discount_by_percent: this.Discount_Percent,
        discount_by_number: this.Discount_Fix,
        minimum_total: this.Discount_Minimum,
        maximum_discount: this.Discount_Maximum,
        limit: this.Discount_Limit
      };
      this.myShopService.UpdatePromoCode(body).subscribe({
        next: (response) => {
          alert("Update complete")
          console.log('Update promocodes:', response);
          this.closePopup();
          this.promocodepage.ngOnInit();
        },
        error: (err) => {
          alert("Update fail")
          console.error('Failed to update promocodes', err);
        }
      });
    }
  }

  CreatePromo() {
    if (this.Discount_Code == "") {
      alert("promo code cannot be blank")
    } else {
      const body = {
        discount_code: this.Discount_Code,
        discount_by_percent: this.Discount_Percent,
        discount_by_number: this.Discount_Fix,
        minimum_total: this.Discount_Minimum,
        maximum_discount: this.Discount_Maximum,
        limit: this.Discount_Limit
      };
      this.myShopService.CreatePromoCode(body).subscribe({
        next: (response) => {
          alert("Create complete")
          console.log('Create promocodes:', response);
          this.closePopup();
          this.promocodepage.ngOnInit();
        },
        error: (err) => {
          alert("Create fail")
          console.error('Failed to create promocodes', err);
        }
      });
    }
  }

  DeactivatePromoCode() {
    this.myShopService.DeactivatePromoCodeByID(this.Discount_ID).subscribe({
        next: (response) => {
          alert("Update complete")
          console.log('Update promocodes:', response);
          this.closePopup();
          this.promocodepage.ngOnInit();
        },
        error: (err) => {
          alert("Update fail")
          console.error('Failed to update promocodes', err);
        }
      });
  }
}
