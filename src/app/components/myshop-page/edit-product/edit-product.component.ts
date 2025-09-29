import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductService } from 'src/app/service/product.service';

@Component({
  selector: 'app-edit-product',
  templateUrl: './edit-product.component.html',
  styleUrls: ['./edit-product.component.css']
})
export class EditProductComponent implements OnInit {
  product: any = {};
  showEdit: boolean = false;

  product_id: number = 0;
  img: string = "";
  product_name: string = "";
  description: string = "";
  selectedSize: string = "";
  price: number = 0;

  constructor(
    private router: Router,
    private productService: ProductService
  ) { }

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
  this.product = JSON.parse(localStorage.getItem('my_product') || '{}');
  this.product_id = this.product.product_id;
  this.img = this.product.img;
  this.product_name = this.product.product_name;
  this.description = this.product.description;
  this.selectedSize = this.product.size;
  this.price = this.product.price;
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
      this.productService.updateProduct(body,this.product_id).subscribe({
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

  goMyShopPage() {
    this.router.navigate(['/MyShop/user']);
  }
}
