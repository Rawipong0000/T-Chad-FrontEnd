import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShopnameEditComponent } from './shopname-edit.component';

describe('ShopnameEditComponent', () => {
  let component: ShopnameEditComponent;
  let fixture: ComponentFixture<ShopnameEditComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ShopnameEditComponent]
    });
    fixture = TestBed.createComponent(ShopnameEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
