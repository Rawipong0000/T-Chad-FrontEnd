// src/app/component/register/register-pop-up.component.ts
import { Component, Output, EventEmitter } from '@angular/core';
import { AuthService } from 'src/app/service/auth.service';

@Component({
  selector: 'app-signup',
  templateUrl: './register-pop-up.component.html',
  styleUrls: ['./register-pop-up.component.css']
})
export class RegisterPopUpComponent {
  @Output() closed = new EventEmitter<void>();
  showFields: boolean = false;
  showValidate: boolean = true;

  email: string = "";
  Name: string = "";
  Surname: string = "";
  Password: string = "";
  Repassword: string = "";

  constructor(private authService: AuthService) {}

  closePopup() {
    this.closed.emit();
  }

  validateEmail() {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!emailPattern.test(this.email)) {
      alert("Invalid email format");
      return;
    }

    this.authService.checkEmail(this.email).subscribe({
      next: () => {
        console.log("Email is available");
        alert("Email is available");
        this.showValidate = false;
        this.showFields = true;
      },
      error: (err) => {
        if (err.status === 400) {
          console.log("Duplicate Email:", err);
          alert("Duplicate Email");
        } else {
          console.error("Unexpected error:", err);
          alert("Server error");
        }
      }
    });
  }

  Sumit_Register() {
    const passwordPattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[^A-Za-z0-9]).{6,}$/;

    if (!passwordPattern.test(this.Password)) {
      alert("Password must contain uppercase, lowercase, and special character.");
      return;
    }

    if (this.Password !== this.Repassword) {
      alert("Re-password not same as password");
      return;
    }

    const body = {
      email: this.email,
      Name: this.Name,
      Surname: this.Surname,
      Password: this.Password
    };

    this.authService.register(body).subscribe({
      next: (response) => {
        console.log('create successful:', response);
        alert(response.message);
        this.closePopup();
      },
      error: (err) => {
        console.error("create failed:", err.error);
        alert("create failed: " + err.error);
      }
    });
  }
}
