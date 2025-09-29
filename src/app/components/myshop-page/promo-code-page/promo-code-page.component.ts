import { Component, OnInit, Input } from '@angular/core';
import { MyShopService } from 'src/app/service/myshop.service';
import { Router } from '@angular/router';
import { DiscountCode } from 'src/app/model/transaction.model';

@Component({
  selector: 'app-promo-code-page',
  templateUrl: './promo-code-page.component.html',
  styleUrls: ['./promo-code-page.component.css']
})
export class PromoCodePageComponent implements OnInit {
  showEdit: boolean = false;
  ID: number = 0;
  promocodes: DiscountCode[] = [];

  constructor(
      private myShopService: MyShopService,
      private router: Router,
    ) { }

  ngOnInit(): void {
      this.myShopService.GetPromoCode().subscribe({
        next: (data) => {
          console.log('Fetched promocodes data:', data);
          this.promocodes = data
        },
        error: (err) => {
          console.error('Failed to fetch promocodes', err);
        }
      });
    }

    ShowLog() {
      console.log('Promo ID:', this.ID);
    }
}
