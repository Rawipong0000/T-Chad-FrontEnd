import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-edit-product-image',
  templateUrl: './edit-product-image.component.html',
  styleUrls: ['./edit-product-image.component.css']
})
export class EditProductImageComponent {
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
