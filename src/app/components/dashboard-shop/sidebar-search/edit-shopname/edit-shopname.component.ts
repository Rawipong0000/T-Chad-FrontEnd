import { Component, Output, EventEmitter } from '@angular/core';
import { MyShopService } from 'src/app/service/myshop.service';
import { SidebarSearchComponent } from '../sidebar-search.component';

@Component({
  selector: 'app-edit-shopname',
  templateUrl: './edit-shopname.component.html',
  styleUrls: ['./edit-shopname.component.css']
})
export class EditShopnameComponent {
  @Output() closed = new EventEmitter<void>();
  
    Shopname: string = "";
  
    constructor(
      private myShopService: MyShopService,
      private sidebarMyshop: SidebarSearchComponent
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
