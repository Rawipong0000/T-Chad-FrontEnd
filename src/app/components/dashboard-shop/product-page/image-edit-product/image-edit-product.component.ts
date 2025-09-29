import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-image-edit-product',
  templateUrl: './image-edit-product.component.html',
  styleUrls: ['./image-edit-product.component.css']
})
export class ImageEditProductComponent {
  @Output() closed = new EventEmitter<void>();
  
    URL: string = "";
  
    constructor(
  
    ) { }
  
    closePopup() {
      this.closed.emit();
    }
  
    @Output() urlSaved = new EventEmitter<string>(); // ✅ ส่ง string ออกไป
  
    SaveURL() {
      if (this.URL.trim() !== "") {
        this.urlSaved.emit(this.URL);   // ✅ ส่งค่าออกไป
        this.closePopup();              // ปิด popup หลังบันทึก
      } else {
        alert("Please enter a URL");
      }
    }
}
