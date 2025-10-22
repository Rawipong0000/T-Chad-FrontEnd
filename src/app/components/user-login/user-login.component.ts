import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/service/auth.service';

@Component({
  selector: 'app-user-login',
  templateUrl: './user-login.component.html',
  styleUrls: ['./user-login.component.css']
})
export class UserLoginComponent {
  username: string = '';
  password: string = '';
  showSignup: boolean = false;

  constructor(
    private authService: AuthService,
    private router: Router
  ) { }

  goToDashboard() {
    this.router.navigate(['/dashboard']);
  }

  onLogin() {

    this.authService.login(this.username, this.password).subscribe({
      next: (response) => {
        console.log('Login success:', response);
        localStorage.setItem("token", response);
        this.goToDashboard();
      },
      error: (err) => {
        console.error('Login failed:', err.error);
        alert('Invalid password: ' + err.error);
      }
    });
  }
}
