import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddPromoCodePageComponent } from './add-promo-code-page.component';

describe('AddPromoCodePageComponent', () => {
  let component: AddPromoCodePageComponent;
  let fixture: ComponentFixture<AddPromoCodePageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AddPromoCodePageComponent]
    });
    fixture = TestBed.createComponent(AddPromoCodePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
