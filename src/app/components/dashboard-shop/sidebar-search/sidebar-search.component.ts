import { Component, Output, EventEmitter, OnInit } from '@angular/core';
import { MyShopService } from 'src/app/service/myshop.service';
import { Users } from 'src/app/model/users.model';
import { Router, NavigationEnd } from '@angular/router';
import { faMagnifyingGlass, faUser } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-sidebar-search',
  templateUrl: './sidebar-search.component.html',
  styleUrls: ['./sidebar-search.component.css']
})
export class SidebarSearchComponent implements OnInit {
  faMagnifyingGlass = faMagnifyingGlass;
  faUser = faUser;
  searchText: string = '';
  priceMin: number | null = null;
  priceMax: number | null = null;
  selectedSizes: string[] = [];
  pageURL: string = "";

  usersInfo: Users | null = null;
  showEdit: boolean = false;

  constructor(
    private myShopService: MyShopService,
    private router: Router,
  ) {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.pageURL = event.urlAfterRedirects;
      }
    });
  }

  @Output() filterChanged = new EventEmitter<any>();

  ngOnInit(): void {
    this.myShopService.getShopName().subscribe({
      next: (users) => {
        console.log('Fetched users data:', users);
        this.usersInfo = users;
      },
      error: (err) => {
        console.error('Failed to fetch users', err);
      }
    })
  }

  toggleSize(size: string) {
    if (this.selectedSizes.includes(size)) {
      this.selectedSizes = this.selectedSizes.filter(s => s !== size);
    } else {
      this.selectedSizes.push(size);
    }
  }

  applyFilter() {
    this.filterChanged.emit({
      searchText: this.searchText,
      priceMin: this.priceMin,
      priceMax: this.priceMax,
      selectedSizes: this.selectedSizes
    });
  }

  goToOrderManage(){
    this.router.navigate(['/MyShop/user/orders-manage-page']);
  }

  goToPromoCode(){
    this.router.navigate(['/MyShop/user/promo-code-page']);
  }
}
