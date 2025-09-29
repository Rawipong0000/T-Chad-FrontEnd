import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { UsersService } from 'src/app/service/users.service';
import { Users } from 'src/app/model/users.model';
import { faUser } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent implements OnInit {
  faUser = faUser;
  profile!: Users;
  showEdit: boolean = false;
  showEditAddress: boolean = false;
  editingBox: Number = 0;
  img: string = "";
  currentStep: number = 1;

  FullAddress: string = "";

  Name: string = "";
  Lastname: string = "";
  Phone: string = "";
  Address: string = "";
  Subdistrict: string = "";
  District: string = "";
  Province: string = "";
  Postal_code: string = "";

  constructor(
    private userService: UsersService,
  ) { }

  ngOnInit(): void {
    this.userService.getUserByID().subscribe({
      next: (data) => {
        console.log('Fetched profile data:', data);
        this.profile = data;
        this.LoadData();
      },
      error: (err) => {
        console.error('Failed to fetch profile', err);
      }
    });
  }

  LoadData() {
    this.Name = this.profile.Name;
    this.Lastname = this.profile.Last_name;
    this.Phone = this.profile.phone || "";
    this.Address = this.profile.address || "";
    this.Subdistrict = this.profile.subdistrict || "";
    this.District = this.profile.district || "";
    this.Province = this.profile.province || "";
    this.Postal_code = this.profile.postal_code || "";
    this.FullAddress = this.Address + ", " + this.Subdistrict + ", " + this.District + ", " + this.Province + ", " + this.Postal_code;
  }

  SaveEditing(editing: string) {
    switch (this.editingBox) {
      case 1:
        this.Name = editing;
        console.log("data saved:", this.Name);
        break;

      case 2:
        this.Lastname = editing;
        console.log("data saved:", this.Lastname);
        break;

      case 3:
        this.Phone = editing;
        console.log("data saved:", this.Phone);
        break;

      default:
        console.warn("Unknown editingBox value:", this.editingBox);
    }
  }

  onAddressUpdate(data: any) {
    this.Address = data.address;
    this.Subdistrict = data.subdistrict;
    this.District = data.district;
    this.Province = data.province;
    this.Postal_code = data.zipcode;
    this.FullAddress = this.Address + ", " + this.Subdistrict + ", " + this.District + ", " + this.Province + ", " + this.Postal_code;
  }

  updateProfile() {
    const body = {
      Name: this.Name,
      Last_name: this.Lastname,
      phone: this.Phone,
      address: this.Address,
      subdistrict: this.Subdistrict,
      district: this.District,
      province: this.Province,
      postal_code: this.Postal_code
    };
    this.userService.updateUser(body).subscribe({
      next: (response) => {
        console.log('Update successful:', response);
        alert('Update successful!');
        this.ngOnInit();
      },
      error: (err) => {
        console.error("update failed:", err.error);
        alert("update failed: " + err.error);
      }
    })
  }
}
