import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { faUser,faHouse } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  constructor(
    private router: Router,
  ) { }
  faUser = faUser;
  faHouse = faHouse;

  showMenu: boolean = false;

  toggleMenu() {
    this.showMenu = !this.showMenu;
  }

  goToDashboard() {
    this.router.navigate(['/dashboard']);
  }

  goToProfile() {
    this.router.navigate(['/profile']);
  }

  goToHistory() {
    this.router.navigate(['/history']);
  }

  goPurchasePage(){
    this.router.navigate(['/purchase-list/user']);
  }

  goMyShopPage(){
    this.router.navigate(['/MyShop/user']);
  }

  Logout() {
    localStorage.removeItem("cart");
    localStorage.removeItem("token");
    this.router.navigate(['/login']);
  }
}
