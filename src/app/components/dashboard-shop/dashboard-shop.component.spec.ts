import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardShopComponent } from './dashboard-shop.component';

describe('DashboardShopComponent', () => {
  let component: DashboardShopComponent;
  let fixture: ComponentFixture<DashboardShopComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DashboardShopComponent]
    });
    fixture = TestBed.createComponent(DashboardShopComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
