import { Component, Output, EventEmitter } from '@angular/core';
import { MyShopService } from 'src/app/service/myshop.service';
import { SidebarMyshopComponent } from '../sidebar-myshop.component';

@Component({
  selector: 'app-shopname-edit',
  templateUrl: './shopname-edit.component.html',
  styleUrls: ['./shopname-edit.component.css']
})
export class ShopnameEditComponent {
  @Output() closed = new EventEmitter<void>();

  Shopname: string = "";

  constructor(
    private myShopService: MyShopService,
    private sidebarMyshop: SidebarMyshopComponent
  ) { }

  closePopup() {
    this.closed.emit();
  }

  EditShopName() {
    const shop = {
      shopname: this.Shopname 
    };

    this.myShopService.editShopName(shop.shopname).subscribe({
      next: (response) => {
        console.log('edit successful:', response);
        alert(response.message);
        this.closePopup();
        this.sidebarMyshop.ngOnInit();
      },
      error: (err) => {
        console.error("edit failed:", err.error);
        alert("edit failed: " + err.error);
      }
    })
  }
}
