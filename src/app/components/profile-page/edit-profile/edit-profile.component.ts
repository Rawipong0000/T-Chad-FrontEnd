import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent {
  @Output() closed = new EventEmitter<void>();

  NewEdit: string = "";

  constructor(

  ) { }

  closePopup() {
    this.closed.emit();
  }

  @Output() editing = new EventEmitter<string>(); // ✅ ส่ง string ออกไป

  SaveEditing() {
    if (this.NewEdit.trim() !== "") {
      this.editing.emit(this.NewEdit);   // ✅ ส่งค่าออกไป
      this.closePopup();              // ปิด popup หลังบันทึก
    } else {
      alert("Please enter new data");
    }
  }
}
