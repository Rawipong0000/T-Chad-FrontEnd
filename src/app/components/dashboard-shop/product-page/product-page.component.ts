import { Component, OnInit } from '@angular/core';
import { Router, NavigationEnd } from '@angular/router';
import { ProductService } from 'src/app/service/product.service';
import { CartItem } from 'src/app/model/productlist.model';

@Component({
  selector: 'app-product-page',
  templateUrl: './product-page.component.html',
  styleUrls: ['./product-page.component.css']
})

export class ProductPageComponent implements OnInit {
  product: any = {};
  pageURL: string = "";
  showEdit: boolean = false;

  product_id: number = 0;
  img: string = "";
  product_name: string = "";
  description: string = "";
  selectedSize: string = "";
  price: number = 0;
  private cartItems: number[] = [];

  constructor(
    private router: Router,
    private productService: ProductService
  ) {
      this.router.events.subscribe(event => {
        if (event instanceof NavigationEnd) {
          this.pageURL = event.urlAfterRedirects;
          console.log('Current URL:', this.pageURL);
          this.loadDashboardProduct(); 
        }
      });
  }

  loadDashboardProduct() {
    this.product = JSON.parse(localStorage.getItem('product') || '{}')
    this.product_name = this.product.product_name
    this.product_id = this.product.product_id;
    this.img = this.product.img;
    this.description = this.product.description;
    this.selectedSize = this.product.size;
    this.price = this.product.price;
  }

  handleSavedURL(url: string) {
    this.img = url;
    console.log("Image URL saved:", this.img);
  }

  toggleSize(size: string) {
    this.selectedSize = size;
  }

  goToDashboard() {
    this.router.navigate(['/dashboard']);
  }

  ngOnInit(): void {



  }

  EditProduct() {
    const body = {
      product_name: this.product_name,
      price: this.price,
      description: this.description,
      size: this.selectedSize,
      img: this.img
    };

    if (this.product_name == "") {
      alert('Product name cannot be blank');
    }
    if (this.selectedSize == "") {
      alert('Please select size')
    }
    if (this.price == 0) {
      alert('Price cannot be 0')
    } else {
      this.productService.updateProduct(body, this.product_id).subscribe({
        next: (response) => {
          console.log('Update successful:', response);
          alert('Update successful!');
          this.goMyShopPage();
        },
        error: (err) => {
          console.error("update failed:", err.error);
          alert("update failed: " + err.error);
        }
      })
    }
  }

  SaveProduct() {
    const body = {
      product_name: this.product_name,
      price: this.price,
      description: this.description,
      size: this.selectedSize,
      img: this.img
    };

    if (this.product_name == "") {
      alert('Product name cannot be blank');
    }
    if (this.selectedSize == "") {
      alert('Please select size')
    }
    if (this.price == 0) {
      alert('Price cannot be 0')
    } else {
      this.productService.createProduct(body).subscribe({
        next: (response) => {
          console.log('Create successful:', response);
          alert('Create successful!');
          this.goMyShopPage();
        },
        error: (err) => {
          console.error("create failed:", err.error);
          alert("create failed: " + err.error);
        }
      })
    }
  }

  goMyShopPage() {
    this.router.navigate(['/MyShop/user']);
  }

  addToCart(item: CartItem) {
    if (localStorage.getItem('cart') !== null) {
      console.log('มี key นี้อยู่ใน localStorage');
      this.cartItems = JSON.parse(localStorage.getItem('cart') || '[]');
      this.cartItems.push(item.id);
    } else {
      this.cartItems.push(item.id);
      console.log('ไม่มี key นี้');
    }
    console.log(this.cartItems)
    localStorage.setItem("product", JSON.stringify(item));
    localStorage.setItem("cart", JSON.stringify(this.cartItems));
  }

  AddPurchase() {
    const item: CartItem = {
      id: this.product.product_id,
      name: this.product.product_name,
      price: this.product.price,
      size: this.product.size,
      seller: this.product.seller,
      img: this.product.img
    };
    this.addToCart(item);
    alert('Added to cart!');
    this.goToDashboard();
  }
}
