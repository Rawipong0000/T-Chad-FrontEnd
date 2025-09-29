import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditShopnameComponent } from './edit-shopname.component';

describe('EditShopnameComponent', () => {
  let component: EditShopnameComponent;
  let fixture: ComponentFixture<EditShopnameComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [EditShopnameComponent]
    });
    fixture = TestBed.createComponent(EditShopnameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
