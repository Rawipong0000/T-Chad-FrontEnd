import { Component, OnInit } from '@angular/core';
import { ProductService } from '../../service/product.service';
import { MyShopService } from 'src/app/service/myshop.service';
import { Productlist } from '../../model/productlist.model';
import { Router, NavigationEnd } from '@angular/router';

@Component({
  selector: 'app-dashboard-shop',
  templateUrl: './dashboard-shop.component.html',
  styleUrls: ['./dashboard-shop.component.css']
})
export class DashboardShopComponent implements OnInit {
  products: Productlist[] = [];
  filteredProducts: Productlist[] = [];
  dataProducts: Productlist[] = [];
  pageURL: string = "";
  pageResource: number = 0;

  constructor(
    private productService: ProductService,
    private myShopService: MyShopService,
    private router: Router,
  ) {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.pageURL = event.urlAfterRedirects;
        console.log('Current URL:', this.pageURL);
      }
    });
  }

  ngOnInit(): void {
    if (this.pageURL == "/MyShop/user") {
      this.myShopService.getMyShopAllProducts().subscribe({
        next: (data) => {
          console.log('Fetched data:', data);
          this.dataProducts = data;
          this.applyCartFilter(); // กรองสินค้าที่ไม่อยู่ใน cart
        },
        error: (err) => {
          console.error('Failed to fetch products', err);
        }
      });
    } else {
      this.productService.getAllProducts().subscribe({
        next: (data) => {
          console.log('Fetched data:', data);
          this.dataProducts = data;
          this.applyCartFilter(); // กรองสินค้าที่ไม่อยู่ใน cart
        },
        error: (err) => {
          console.error('Failed to fetch products', err);
        }
      });
    }
  }

  applyCartFilter() {
    const cartIds = this.getCartProductIds()
    console.log("cartIds = ", cartIds)
    this.products = this.dataProducts.filter(p => !cartIds.includes(p.product_id));
    console.log("this product = ", this.products)
  }

  goProductPage(product: Productlist) {
    localStorage.setItem("product", JSON.stringify(product));
    this.router.navigate(['/dashboard/product', product.product_id]);
  }

  getCartProductIds(): number[] {
    const InCartID = JSON.parse(localStorage.getItem('cart') || '[]');
    console.log(InCartID);
    return JSON.parse(localStorage.getItem('cart') || '[]');
  }

  onFilterChanged(filter: any) {
    const cartIds = this.getCartProductIds(); // ใช้ร่วมกับ filter
    this.products = this.dataProducts.filter(product => {
      const matchSearch = !filter.searchText || product.product_name.toLowerCase().includes(filter.searchText.toLowerCase());
      const matchPriceMin = filter.priceMin == null || product.price >= filter.priceMin;
      const matchPriceMax = filter.priceMax == null || product.price <= filter.priceMax;
      const matchSize = !filter.selectedSizes.length || filter.selectedSizes.includes(product.size);
      const notInCart = !cartIds.includes(product.product_id);
      return matchSearch && matchPriceMin && matchPriceMax && matchSize && notInCart;
    });
  }

  editProduct(product: Productlist) {
    localStorage.setItem("product", JSON.stringify(product));
    this.router.navigate(['/MyShop/user/edit-product', product.product_id]);
  }

  addProduct() {
    localStorage.setItem("product", JSON.stringify({}));
    this.router.navigate(['/MyShop/user/add-product']);
  }
}

