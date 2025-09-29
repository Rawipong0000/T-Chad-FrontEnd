import { Component, Output, EventEmitter, OnInit } from '@angular/core';
import { Users } from 'src/app/model/users.model';
import { MyShopService } from 'src/app/service/myshop.service';

@Component({
  selector: 'app-sidebar-myshop',
  templateUrl: './sidebar-myshop.component.html',
  styleUrls: ['./sidebar-myshop.component.css']
})
export class SidebarMyshopComponent implements OnInit {
  usersInfo: Users | null = null;
  showEdit: boolean = false;

  searchText: string = '';
  priceMin: number | null = null;
  priceMax: number | null = null;
  selectedSizes: string[] = [];

  constructor(
    private myShopService: MyShopService,
  ) { }

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
}
