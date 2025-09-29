import { Component, Output, Input, EventEmitter, OnInit } from '@angular/core';
import { UsersService } from 'src/app/service/users.service';
import { Province, District, Subdistrict } from 'src/app/model/users.model';

@Component({
  selector: 'app-edit-address',
  templateUrl: './edit-address.component.html',
  styleUrls: ['./edit-address.component.css']
})
export class EditAddressComponent implements OnInit {
  @Output() closed = new EventEmitter<void>();
  @Input() step = 1;

  Province_list: Province[] = [];
  selectedProvince: Province | null = null;
  District_list: District[] = [];
  selectedDistrict: District | null = null;
  Subdistrict_list: Subdistrict[] = [];
  selectedSubdistrict: Subdistrict | null = null;

  Step: number = this.step;

  Address: string = "";
  ProvinceEdit: string = "";
  DistrictEdit: string = "";
  SubdistrictEdit: string = "";
  ZipcodeEdit: string = "";

  constructor(
    private userService: UsersService,
  ) { }

  ngOnInit(): void {
    this.userService.getProvince().subscribe({
      next: (data) => {
        console.log('Fetched province data:', data);
        this.Province_list = data;
      },
      error: (err) => {
        console.error('Failed to fetch province', err);
      }
    });
  }

  StepRecive(step: number){
    this.Step = step;
  }

  MainAddressEdit() {
    switch (this.Step) {
      case 1:
        const ProvinceID = this.selectedProvince?.province_id ?? 0;
        console.log("ProvinceID:", ProvinceID);
        this.userService.getDistrict(ProvinceID).subscribe({
          next: (data) => {
            this.ProvinceEdit = this.selectedProvince?.name_th ?? "";
            console.log("data saved:", this.ProvinceEdit);
            console.log('Fetched district data:', data);
            this.District_list = data;
            this.Step = 2;
          },
          error: (err) => {
            console.error('Failed to fetch district', err);
            this.Step = 1;
          }
        });
        break;

      case 2:
        const DistrictID = this.selectedDistrict?.district_id ?? 0;
        console.log("DistrictID:", DistrictID);
        this.userService.getSubdistrict(DistrictID).subscribe({
          next: (data) => {
            this.DistrictEdit = this.selectedDistrict?.name_th ?? "";
            console.log("data saved:", this.DistrictEdit);
            console.log('Fetched subdistrict data:', data);
            this.Subdistrict_list = data;
            this.Step = 3;
          },
          error: (err) => {
            console.error('Failed to fetch subdistrict', err);
            this.Step = 2;
          }
        });
        break;

      case 3:
        this.SubdistrictEdit = this.selectedSubdistrict?.name_th ?? "";
        console.log("data saved:", this.SubdistrictEdit);
        this.ZipcodeEdit = this.selectedSubdistrict?.zipcode ?? "";
        console.log("data saved:", this.ZipcodeEdit);
        this.Step = 4;
        break;

      case 4:
        console.log("data saved:", this.Address);
        this.sendAddress()
        this.Step = 1;
        this.closePopup();
        break;

      default:
        console.warn("Unknown editingBox value:");
    }
  }

  closePopup() {
    this.closed.emit();
  }

  @Output() addressChange = new EventEmitter<{
    address: string;
    province: string;
    district: string;
    subdistrict: string;
    zipcode: string;
  }>();

  sendAddress() {
    this.addressChange.emit({
      address: this.Address,
      province: this.ProvinceEdit,
      district: this.DistrictEdit,
      subdistrict: this.SubdistrictEdit,
      zipcode: this.ZipcodeEdit
    });
  }
}
