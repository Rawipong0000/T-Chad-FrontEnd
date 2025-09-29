import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PromoCodePageComponent } from './promo-code-page.component';

describe('PromoCodePageComponent', () => {
  let component: PromoCodePageComponent;
  let fixture: ComponentFixture<PromoCodePageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PromoCodePageComponent]
    });
    fixture = TestBed.createComponent(PromoCodePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
